package yamlpath

import (
	"errors"
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/compile"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Option is a configuration option that may be specified during compilation to
// customize the behavior of the YAMLPath expression. Options are specified as
// variadic arguments to the [Compile] function.
type Option interface {
	apply(*compile.Config) error
}

type option func(c *compile.Config) error

func (o option) apply(c *compile.Config) error {
	return o(c)
}

var _ Option = (*option)(nil)

var (
	collectionType = reflect.TypeFor[Collection]()
)

// ExternalConstant is a type constraint for values that can be used as external
// constants in YAMLPath expressions. It includes all types that can be
// represented as a YAML node, such as integers, floats, strings, and
// collections of YAML nodes. The type must also be convertible to a YAML node
// or a slice of YAML nodes.
type ExternalConstant interface {
	constraints.Float | constraints.Integer | ~bool | ~string | *yaml.Node | []*yaml.Node |
		Collection
}

// WithConstant is an [Option] that adds an external constant to the
// YAMLPath expression. The constant can be used in the expression by
// referencing it with the name provided.
func WithConstant[T ExternalConstant](name string, value T) Option {
	return option(func(c *compile.Config) error {
		if c.Constants == nil {
			c.Constants = make(map[string][]*yaml.Node)
		}
		return setExternalConstant(c.Constants, name, value)
	})
}

func setExternalConstant[T ExternalConstant](constants map[string][]*yaml.Node, name string, value T) error {
	if _, ok := constants[name]; ok {
		return fmt.Errorf("constant %q already defined", name)
	}

	set := false
	switch v := any(value).(type) {
	case *yaml.Node:
		constants[name] = []*yaml.Node{v}
		set = true
	case []*yaml.Node:
		constants[name] = v
		set = true
	case Collection:
		constants[name] = v
		set = true
	}
	if set {
		return nil
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		constants[name] = yamlconv.Ints(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		constants[name] = yamlconv.Ints(rv.Uint())
	case reflect.Float32, reflect.Float64:
		constants[name] = yamlconv.Floats(rv.Float())
	case reflect.Bool:
		constants[name] = yamlconv.Bools(rv.Bool())
	case reflect.String:
		constants[name] = yamlconv.Strings(rv.String())
	}

	return nil
}

// WithFunction is an option that adds a custom function for the YAMLPath
// evaluation to be able to call. The specified function must have the following
//
//   - At least 1 argument, where the first argument is a [Collection].
//     This will represent the input nodes provided at the time of evaluation.
//   - At least 1 return type, where the first return type is a [Collection].
//     This will represent the output nodes as the result of the function.
//   - If a second return type is provided, the type must be an [error] and
//     indicate whether the function was successful or not. If a non-nil error
//     is returned, the value will be ignored and the function will be considered
//     to have failed.
//   - If additional arguments are provided as inputs to the function, any
//     parameters passed to the function will automatically be decoded to the
//     type of the argument.
//
// For example, a function defined like:
//
//	WithFunction("selected_prefixed", func(c Collection, str string) Collection {
//	    var result Collection
//	    for _, n := range c {
//	        if n.Kind != yaml.ScalarNode || !strings.HasPrefix(n.Value, str) {
//	            continue
//	        }
//	        result = append(result, n)
//	    }
//	    return result
//	})
//
// Can be called in a YAMLPath expression like:
//
//	$.selected_prefixed("foo")
//
// Variadic functions are supported as well.
func WithFunction(name string, fn any) Option {
	return option(func(c *compile.Config) error {
		rv := reflect.ValueOf(fn)
		rt := rv.Type()
		if err := validateFunc(name, rt); err != nil {
			return err
		}

		ar := arity.Exactly(rt.NumIn() - 1)
		if rt.IsVariadic() {
			ar = arity.AtLeast(rt.NumIn() - 2)
		}
		var getValues []toInputFunc

		for i := 1; i < rt.NumIn(); i++ {
			argType := rt.In(i)
			if i == (rt.NumIn()-1) && rt.IsVariadic() {
				getValues = append(getValues, toVariadicInput(argType))
				continue
			}
			if argType.Kind() == reflect.Slice {
				if argType.ConvertibleTo(collectionType) {
					getValues = append(getValues, toCollectionInput(argType))
					continue
				}
			}
			getValues = append(getValues, toInput(argType))
		}

		fn := func(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
			if err := ar.Check(len(params)); err != nil {
				return nil, err
			}

			input := ctx.Current()
			args := make([]reflect.Value, 0, len(params)+1)
			args = append(args, reflect.ValueOf(input))
			for i, getValue := range getValues {
				err := getValue(&args, ctx, params[i:]...)
				if err != nil {
					return nil, err
				}
			}

			results := rv.Call(args)
			if len(results) == 2 {
				if err, ok := results[1].Interface().(error); ok && err != nil {
					return nil, err
				}
			}
			resultt := results[0].Type()
			if resultt.ConvertibleTo(collectionType) {
				return results[0].Convert(collectionType).Interface().(Collection), nil
			}
			if resultt == reflect.TypeFor[*yaml.Node]() {
				return []*yaml.Node{results[0].Interface().(*yaml.Node)}, nil
			}
			node := &yaml.Node{}
			err := node.Encode(results[0].Interface())
			if err != nil {
				return nil, err
			}
			return []*yaml.Node{node}, nil
		}

		c.Table.Add(name, fn).SetArity(ar)

		return nil
	})
}

type toInputFunc = func(out *[]reflect.Value, ctx invocation.Context, params ...invocation.Parameter) error

func toCollectionInput(rt reflect.Type) toInputFunc {
	return func(out *[]reflect.Value, ctx invocation.Context, params ...invocation.Parameter) error {
		collection, err := params[0].GetArg(ctx)
		if err != nil {
			return err
		}
		*out = append(*out, reflect.ValueOf(Collection(collection)).Convert(rt))
		return nil
	}
}

func toVariadicInput(rt reflect.Type) toInputFunc {
	return func(out *[]reflect.Value, ctx invocation.Context, params ...invocation.Parameter) error {
		for _, param := range params {
			collection, err := param.GetArg(ctx)
			if err != nil {
				return err
			}
			routv, err := toValue(rt.Elem(), collection)
			if err != nil {
				return err
			}
			*out = append(*out, routv)
		}
		return nil
	}
}

func toInput(rt reflect.Type) toInputFunc {
	return func(out *[]reflect.Value, ctx invocation.Context, params ...invocation.Parameter) error {
		collection, err := params[0].GetArg(ctx)
		if err != nil {
			return err
		}
		rv, err := toValue(rt, collection)
		if err != nil {
			return err
		}
		*out = append(*out, rv)
		return nil
	}
}

func toValue(rt reflect.Type, collection Collection) (reflect.Value, error) {
	rv := reflect.New(rt).Elem()

	if rv.Kind() != reflect.Slice && len(collection) == 0 {
		return reflect.Value{}, fmt.Errorf("function: %q expected a non-empty collection", rt)
	}
	outv := rv.Addr().Interface()
	if err := collection[0].Decode(outv); err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(outv).Elem(), nil
}

func validateFunc(name string, rt reflect.Type) error {
	if rt.Kind() != reflect.Func {
		return fmt.Errorf("function %q is %w", name, errs.ErrNotAFunction)
	}

	var result []error
	if rt.NumIn() < 1 {
		result = append(result, fmt.Errorf("function %q %w", name, errs.ErrFuncTooFewArguments))
	}
	if rt.NumIn() >= 1 && rt.In(0) != collectionType {
		result = append(result, fmt.Errorf("function %q %w", name, errs.ErrBadFirstArgument))
	}

	if rt.NumOut() == 0 {
		result = append(result, fmt.Errorf("function %q %w must have at least one return value", name, errs.ErrBadReturnSignature))
	} else if rt.NumOut() > 2 {
		result = append(result, fmt.Errorf("function %q %w must have at most two return values", name, errs.ErrBadReturnSignature))
	}
	if rt.NumOut() == 2 && !rt.Out(1).Implements(reflect.TypeFor[error]()) {
		result = append(result, fmt.Errorf("function: %q %w second return type must be an error", name, errs.ErrBadReturnSignature))
	}
	return errors.Join(result...)
}
