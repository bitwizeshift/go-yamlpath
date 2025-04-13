/*
Package invocationtest provides test-doubles for the invocation package
*/
package invocationtest

import (
	"golang.org/x/exp/constraints"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

type result struct {
	value []*yaml.Node
	err   error
}

// FakeParameter is a fake parameter that can be used in tests.
// This implements the [invocation.Parameter] interface.
type FakeParameter struct {
	expr    func(invocation.Context) ([]*yaml.Node, error)
	results []*result
}

// AddSuccess sets up the parameter to return a successful value on the next
// available invocation.
func (fp *FakeParameter) AddSuccess(nodes ...*yaml.Node) *FakeParameter {
	fp.results = append(fp.results, &result{
		value: nodes,
	})
	return fp
}

// AddError sets up the parameter to return an error value on the next available
// invocation.
func (fp *FakeParameter) AddError(err error) *FakeParameter {
	fp.results = append(fp.results, &result{
		err: err,
	})
	return fp
}

func (fp *FakeParameter) next(invocation.Context) ([]*yaml.Node, error) {
	if len(fp.results) == 0 {
		return nil, nil
	}
	next := fp.results[0]
	fp.results = fp.results[1:]
	return next.value, next.err
}

// GetArg returns the value of this parameter given the context.
func (fp *FakeParameter) GetArg(ctx invocation.Context) ([]*yaml.Node, error) {
	return fp.expr(ctx)
}

var _ invocation.Parameter = (*FakeParameter)(nil)

// NewParameter creates a new [FakeParameter] with the given function behavior.
func NewParameter() *FakeParameter {
	fp := &FakeParameter{}
	fp.expr = fp.next
	return fp
}

// SuccessParameter creates a new [FakeParameter] that returns the given nodes.
func SuccessParameter(nodes ...*yaml.Node) *FakeParameter {
	return NewParameter().AddSuccess(nodes...)
}

// Int creates a new [FakeParameter] that returns the given integer value.
func Int[T constraints.Integer](i T) *FakeParameter {
	return SuccessParameter(yamlconv.Int(i))
}

// String creates a new [FakeParameter] that returns the given string value.
func String(s string) *FakeParameter {
	return SuccessParameter(yamlconv.String(s))
}

// Bool creates a new [FakeParameter] that returns the given boolean value.
func Bool(b bool) *FakeParameter {
	return SuccessParameter(yamlconv.Bool(b))
}

// Float creates a new [FakeParameter] that returns the given float value.
func Float[T constraints.Float](f T) *FakeParameter {
	return SuccessParameter(yamlconv.Float(f))
}

// Number creates a new [FakeParameter] that returns the given number value.
func Number[T constraints.Integer | constraints.Float](n T) *FakeParameter {
	return SuccessParameter(yamlconv.Number(n))
}

// Node creates a new [FakeParameter] that returns the given node value.
func Node[T yamlconv.Primitive](v T) *FakeParameter {
	return SuccessParameter(yamlconv.Node(v))
}

// Sequence creates a new [FakeParameter] that returns a YAML sequence node
func Sequence(nodes ...*yaml.Node) *FakeParameter {
	return SuccessParameter(yamlconv.Sequence(nodes...))
}

// Ints creates a new [FakeParameter] that returns the given integer values as
// a collection of nodes.
func Ints[T constraints.Integer](n ...T) *FakeParameter {
	return SuccessParameter(yamlconv.Ints(n...)...)
}

// Floats creates a new [FakeParameter] that returns the given float values as
// a collection of nodes.
func Floats[T constraints.Float](f ...T) *FakeParameter {
	return SuccessParameter(yamlconv.Floats(f...)...)
}

// Numbers creates a new [FakeParameter] that returns the given number values
// as a collection of nodes.
func Numbers[T constraints.Integer | constraints.Float](n ...T) *FakeParameter {
	return SuccessParameter(yamlconv.Numbers(n...)...)
}

// Strings creates a new [FakeParameter] that returns the given string values
// as a collection of nodes.
func Strings(s ...string) *FakeParameter {
	return SuccessParameter(yamlconv.Strings(s...)...)
}

// Bools creates a new [FakeParameter] that returns the given boolean values
// as a collection of nodes.
func Bools(b ...bool) *FakeParameter {
	return SuccessParameter(yamlconv.Bools(b...)...)
}

// Nodes creates a new [FakeParameter] that returns the given primitive
// values as a collection of nodes.
func Nodes[T yamlconv.Primitive](v ...T) *FakeParameter {
	return SuccessParameter(yamlconv.Nodes(v...)...)
}

// Error creates a new [FakeParameter] that always returns an error.
func Error(err error) *FakeParameter {
	return NewParameter().AddError(err)
}
