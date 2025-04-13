package yamlhash_test

import (
	"testing"

	"rodusek.dev/pkg/yamlpath/internal/yamlhash"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestHash(t *testing.T) {
	node := yamltest.MustParseDocument(`{"a": 1, "b": [1, 2, 3]}`)

	hash := yamlhash.Hash(node.Content[0])

	if got, want := hash, "fZmDQCLyWU2kb+noo3J6wiMTI7Y="; got != want {
		t.Errorf("Hash() = %v, want %v", got, want)
	}
}
