package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

type UnionExpr struct {
	Union Union
}

func (e *UnionExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node
	for _, node := range nodes {
		result = append(result, e.Union.Index(node)...)
	}
	return result, nil
}

var _ Expr = (*UnionExpr)(nil)
