package expr

import (
	"regexp"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// MatchExpr represents a regular expression match expression in YAMLPath.
type MatchExpr struct {
	Regex *regexp.Regexp
	Expr  Expr
}

// Eval evaluates the match expression against the given nodes.
func (e *MatchExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, err := e.Expr.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(left) != 1 {
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}
	node := left[0]

	if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}

	if e.Regex.MatchString(node.Value) {
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

var _ Expr = (*MatchExpr)(nil)
