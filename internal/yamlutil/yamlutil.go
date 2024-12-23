/*
Package yamlutil provides some simple reusable utilities for working with
the yaml library.
*/
package yamlutil

import "gopkg.in/yaml.v3"

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
