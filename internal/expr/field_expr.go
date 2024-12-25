package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type FieldExpression struct {
	Names []string
}

func (e *FieldExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes = yamlutil.Normalize(nodes...)
	for _, n := range nodes {
		if n.Kind != yaml.MappingNode {
			continue
		}
		for _, name := range e.Names {
			for i := 0; (i + 1) < len(n.Content); i += 2 {
				if n.Content[i].Value == name {
					result = append(result, n.Content[i+1])
				}
			}
		}
	}
	return result, nil
}
