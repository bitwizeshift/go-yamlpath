package expr

import (
	"context"
	"regexp"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// MatchExpr represents a regular expression match expression in the AST.
type MatchExpr struct {
	Regex *regexp.Regexp
	Expr  Expr
}

// Eval evaluates the match expression against the given nodes.
func (e *MatchExpr) Eval(ctx context.Context, nodes []*yaml.Node) ([]*yaml.Node, error) {
	left, err := e.Expr.Eval(ctx, nodes)
	if err != nil {
		return nil, err
	}
	if len(left) != 1 {
		return []*yaml.Node{yamlutil.False}, nil
	}
	node := left[0]

	if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
		return []*yaml.Node{yamlutil.False}, nil
	}

	if e.Regex.MatchString(node.Value) {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}

var _ Expr = (*MatchExpr)(nil)
