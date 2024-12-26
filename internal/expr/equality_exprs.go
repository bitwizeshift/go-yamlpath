package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type EqualityExpr struct {
	Left, Right Expr
}

func (e *EqualityExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}

	if yamlutil.EqualRange(left, right) {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}

type InequalityExpr struct {
	Left, Right Expr
}

func (e *InequalityExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}

	if !yamlutil.EqualRange(left, right) {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}