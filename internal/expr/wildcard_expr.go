package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type WildcardExpression struct{}

func (*WildcardExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
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
