/*
Package yamlpath provides an implementation of the JSONPath specification for
YAML, leveraging the go-yaml library.
*/
package yamlpath

import (
	"encoding"
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/compile"
	"rodusek.dev/pkg/yamlpath/internal/expr"
)

var (
	// ErrSyntax is returned when an error occurs during parsing.
	ErrSyntax = compile.ErrSyntax

	// ErrCompile is returned when an error occurs during compilation.
	ErrCompile = compile.ErrCompile
)

// CompileError represents an error that occurs during compilation.
type CompileError = compile.CompileError

// YAMLPath is the representation of a compiled YAMLPath expression.
//
// A YAMLPath is safe for concurrent use by multiple goroutines,
// except for if custom functions have been provided that are themselves unsuafe
// for concurrent use.
type YAMLPath struct {
	path       string
	expression expr.Expr
}

// Compile parses a YAMLPath expressions and returns, if successful, a
// [YAMLPath] object that can be used to match against [yaml.Node]s.
//
// On error, this will return a [CompileError] that can be used to determine
// the cause of the error, and the location where the error occurs.
func Compile(path string) (*YAMLPath, error) {
	cfg := &compile.Config{
		Table: tableV1,
	}
	expr, err := compile.NewTree(path, cfg)
	if err != nil {
		return nil, err
	}
	result := &YAMLPath{
		path:       path,
		expression: expr,
	}
	return result, nil
}

// MustCompile compiles a YAMLPath expression. It panics if an error occurs.
//
// This is a convenience wrapper around [Compile].
func MustCompile(str string) *YAMLPath {
	yp, err := Compile(str)
	if err != nil {
		panic(err)
	}
	return yp
}

// String returns the string representation of the YAMLPath.
func (yp *YAMLPath) String() string {
	if yp == nil {
		return ""
	}
	return yp.path
}

var _ fmt.Stringer = (*YAMLPath)(nil)

// Match evaluates the YAMLPath expression against the given YAML node,
// returning all matching subnodes found. If an evaluation error occurs
// during the matching process, it is returned and the collection will be nil.
func (yp *YAMLPath) Match(node *yaml.Node) (Collection, error) {
	if yp == nil || yp.expression == nil {
		return nil, nil
	}
	var input []*yaml.Node
	if node != nil {
		input = []*yaml.Node{node}
	}

	ctx := expr.NewContext(input)
	return yp.expression.Eval(ctx)
}

// MustMatch evaluates the YAMLPath expression against the given YAML node,
// returning all matches. It panics if an error occurs.
func (yp *YAMLPath) MustMatch(node *yaml.Node) Collection {
	nodes, err := yp.Match(node)
	if err != nil {
		panic(err)
	}
	return nodes
}

// UnmarshalText unmarshals the YAMLPath expression from text.
//
// This enables YAMLPath to be used as a field in structs that may be unmarshaled
// in configs like YAML or JSON.
func (yp *YAMLPath) UnmarshalText(text []byte) error {
	ypath, err := Compile(string(text))
	if err != nil {
		return err
	}
	yp.path = ypath.path
	yp.expression = ypath.expression
	return nil
}

var _ encoding.TextUnmarshaler = (*YAMLPath)(nil)

// MarshalText marshals the YAMLPath expression to text.
func (yp *YAMLPath) MarshalText() ([]byte, error) {
	if yp == nil {
		return nil, nil
	}
	return []byte(yp.path), nil
}

// Equal returns true if the two YAMLPath expressions are equal.
//
// Equality is determined by evaluating whether the two paths are equal. This
// makes nil YAMLPath equivalent to zero value YAMLPath objects.
func (yp *YAMLPath) Equal(other *YAMLPath) bool {
	if yp == nil && other == nil {
		return true
	}
	if yp == nil && other.path == "" || yp.path == "" && other == nil {
		return true
	}
	return yp.path == other.path
}

var _ encoding.TextMarshaler = (*YAMLPath)(nil)

// Match evaluates the YAMLPath expression against the given YAML node.
func Match(path string, node *yaml.Node) (Collection, error) {
	yp, err := Compile(path)
	if err != nil {
		return nil, err
	}
	return yp.Match(node)
}

// MustMatch evaluates the YAMLPath expression against the given YAML node. It
// panics if an error occurs.
func MustMatch(path string, node *yaml.Node) Collection {
	nodes, err := Match(path, node)
	if err != nil {
		panic(err)
	}
	return nodes
}
