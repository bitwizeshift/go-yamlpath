package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
		if yamlutil.IsTruthy(args...) {
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
