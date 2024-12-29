package expr

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type ArithmeticExpr struct {
	Left, Right Expr
	Operation   func(lhs, rhs decimal.Decimal) decimal.Decimal
}

func (e *ArithmeticExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, right, err := e.eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(left) > 1 || len(right) > 1 {
		return nil, fmt.Errorf("multiplicative expressions must have exactly one left and right value")
	}
	if len(left) == 0 || len(right) == 0 {
		return nil, nil
	}
	lhs, rhs := left[0], right[0]
	if lhs.Kind != yaml.ScalarNode || rhs.Kind != yaml.ScalarNode {
		return nil, fmt.Errorf("multiplicative expressions must have scalar nodes")
	}
	if lhs.Tag != "!!int" && lhs.Tag != "!!float" || rhs.Tag != "!!float" && rhs.Tag != "!!int" {
		return nil, fmt.Errorf("multiplicative expressions must have scalar nodes of type int or float")
	}
	lv, err := decimal.NewFromString(lhs.Value)
	if err != nil {
		return nil, err
	}
	rv, err := decimal.NewFromString(rhs.Value)
	if err != nil {
		return nil, err
	}
	prod := e.Operation(lv, rv)

	return []*yaml.Node{yamlutil.Number(prod.String())}, nil
}

func (e *ArithmeticExpr) eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, []*yaml.Node, error) {
	left, err := e.Left.Eval(ctx, nodes)
	if err != nil {
		return nil, nil, err
	}
	right, err := e.Right.Eval(ctx, nodes)
	if err != nil {
		return nil, nil, err
	}
	return left, right, nil
}