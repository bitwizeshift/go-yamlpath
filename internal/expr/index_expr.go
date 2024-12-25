package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

type IndexExpr struct {
	Indices []int64
}

func (i *IndexExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes = yamlutil.Normalize(nodes...)
	for _, n := range nodes {
		if n.Kind != yaml.SequenceNode {
			continue
		}
		for _, index := range i.Indices {
			if index < 0 {
				index += int64(len(n.Content))
			}
			if index < 0 || index >= int64(len(n.Content)) {
				continue
			}
			result = append(result, n.Content[index])
		}
	}
	return result, nil
}
