package yamlpathctx

import (
	"context"

	"gopkg.in/yaml.v3"
)

type ctxKey int

const (
	ctxKeyRoot ctxKey = iota
	ctxKeyCurrent
)

// SetRoot sets the root node in the context.
func SetRoot(ctx context.Context, root []*yaml.Node) context.Context {
	return context.WithValue(ctx, ctxKeyRoot, root)
}

// GetRoot returns the root node from the context.
func GetRoot(ctx context.Context) []*yaml.Node {
	root := ctx.Value(ctxKeyRoot)
	if root == nil {
		return nil
	}
	return root.([]*yaml.Node)
}

// SetCurrent sets the current node in the context.
func SetCurrent(ctx context.Context, current []*yaml.Node) context.Context {
	return context.WithValue(ctx, ctxKeyCurrent, current)
}

// GetCurrent returns the current node from the context.
func GetCurrent(ctx context.Context) []*yaml.Node {
	current := ctx.Value(ctxKeyCurrent)
	if current == nil {
		return nil
	}
	return current.([]*yaml.Node)
}
