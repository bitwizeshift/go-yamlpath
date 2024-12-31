package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// InExpr checks for the presence of a single value in a collection of values.
type InExpr struct {
	Left, Right Expr
}

// Eval evaluates the in expression by checking if the left-hand side is in the
// right-hand side.
func (e *InExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(left) == 0 {
		return nil, nil
	}
	if len(left) != 1 {
		return nil, NewSingletonError("operator 'in'", len(left))
	}

	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	// All non-empty sets contain the empty set
	if len(right) == 0 {
		return nil, nil
	}
	right = e.unwrap(right)

	l := left[0]
	for _, r := range right {
		if yamlutil.Equal(l, r) {
			return []*yaml.Node{yamlutil.True}, nil
		}
	}
	return []*yaml.Node{yamlutil.False}, nil
}

// unwrap nodes by examining whether the result is a single sequence node and,
// if it is, unwrapping it so that the "in" operator may compare against the
// individual elements of the sequence.
func (e *InExpr) unwrap(nodes []*yaml.Node) []*yaml.Node {
	if len(nodes) == 1 && nodes[0].Kind == yaml.SequenceNode {
		return nodes[0].Content
	}
	return nodes
}

var _ Expr = (*InExpr)(nil)
