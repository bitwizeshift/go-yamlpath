package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

// ValueExpr is an expression that always returns the specified value.
// This is primarily used to return literals that are specified in the
// path.
type ValueExpr struct {
	Node *yaml.Node
}

// Eval evaluates the expression and returns the specified value.
func (e *ValueExpr) Eval(context.Context, []*yaml.Node) ([]*yaml.Node, error) {
	return []*yaml.Node{e.Node}, nil
}

var _ Expr = (*ValueExpr)(nil)
