package expr

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// ScriptExpr represents a script expression that can be used to create
// computed fields in a YAMLPath.
type ScriptExpr struct {
	Expr Expr
}

// Eval evaluates the script expression by evaluating the expression and
// using the result as a key to extract a value from the input nodes.
func (e *ScriptExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	scriptNodes, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if len(scriptNodes) != 1 {
		return nil, NewSingletonError("script operator '(...)'", len(scriptNodes))
	}
	node := scriptNodes[0]

	if node.Kind != yaml.ScalarNode {
		return nil, NewKindError("script operator '(...)'", yaml.ScalarNode, node.Kind)
	}

	expr, err := e.createExpr(scriptNodes[0])
	if err != nil {
		return nil, err
	}
	return expr.Eval(ctx)
}

func (e *ScriptExpr) createExpr(node *yaml.Node) (Expr, error) {
	switch node.Tag {
	case "!!int":
		key, err := yamlutil.ToInt(node)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrEval, err)
		}
		return &IndexExpr{Indices: []int64{int64(key)}}, nil

	case "!!str":
		return &FieldExpr{Fields: []string{node.Value}}, nil
	}
	return nil, NewTagError("script operator '(...)'", "!!str or !!int", node.Tag)
}

var _ Expr = (*ScriptExpr)(nil)
