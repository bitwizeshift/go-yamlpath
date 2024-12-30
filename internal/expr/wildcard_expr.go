package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// WildcardExpr is a representation of the `.*` and `[*]` expressions in
// YAMLPath, which will select all fields or sequence elements from the
// current node.
type WildcardExpr struct{}

// Eval evaluates the wildcard expression against the provided nodes.
func (*WildcardExpr) Eval(_ context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes = yamlutil.Normalize(nodes...)
	for _, n := range nodes {
		switch n.Kind {
		case yaml.MappingNode:
			for i := 0; (i + 1) < len(n.Content); i += 2 {
				result = append(result, n.Content[i+1])
			}
		case yaml.SequenceNode:
			result = append(result, n.Content...)
		}
	}
	return result, nil
}
