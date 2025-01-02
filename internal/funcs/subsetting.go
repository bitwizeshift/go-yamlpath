package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// Single evaluates the current node and returns an error if the current node
// is not a singleton.
func Single(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	if len(current) == 0 {
		return nil, nil
	}
	if len(current) == 1 {
		return current, nil
	}
	return nil, errs.NewSingletonError("single()", current)
}

// First returns the first n elements of the current collection.
// If no n is specified, it returns the first element.
func First(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	n := 1
	if len(params) == 1 {
		v, err := paramToInt(ctx, "first()", params[0])
		if err != nil {
			return nil, err
		}

		n = v
	}
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	n = min(n, len(current))
	return current[:n], nil
}

// Last returns the current collection with all but the last n elements removed.
// If no n is specified, it returns the last element.
func Last(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	n := 1
	if len(params) == 1 {
		v, err := paramToInt(ctx, "last()", params[0])
		if err != nil {
			return nil, err
		}
		n = v
	}

	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	n = max(len(current)-n, 0)
	return current[n:], nil
}

// Skip returns the current collection with the first n elements removed.
// If the current collection is empty, it returns an empty collection.
// If n is negative, it skips the last n elements.
func Skip(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	n, err := paramToInt(ctx, "skip()", params[0])
	if err != nil {
		return nil, err
	}

	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	if n < 0 {
		n = -n
		n = max(len(current)-n, 0)
		return current[:n], nil
	}
	n = min(n, len(current))

	return current[n:], nil
}
