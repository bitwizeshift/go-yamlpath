package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// EqualityExpr is a representation of the '==' operator in the YAMLPath
// grammar. This enables comparison across both left and right sub-expressions.
type EqualityExpr struct {
	Left, Right Expr
}

// Eval evaluates the equality expression against the given nodes.
func (e *EqualityExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if yamlutil.EqualRange(left, right) {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}
