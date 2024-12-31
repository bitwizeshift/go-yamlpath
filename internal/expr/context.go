package expr

import (
	"context"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Context represents the current context of the expression evaluation.
type Context struct {
	root    []*yaml.Node
	current []*yaml.Node
	ctx     context.Context
}

// NewContext creates a new context with the given root node.
func NewContext(root []*yaml.Node) *Context {
	root = yamlutil.Normalize(root...)
	return &Context{
		root:    root,
		current: root,
	}
}

// Root returns the root node of the context.
func (c *Context) Root() []*yaml.Node {
	return c.root
}

// Current returns the current node of the context.
func (c *Context) Current() []*yaml.Node {
	return c.current
}

// WithContext sets the current node of the context.
func (c *Context) WithContext(ctx context.Context) *Context {
	c.ctx = ctx
	return c
}

// SubContext creates a new context with a different 'current' node.
func (c *Context) SubContext(current []*yaml.Node) *Context {
	return &Context{
		root:    c.root,
		current: current,
	}
}
