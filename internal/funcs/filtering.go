package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Where performs a filter operation, similar to [?()] in JSONPath.
// It takes a single expression argument which is evaluated for each element in
// the current collection. If the expression returns true, the element is
// included in the result.
func Where(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	var result []*yaml.Node
	current := ctx.Current()

	for _, node := range current {
		args, err := params[0].GetArg(ctx.NewContext([]*yaml.Node{node}))
		if err != nil {
			return nil, err
		}
		if yamlconv.IsTruthy(args...) {
			result = append(result, node)
		}
	}
	return result, nil
}

// Transform performs a map operation, similar to [()] in JSONPath.
func Transform(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	var result []*yaml.Node
	current := ctx.Current()

	for _, node := range current {
		args, err := params[0].GetArg(ctx.NewContext([]*yaml.Node{node}))
		if err != nil {
			return nil, err
		}
		result = append(result, args...)
	}
	return result, nil
}

// Keys returns the keys of the current mapping node.
func Keys(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	var result []*yaml.Node
	current := ctx.Current()

	for _, node := range current {
		if node.Kind != yaml.MappingNode {
			continue
		}
		for i := 0; (i + 1) < len(node.Content); i += 2 {
			result = append(result, node.Content[i])
		}
	}
	return result, nil
}

// Select returns the selected keys from the current mapping or the selected
// indices from the current sequence.
func Select(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	var (
		keys    []string
		indices []int64
	)
	for _, p := range params {
		input, err := p.GetArg(ctx)
		if err != nil {
			return nil, err
		}
		if len(input) == 0 {
			continue
		}
		if len(input) > 1 {
			return nil, errs.NewSingletonError("select()", input)
		}
		node := input[0]
		if node.Kind != yaml.ScalarNode {
			return nil, errs.NewKindError("select()", node, yaml.ScalarNode)
		}
		switch node.Tag {
		case "!!int":
			index, err := yamlconv.ParseInt(node)
			if err != nil {
				return nil, errs.NewEvalError(err)
			}
			indices = append(indices, index)
		case "!!str":
			keys = append(keys, node.Value)
		default:
			return nil, errs.NewTagError("select()", node, "!!int", "!!str")
		}
	}

	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	var result []*yaml.Node
	for _, node := range current {
		switch node.Kind {
		case yaml.MappingNode:
			result = append(result, selectKeys(node, keys)...)
		case yaml.SequenceNode:
			result = append(result, selectIndices(node, indices)...)
		}
	}

	return result, nil
}

func selectKeys(node *yaml.Node, keys []string) []*yaml.Node {
	var result []*yaml.Node
	for i := 0; i < len(node.Content); i += 2 {
		key := node.Content[i]
		value := node.Content[i+1]
		for _, k := range keys {
			if key.Value == k {
				result = append(result, value)
				break
			}
		}
	}
	return result
}

func selectIndices(node *yaml.Node, indices []int64) []*yaml.Node {
	var result []*yaml.Node
	for _, index := range indices {
		if index < 0 {
			index = int64(len(node.Content)) + index
		}
		if index < 0 || index >= int64(len(node.Content)) {
			continue
		}
		result = append(result, node.Content[index])
	}
	return result
}
