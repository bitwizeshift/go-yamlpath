package yamlpath

import (
	"rodusek.dev/pkg/yamlpath/internal/compile"
	"rodusek.dev/pkg/yamlpath/internal/errs"
)

var (
	// ErrSyntax is returned when an error occurs during parsing.
	ErrSyntax = compile.ErrSyntax

	// ErrCompile is returned when an error occurs during compilation.
	ErrCompile = compile.ErrCompile

	// ErrEval is returned when an error occurs during evaluation.
	ErrEval = errs.ErrEval

	// ErrNotSingleton is returned when a collection is expected to contain
	// only a single node, but contains more or less than one node.
	//
	// More details can be retrieved from these errors by using [errors.As] with
	// [NotSingletonError].
	ErrNotSingleton = errs.ErrNotSingleton

	// ErrBadKind is returned when an error occurs during evaluation due to
	// a bad YAML node kind.
	//
	// More details can be retrieved from these errors by using [errors.As] with
	// [KindError].
	ErrBadKind = errs.ErrBadKind

	// ErrBadTag is returned when an error occurs during evaluation due to
	// a bad YAML node tag.
	//
	// More details can be retrieved from these errors by using [errors.As] with
	// [TagError].
	ErrBadTag = errs.ErrBadTag
)

// CompileError represents an error that occurs during compilation.
type CompileError = compile.CompileError

// NotSingletonError is an evaluation error returned when an expression is
// expecting a singleton value, but receives a collection with more or less
// than one node.
//
// This error can be tested for with [errors.Is] and retrieved for with
// [errors.As]. Errors of this type will also satisfy [ErrNotSingleton] and
// [ErrEval].
type NotSingletonError = errs.NotSingletonError

// KindError is an evaluation error returned when an expression is expecting a
// specific YAML node kind, but receives a node with a different kind.
//
// This error can be tested for with [errors.Is] and retrieved for with
// [errors.As]. Errors of this type will also satisfy [ErrBadKind] and [ErrEval].
type KindError = errs.KindError

// TagError is an evaluation error returned when an expression is expecting a
// specific YAML node tag, but receives a node with a different tag.
//
// This error can be tested for with [errors.Is] and retrieved for with
// [errors.As]. Errors of this type will also satisfy [ErrBadTag] and [ErrEval].
type TagError = errs.TagError

type FuncError struct {
	FuncName string
	Reason   error
}
