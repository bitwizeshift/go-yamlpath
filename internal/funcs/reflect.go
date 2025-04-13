package funcs

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Reflect returns a reflected YAML representation of the current nodes.
func Reflect(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	var result []*yaml.Node
	for _, node := range current {
		n, err := reflectNode(node)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}

func reflectNode(node *yaml.Node) (*yaml.Node, error) {
	switch node.Kind {
	case yaml.DocumentNode:
		return reflectNode(node.Content[0])
	case yaml.SequenceNode:
		var nodes []*yaml.Node
		for _, child := range node.Content {
			n, err := reflectNode(child)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, n)
		}
		return toNode(&sequenceNode{
			Kind:    "sequence",
			Tag:     node.Tag,
			Entries: nodes,
			Source:  toSource(node),
		})
	case yaml.MappingNode:
		var nodes []*yaml.Node
		for i := 0; (i + 1) < len(node.Content); i += 2 {
			key := node.Content[i]
			value, err := reflectNode(node.Content[i+1])
			if err != nil {
				return nil, err
			}
			entry := &mappingNodeEntry{
				Key:    key.Value,
				Value:  value,
				Source: toSource(key),
			}
			node, err := toNode(entry)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, node)
		}

		return toNode(&mappingNode{
			Kind:    "mapping",
			Tag:     node.Tag,
			Entries: nodes,
			Source:  toSource(node),
		})
	case yaml.ScalarNode:
		return toNode(&scalarNode{
			Kind:   "scalar",
			Tag:    node.Tag,
			Value:  node,
			Source: toSource(node),
		})
	}
	return nil, fmt.Errorf("unknown node kind: %v", node.Kind)
}

func toSource(node *yaml.Node) *source {
	if node == nil || node.Line == 0 {
		return nil
	}
	return &source{
		Line:   node.Line,
		Column: node.Column,
	}
}

type source struct {
	Line   int `yaml:"line"`
	Column int `yaml:"column"`
}

type scalarNode struct {
	Kind   string     `yaml:"kind"`
	Tag    string     `yaml:"tag"`
	Value  *yaml.Node `yaml:"value"`
	Source *source    `yaml:"source,omitempty"`
}

type sequenceNode struct {
	Kind    string       `yaml:"kind"`
	Tag     string       `yaml:"tag"`
	Entries []*yaml.Node `yaml:"entries"`
	Source  *source      `yaml:"source,omitempty"`
}

type mappingNode struct {
	Kind    string       `yaml:"kind"`
	Tag     string       `yaml:"tag"`
	Entries []*yaml.Node `yaml:"entries"`
	Source  *source      `yaml:"source,omitempty"`
}

type mappingNodeEntry struct {
	Key    string     `yaml:"key"`
	Value  *yaml.Node `yaml:"value"`
	Source *source    `yaml:"source,omitempty"`
}

func toNode[T any](v *T) (*yaml.Node, error) {
	data, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	var node yaml.Node
	if err := yaml.Unmarshal(data, &node); err != nil {
		return nil, err
	}
	return yamlconv.FlattenDocuments(&node)[0], nil
}

// IsString checks if the current node is a string.
func IsString(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.ScalarNode, "!!str")
}

// IsInteger checks if the current node is an integer.
func IsInteger(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.ScalarNode, "!!int")
}

// IsFloat checks if the current node is a float.
func IsFloat(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.ScalarNode, "!!float")
}

// IsBoolean checks if the current node is a boolean.
func IsBoolean(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.ScalarNode, "!!bool")
}

// IsNull checks if the current node is null.
func IsNull(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.ScalarNode, "!!null")
}

// IsScalar checks if the current node is a scalar.
func IsScalar(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) == 1 {
		node := current[0]
		if node.Kind == yaml.ScalarNode {
			return []*yaml.Node{yamlconv.Bool(true)}, nil
		}
		return []*yaml.Node{yamlconv.Bool(false)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// IsSequence checks if the current node is a sequence.
func IsSequence(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.SequenceNode, "!!seq")
}

// IsMapping checks if the current node is a mapping.
func IsMapping(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	return hasTag(ctx, yaml.MappingNode, "!!map")
}

func hasTag(ctx invocation.Context, kind yaml.Kind, tag string) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) == 1 {
		node := current[0]
		return []*yaml.Node{
			yamlconv.Bool(node.Kind == kind && node.Tag == tag),
		}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}
