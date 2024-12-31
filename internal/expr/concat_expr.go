package expr

import (
	"fmt"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type ConcatExpr struct {
	Left, Right Expr
}

func (e *ConcatExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	left, right, err := e.eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(left) > 1 {
		return nil, NewSingletonError("concatenation", len(left))
	}
	if len(right) > 1 {
		return nil, NewSingletonError("concatenation", len(right))
	}
	if len(left) == 0 || len(right) == 0 {
		return nil, nil
	}
	lhs, rhs := left[0], right[0]
	if lhs.Kind != yaml.ScalarNode {
		return nil, NewKindError("concatenation", lhs.Kind, yaml.ScalarNode)
	}
	if rhs.Kind != yaml.ScalarNode {
		return nil, NewKindError("concatenation", rhs.Kind, yaml.ScalarNode)
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
		return []*yaml.Node{yamlutil.Number(prod.String())}, nil
	} else if lhs.Tag == "!!str" && rhs.Tag == "!!str" {
		concat := lhs.Value + rhs.Value
		return []*yaml.Node{yamlutil.String(concat)}, nil
	}
	return nil, fmt.Errorf("%w: %s and %s are not compatible for '+' operator", ErrEval, lhs.Tag, rhs.Tag)
}

func (e *ConcatExpr) eval(ctx *Context) ([]*yaml.Node, []*yaml.Node, error) {
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
