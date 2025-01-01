package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// RootExpr represents a root expression, either '$' or '@' in the path.
type RootExpr struct {
	Root string
}

// Eval evaluates the root expression.
func (e *RootExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	if e.Root == "$" {
		return ctx.Root(), nil
	}
	if e.Root == "@" {
		return ctx.Current(), nil
	}
	return nil, errs.NewUnsupportedErrorf("unhandled root expression %q", e.Root)
}

var _ Expr = (*RootExpr)(nil)
