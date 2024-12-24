package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

type RecursiveDescentExpression struct{}

func (r *RecursiveDescentExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node
	for _, node := range nodes {
		result = append(result, r.flatten(node)...)
	}
	return result, nil
}

func (r *RecursiveDescentExpression) flatten(node *yaml.Node) []*yaml.Node {
	var result []*yaml.Node
	switch node.Kind {
	case yaml.DocumentNode:
		return r.flatten(node.Content[0])
	case yaml.SequenceNode:
		result = append(result, node)
		for _, n := range node.Content {
			result = append(result, r.flatten(n)...)
		}
	case yaml.MappingNode:
		result = append(result, node)
		for i := 0; (i + 1) < len(node.Content); i += 2 {
			result = append(result, r.flatten(node.Content[i+1])...)
		}
	default:
		result = append(result, node)
	}
	return result
}