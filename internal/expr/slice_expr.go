package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// SliceExpr represents a slice expression
type SliceExpr struct {
	Slice *Slice
}

// Eval evaluates the slice expression against the provided nodes
// and returns the resulting nodes that match the slice expression.
func (s *SliceExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes = yamlutil.Normalize(nodes...)
	for _, n := range nodes {
		if n.Kind != yaml.SequenceNode {
			continue
		}
		end := min(len(n.Content), s.Slice.End)
		start := max(0, s.Slice.Start)
		step := s.Slice.Step
		for i := start; i < end; i += step {
			result = append(result, n.Content[i])
		}
	}
	return result, nil
}
