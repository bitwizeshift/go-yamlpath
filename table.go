package yamlpath

import (
	"sync"

	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/arity"
)

// tableV1 creates the function table for the V1 version of this library.
var tableV1 = sync.OnceValue(func() *invocation.Table {
	table := invocation.NewTable()

	// Existence Functions
	table.Add("empty", funcs.Empty).SetArity(arity.None())
	table.Add("exists", funcs.Exists).SetArity(arity.AtMost(1))
	table.Add("count", funcs.Count).SetArity(arity.None())
	table.Add("distinct", funcs.Distinct).SetArity(arity.None())
	table.Add("isDistinct", funcs.IsDistinct).SetArity(arity.None())
	table.Add("all", funcs.All).SetArity(arity.Exactly(1))
	table.Add("any", funcs.Any).SetArity(arity.Exactly(1))
	table.Add("allTrue", funcs.AllTrue).SetArity(arity.None())
	table.Add("anyTrue", funcs.AnyTrue).SetArity(arity.None())
	table.Add("allFalse", funcs.AllFalse).SetArity(arity.None())
	table.Add("anyFalse", funcs.AnyFalse).SetArity(arity.None())

	// Filter functions
	table.Add("where", funcs.Where).SetArity(arity.Exactly(1))
	table.Add("transform", funcs.Transform).SetArity(arity.Exactly(1))
	table.Add("keys", funcs.Keys).SetArity(arity.None())
	table.Add("select", funcs.Select).SetArity(arity.Any())

	// Subsetting functions
	table.Add("first", funcs.First).SetArity(arity.AtMost(1))
	table.Add("last", funcs.Last).SetArity(arity.AtMost(1))
	table.Add("skip", funcs.Skip).SetArity(arity.Exactly(1))
	table.Add("single", funcs.Single).SetArity(arity.None())

	// String functions
	table.Add("lower", funcs.Lower).SetArity(arity.None())
	table.Add("upper", funcs.Upper).SetArity(arity.None())

	// Conversion functions
	table.Add("toBoolean", funcs.ToBoolean).SetArity(arity.None())
	table.Add("toString", funcs.ToString).SetArity(arity.None())
	table.Add("toNumber", funcs.ToNumber).SetArity(arity.None())
	table.Add("toSequence", funcs.ToSequence).SetArity(arity.None())

	// Reflection
	table.Add("reflect", funcs.Reflect).SetArity(arity.None())
	return table
})
