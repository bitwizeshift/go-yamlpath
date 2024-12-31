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

// Select performs a map operation, similar to [()] in JSONPath.
func Select(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
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
