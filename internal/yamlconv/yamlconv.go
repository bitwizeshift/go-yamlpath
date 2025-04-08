/*
Package yamlconv provides functions to convert between YAML and native Go
formats.
*/
package yamlconv

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
)

// Primitive is a type constraint for primitive types that can be converted to
// YAML nodes.
type Primitive interface {
	constraints.Integer | constraints.Float | ~string | ~bool
}

const (
	yamlStringTag = "!!str"
	yamlIntTag    = "!!int"
	yamlFloatTag  = "!!float"
	yamlBoolTag   = "!!bool"
	yamlSeqTag    = "!!seq"
)

// Int converts an integer to a YAML node.
func Int[T constraints.Integer](n T) *yaml.Node {
	return RawInt(strconv.FormatInt(int64(n), 10))
}

// RawInt creates a YAML node with a raw integer value.
func RawInt(s string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   yamlIntTag,
		Value: s,
	}
}

// Float converts a float64 to a YAML node.
func Float[T constraints.Float](f T) *yaml.Node {
	return RawFloat(strconv.FormatFloat(float64(f), 'f', -1, 64))
}

// RawFloat creates a YAML node with a raw float value.
func RawFloat(s string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   yamlFloatTag,
		Value: s,
	}
}

// Number converts a number (int or float) to a YAML node.
func Number[T constraints.Integer | constraints.Float](n T) *yaml.Node {
	rv := reflect.ValueOf(n)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Int(int64(rv.Uint()))
	case reflect.Float32, reflect.Float64:
		return Float(rv.Float())
	default:
		panic("unreachable")
	}
}

// RawNumber creates a YAML node with a raw number string value.
func RawNumber(s string) *yaml.Node {
	tag := yamlIntTag
	if strings.Contains(s, ".") || strings.Contains(s, "e") || strings.Contains(s, "E") {
		tag = yamlFloatTag
	}
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   tag,
		Value: s,
	}
}

// String converts a string to a YAML node.
func String(s string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   yamlStringTag,
		Value: s,
	}
}

// Bool converts a boolean to a YAML node.
func Bool(b bool) *yaml.Node {
	return RawBool(strconv.FormatBool(b))
}

// RawBool creates a YAML node with a raw boolean value.
func RawBool(s string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   yamlBoolTag,
		Value: s,
	}
}

// Null creates a YAML node representing a null value.
func Null() *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!null",
		Value: "",
	}
}

// Node converts a primitive type to a YAML node.
func Node[T Primitive](v T) *yaml.Node {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Int(int64(rv.Uint()))
	case reflect.Float32, reflect.Float64:
		return Float(rv.Float())
	case reflect.String:
		return String(rv.String())
	case reflect.Bool:
		return Bool(rv.Bool())
	default:
		panic("unreachable")
	}
}

// Ints converts a list of integers to a list of YAML nodes.
func Ints[T constraints.Integer](n ...T) []*yaml.Node {
	return toNodes(Int, n...)
}

// Floats converts a list of floats to a list of YAML nodes.
func Floats[T constraints.Float](f ...T) []*yaml.Node {
	return toNodes(Float, f...)
}

// Numbers converts a list of numbers (int or float) to a list of YAML nodes.
func Numbers[T constraints.Integer | constraints.Float](n ...T) []*yaml.Node {
	return toNodes(Number, n...)
}

// Strings converts a list of strings to a list of YAML nodes.
func Strings(s ...string) []*yaml.Node {
	return toNodes(String, s...)
}

// Bools converts a list of booleans to a list of YAML nodes.
func Bools(b ...bool) []*yaml.Node {
	return toNodes(Bool, b...)
}

// Nodes converts a list of primitive types to a list of YAML nodes.
func Nodes[T Primitive](v ...T) []*yaml.Node {
	return toNodes(Node, v...)
}

func toNodes[T any](fn func(T) *yaml.Node, v ...T) []*yaml.Node {
	nodes := make([]*yaml.Node, len(v))
	for i, val := range v {
		nodes[i] = fn(val)
	}
	return nodes
}

// Sequence creates a YAML sequence node with the given child nodes.
func Sequence(nodes ...*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     yamlSeqTag,
		Content: nodes,
	}
}

// Document creates a YAML document node with the given child nodes.
func Document(nodes ...*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.DocumentNode,
		Content: nodes,
	}
}

// ParseString parses a YAML node as a string.
func ParseString(node *yaml.Node) (string, error) {
	if node.Kind != yaml.ScalarNode {
		return "", errs.NewKindError("", node, yaml.ScalarNode)
	}
	if tag := yamlStringTag; node.Tag != tag {
		return "", errs.NewTagError("", node, tag)
	}
	return node.Value, nil
}

// ParseInt parses a YAML node as an integer.
func ParseInt(node *yaml.Node) (int64, error) {
	if node.Kind != yaml.ScalarNode {
		return 0, errs.NewKindError("", node, yaml.ScalarNode)
	}
	if tag := yamlIntTag; node.Tag != tag {
		return 0, errs.NewTagError("", node, tag)
	}
	return strconv.ParseInt(node.Value, 10, 64)
}

// ParseBool parses a YAML node as a boolean.
func ParseBool(node *yaml.Node) (bool, error) {
	if node.Kind != yaml.ScalarNode {
		return false, errs.NewKindError("", node, yaml.ScalarNode)
	}
	if tag := yamlBoolTag; node.Tag != tag {
		return false, errs.NewTagError("", node, tag)
	}
	return strconv.ParseBool(node.Value)
}

// ParseFloat parses a YAML node as a float.
func ParseFloat(node *yaml.Node) (float64, error) {
	if node.Kind != yaml.ScalarNode {
		return 0, errs.NewKindError("", node, yaml.ScalarNode)
	}
	if tag := yamlFloatTag; node.Tag != tag {
		return 0, errs.NewTagError("", node, tag)
	}
	return strconv.ParseFloat(node.Value, 64)
}

// ParseDecimal parses a yaml node to a decimal.
func ParseDecimal(node *yaml.Node) (decimal.Decimal, error) {
	if node.Kind != yaml.ScalarNode {
		return decimal.Zero, errs.NewKindError("", node, yaml.ScalarNode)
	}
	if node.Tag != yamlIntTag && node.Tag != yamlFloatTag {
		return decimal.Zero, errs.NewTagError("", node, yamlIntTag, yamlFloatTag)
	}
	return decimal.NewFromString(node.Value)
}

// FlattenDocuments normalizes a list of yaml nodes by flattening any document
// nodes into their content nodes.
func FlattenDocuments(node ...*yaml.Node) []*yaml.Node {
	var result []*yaml.Node
	for _, n := range node {
		if n.Kind == yaml.DocumentNode {
			result = append(result, FlattenDocuments(n.Content...)...)
		} else {
			result = append(result, n)
		}
	}
	return result
}

// IsTruthy returns true if any of the nodes are truthy.
// Truthiness is determined by the following algorithm:
//
//  1. If no nodes are provided, the result is false
//  2. If more than one node is provided, the result is true
//  3. If a single node is provided, then:
//     a. If the node is not a boolean, the result is true
//     b. If the node is a boolean, the result is the boolean value
//
// This makes it possible to test for conditional existence of boolean nodes
// without having to check for nil nodes.
func IsTruthy(nodes ...*yaml.Node) bool {
	if len(nodes) == 0 {
		return false
	}
	if len(nodes) > 1 {
		return true
	}
	b, err := ParseBool(nodes[0])
	if err != nil {
		return true
	}
	return b
}
