package expr_test

import (
	"context"
	"testing"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// ExprFunc is a function that implements the Expr interface
type ExprFunc func(ctx context.Context, node []*yaml.Node) ([]*yaml.Node, error)

// Eval calls the function
func (f ExprFunc) Eval(ctx context.Context, node []*yaml.Node) ([]*yaml.Node, error) {
	return f(ctx, node)
}

var _ expr.Expr = (*ExprFunc)(nil)

// ExprReturnsError returns an Expr that returns an error.
func ExprReturnsError(err error) expr.Expr {
	return ExprFunc(func(ctx context.Context, node []*yaml.Node) ([]*yaml.Node, error) {
		return nil, err
	})
}

// ExprReturnsNodes returns an Expr that returns the given nodes.
func ExprReturnsNodes(nodes ...*yaml.Node) expr.Expr {
	return ExprFunc(func(ctx context.Context, node []*yaml.Node) ([]*yaml.Node, error) {
		return nodes, nil
	})
}

func YAML(t *testing.T, s string) *yaml.Node {
	return yamlutil.Normalize(Document(t, s))[0]
}

func Document(t *testing.T, s string) *yaml.Node {
	var n yaml.Node
	if err := yaml.Unmarshal([]byte(s), &n); err != nil {
		t.Fatal(err)
	}
	return &n
}
