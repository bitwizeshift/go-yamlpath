/*
Package errs is an internal package that provides shared errors across the
go-yamlpath library.

This package largely exists so that errors can be reused across both the
[expr], [yamlpath], and [funcs] packages without causing import cycles.
*/
package errs

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	// ErrEval is returned when an error occurs during evaluation of an
	// expression.
	ErrEval = errors.New("yamlpath eval")

	// ErrBadKind is returned when a node encounters an unexpected kind during
	// evaluation.
	//
	// This value is a sentinel value to test if the error is caused by a bad
	// kind; to get actual details, use [errors.Is] with [KindError].
	ErrBadKind = errors.New("bad kind")

	// ErrBadTag is returned when a node encounters an unexpected tag during
	// evaluation.
	//
	// This value is a sentinel value to test if the error is caused by a bad
	// tag; to get actual details, use [errors.Is] with [TagError].
	ErrBadTag = errors.New("bad tag")

	// ErrNotSingleton is returned when a collection is not a singleton
	ErrNotSingleton = errors.New("collection is not singleton")

	// ErrUnsupported is returned when an operation is specified that is either
	// not supported or not yet implemented.
	ErrUnsupported = errors.New("unsupported operation")

	// ErrIncompatible is returned when two operands are incompatible.
	ErrIncompatible = errors.New("incompatible operands")

	// ErrNotAFunction is an error returned when a non-function is provided to
	// an option expecting a function
	ErrNotAFunction = errors.New("not a function")

	// ErrFuncTooFewArguments is an error returned when a function is passed to
	// WithFunction with too few arguments.
	ErrFuncTooFewArguments = errors.New("must have at least one argument")

	// ErrBadFirstArgument is an error returned when a function is passed to
	// WithFunction with a first argument that is not a collection of nodes.
	ErrBadFirstArgument = errors.New("first argument must be a collection of nodes")

	// ErrBadReturnSignature is an error returned when a function is passed to
	// WithFunction with either too few or too many return types.
	ErrBadReturnSignature = errors.New("bad return signature")
)

// KindError represents an error that occurs when a node encounters an
// unexpected kind during evaluation.
type KindError struct {
	Source   string
	Expected []yaml.Kind
	Node     *yaml.Node
}

// NewKindError creates a new error that indicates that a node has an unexpected kind.
func NewKindError(source string, got *yaml.Node, want ...yaml.Kind) error {
	return &KindError{
		Source:   source,
		Expected: want,
		Node:     got,
	}
}

// Error returns a string representation of the error.
func (e *KindError) Error() string {
	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, ErrEval.Error())
	_, _ = fmt.Fprint(&sb, ":")
	if e.Source != "" {
		_, _ = fmt.Fprintf(&sb, " %s", e.Source)
	}
	_, _ = fmt.Fprint(&sb, " expected YAML node with kind")
	_, _ = fmt.Fprintf(&sb, " '%s'", kindToString(e.Expected[0]))
	if len(e.Expected) > 1 {
		for i := 1; i < len(e.Expected)-1; i++ {
			_, _ = fmt.Fprintf(&sb, ", '%s'", kindToString(e.Expected[i]))
		}
		if len(e.Expected) > 2 {
			_, _ = fmt.Fprint(&sb, ",")
		}
		_, _ = fmt.Fprintf(&sb, " or '%s'", kindToString(e.Expected[len(e.Expected)-1]))
	}
	_, _ = fmt.Fprintf(&sb, ", but received '%s'", kindToString(e.Node.Kind))
	return sb.String()
}

// Unwrap returns the errors that caused this error.
func (e *KindError) Unwrap() []error {
	return []error{ErrEval, ErrBadKind}
}

var _ error = (*KindError)(nil)

// TagError represents an error that occurs when a node encounters an unexpected
// tag.
type TagError struct {
	Source   string
	Expected []string
	Node     *yaml.Node
}

