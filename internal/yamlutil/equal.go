package yamlutil

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

// EqualRange compares two ranges of yaml nodes for equality.
// Documents, source locations, and comments are ignored.
func EqualRange(lhs, rhs []*yaml.Node) bool {
	lhs, rhs = Normalize(lhs...), Normalize(rhs...)

	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if !Equal(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

// Equal compares two yaml nodes for equality.
func Equal(lhs, rhs *yaml.Node) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	lhs, rhs = Normalize(lhs)[0], Normalize(rhs)[0]

	if lhs.Kind != rhs.Kind {
		return false
	}
	if (lhs.Tag == "!!int" || lhs.Tag == "!!float") && (rhs.Tag == "!!int" || rhs.Tag == "!!float") {
		left, err1 := decimal.NewFromString(lhs.Value)
		right, err2 := decimal.NewFromString(rhs.Value)
		if err1 != nil || err2 != nil {
			return lhs.Value == rhs.Value
		}
		return left.Equal(right) && left.NumDigits() == right.NumDigits()
	}
	if lhs.Tag != rhs.Tag || lhs.Value != rhs.Value {
		return false
	}
	if len(lhs.Content) != len(rhs.Content) {
		return false
	}
	for i := range lhs.Content {
		if !Equal(lhs.Content[i], rhs.Content[i]) {
			return false
		}
	}
	return true
}
