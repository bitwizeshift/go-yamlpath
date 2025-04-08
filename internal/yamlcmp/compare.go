package yamlcmp

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// LessRange compares two ranges of yaml nodes for less than.
func LessRange(lhs, rhs []*yaml.Node) (bool, error) {
	cmp, err := CompareRange(lhs, rhs)
	return cmp < 0, err
}

// CompareRange compares two ranges of yaml nodes for less than.
func CompareRange(lhs, rhs []*yaml.Node) (int, error) {
	lhs, rhs = yamlconv.FlattenDocuments(lhs...), yamlconv.FlattenDocuments(rhs...)

	if len(lhs) != len(rhs) {
		return len(lhs) - len(rhs), nil
	}
	for i := range lhs {
		cmp, err := Compare(lhs[i], rhs[i])
		if err != nil {
			return 0, err
		}
		if cmp == 0 {
			continue
		}
		return cmp, nil
	}
	return 0, nil
}

func isNumber(node *yaml.Node) bool {
	return node.Tag == "!!int" || node.Tag == "!!float"
}

// Compare compares two yaml nodes for less than.
func Compare(lhs, rhs *yaml.Node) (int, error) {
	if lhs == rhs {
		return 0, nil
	}
	if lhs == nil {
		return -1, nil
	}
	if rhs == nil {
		return 1, nil
	}

	if lhs.Kind != rhs.Kind {
		return 0, fmt.Errorf("cannot compare different kinds: %v and %v", lhs.Kind, rhs.Kind)
	}
	switch lhs.Kind {
	case yaml.ScalarNode:
		return compareScalar(lhs, rhs)
	case yaml.SequenceNode:
		return compareSequence(lhs, rhs)
	}
	return 0, fmt.Errorf("cannot compare kind: %v", lhs.Kind)
}

func compareScalar(lhs, rhs *yaml.Node) (int, error) {
	if isNumber(lhs) && isNumber(rhs) {
		l, err := decimal.NewFromString(lhs.Value)
		if err != nil {
			return 0, err
		}
		r, err := decimal.NewFromString(rhs.Value)
		if err != nil {
			return 0, err
		}
		return l.Compare(r), nil
	}
	if lhs.Tag != rhs.Tag {
		return 0, fmt.Errorf("cannot compare scalar tag: %s", lhs.Tag)
	}

	switch lhs.Tag {
	case "!!bool":
		l, err := yamlconv.ParseBool(lhs)
		if err != nil {
			return 0, err
		}
		r, err := yamlconv.ParseBool(rhs)
		if err != nil {
			return 0, err
		}
		if l == r {
			return 0, nil
		}
		if !l {
			return -1, nil
		}
		return 1, nil
	case "!!str":
		return strings.Compare(lhs.Value, rhs.Value), nil
	case "!!null":
		return 0, nil
	}

	// Should be unreachable in practice?
	return 0, fmt.Errorf("cannot compare scalar tags: %s", lhs.Tag)
}

func compareSequence(lhs, rhs *yaml.Node) (int, error) {
	if len(lhs.Content) != len(rhs.Content) {
		return len(lhs.Content) - len(rhs.Content), nil
	}
	for i := range lhs.Content {
		cmp, err := Compare(lhs.Content[i], rhs.Content[i])
		if err != nil {
			return 0, err
		}
		if cmp == 0 {
			continue
		}
		return cmp, nil
	}
	return 0, nil
}
