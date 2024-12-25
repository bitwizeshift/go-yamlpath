package expr

import "gopkg.in/yaml.v3"

type Union []any

func (u *Union) Index(node *yaml.Node) []*yaml.Node {
	if node.Kind == yaml.SequenceNode {
		return u.indexInt(node)
	}
	if node.Kind == yaml.MappingNode {
		return u.indexMap(node)
	}

	return nil
}

func (u Union) indexInt(node *yaml.Node) []*yaml.Node {
	var result []*yaml.Node
	for _, index := range u {
		if i, ok := index.(int64); ok {
			if i < 0 || i >= int64(len(node.Content)) {
				continue
			}
			result = append(result, node.Content[i])
		}
	}
	return result
}

func (u Union) indexMap(node *yaml.Node) []*yaml.Node {
	var result []*yaml.Node
	for _, index := range u {
		if key, ok := index.(string); ok {
			for i := 0; (i + 1) < len(node.Content); i += 2 {
				if node.Content[i].Value == key {
					result = append(result, node.Content[i+1])
				}
			}
		}
	}
	return result
}
