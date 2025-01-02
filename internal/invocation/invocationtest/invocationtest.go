/*
Package invocationtest provides test-doubles for the invocation package
*/
package invocationtest

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

type result struct {
	value []*yaml.Node
	err   error
}

// FakeParameter is a fake parameter that can be used in tests.
// This implements the [invocation.Parameter] interface.
type FakeParameter struct {
	expr    func(invocation.Context) ([]*yaml.Node, error)
	results []*result
}

// AddSuccess sets up the parameter to return a successful value on the next
// available invocation.
func (fp *FakeParameter) AddSuccess(nodes ...*yaml.Node) *FakeParameter {
	fp.results = append(fp.results, &result{
		value: nodes,
	})
	return fp
}

// AddError sets up the parameter to return an error value on the next available
// invocation.
func (fp *FakeParameter) AddError(err error) *FakeParameter {
	fp.results = append(fp.results, &result{
		err: err,
	})
	return fp
}

func (fp *FakeParameter) next(invocation.Context) ([]*yaml.Node, error) {
	if len(fp.results) == 0 {
		return nil, nil
	}
	next := fp.results[0]
	fp.results = fp.results[1:]
	return next.value, next.err
}

// GetArg returns the value of this parameter given the context.
func (fp *FakeParameter) GetArg(ctx invocation.Context) ([]*yaml.Node, error) {
	return fp.expr(ctx)
}

var _ invocation.Parameter = (*FakeParameter)(nil)

// NewParameter creates a new [FakeParameter] with the given function behavior.
func NewParameter() *FakeParameter {
	fp := &FakeParameter{}
	fp.expr = fp.next
	return fp
}

// SuccessParameter creates a new [FakeParameter] that returns the given nodes.
func SuccessParameter(nodes ...*yaml.Node) *FakeParameter {
	return NewParameter().AddSuccess(nodes...)
}

// ErrorParameter creates a new [FakeParameter] that always returns an error.
func ErrorParameter(err error) *FakeParameter {
	return NewParameter().AddError(err)
}
