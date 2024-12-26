package expr

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// InExpr checks for the presence of a single value in a collection of values.
type InExpr struct {
	Left, Right Expr
}

// Eval evaluates the in expression by checking if the left-hand side is in the
// right-hand side.
func (e *InExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}

	if len(left) != 1 {
		return nil, fmt.Errorf("in operator requires exactly one left-hand value")
	}
	l := left[0]
	for _, r := range right {
		if yamlutil.Equal(l, r) {
			return []*yaml.Node{yamlutil.True}, nil
		}
	}
	return []*yaml.Node{yamlutil.False}, nil
}
