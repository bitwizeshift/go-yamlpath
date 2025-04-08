package expr

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// ConcatExpr represents an expression that concatenates two expressions
// together using the '+' operator.
//
// For strings, this produces a true concatenation of the strings. For
// integers, this is the addition operator. If either side of the operator
// is not compatible, this will return an error.
type ConcatExpr struct {
	Left, Right Expr
}

func (e *ConcatExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, right, err := e.eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(left) > 1 {
		return nil, errs.NewSingletonError("operator + lhs", left)
	}
	if len(right) > 1 {
		return nil, errs.NewSingletonError("operator + rhs", right)
	}
	if len(left) == 0 || len(right) == 0 {
		return nil, nil
	}
	lhs, rhs := left[0], right[0]
	if lhs.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("operator + lhs", lhs, yaml.ScalarNode)
	}
	if rhs.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("operator + rhs", rhs, yaml.ScalarNode)
	}
	if (lhs.Tag == "!!int" || lhs.Tag == "!!float") && (rhs.Tag == "!!float" || rhs.Tag == "!!int") {
		lv, err := decimal.NewFromString(lhs.Value)
		if err != nil {
			return nil, err
		}
		rv, err := decimal.NewFromString(rhs.Value)
		if err != nil {
			return nil, err
		}
		prod := lv.Add(rv)
		return []*yaml.Node{yamlconv.RawNumber(prod.String())}, nil
	} else if lhs.Tag == "!!str" && rhs.Tag == "!!str" {
		concat := lhs.Value + rhs.Value
		return []*yaml.Node{yamlconv.String(concat)}, nil
	}
	return nil, errs.NewIncompatibleError("operator +", lhs, rhs)
}

func (e *ConcatExpr) eval(ctx invocation.Context) ([]*yaml.Node, []*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, nil, err
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, nil, err
	}
	return left, right, nil
}

var _ Expr = (*ConcatExpr)(nil)
