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
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
	expression expr.Expr
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

// EvalBool evaluates the YAMLPath expression against the given YAML node and
// returns the result as a boolean. If the path returns more than one node, or
// if the node is not a boolean, an error is returned.
func (yp *YAMLPath) EvalBool(node *yaml.Node) (bool, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return false, err
	}
	if len(nodes) == 0 {
		return false, nil
	}
	if len(nodes) > 1 {
		return false, fmt.Errorf("expected single node, but got %d", len(nodes))
	}
	b, err := yamlutil.ToBool(node)
	if err != nil {
		return false, err
	}
	return b, nil
}

// EvalBools evaluates the YAMLPath expression against the given YAML node and
// returns the result as a slice of booleans. If any node is not a boolean, an
// error is returned.
func (yp *YAMLPath) EvalBools(node *yaml.Node) ([]bool, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return nil, err
	}
	var result []bool
	for _, node := range nodes {
		b, err := yamlutil.ToBool(node)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

// EvalString evaluates the YAMLPath expression against the given YAML node and
// returns the result as a string. If the path returns more than one node, an
// error is returned.
func (yp *YAMLPath) EvalString(node *yaml.Node) (string, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return "", err
	}
	if len(nodes) == 0 {
		return "", nil
	}
	if len(nodes) > 1 {
		return "", fmt.Errorf("expected single node, but got %d", len(nodes))
	}
	return yamlutil.ToString(nodes[0])
}

// EvalStrings evaluates the YAMLPath expression against the given YAML node and
// returns the result as a slice of strings.
func (yp *YAMLPath) EvalStrings(node *yaml.Node) ([]string, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, node := range nodes {
		s, err := yamlutil.ToString(node)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

// EvalInt evaluates the YAMLPath expression against the given YAML node and
// returns the result as an integer. If the path returns more than one node, an
// error is returned.
func (yp *YAMLPath) EvalInt(node *yaml.Node) (int, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return 0, err
	}
	if len(nodes) == 0 {
		return 0, nil
	}
	if len(nodes) > 1 {
		return 0, fmt.Errorf("expected single node, but got %d", len(nodes))
	}
	return yamlutil.ToInt(nodes[0])
}

// EvalInts evaluates the YAMLPath expression against the given YAML node and
// returns the result as a slice of integers.
func (yp *YAMLPath) EvalInts(node *yaml.Node) ([]int, error) {
	nodes, err := yp.Eval(node)
	if err != nil {
		return nil, err
	}
	var result []int
	for _, node := range nodes {
		i, err := yamlutil.ToInt(node)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
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
