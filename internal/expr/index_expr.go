package expr

import (
	"gopkg.in/yaml.v3"
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
	Indices []int64
}

// Eval evaluates the index expression against the given nodes.
func (i *IndexExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
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

var _ Expr = (*IndexExpr)(nil)
