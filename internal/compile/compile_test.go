package compile_test

import (
	"testing"

	"rodusek.dev/pkg/yamlpath/internal/compile"
)

func TestNewTree(t *testing.T) {
	compile.NewTree("$.store.book[0].title")
}
