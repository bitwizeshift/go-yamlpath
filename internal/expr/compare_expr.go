package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Comparator is a comparator function that represents a comparison operator.
type Comparator func(i int) bool

var (
	// CompareLess represents the '<' operator
	CompareLess Comparator = func(i int) bool { return i < 0 }
	// CompareGreater represents the '>' operator
	CompareGreater Comparator = func(i int) bool { return i > 0 }
	// CompareLessEqual represents the '<=' operator
	CompareLessEqual Comparator = func(i int) bool { return i <= 0 }
	// CompareGreaterEqual represents the '>=' operator
	CompareGreaterEqual Comparator = func(i int) bool { return i >= 0 }
)

// CompareExpr is a representation of the various comparison operators for
// YAMLPath ('<=', '<', '>', '>='). Which operator is used here depends solely
// on the Compare field, which may be one of the [Comparator] definitions.
type CompareExpr struct {
	Left, Right Expr
	Compare     Comparator
}

func (e *CompareExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}

	cmp, err := yamlutil.CompareRange(left, right)
	if err != nil {
		return nil, err
	}
	if e.Compare(cmp) {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}

var _ Expr = (*CompareExpr)(nil)
