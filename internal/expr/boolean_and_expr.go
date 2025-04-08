package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
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
	if !yamlconv.IsTruthy(left...) {
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}

	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if !yamlconv.IsTruthy(right...) {
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}

	return []*yaml.Node{yamlconv.Bool(true)}, nil
}
