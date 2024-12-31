package expr

import (
	"gopkg.in/yaml.v3"
)

// FieldExpr extracts the fields from the nodes that match the named fields.
type FieldExpr struct {
	Fields []string
}

// Eval evaluates the field expression by extracting the fields from the nodes
// that match the named fields.
func (e *FieldExpr) Eval(ctx *Context) ([]*yaml.Node, error) {
	var result []*yaml.Node

	nodes := ctx.Current()
	for _, n := range nodes {
		if n.Kind != yaml.MappingNode {
			continue
		}
		for _, name := range e.Fields {
			for i := 0; (i + 1) < len(n.Content); i += 2 {
				if n.Content[i].Value == name {
					result = append(result, n.Content[i+1])
				}
			}
		}
	}
	return result, nil
}

var _ Expr = (*FieldExpr)(nil)
