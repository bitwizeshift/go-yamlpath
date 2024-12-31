package invocation

import (
	"context"

	"gopkg.in/yaml.v3"
)

// Context is a representation of the current context of the yamlpath
// expression.
type Context interface {
	// Root returns the root node that started the full YAMLPath.
	Root() []*yaml.Node

	// Current returns the current node in the YAMLPath expression.
	// This is typically the node that should be operated on.
	Current() []*yaml.Node

	// NewContext returns a new context with the given node as the current node.
	NewContext([]*yaml.Node) Context

	// WithContext returns a new context with the given context.
	WithContext(context.Context) Context

	// GetContext retrieves the context from the invocation context.
	Context() context.Context
}

// Parameter is a representation of a parameter that can be used in a yamlpath
// expression.
type Parameter interface {
	// GetArg retrieves the argument from the parameter.
	GetArg(Context) ([]*yaml.Node, error)
}

// Func is a representation of a function that can be used in a yamlpath
// expression.
type Func interface {
	// Invoke invokes the function with the given context and arguments.
	Invoke(ctx Context, params ...Parameter) ([]*yaml.Node, error)

	// TestArity tests the arity of the function with the given number of
	// arguments.
	TestArity(n int) error
}
