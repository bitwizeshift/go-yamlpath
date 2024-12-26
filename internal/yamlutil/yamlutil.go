/*
Package yamlutil provides some simple reusable utilities for working with
the yaml library.
*/
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

	// Null is a yaml node representing a null value
	Null = &yaml.Node{Kind: yaml.ScalarNode, Value: "", Tag: "!!null"}
)

// IsTruthy returns true if any of the nodes are truthy.
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

func FromBool(b bool) *yaml.Node {
	if b {
		return True
	}
	return False
}

func ToBool(node *yaml.Node) (bool, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!bool" {
		return false, fmt.Errorf("expected boolean node, but got %s", node.Tag)
	}
	return node.Value == "true", nil
}

// String returns a yaml node representing a string.
func String(s string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!str"}
}

func FromString(s string) *yaml.Node {
	return String(s)
}

func ToString(node *yaml.Node) (string, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
		return "", fmt.Errorf("expected string node, but got %s", node.Tag)
	}
	return node.Value, nil
}

// Number returns a yaml node representing an integer.
// The input is assumed to be a valid integer string.
func Number(s string) *yaml.Node {
	if strings.Contains(s, ".") {
		return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!float"}
	}
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!int"}
}

func FromInt(i int) *yaml.Node {
	return Number(strconv.Itoa(i))
}

func ToInt(node *yaml.Node) (int, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!int" {
		return 0, fmt.Errorf("expected integer node, but got %s", node.Tag)
	}
	return strconv.Atoi(node.Value)
}

func FromFloat(f decimal.Decimal) *yaml.Node {
	return Number(f.String())
}

func ToFloat(node *yaml.Node) (decimal.Decimal, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!float" {
		return decimal.Zero, fmt.Errorf("expected float node, but got %s", node.Tag)
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
