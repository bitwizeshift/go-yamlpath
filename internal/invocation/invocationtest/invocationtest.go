/*
Package invocationtest provides test-doubles for the invocation package
*/
package invocationtest

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// FakeParameter is a fake parameter that can be used in tests.
// This implements the [invocation.Parameter] interface.
type FakeParameter struct {
	expr func(invocation.Context) ([]*yaml.Node, error)
}

// GetArg returns the value of this parameter given the context.
func (f *FakeParameter) GetArg(ctx invocation.Context) ([]*yaml.Node, error) {
	return f.expr(ctx)
}

var _ invocation.Parameter = (*FakeParameter)(nil)

// NewParameter creates a new [FakeParameter] with the given function behavior.
func NewParameter(fn func(invocation.Context) ([]*yaml.Node, error)) *FakeParameter {
	return &FakeParameter{expr: fn}
}

// SuccessParameter creates a new [FakeParameter] that returns the given nodes.
func SuccessParameter(nodes ...*yaml.Node) *FakeParameter {
	return NewParameter(func(invocation.Context) ([]*yaml.Node, error) {
		return nodes, nil
	})
}

// ErrorParameter creates a new [FakeParameter] that always returns an error.
func ErrorParameter(err error) *FakeParameter {
	return NewParameter(func(invocation.Context) ([]*yaml.Node, error) {
		return nil, err
	})
}
