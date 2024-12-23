/*
Package yamlpath provides an implementation of the JSONPath specification for
YAML, leveraging the go-yaml library.
*/
package yamlpath

import (
	"context"
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

// YAMLPath represents a compiled YAMLPath expression.
type YAMLPath struct {
	path       string
	expression expr.Expression
}

// Compile compiles a YAMLPath expression.
func Compile(path string) (*YAMLPath, error) {
	expr, err := compile.NewTree(path)
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

// Eval evaluates the YAMLPath expression against the given YAML node.
func (yp *YAMLPath) Eval(node *yaml.Node) ([]*yaml.Node, error) {
	if yp == nil {
		return nil, nil
	}
	var input []*yaml.Node
	if node != nil {
		input = []*yaml.Node{node}
	}

	ctx := context.Background()
	return yp.expression.Eval(ctx, input)
}

// MustEval evaluates the YAMLPath expression against the given YAML node. It
// panics if an error occurs.
func (yp *YAMLPath) MustEval(node *yaml.Node) []*yaml.Node {
	nodes, err := yp.Eval(node)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Eval evaluates the YAMLPath expression against the given YAML node.
func Eval(path string, node *yaml.Node) ([]*yaml.Node, error) {
	yp, err := Compile(path)
	if err != nil {
		return nil, err
	}
	return yp.Eval(node)
}

// MustEval evaluates the YAMLPath expression against the given YAML node. It
// panics if an error occurs.
func MustEval(path string, node *yaml.Node) []*yaml.Node {
	nodes, err := Eval(path, node)
	if err != nil {
		panic(err)
	}
	return nodes
}
