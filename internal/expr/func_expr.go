package expr

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// FuncExpr represents an expression that is implemented by a function
// that takes a context and returns a list of nodes.
type FuncExpr struct {
	Parameters []invocation.Parameter
	Func       invocation.Func
}

// Eval evaluates the expression by invoking the function with the given
// context and parameters.
func (e *FuncExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	return e.Func.Invoke(ctx, e.Parameters...)
}
