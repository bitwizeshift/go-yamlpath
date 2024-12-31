package yamlpath

import (
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
)

// tableV1 is the function table for V1 of the YAMLPath library.
var tableV1 *invocation.Table

func init() {
	tableV1 = invocation.NewTable()

	// Existence Functions
	tableV1.Add("empty", funcs.Empty).SetArity(arity.None())
	tableV1.Add("exists", funcs.Exists).SetArity(arity.AtMost(1))

	// Filter functions
	tableV1.Add("where", funcs.Where).SetArity(arity.Exactly(1))
}
