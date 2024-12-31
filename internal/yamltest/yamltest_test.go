package yamltest_test

import (
	"testing"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestMustParseNode_SuccessfulParse(t *testing.T) {
	node := yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)

	if node.Kind != yaml.MappingNode {
		t.Errorf("MustParseNode() = %v, want %v", node.Kind, yaml.MappingNode)
	}
}

func TestMustParseNode_UnsuccessfulParse(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseNode() did not panic")
		}
	}()

	yamltest.MustParseNode(`{"name": "Alice", "age": 30`)
}

func TestMustParseDocument_SuccessfulParse(t *testing.T) {
	node := yamltest.MustParseDocument(`{"name": "Alice", "age": 30}`)

	if node.Kind != yaml.DocumentNode {
		t.Errorf("MustParseDocument() = %v, want %v", node.Kind, yaml.DocumentNode)
	}
}

func TestMustParseDocument_UnsuccessfulParse(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseDocument() did not panic")
		}
	}()

	yamltest.MustParseDocument(`{"name": "Alice", "age": 30`)
}
