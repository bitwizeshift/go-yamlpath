package expr

import (
	"context"
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
func (e *ScriptExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	scriptNodes, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}

	if len(scriptNodes) != 1 {
		return nil, fmt.Errorf("script expression must evaluate to a single value")
	}

	key, err := e.toKey(scriptNodes[0])
	if err != nil {
		return nil, err
	}

	switch key := key.(type) {
	case int:
		expr := &IndexExpr{Indices: []int64{int64(key)}}
		return expr.Eval(ctx, nodes)
	case string:
		expr := &FieldExpr{Names: []string{key}}
		return expr.Eval(ctx, nodes)
	}
	return nil, fmt.Errorf("script expression must evaluate to an integer or string")
}

func (e *ScriptExpr) toKey(node *yaml.Node) (any, error) {
	if node.Kind != yaml.ScalarNode {
		return nil, fmt.Errorf("script expression must evaluate to a scalar")
	}
	if node.Tag == "!!int" {
		return yamlutil.ToInt(node)
	}
	if node.Tag == "!!str" {
		return yamlutil.ToString(node)
	}
	return nil, fmt.Errorf("script expression must evaluate to an integer or string")
}
