package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

// PrefixPlusExpr is a representation of the unary '+' operator in YAMLPath.
//
// This is effectively an identity expression that will only return the value
// back to the caller -- with the one exception that it will only operate on
// numeric values.
type PrefixPlusExpr struct {
	Expr Expr
}

// Eval evaluates the unary '+' operator against the provided nodes.
func (e *PrefixPlusExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) != 1 {
		return nil, NewSingletonError("operator prefix '+'", len(result))
	}
	node := result[0]
	if node.Kind != yaml.ScalarNode {
		return nil, NewKindError("operator prefix '+'", yaml.ScalarNode, node.Kind)
	}
	if node.Tag != "!!int" && node.Tag != "!!float" {
		return nil, NewTagError("operator prefix '+'", "numeric", node.Tag)
	}
	return []*yaml.Node{node}, nil
}

var _ Expr = (*PrefixPlusExpr)(nil)
