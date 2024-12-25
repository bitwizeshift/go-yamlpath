package expr

import (
	"math"
	"strconv"
	"strings"
)

// Slice represents a slice expression in the form of `start[:end[:step]]` from
// the JSONPath specification.
type Slice struct {
	Start, End, Step int
}

// ParseSlice parses a slice expression from a string in the form of
// `start[:end[:step]]`.
func ParseSlice(s string) (*Slice, error) {
	var result Slice
	parts := strings.Split(s, ":")
	ints, err := intsFromStrings(parts)
	if err != nil {
		return nil, err
	}
	if len(ints) > 0 && ints[0] != nil {
		result.Start = *ints[0]
	} else {
		result.Start = 0
	}

	if len(ints) > 1 && ints[1] != nil {
		result.End = *ints[1]
	} else {
		result.End = math.MaxInt
	}
	if len(ints) > 2 && ints[2] != nil {
		result.Step = *ints[2]
	} else {
		result.Step = 1
	}
	return &result, nil
}

func intsFromStrings(s []string) ([]*int, error) {
	var result []*int
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v == "" {
			result = append(result, nil)
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, &i)
	}
	return result, nil
}
