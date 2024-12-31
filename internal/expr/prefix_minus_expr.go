package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type PrefixMinusExpr struct {
	Expr Expr
}

func (e *PrefixMinusExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	rest, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(rest) == 0 {
		return nil, nil
	}
	if len(rest) != 1 {
		return nil, NewSingletonError("operator prefix '-'", len(rest))
	}
	node := rest[0]
	if node.Kind != yaml.ScalarNode {
		return nil, NewKindError("operator prefix '-'", yaml.ScalarNode, node.Kind)
	}

	var result []*yaml.Node
	switch node.Tag {
	case "!!int", "!!float":
		if node.Value[0] == '-' {
			result = append(result, yamlutil.Number(node.Value[1:]))
		} else {
			result = append(result, yamlutil.Number("-"+node.Value))
		}
	default:
		return nil, NewTagError("operator prefix '-'", "numeric", node.Tag)
	}
	return result, nil
}

var _ Expr = (*PrefixMinusExpr)(nil)
