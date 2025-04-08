/*
Package yamlpathtest provides test-double utilities to aid in testing code that
leverages the yamlpath library.

This makes it possible to control what concrete [yamlpath.YAMLPath] instances
return without requiring specific yaml constructs to do so.
*/
package yamlpathtest

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Collection returns a [yamlpath.YAMLPath] object that returns the given nodes.
func Collection(nodes ...*yaml.Node) *yamlpath.YAMLPath {
	nodes = yamlconv.FlattenDocuments(nodes...)
	fn := func(yamlpath.Collection) yamlpath.Collection {
		return nodes
	}
	return compile(fn)
}

// String returns a [yamlpath.YAMLPath] object that returns the given string.
func String(s string) *yamlpath.YAMLPath {
	node := yamlconv.String(s)
	fn := func(yamlpath.Collection) *yaml.Node {
		return node
	}
	return compile(fn)
}

// StringSequence returns a [yamlpath.YAMLPath] object that returns a sequence
// containing the given strings.
func StringSequence(strings ...string) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) []string {
		return strings
	}
	return compile(fn)
}

// Bool returns a [yamlpath.YAMLPath] object that returns the given boolean.
func Bool(b bool) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) bool {
		return b
	}
	return compile(fn)
}

// BoolSequence returns a [yamlpath.YAMLPath] object that returns a sequence
// containing the given booleans.
func BoolSequence(bools ...bool) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) []bool {
		return bools
	}
	return compile(fn)
}

// Float returns a [yamlpath.YAMLPath] object that returns the given float.
func Float(float float64) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) float64 {
		return float
	}
	return compile(fn)
}

// FloatSequence returns a [yamlpath.YAMLPath] object that returns a sequence
// containing the given floats.
func FloatSequence(floats ...float64) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) []float64 {
		return floats
	}
	return compile(fn)
}

// Int returns a [yamlpath.YAMLPath] object that returns the given integer.
func Int(i int) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) int {
		return i
	}
	return compile(fn)
}

// IntSequence returns a [yamlpath.YAMLPath] object that returns a sequence
// containing the given integers.
func IntSequence(ints ...int) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) []int {
		return ints
	}
	return compile(fn)
}

// Object returns a [yamlpath.YAMLPath] object that returns the given value.
func Object(v any) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) any {
		return v
	}
	return compile(fn)
}

// ObjectSequence returns a [yamlpath.YAMLPath] object that returns a sequence
// containing the given values.
func ObjectSequence(vs ...any) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) []any {
		return vs
	}
	return compile(fn)
}

// Null returns a [yamlpath.YAMLPath] object that returns a null value.
func Null() *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) *yaml.Node {
		return yamlconv.Null()
	}
	return compile(fn)
}

// Error returns a [yamlpath.YAMLPath] object that always returns the given
// error.
func Error(err error) *yamlpath.YAMLPath {
	fn := func(yamlpath.Collection) (yamlpath.Collection, error) {
		return nil, err
	}
	return compile(fn)
}

var (
	// True is a YAMLPath that always returns a boolean truthy value.
	True = Bool(true)

	// False is a YAMLPath that always returns a boolean falsey value.
	False = Bool(false)
)

func compile(fn any) *yamlpath.YAMLPath {
	yp := yamlpath.MustCompile("$.return()",
		yamlpath.WithFunction("return", fn),
	)
	return yp
}
