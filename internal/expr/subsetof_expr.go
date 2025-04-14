package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
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
	left = e.unwrap(left)
	right = e.unwrap(right)

	for _, l := range left {
		found := false
		for _, r := range right {
			if yamlcmp.Equal(l, r) {
				found = true
				break
			}
		}
		if !found {
			return yamlconv.Bools(false), nil
		}
	}
	return yamlconv.Bools(true), nil
}

// unwrap nodes by examining whether the result is a single sequence node and,
// if it is, unwrapping it so that the "in" operator may compare against the
// individual elements of the sequence.
func (e *SubsetOfExpr) unwrap(nodes []*yaml.Node) []*yaml.Node {
	if len(nodes) == 1 && nodes[0].Kind == yaml.SequenceNode {
		return nodes[0].Content
	}
	return nodes
}

var _ Expr = (*SubsetOfExpr)(nil)
