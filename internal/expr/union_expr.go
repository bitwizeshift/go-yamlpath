package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// UnionExpr represents a union of expressions that are evaluated
// in sequence to produce a single result collection.
type UnionExpr []Expr

// Eval evaluates the union of expressions in sequence, returning the
// collection of results from each expression.
func (e UnionExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var results []*yaml.Node
	for _, expr := range e {
		nodes, err := expr.Eval(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, nodes...)
	}
	return results, nil
}
