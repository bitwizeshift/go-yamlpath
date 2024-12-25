package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type NegationExpr struct {
	Expr Expr
}

func (e *NegationExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return []*yaml.Node{yamlutil.True}, nil
	}

	if len(result) > 1 {
		return []*yaml.Node{yamlutil.False}, nil
	}
	b, err := yamlutil.ToBool(result[0])
	if err != nil {
		return []*yaml.Node{yamlutil.False}, nil
	}
	return []*yaml.Node{yamlutil.FromBool(!b)}, nil
}
