package funcs

import (
	"strconv"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
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
			d, err := yamlconv.ParseDecimal(node)
			if err != nil {
				return nil, err
			}
			if d.Equal(decimal.Zero) {
				return []*yaml.Node{yamlconv.Bool(false)}, nil
			}
			if d.Equal(decimal.NewFromInt(1)) {
				return []*yaml.Node{yamlconv.Bool(true)}, nil
			}
		case "!!str":
			switch node.Value {
			case "true", "1", "1.0":
				return []*yaml.Node{yamlconv.Bool(true)}, nil
			case "false", "0", "0.0":
				return []*yaml.Node{yamlconv.Bool(false)}, nil
			}
		}
	}
	return nil, nil
}

// ConvertsToBoolean checks if the current node can be converted to a boolean
// value.
func ConvertsToBoolean(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("convertsToBoolean()", current)
	}
	node := current[0]

	if node.Kind != yaml.ScalarNode {
		return yamlconv.Bools(false), nil
	}
	switch node.Tag {
	case "!!bool":
		return yamlconv.Bools(true), nil
	case "!!int", "!!float":
		d, err := yamlconv.ParseDecimal(node)
		if err == nil && (d.Equal(decimal.Zero) || d.Equal(decimal.NewFromInt(1))) {
			return yamlconv.Bools(true), nil
		}
	case "!!str":
		switch node.Value {
		case "true", "1", "1.0", "false", "0", "0.0":
			return yamlconv.Bools(true), nil
		}
	}
	return yamlconv.Bools(false), nil
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
		return []*yaml.Node{yamlconv.String(node.Value)}, nil
	}
	return nil, nil
}

// ConvertsToString checks if the current node can be converted to a string
// value.
func ConvertsToString(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("convertsToString()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	case yaml.MappingNode, yaml.SequenceNode:
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
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
			b, err := yamlconv.ParseBool(node)
			if err != nil {
				return nil, nil
			}
			if b {
				return []*yaml.Node{yamlconv.Number(1)}, nil
			}
			return []*yaml.Node{yamlconv.Number(0)}, nil
		case "!!str":
			if _, err := decimal.NewFromString(node.Value); err == nil {
				return []*yaml.Node{yamlconv.NumberString(node.Value)}, nil
			}
		}
	}
	return nil, nil
}

// ConvertsToNumber checks if the current node can be converted to a number
// value.
func ConvertsToNumber(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("convertsToNumber()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!int", "!!float", "!!bool":
			return yamlconv.Bools(true), nil
		case "!!str":
			if _, err := decimal.NewFromString(node.Value); err == nil {
				return yamlconv.Bools(true), nil
			}
		}
	}
	return yamlconv.Bools(false), nil
}

// ToInteger converts the current node to an integer value.
func ToInteger(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("toInteger()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!int":
			return []*yaml.Node{node}, nil
		case "!!bool":
			if node.Value == "true" {
				return []*yaml.Node{yamlconv.Int(1)}, nil
			}
			return []*yaml.Node{yamlconv.Int(0)}, nil
		case "!!str":
			i, err := strconv.ParseInt(node.Value, 10, 64)
			if err != nil {
				return nil, nil
			}
			return []*yaml.Node{yamlconv.Int(i)}, nil
		}
	}
	return nil, nil
}

// ConvertsToInteger checks if the current node can be converted to an integer
// value.
func ConvertsToInteger(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("convertsToInteger()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!int", "!!bool":
			return yamlconv.Bools(true), nil
		case "!!str":
			if _, err := strconv.ParseInt(node.Value, 10, 64); err == nil {
				return yamlconv.Bools(true), nil
			}
		}
	}
	return yamlconv.Bools(false), nil
}

// ToFloat converts the current node to a float value.
func ToFloat(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("toFloat()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!float":
			return []*yaml.Node{node}, nil
		case "!!int":
			_, err := yamlconv.ParseInt(node)
			if err != nil {
				return nil, nil
			}
			return []*yaml.Node{yamlconv.FloatString(node.Value)}, nil
		case "!!bool":
			if node.Value == "true" {
				return []*yaml.Node{yamlconv.Float(1.0)}, nil
			}
			return []*yaml.Node{yamlconv.Float(0.0)}, nil
		case "!!str":
			f, err := strconv.ParseFloat(node.Value, 64)
			if err != nil {
				return nil, nil
			}
			return []*yaml.Node{yamlconv.Float(f)}, nil
		}
	}
	return nil, nil
}

// ConvertsToFloat checks if the current node can be converted to a float
// value.
func ConvertsToFloat(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("convertsToFloat()", current)
	}

	node := current[0]
	switch node.Kind {
	case yaml.ScalarNode:
		switch node.Tag {
		case "!!float", "!!int", "!!bool":
			return yamlconv.Bools(true), nil
		case "!!str":
			if _, err := decimal.NewFromString(node.Value); err == nil {
				return yamlconv.Bools(true), nil
			}
		}
	}
	return yamlconv.Bools(false), nil
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
