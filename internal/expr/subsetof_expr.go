package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// SubsetOfExpr checks if the left-hand side is a subset of the right-hand side
// of the expression.
type SubsetOfExpr struct {
	Left, Right Expr
}

// Eval evaluates the subset-of expression by checking if the left-hand side
// is a subset of the right-hand side.
func (e *SubsetOfExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	for _, l := range left {
		found := false
		for _, r := range right {
			if yamlutil.Equal(l, r) {
				found = true
				break
			}
		}
		if !found {
			return []*yaml.Node{yamlutil.False}, nil
		}
	}
	return []*yaml.Node{yamlutil.True}, nil
}

var _ Expr = (*SubsetOfExpr)(nil)
