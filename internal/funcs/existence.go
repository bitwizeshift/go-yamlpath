package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Empty returns true if the current node is empty, false otherwise.
func Empty(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}

// Exists returns true if the current node exists, false otherwise.
func Exists(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	if len(params) == 0 {
		current := ctx.Current()
		if len(current) > 0 {
			return []*yaml.Node{yamlutil.True}, nil
		}
		return []*yaml.Node{yamlutil.False}, nil
	}
	next, err := Where(ctx, params...)
	if err != nil {
		return nil, err
	}

	if len(next) > 0 {
		return []*yaml.Node{yamlutil.True}, nil
	}
	return []*yaml.Node{yamlutil.False}, nil
}
