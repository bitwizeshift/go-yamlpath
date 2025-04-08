package yamlpath

import (
	"fmt"

	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Collection represents a collection of YAML nodes that can be retrieved from
// a YAMLPath evaluation result.
type Collection []*yaml.Node

// Nodes converts this Collection into a slice of [yaml.Node] objects.
func (c Collection) Nodes() []*yaml.Node {
	return c
}

// IsTruthy evaluates whether this collection may be considered as having a
// truthful representation.
//
// Truthiness of a collection is determined by the following rules:
//
//   - An empty collection is considered falsy
//   - A collection with multiple nodes is considered truthy
//   - A collection with a single node that is not a boolean is considered truthy
//   - A collection with a single node that is a boolean is considered truthy if
//     the boolean value is true
func (c Collection) IsTruthy() bool {
	return yamlconv.IsTruthy(c...)
}

// Decoder returns a new Decoder that can be used to decode the collection
// into target objects.
func (c Collection) Decoder() *Decoder {
	return NewDecoder(c)
}

// IsEmpty returns true if the collection is empty.
func (c Collection) IsEmpty() bool {
	return len(c) == 0
}

// Singleton evaluates whether the collection contains only a single node,
// and returns that node if it does. If it does not, an [ErrNotSingleton]
// error is raised.
func (c Collection) Singleton() (*yaml.Node, error) {
	if len(c) != 1 {
		return nil, fmt.Errorf("%w; got %d", ErrNotSingleton, len(c))
	}
	return c[0], nil
}

// SingletonString evaluates whether the collection contains only a single
// node, and returns the string value of that node if it does. If it does not,
// an [ErrNotSingleton] error is raised.
func (c Collection) SingletonString() (v string, err error) {
	return decodeSingleton[string](c)
}

// SingletonBool evaluates whether the collection contains only a single
// node, and returns the boolean value of that node if it does. If it does not,
// an [ErrNotSingleton] error is raised.
func (c Collection) SingletonBool() (v bool, err error) {
	return decodeSingleton[bool](c)
}

// SingletonInt evaluates whether the collection contains only a single
// node, and returns the integer value of that node if it does. If it does not,
// an [ErrNotSingleton] error is raised.
func (c Collection) SingletonInt() (v int, err error) {
	return decodeSingleton[int](c)
}

// SingletonFloat64 evaluates whether the collection contains only a single
// node, and returns the float value of that node if it does. If it does not,
// an [ErrNotSingleton] error is raised.
func (c Collection) SingletonFloat64() (v float64, err error) {
	return decodeSingleton[float64](c)
}

// SingletonDecimal evaluates whether the collection contains only a single
// node, and returns the decimal value of that node if it does. If it does not,
// an [ErrNotSingleton] error is raised.
func (c Collection) SingletonDecimal() (v decimal.Decimal, err error) {
	return decodeSingleton[decimal.Decimal](c)
}

// Strings decodes all nodes in the collection into a slice of strings.
// If any of the nodes are unable to be marshaled into a string, an error
// is returned.
func (c Collection) Strings() ([]string, error) {
	return decodeAll[string](c)
}

// Bools decodes all nodes in the collection into a slice of booleans.
// If any of the nodes are unable to be marshaled into a boolean, an error
// is returned.
func (c Collection) Bools() ([]bool, error) {
	return decodeAll[bool](c)
}

// Ints decodes all nodes in the collection into a slice of integers.
// If any of the nodes are unable to be marshaled into an integer, an error
// is returned.
func (c Collection) Ints() ([]int, error) {
	return decodeAll[int](c)
}

// Float64s decodes all nodes in the collection into a slice of floats.
// If any of the nodes are unable to be marshaled into a float, an error
// is returned.
func (c Collection) Float64s() ([]float64, error) {
	return decodeAll[float64](c)
}

// Decimals decodes all nodes in the collection into a slice of decimals.
// If any of the nodes are unable to be marshaled into a decimal, an error
// is returned.
func (c Collection) Decimals() ([]decimal.Decimal, error) {
	return decodeAll[decimal.Decimal](c)
}

// Equal compares this collection to another collection for equality.
//
// This does a deep comparison of the YAML nodes, and introspects into whether
// the semantic values of the nodes are equal (e.g. YAML scalar floats may be
// in various notations, but represent the same value).
//
// This comparison follows the same behavior as the `==` operator.
func (c Collection) Equal(other Collection) bool {
	return yamlcmp.EqualRange(c, other)
}

func decodeSingleton[T any](nodes Collection) (T, error) {
	var value T
	node, err := nodes.Singleton()
	if err != nil {
		return value, err
	}
	err = node.Decode(&value)
	return value, err
}

func decodeAll[T any](nodes Collection) ([]T, error) {
	var values []T
	for _, node := range nodes {
		var value T
		if err := node.Decode(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}
