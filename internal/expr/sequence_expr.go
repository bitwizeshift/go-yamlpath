package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

type SequenceExpression []Expression

func (s SequenceExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var err error
	for _, expr := range s {
		nodes, err = expr.Eval(ctx, nodes)
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

var _ Expression = (*SequenceExpression)(nil)
