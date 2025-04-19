package funcs

import (
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
)

// Keys returns the keys of the current mapping node.
func Keys(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	var result []*yaml.Node
	current := ctx.Current()

	for _, node := range current {
		if node.Kind != yaml.MappingNode {
			continue
		}
		for i := 0; (i + 1) < len(node.Content); i += 2 {
			result = append(result, node.Content[i])
		}
	}
	return result, nil
}

// Children returns the immediate children of the input mapping node.
// If the input collection is empty, it returns an empty collection.
// If the input collection is not a singleton value, an error is raised to the
// calling environment.
// If the input is not a mapping node, it returns an empty collection.
func Children(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("children()", current)
	}
	node := current[0]
	if node.Kind != yaml.MappingNode {
		return nil, nil
	}

	result := make([]*yaml.Node, 0, len(node.Content)/2)
	for i := 0; (i + 1) < len(node.Content); i += 2 {
		result = append(result, node.Content[i+1])
	}
	return result, nil
}

// Descendants returns all descendants of the current node.
// It traverses the YAML tree and collects all nodes, including
// sequences, mappings, and scalar nodes, and adds it to the output collection.
func Descendants(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("descendants()", current)
	}
	var result []*yaml.Node
	descendants(&result, true, current...)
	return result, nil
}

func descendants(result *[]*yaml.Node, first bool, current ...*yaml.Node) {
	for _, node := range current {
		switch node.Kind {
		case yaml.DocumentNode:
			descendants(result, first, node.Content...)
		case yaml.SequenceNode:
			if !first {
				*result = append(*result, node)
			}
			descendants(result, false, node.Content...)
		case yaml.MappingNode:
			if !first {
				*result = append(*result, node)
			}
			for i := 0; (i + 1) < len(node.Content); i += 2 {
				descendants(result, false, node.Content[i+1])
			}
		case yaml.AliasNode:
			descendants(result, first, node.Content...)
		default:
			if !first {
				*result = append(*result, node)
			}
		}
	}
}
