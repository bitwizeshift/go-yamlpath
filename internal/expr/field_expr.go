package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type FieldExpression struct {
	Name string
}

func (f *FieldExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes = yamlutil.Normalize(nodes...)
	for _, n := range nodes {
		switch n.Kind {
		case yaml.MappingNode:
			for i := 0; (i + 1) < len(n.Content); i += 2 {
				if f.Name == "*" || n.Content[i].Value == f.Name {
					result = append(result, n.Content[i+1])
				}
			}
		case yaml.SequenceNode:
			if f.Name == "*" {
				result = append(result, n.Content...)
			}
		}
	}
	return result, nil
}
