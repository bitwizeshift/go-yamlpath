package funcs

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
				Key:   key.Value,
				Value: value,
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
		})
	case yaml.ScalarNode:
		return toNode(&scalarNode{
			Kind:  "scalar",
			Tag:   node.Tag,
			Value: node,
		})
	}
	return nil, fmt.Errorf("unknown node kind: %v", node.Kind)
}

type scalarNode struct {
	Kind  string     `yaml:"kind"`
	Tag   string     `yaml:"tag"`
	Value *yaml.Node `yaml:"value"`
}

type sequenceNode struct {
	Kind    string       `yaml:"kind"`
	Tag     string       `yaml:"tag"`
	Entries []*yaml.Node `yaml:"entries"`
}

type mappingNode struct {
	Kind    string       `yaml:"kind"`
	Tag     string       `yaml:"tag"`
	Entries []*yaml.Node `yaml:"entries"`
}

type mappingNodeEntry struct {
	Key   string     `yaml:"key"`
	Value *yaml.Node `yaml:"value"`
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
	return yamlutil.Normalize(&node)[0], nil
}
