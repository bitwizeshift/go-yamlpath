package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// IndexExpr is a representation of the indexing operator `[...]` in YAMLPath
// for numeric indexing of sequences.
//
// This implements both union-indexing, which allows multiple selection, as well
// as just individual indexing. If an index is out-of-bounds, it is not selected
// -- no error is returned.
//
// Negative indices allow selecting from the reverse side.
type IndexExpr struct {
	Indices Expr
}

// Eval evaluates the index expression against the given nodes.
func (i *IndexExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()

	next, err := i.Indices.Eval(ctx)
	if err != nil {
		return nil, err
	}
	var indices []int64
	for _, n := range next {
		index, err := yamlconv.ParseInt(n)
		if err != nil {
			return nil, errs.IncludeSource(err, "index expression")
		}
		indices = append(indices, index)
	}

	for _, n := range nodes {
		for _, index := range indices {
			if n.Kind != yaml.SequenceNode {
				continue
			}
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

var _ Expr = (*IndexExpr)(nil)
