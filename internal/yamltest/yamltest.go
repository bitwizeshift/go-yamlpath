/*
Package yamltest provides test-doubles for [yaml.Node] implementations.
*/
package yamltest

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

// IgnoreMetaFields returns a [cmp.Option] that ignores all of the meta-specific
// fields of a [yaml.Node], such as the comments, source-location, or style.
// This allows [cmp.Diff] and [cmp.Equal] to focus on the important fields that
// matter like the value, kind, tag, and content.
func IgnoreMetaFields() cmp.Option {
	return cmpopts.IgnoreFields(
		yaml.Node{},
		"LineComment",
		"HeadComment",
		"FootComment",
		"Line",
		"Column",
		"Style",
	)
}
