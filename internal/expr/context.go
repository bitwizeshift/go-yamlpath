package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// ctx represents the current context of the expression evaluation.
type ctx struct {
	root    []*yaml.Node
	current []*yaml.Node
	ctx     context.Context
}

// NewContext creates a new context with the given root node.
func NewContext(root []*yaml.Node) *ctx {
	root = yamlutil.Normalize(root...)
	return &ctx{
		root:    root,
		current: root,
	}
}

// Root returns the root node of the context.
func (c *ctx) Root() []*yaml.Node {
	return c.root
}

// Current returns the current node of the context.
func (c *ctx) Current() []*yaml.Node {
	return c.current
}

// WithContext sets the current node of the context.
func (c *ctx) WithContext(ctx context.Context) invocation.Context {
	c.ctx = ctx
	return c
}

func (c *ctx) Context() context.Context {
	return c.ctx
}

// NewContext creates a new context with a different 'current' node.
func (c *ctx) NewContext(current []*yaml.Node) invocation.Context {
	return &ctx{
		root:    c.root,
		current: current,
	}
}

var _ invocation.Context = (*ctx)(nil)
