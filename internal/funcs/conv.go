package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func singleton(name string, ctx invocation.Context) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError(name, current)
	}
	return current, nil
}

func paramToString(name string, ctx invocation.Context, param invocation.Parameter) (string, error) {
	nodes, err := param.GetArg(ctx)
	if err != nil {
		return "", err
	}
	return toString(name, nodes)
}

func toString(name string, nodes []*yaml.Node) (string, error) {
	node := nodes[0]
	if node.Kind != yaml.ScalarNode {
		return "", errs.NewKindError(name, node, yaml.ScalarNode)
	}
	if node.Tag != "!!str" {
		return "", errs.NewTagError(name, node, "!!str")
	}
	return node.Value, nil
}

func toInt(name string, nodes []*yaml.Node) (int64, error) {
	node := nodes[0]
	if node.Kind != yaml.ScalarNode {
		return 0, errs.NewKindError(name, node, yaml.ScalarNode)
	}
	if node.Tag != "!!int" {
		return 0, errs.NewTagError(name, node, "!!int")
	}
	return yamlconv.ParseInt(node)
}

func toBool(name string, nodes []*yaml.Node) (bool, error) {
	node := nodes[0]
	if node.Kind != yaml.ScalarNode {
		return false, errs.NewKindError(name, node, yaml.ScalarNode)
	}
	if node.Tag != "!!bool" {
		return false, errs.NewTagError(name, node, "!!bool")
	}
	return yamlconv.ParseBool(node)
}