// NewTagError creates a new error that indicates that a node has an unexpected tag.
func NewTagError(source string, node *yaml.Node, want ...string) error {
	return &TagError{
		Source:   source,
		Expected: want,
		Node:     node,
	}
}

// Error returns a string representation of the error.
func (e *TagError) Error() string {
	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, ErrEval.Error())
	_, _ = fmt.Fprintf(&sb, ": %s expected YAML node with tag", e.Source)
	_, _ = fmt.Fprintf(&sb, " '%s'", e.Expected[0])
	if len(e.Expected) > 1 {
		for i := 1; i < len(e.Expected)-1; i++ {
			_, _ = fmt.Fprintf(&sb, ", '%s'", e.Expected[i])
		}
		if len(e.Expected) > 2 {
			_, _ = fmt.Fprint(&sb, ",")
		}
		_, _ = fmt.Fprintf(&sb, " or '%s'", e.Expected[len(e.Expected)-1])
	}
	_, _ = fmt.Fprintf(&sb, ", but received '%s'", e.Node.Tag)
	return sb.String()
}

// Unwrap returns the errors that caused this error.
func (e *TagError) Unwrap() []error {
	return []error{ErrEval, ErrBadTag}
}

var _ error = (*TagError)(nil)

// NotSingletonError represents an error that occurs during evaluation of an expression.
type NotSingletonError struct {
	Source string
	Nodes  []*yaml.Node
}

// NewSingletonError creates a new error that indicates that a collection is not a singleton.
func NewSingletonError(source string, nodes []*yaml.Node) error {
	return &NotSingletonError{
		Source: source,
		Nodes:  nodes,
	}
}

// Unwrap returns the errors that caused this error.
func (e *NotSingletonError) Unwrap() []error {
	return []error{ErrEval, ErrNotSingleton}
}

// Error returns a string representation of the error.
func (e *NotSingletonError) Error() string {
	var sb strings.Builder
	_, _ = fmt.Fprint(&sb, ErrEval.Error())
	_, _ = fmt.Fprintf(&sb, ": %s expected YAML singleton collection,", e.Source)
	_, _ = fmt.Fprintf(&sb, " but received %d nodes", len(e.Nodes))
	return sb.String()
}

var _ error = (*NotSingletonError)(nil)

// NewUnsupportedErrorf creates a new error that indicates that an operation is
// not supported or not yet implemented.
func NewUnsupportedErrorf(format string, args ...any) error {
	return fmt.Errorf("%w: %w: %s", ErrEval, ErrUnsupported, fmt.Sprintf(format, args...))
}

// NewIncompatibleError creates a new error that indicates that two operands
// are incompatible.
func NewIncompatibleError(source string, lhs, rhs *yaml.Node) error {
	return fmt.Errorf("%w: %s has %w (%s,%s) and (%s,%s)",
		ErrEval,
		source,
		ErrIncompatible,
		kindToString(lhs.Kind), lhs.Tag,
		kindToString(rhs.Kind), rhs.Tag,
	)
}

// NewEvalError creates a new error that indicates that an error occurred during
// evaluation of an expression. If the error is already an evaluation error, it
// is returned as-is.
func NewEvalError(err error) error {
	if errors.Is(err, ErrEval) {
		return err
	}
	return fmt.Errorf("%w: %w", ErrEval, err)
}

func kindToString(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.MappingNode:
		return "mapping"
	case yaml.ScalarNode:
		return "scalar"
	case yaml.SequenceNode:
		return "sequence"
	case yaml.AliasNode:
		return "alias"
	}
	return fmt.Sprintf("unknown(%d)", kind)
}

// IncludeSource adds a source to the error if it is an [KindError], or
// [TagError], or [NotSingletonError].
func IncludeSource(err error, source string) error {
	switch err := err.(type) {
	case *NotSingletonError:
		err.Source = source
	case *KindError:
		err.Source = source
	case *TagError:
		err.Source = source
	}
	return err
}
