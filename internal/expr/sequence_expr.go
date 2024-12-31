package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// SequenceExpr is a representation of a sequence of expressions in YAMLPath.
//
// Almost every path is composed of sequence expressions that build up the
// path. For example, the path `$.foo.bar` is composed of three sequence
// expressions: `$`, `.foo`, and `.bar`.
type SequenceExpr []Expr

// Eval evaluates the sequence of expressions.
func (s SequenceExpr) Eval(ctx invocation.Context) (nodes []*yaml.Node, err error) {
	for _, expr := range s {
		nodes, err = expr.Eval(ctx)
		if err != nil {
			return nil, err
		}
		ctx = ctx.NewContext(nodes)
	}
	return nodes, nil
}

// Append appends an expression to the sequence.
func (s *SequenceExpr) Append(expr Expr) {
	if seq, ok := expr.(SequenceExpr); ok {
		*s = append(*s, seq...)
	} else {
		*s = append(*s, expr)
	}
}

var _ Expr = (*SequenceExpr)(nil)
