package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

type PrefixMinusExpr struct {
	Expr Expr
}

func (e *PrefixMinusExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	rest, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(rest) == 0 {
		return nil, nil
	}
	if len(rest) != 1 {
		return nil, errs.NewSingletonError("operator prefix '-'", rest)
	}
	node := rest[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("operator prefix '-'", node, yaml.ScalarNode)
	}

	var result []*yaml.Node
	switch node.Tag {
	case "!!int", "!!float":
		if node.Value[0] == '-' {
			result = append(result, yamlconv.RawNumber(node.Value[1:]))
		} else {
			result = append(result, yamlconv.RawNumber("-"+node.Value))
		}
	default:
		return nil, errs.NewTagError("operator prefix '-'", node, "!!int", "!!float")
	}
	return result, nil
}

var _ Expr = (*PrefixMinusExpr)(nil)
