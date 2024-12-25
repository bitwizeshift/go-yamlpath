/*
Package yamlutil provides some simple reusable utilities for working with
the yaml library.
*/
package yamlutil

import (
	"fmt"
	"strconv"

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

// Number returns a yaml node representing an integer.
// The input is assumed to be a valid integer string.
func Number(s string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!int"}
}

// String returns a yaml node representing a string.
func String(s string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: s, Tag: "!!str"}
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

func ToString(node *yaml.Node) (string, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!str" {
		return "", fmt.Errorf("expected string node, but got %s", node.Tag)
	}
	return node.Value, nil
}

func ToInt(node *yaml.Node) (int, error) {
	if node.Kind != yaml.ScalarNode || node.Tag != "!!int" {
		return 0, fmt.Errorf("expected integer node, but got %s", node.Tag)
	}
	return strconv.Atoi(node.Value)
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

// EqualRange compares two ranges of yaml nodes for equality.
// Documents, source locations, and comments are ignored.
func EqualRange(got, want []*yaml.Node) bool {
	got, want = Normalize(got...), Normalize(want...)

	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if !Equal(got[i], want[i]) {
			return false
		}
	}
	return true
}

// Equal compares two yaml nodes for equality.
func Equal(got, want *yaml.Node) bool {
	got, want = Normalize(got)[0], Normalize(want)[0]

	if got.Kind != want.Kind {
		return false
	}
	if got.Tag != want.Tag {
		return false
	}
	if got.Value != want.Value {
		return false
	}
	if len(got.Content) != len(want.Content) {
		return false
	}
	for i := range got.Content {
		if !Equal(got.Content[i], want.Content[i]) {
			return false
		}
	}
	return true
}
