package expr

import (
	"context"

	"gopkg.in/yaml.v3"
)

// Expression represents an expression that can be evaluated against a YAML
// node.
type Expression interface {
	// Eval evaluates the expression against the given context.
	Eval(ctx context.Context, node []*yaml.Node) ([]*yaml.Node, error)
}