package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type FilterExpr struct {
	Expr Expr
}

func (f *FilterExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node
	for _, node := range nodes {
		filtered, err := f.Expr.Eval(ctx, []*yaml.Node{node})
		if err != nil {
			return nil, err
		}

		if yamlutil.IsTruthy(filtered...) {
			result = append(result, node)
		}
	}
	return result, nil
}
