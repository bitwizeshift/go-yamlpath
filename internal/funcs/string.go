package funcs

import (
	"strings"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Lower converts a singleton string input into lowercase.
func Lower(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("lower()", current)
	}

	node := current[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("lower()", node, yaml.ScalarNode)
	}
	if node.Tag != "!!str" {
		return nil, errs.NewTagError("lower()", node, "!!str")
	}

	return []*yaml.Node{yamlconv.String(strings.ToLower(node.Value))}, nil
}

// Upper converts a singleton string input into uppercase.
func Upper(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("upper()", current)
	}

	node := current[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("upper()", node, yaml.ScalarNode)
	}
	if node.Tag != "!!str" {
		return nil, errs.NewTagError("upper()", node, "!!str")
	}
	return []*yaml.Node{yamlconv.String(strings.ToUpper(node.Value))}, nil
}
