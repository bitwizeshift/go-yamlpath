package expr

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	// ErrEval is returned when an error occurs during evaluation of an
	// expression.
	ErrEval = errors.New("yamlpath eval")
)

func NewSingletonError(operator string, count int) error {
	return fmt.Errorf(
		"%w: %s requires singleton node, but received %d nodes",
		ErrEval,
		operator,
		count,
	)
}

func NewKindError(operator string, got, want yaml.Kind) error {
	return fmt.Errorf(
		"%w: %s may only operate on %s nodes, but received %s",
		ErrEval,
		operator,
		kindToString(got),
		kindToString(want),
	)
}

func NewTagError(operator string, gotTag, wantTag string) error {
	return fmt.Errorf(
		"%w: %s may only operate on '%s' tagged scalar values, but received tag '%s'",
		ErrEval,
		operator,
		wantTag,
		gotTag,
	)
}

func kindToString(kind yaml.Kind) string {
	switch kind {
	case yaml.DocumentNode:
		return "document"
	case yaml.MappingNode:
		return "mapping"
	case yaml.ScalarNode:
		return "scalar"
	case yaml.SequenceNode:
		return "sequence"
	}
	return "unknown"
}
