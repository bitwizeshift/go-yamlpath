package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

type UnionExpression struct {
	Union Union
}

func (e *UnionExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node
	for _, node := range nodes {
		result = append(result, e.Union.Index(node)...)
	}
	return result, nil
}

var _ Expression = (*UnionExpression)(nil)
