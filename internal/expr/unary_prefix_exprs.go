package expr

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type PrefixPlusExpr struct {
	Expr Expr
}

func (e *PrefixPlusExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	result, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("expected one node, but got %d", len(result))
	}
	node := result[0]
	if node.Kind != yaml.ScalarNode {
		return nil, fmt.Errorf("expected a scalar node, but got %v", node.Kind)
	}
	if node.Tag != "!!int" && node.Tag != "!!float" {
		return nil, fmt.Errorf("expected a number node, but got %s", node.Tag)
	}
	return []*yaml.Node{node}, nil
}

var _ Expr = (*PrefixPlusExpr)(nil)

type PrefixMinusExpr struct {
	Expr Expr
}

func (e *PrefixMinusExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	rest, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(rest) == 0 {
		return nil, nil
	}
	if len(rest) != 1 {
		return nil, fmt.Errorf("expected one node, but got %d", len(rest))
	}
	node := rest[0]
	if node.Kind != yaml.ScalarNode {
		return nil, fmt.Errorf("expected a scalar node, but got %v", node.Kind)
	}

	var result []*yaml.Node
	switch node.Tag {
	case "!!int":
		result = append(result, yamlutil.Number("-"+node.Value))
	case "!!float":
		if node.Value[0] == '-' {
			result = append(result, yamlutil.Number(node.Value[1:]))
		} else {
			result = append(result, yamlutil.Number("-"+node.Value))
		}
	default:
		return nil, fmt.Errorf("expected a number node, but got %s", node.Tag)
	}
	return result, nil
}
