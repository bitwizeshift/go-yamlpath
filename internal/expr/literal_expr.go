package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// LiteralExpr is an expression that always returns the specified value.
// This is primarily used to return literals that are specified in the
// path.
type LiteralExpr struct {
	Nodes []*yaml.Node
}

// Eval evaluates the expression and returns the specified value.
func (e *LiteralExpr) Eval(invocation.Context) ([]*yaml.Node, error) {
	return e.Nodes, nil
}

var _ Expr = (*LiteralExpr)(nil)
