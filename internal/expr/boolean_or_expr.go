package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// BooleanOrExpr represents a boolean OR expression.
type BooleanOrExpr struct {
	Left, Right Expr
}

// Eval evaluates the boolean OR expression against the given context.
func (e *BooleanOrExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if yamlutil.IsTruthy(left...) {
		return []*yaml.Node{yamlutil.True}, nil
	}

	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if yamlutil.IsTruthy(right...) {
		return []*yaml.Node{yamlutil.True}, nil
	}

	return []*yaml.Node{yamlutil.False}, nil
}
