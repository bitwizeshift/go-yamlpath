package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// WildcardBracketExpr is a representation of the `[*]` expression in
// YAMLPath, which will select all fields of a sequence of elements from the
// current node.
type WildcardBracketExpr struct{}

// Eval evaluates the wildcard expression against the provided nodes.
func (*WildcardBracketExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
	for _, n := range nodes {
		if n.Kind != yaml.SequenceNode {
			continue
		}
		result = append(result, n.Content...)
	}
	return result, nil
}

var _ Expr = (*WildcardExpr)(nil)
