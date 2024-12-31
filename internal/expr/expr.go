package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// Expr represents an expression that can be evaluated against a YAML
// node.
type Expr interface {
	// Eval evaluates the expression against the given context.
	Eval(ctx invocation.Context) ([]*yaml.Node, error)
}
