/*
Package invocation provides a mechanism for defining and invoking functions in
a yamlpath expression.
*/
package invocation

import (
	"errors"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
)

// ErrUnknownFunc is an error that is returned when a function is not found in
// the function table.
var ErrUnknownFunc = errors.New("unknown function")

// Table is an invocation table that maps function names to function definitions.
type Table struct {
	entries map[string]*Entry

	// parent is the parent table that this table is derived from.
	parent *Table
}

// Entry is a function entry in the function table.
type Entry struct {
	fn    funcEntry
	arity arity.Arity
}

// SetArity sets the arity of the function entry.
func (e *Entry) SetArity(a arity.Arity) {
	e.arity = a
}

// TestArity tests the arity of the function with the given number of arguments.
func (e *Entry) TestArity(n int) error {
	return e.arity.Check(n)
}

// Invoke invokes the function with the given context and arguments.
func (e *Entry) Invoke(ctx Context, params ...Parameter) ([]*yaml.Node, error) {
	if err := e.TestArity(len(params)); err != nil {
		return nil, err
	}

	return e.fn(ctx, params...)
}

var _ Func = (*Entry)(nil)

// NewTable creates a new function table.
func NewTable() *Table {
	return &Table{
		entries: make(map[string]*Entry),
	}
}

// New creates a new function table from the current parent table.
func (t *Table) New() *Table {
	return &Table{
		entries: make(map[string]*Entry),
		parent:  t,
	}
}

// Add adds a function to the table.
// By default, functions have an arity of zero -- meaning no arguments may be
// provided to them. To set the arity of the function, use the [Entry.SetArity]
// method on the returned [Entry].
func (t *Table) Add(name string, fn funcEntry) *Entry {
	entry := &Entry{
		fn:    fn,
		arity: arity.None(),
	}
	t.entries[name] = entry
	return entry
}

// Lookup performs a function lookup in the table, and if the function is not
// found, it will recursively search the parent table. If the function is not
// found in the parent table, ok will be set to false and fn will be nil.
func (t *Table) Lookup(name string) (fn Func, ok bool) {
	if t == nil || t.entries == nil {
		return nil, false
	}
	entry, ok := t.entries[name]
	if !ok {
		return t.parent.Lookup(name)
	}

	return entry, true
}

type funcEntry = func(ctx Context, params ...Parameter) ([]*yaml.Node, error)
