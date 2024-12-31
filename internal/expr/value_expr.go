package expr

import (
	"gopkg.in/yaml.v3"
)

// ValueExpr is an expression that always returns the specified value.
// This is primarily used to return literals that are specified in the
// path.
type ValueExpr struct {
	Nodes []*yaml.Node
}

// Eval evaluates the expression and returns the specified value.
func (e *ValueExpr) Eval(*Context) ([]*yaml.Node, error) {
	return e.Nodes, nil
}

var _ Expr = (*ValueExpr)(nil)
