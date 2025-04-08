package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// NegationExpr is a representation of the '!' operator in YAMLPath expressions.
// It negates the result of the sub-expression.
type NegationExpr struct {
	Expr Expr
}

// Eval evaluates the '!' operator against the provided nodes.
func (e *NegationExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if yamlconv.IsTruthy(result...) {
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(true)}, nil
}
