package expr

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlpathctx"
)

// RootExpression represents a root expression, either '$' or '@' in the path.
type RootExpression struct {
	Root string
}

// Eval evaluates the root expression.
func (e *RootExpression) Eval(ctx context.Context, _ []*yaml.Node) ([]*yaml.Node, error) {
	if e.Root == "$" {
		return yamlpathctx.GetRoot(ctx), nil
	}
	if e.Root == "@" {
		return yamlpathctx.GetCurrent(ctx), nil
	}
	return nil, fmt.Errorf("unsupported root expression: %s", e.Root)
}
