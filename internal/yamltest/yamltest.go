/*
Package yamltest provides test-doubles for [yaml.Node] implementations.
*/
package yamltest

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// MustParseNode unmarshals the given yaml string into a yaml.Node
func MustParseNode(data string) *yaml.Node {
	return yamlconv.FlattenDocuments(MustParseDocument(data))[0]
}

// MustParseDocument unmarshals the given yaml string into a yaml.Node without normalizing
// away the 'document' kind.
func MustParseDocument(data string) *yaml.Node {
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(data), &node); err != nil {
		panic(fmt.Sprintf("failed to unmarshal yaml: %v", err))
	}
	return &node
}
