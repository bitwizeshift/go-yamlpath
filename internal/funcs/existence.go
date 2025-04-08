package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Empty returns true if the current node is empty, false otherwise.
func Empty(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// Exists returns true if the current node exists, false otherwise.
func Exists(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	if len(params) == 0 {
		current := ctx.Current()
		if len(current) > 0 {
			return []*yaml.Node{yamlconv.Bool(true)}, nil
		}
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}
	next, err := Where(ctx, params...)
	if err != nil {
		return nil, err
	}

	if len(next) > 0 {
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// Count returns the number of elements in the current collection. If the
// current collection is empty, it returns 0.
func Count(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	return []*yaml.Node{yamlconv.Number(len(current))}, nil
}

// Distinct returns a collection with all duplicate elements removed.
func Distinct(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	// seen is a map of hash -> nodes with that hash.
	// It is extremely unlikely to find a sha1 hash collision in practice, but
	// it's still accounted for by storing a slice of nodes with the same hash.
	seen := make(map[string][]*yaml.Node, len(current))
	var result []*yaml.Node

	for _, node := range current {
		key := yamlutil.Hash(node)
		others, ok := seen[key]
		if !ok {
			seen[key] = append(seen[key], node)
			result = append(result, node)
			continue
		}

		if !contains(others, node) {
			seen[key] = append(others, node)
			result = append(result, node)
		}
	}
	return result, nil
}

func contains(nodes []*yaml.Node, node *yaml.Node) bool {
	for _, n := range nodes {
		if yamlcmp.Equal(n, node) {
			return true
		}
	}
	return false
}

// IsDistinct returns true if the current collection contains only distinct
// elements, false otherwise.
func IsDistinct(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	// Note: 'Distinct' does not return an error, so we can ignore it here.
	distinct, _ := Distinct(ctx)

	if len(current) == len(distinct) {
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// All returns true if all nodes satisfy the 'criteria' parameter expression.
// This applies the singleton-boolean rule to each node to determine if the
// node is truthy.
// If the current collection is empty, it returns true.
func All(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	for _, node := range current {
		args, err := params[0].GetArg(ctx.NewContext([]*yaml.Node{node}))
		if err != nil {
			return nil, err
		}
		if !yamlconv.IsTruthy(args...) {
			return []*yaml.Node{yamlconv.Bool(false)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(true)}, nil
}

// Any returns true if any node satisfies the 'criteria' parameter expression.
// This applies the singleton-boolean rule to each node to determine if the
// node is truthy.
// If the current collection is empty, it returns false.
func Any(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	for _, node := range current {
		args, err := params[0].GetArg(ctx.NewContext([]*yaml.Node{node}))
		if err != nil {
			return nil, err
		}
		if yamlconv.IsTruthy(args...) {
			return []*yaml.Node{yamlconv.Bool(true)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// AllTrue evaluates the input collection and returns true if all elements are
// boolean nodes with the value of true. It returns false if any element is not
// a boolean node, or if any element is a boolean node with the value of false.
func AllTrue(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	for _, node := range current {
		b, err := yamlconv.ParseBool(node)
		if err != nil {
			return []*yaml.Node{yamlconv.Bool(false)}, nil
		}
		if !b {
			return []*yaml.Node{yamlconv.Bool(false)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(true)}, nil
}

// AnyTrue evaluates the input collection and returns true if any element is a
// boolean node with the value of true. It returns false if all elements are not
// boolean nodes, or if all elements are boolean nodes with the value of false.
func AnyTrue(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	for _, node := range current {
		b, err := yamlconv.ParseBool(node)
		if err != nil {
			continue
		}
		if b {
			return []*yaml.Node{yamlconv.Bool(true)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// AllFalse evaluates the input collection and returns true if all elements are
// boolean nodes with the value of false. It returns false if any element is not
// a boolean node, or if any element is a boolean node with the value of true.
func AllFalse(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	for _, node := range current {
		b, err := yamlconv.ParseBool(node)
		if err != nil {
			return []*yaml.Node{yamlconv.Bool(false)}, nil
		}
		if b {
			return []*yaml.Node{yamlconv.Bool(false)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(true)}, nil
}

// AnyFalse evaluates the input collection and returns true if any element is a
// boolean node with the value of false. It returns false if all elements are not
// boolean nodes, or if all elements are boolean nodes with the value of true.
func AnyFalse(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	for _, node := range current {
		b, err := yamlconv.ParseBool(node)
		if err != nil {
			continue
		}
		if !b {
			return []*yaml.Node{yamlconv.Bool(true)}, nil
		}
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}
