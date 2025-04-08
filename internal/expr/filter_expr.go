package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// FilterExpr is a representation of the filter operator `[?(...)]` in YAMLPath.
type FilterExpr struct {
	Expr Expr
}

// Eval evaluates the filter expression against the provided nodes.
func (f *FilterExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
	for _, node := range nodes {
		filtered, err := f.Expr.Eval(ctx.NewContext([]*yaml.Node{node}))
		if err != nil {
			return nil, err
		}

		if yamlconv.IsTruthy(filtered...) {
			result = append(result, node)
		}
	}
	return result, nil
}
