package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlpathctx"
)

type SequenceExpression []Expression

func (s SequenceExpression) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	ctx = yamlpathctx.SetRoot(ctx, nodes)

	var err error
	for _, expr := range s {
		nodes, err = expr.Eval(yamlpathctx.SetCurrent(ctx, nodes), nodes)
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// Append appends an expression to the sequence.
func (s *SequenceExpression) Append(expr Expression) {
	if seq, ok := expr.(SequenceExpression); ok {
		*s = append(*s, seq...)
	} else {
		*s = append(*s, expr)
	}
}

var _ Expression = (*SequenceExpression)(nil)
