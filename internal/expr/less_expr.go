package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type LessExpr struct {
	Left, Right Expr
}

func (e *LessExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}

	less, err := yamlutil.LessRange(left, right)
	if err != nil {
		return nil, err
	}
	if less {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}
