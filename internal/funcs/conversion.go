package funcs

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// ToBoolean converts the current node to a boolean value.
func ToBoolean(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("toBoolean()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!bool":
			return []*yaml.Node{node}, nil
		case "!!int", "!!float":
			d, err := yamlutil.ToDecimal(node)
			if err != nil {
				return nil, err
			}
			if d.Equal(decimal.Zero) {
				return []*yaml.Node{yamlutil.False}, nil
			}
			if d.Equal(decimal.NewFromInt(1)) {
				return []*yaml.Node{yamlutil.True}, nil
			}
		case "!!str":
			switch node.Value {
			case "true", "1", "1.0":
				return []*yaml.Node{yamlutil.True}, nil
			case "false", "0", "0.0":
				return []*yaml.Node{yamlutil.False}, nil
			}
		}
	}
	return nil, nil
}

// ToString converts the current node to a string value.
func ToString(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("toString()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		return []*yaml.Node{yamlutil.String(node.Value)}, nil
	}
	return nil, nil
}

// ToNumber converts the current node to a number value.
func ToNumber(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("toNumber()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!int", "!!float":
			return []*yaml.Node{node}, nil
		case "!!bool":
			b, err := yamlutil.ToBool(node)
			if err != nil {
				return nil, nil
			}
			if b {
				return []*yaml.Node{yamlutil.Number("1")}, nil
			}
			return []*yaml.Node{yamlutil.Number("0")}, nil
		case "!!str":
			if _, err := decimal.NewFromString(node.Value); err == nil {
				return []*yaml.Node{yamlutil.Number(node.Value)}, nil
			}
		}
	}
	return nil, nil
}

// ToSequence converts the current node collection into a sequence value.
func ToSequence(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	if len(current) == 0 {
		return nil, nil
	}

	node := &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     "!!seq",
		Content: current,
	}

	return []*yaml.Node{node}, nil
}
