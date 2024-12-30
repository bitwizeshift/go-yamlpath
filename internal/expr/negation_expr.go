package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// NegationExpr is a representation of the '!' operator in YAMLPath expressions.
// It negates the result of the sub-expression.
type NegationExpr struct {
	Expr Expr
}

// Eval evaluates the '!' operator against the provided nodes.
func (e *NegationExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if yamlutil.IsTruthy(result...) {
		return []*yaml.Node{yamlutil.False}, nil
	}
	return []*yaml.Node{yamlutil.True}, nil
}
