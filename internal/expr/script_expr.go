package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// ScriptExpr represents a script expression that can be used to create
// computed fields in a YAMLPath.
type ScriptExpr struct {
	Expr Expr
}

// Eval evaluates the script expression by evaluating the expression and
// using the result as a key to extract a value from the input nodes.
func (e *ScriptExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	scriptNodes, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}

	if len(scriptNodes) != 1 {
		return nil, errs.NewSingletonError("script operator '(...)'", scriptNodes)
	}
	node := scriptNodes[0]

	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("script operator '(...)'", node, yaml.ScalarNode)
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
			return nil, errs.NewEvalError(err)
		}
		return &IndexExpr{Indices: []int64{int64(key)}}, nil

	case "!!str":
		return &FieldExpr{Fields: []string{node.Value}}, nil
	}
	return nil, errs.NewTagError("script operator '(...)'", node, "!!str", "!!int")
}

var _ Expr = (*ScriptExpr)(nil)
