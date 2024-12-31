package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// BooleanAndExpr represents a boolean AND expression.
type BooleanAndExpr struct {
	Left, Right Expr
}

// Eval evaluates the boolean AND expression against the given context.
func (e *BooleanAndExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if !yamlutil.IsTruthy(left...) {
		return []*yaml.Node{yamlutil.False}, nil
	}

	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if !yamlutil.IsTruthy(right...) {
		return []*yaml.Node{yamlutil.False}, nil
	}

	return []*yaml.Node{yamlutil.True}, nil
}
