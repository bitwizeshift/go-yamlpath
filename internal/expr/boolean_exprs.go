package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// BooleanAndExpr represents a boolean AND expression.
type BooleanAndExpr struct {
	Left, Right Expr
}

// Eval evaluates the boolean AND expression against the given context.
func (e *BooleanAndExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if !yamlutil.IsTruthy(left...) {
		return []*yaml.Node{yamlutil.False}, nil
	}

	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if !yamlutil.IsTruthy(right...) {
		return []*yaml.Node{yamlutil.False}, nil
	}

	return []*yaml.Node{yamlutil.True}, nil
}

// BooleanOrExpr represents a boolean OR expression.
type BooleanOrExpr struct {
	Left, Right Expr
}

// Eval evaluates the boolean OR expression against the given context.
func (e *BooleanOrExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if yamlutil.IsTruthy(left...) {
		return []*yaml.Node{yamlutil.True}, nil
	}

	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if yamlutil.IsTruthy(right...) {
		return []*yaml.Node{yamlutil.True}, nil
	}

	return []*yaml.Node{yamlutil.False}, nil
}
