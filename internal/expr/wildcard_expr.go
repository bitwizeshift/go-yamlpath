package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// WildcardExpr is a representation of the `.*` expression in
// YAMLPath, which will select all fields or sequence elements from the
// current node.
type WildcardExpr struct{}

// Eval evaluates the wildcard expression against the provided nodes.
func (*WildcardExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
	for _, n := range nodes {
		if n.Kind != yaml.MappingNode {
			continue
		}
		for i := 0; (i + 1) < len(n.Content); i += 2 {
			result = append(result, n.Content[i+1])
		}
	}
	return result, nil
}

var _ Expr = (*WildcardExpr)(nil)
