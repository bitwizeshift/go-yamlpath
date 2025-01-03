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
	tableV1.Add("count", funcs.Count).SetArity(arity.None())
	tableV1.Add("distinct", funcs.Distinct).SetArity(arity.None())
	tableV1.Add("isDistinct", funcs.IsDistinct).SetArity(arity.None())
	tableV1.Add("all", funcs.All).SetArity(arity.Exactly(1))
	tableV1.Add("any", funcs.Any).SetArity(arity.Exactly(1))
	tableV1.Add("allTrue", funcs.AllTrue).SetArity(arity.None())
	tableV1.Add("anyTrue", funcs.AnyTrue).SetArity(arity.None())
	tableV1.Add("allFalse", funcs.AllFalse).SetArity(arity.None())
	tableV1.Add("anyFalse", funcs.AnyFalse).SetArity(arity.None())

	// Filter functions
	tableV1.Add("where", funcs.Where).SetArity(arity.Exactly(1))
	tableV1.Add("select", funcs.Select).SetArity(arity.Exactly(1))

	// Subsetting functions
	tableV1.Add("first", funcs.First).SetArity(arity.AtMost(1))
	tableV1.Add("last", funcs.Last).SetArity(arity.AtMost(1))
	tableV1.Add("skip", funcs.Skip).SetArity(arity.Exactly(1))
	tableV1.Add("single", funcs.Single).SetArity(arity.None())

	// String functions
	tableV1.Add("lower", funcs.Lower).SetArity(arity.None())
	tableV1.Add("upper", funcs.Upper).SetArity(arity.None())

	// Conversion functions
	tableV1.Add("toBoolean", funcs.ToBoolean).SetArity(arity.None())
	tableV1.Add("toString", funcs.ToString).SetArity(arity.None())
	tableV1.Add("toNumber", funcs.ToNumber).SetArity(arity.None())
}
