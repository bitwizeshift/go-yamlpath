package yamlutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

var (
	// True is a yaml node representing a boolean true value
	True = Boolean("true")

	// False is a yaml node representing a boolean false value
	False = Boolean("false")
)

// Null returns a yaml node representing a null value.
func Null() *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: "", Tag: "!!null"}
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
	b, err := ToBool(nodes[0])
	if err != nil {
		return true
	}
	return b
}

// Boolean returns a yaml node representing a boolean.
// The input is assumed to be a valid boolean string.
func Boolean(s string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!bool"}
}

// String returns a yaml node representing a string.
func String(s string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!str"}
}

// Number returns a yaml node representing an integer.
// The input is assumed to be a valid integer string.
func Number(s string) *yaml.Node {
	if strings.Contains(s, ".") || strings.Contains(s, "e") || strings.Contains(s, "E") {
		return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!float"}
	}
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!int"}
}

// ToBool converts a yaml node to a boolean.
func ToBool(node *yaml.Node) (bool, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!bool" {
		return false, fmt.Errorf("expected boolean node, but got %s", node.Tag)
	}
	return strconv.ParseBool(node.Value)
}

// ToString converts a yaml node to a string.
func ToString(node *yaml.Node) (string, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
		return "", fmt.Errorf("expected string node, but got %s", node.Tag)
	}
	return node.Value, nil
}

// ToInt converts a yaml node to an integer.
func ToInt(node *yaml.Node) (int, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!int" {
		return 0, fmt.Errorf("expected integer node, but got %s", node.Tag)
	}
	return strconv.Atoi(node.Value)
}

// ToDecimal converts a yaml node to a decimal.
func ToDecimal(node *yaml.Node) (decimal.Decimal, error) {
	if node.Kind != yaml.ScalarNode || !isNumber(node) {
		return decimal.Zero, fmt.Errorf("expected numeric node, but got %v", node.Tag)
	}
	return decimal.NewFromString(node.Value)
}

// Normalize normalizes a list of yaml nodes by flattening any document nodes
// into their content nodes.
func Normalize(node ...*yaml.Node) []*yaml.Node {
	var result []*yaml.Node
	for _, n := range node {
		if n.Kind == yaml.DocumentNode {
			result = append(result, n.Content...)
		} else {
			result = append(result, n)
		}
	}
	return result
}
