package yamlutil

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// LessRange compares two ranges of yaml nodes for less than.
func LessRange(lhs, rhs []*yaml.Node) (bool, error) {
	lhs, rhs = Normalize(lhs...), Normalize(rhs...)

	if len(lhs) != len(rhs) {
		return len(lhs) < len(rhs), nil
	}
	for i := range lhs {
		l, err := Less(lhs[i], rhs[i])
		if err != nil {
			return false, err
		}
		if l {
			return true, nil
		}
	}
	return false, nil
}

// Less compares two yaml nodes for less than.
func Less(lhs, rhs *yaml.Node) (bool, error) {
	if lhs.Kind != rhs.Kind {
		return false, fmt.Errorf("cannot compare different kinds: %v and %v", lhs.Kind, rhs.Kind)
	}
	if lhs.Tag != rhs.Tag {
		return false, fmt.Errorf("cannot compare different tags: %s and %s", lhs.Tag, rhs.Tag)
	}
	switch lhs.Kind {
	case yaml.ScalarNode:
		return lessScalar(lhs, rhs)
	case yaml.SequenceNode:
		return lessSequence(lhs, rhs)
	}
	return false, fmt.Errorf("cannot compare kind: %v", lhs.Kind)
}

func lessScalar(lhs, rhs *yaml.Node) (bool, error) {
	switch lhs.Tag {
	case "!!int":
		l, err := ToInt(lhs)
		if err != nil {
			return false, err
		}
		r, err := ToInt(rhs)
		if err != nil {
			return false, err
		}
		return l < r, nil
	case "!!float":
		l, err := ToFloat(lhs)
		if err != nil {
			return false, err
		}
		r, err := ToFloat(rhs)
		if err != nil {
			return false, err
		}
		return l.LessThan(r), nil
	case "!!bool":
		l, err := ToBool(lhs)
		if err != nil {
			return false, err
		}
		r, err := ToBool(rhs)
		if err != nil {
			return false, err
		}
		return l == false && r == true, nil
	case "!!str":
		return lhs.Value < rhs.Value, nil
	}
	return false, fmt.Errorf("cannot compare scalar tag: %s", lhs.Tag)
}

func lessSequence(lhs, rhs *yaml.Node) (bool, error) {
	if len(lhs.Content) != len(rhs.Content) {
		return len(lhs.Content) < len(rhs.Content), nil
	}
	for i := range lhs.Content {
		if Equal(lhs, rhs) {
			continue
		}
		return Less(lhs.Content[i], rhs.Content[i])
	}
	return false, nil
}
