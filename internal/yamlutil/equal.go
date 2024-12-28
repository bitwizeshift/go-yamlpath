package yamlutil

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
)

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
	if got.Tag == "!!int" || got.Tag == "!!float" {
		lhs, err1 := decimal.NewFromString(got.Value)
		rhs, err2 := decimal.NewFromString(want.Value)
		if err1 != nil || err2 != nil {
			return got.Value == want.Value
		}
		return lhs.Equal(rhs) && lhs.NumDigits() == rhs.NumDigits()
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
