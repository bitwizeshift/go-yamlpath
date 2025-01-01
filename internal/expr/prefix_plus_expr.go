package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
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
func (e *PrefixPlusExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) != 1 {
		return nil, errs.NewSingletonError("operator prefix '+'", result)
	}
	node := result[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("operator prefix '+'", node, yaml.ScalarNode)
	}
	if node.Tag != "!!int" && node.Tag != "!!float" {
		return nil, errs.NewTagError("operator prefix '+'", node, "!!int", "!!float")
	}
	return []*yaml.Node{node}, nil
}

var _ Expr = (*PrefixPlusExpr)(nil)
