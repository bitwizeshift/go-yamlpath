package yamlutil_test

import (
	"testing"

	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestHash(t *testing.T) {
	node := yamltest.MustParseDocument(`{"a": 1, "b": [1, 2, 3]}`)

	hash := yamlutil.Hash(node.Content[0])

	if got, want := hash, "fZmDQCLyWU2kb+noo3J6wiMTI7Y="; got != want {
		t.Errorf("Hash() = %v, want %v", got, want)
	}
}
