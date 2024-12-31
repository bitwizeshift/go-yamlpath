package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// SliceExpr is a representation of the `[start:end:step]` slice expression in
// YAMLPath, which will select a range of elements from a sequence node.
// Any non-sequence node objects are filtered and ignored.
type SliceExpr struct {
	Slice *Slice
}

// Eval evaluates the slice expression against the provided nodes
// and returns the resulting nodes that match the slice expression.
func (s *SliceExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
	for _, n := range nodes {
		if n.Kind != yaml.SequenceNode {
			continue
		}
		end := min(len(n.Content), s.Slice.End)
		start := max(0, s.Slice.Start)
		step := s.Slice.Step
		if step == 0 {
			step = 1
		}
		for i := start; i < end; i += step {
			result = append(result, n.Content[i])
		}
	}
	return result, nil
}

var _ Expr = (*SliceExpr)(nil)
