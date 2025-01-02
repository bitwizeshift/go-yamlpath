package main

import (
	"os"
	"testing"
)

func TestMain(*testing.T) {
	os.Args = []string{"yamlpath", "-i", ".github/workflows/build-and-test.yaml", "$.jobs.*"}
	os.Chdir("/Users/bitwize/Development/rodusek.dev/pkg/go-yamlpath")
	main()
}
