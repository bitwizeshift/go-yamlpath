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

func paramToString(ctx invocation.Context, param invocation.Parameter) (string, error) {
	nodes, err := param.GetArg(ctx)
	if err != nil {
		return "", err
	}
	return yamlconv.ParseString(nodes...)
}
