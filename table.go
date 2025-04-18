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
	table.Add("startsWith", funcs.StartsWith).SetArity(arity.Exactly(1))
	table.Add("endsWith", funcs.EndsWith).SetArity(arity.Exactly(1))
	table.Add("contains", funcs.Contains).SetArity(arity.Exactly(1))
	table.Add("indexOf", funcs.IndexOf).SetArity(arity.Exactly(1))
	table.Add("substring", funcs.Substring).SetArity(arity.ClosedRange(1, 2))
	table.Add("replace", funcs.Replace).SetArity(arity.Exactly(2))
	table.Add("length", funcs.Length).SetArity(arity.None())
	table.Add("split", funcs.Split).SetArity(arity.Exactly(1))
	table.Add("toChars", funcs.ToChars).SetArity(arity.None())
	table.Add("matches", funcs.Matches).SetArity(arity.Exactly(1))
	table.Add("replaceMatches", funcs.ReplaceMatches).SetArity(arity.Exactly(2))

	// Math functions
	table.Add("abs", funcs.Abs).SetArity(arity.None())
	table.Add("ceil", funcs.Ceil).SetArity(arity.None())
	table.Add("floor", funcs.Floor).SetArity(arity.None())
	table.Add("exp", funcs.Exp).SetArity(arity.None())
	table.Add("ln", funcs.Ln).SetArity(arity.None())
	table.Add("log", funcs.Log).SetArity(arity.None())
	table.Add("pow", funcs.Pow).SetArity(arity.Exactly(1))
	table.Add("round", funcs.Round).SetArity(arity.ClosedRange(0, 1))
	table.Add("truncate", funcs.Truncate).SetArity(arity.None())
	table.Add("min", funcs.Min).SetArity(arity.Any())
	table.Add("max", funcs.Max).SetArity(arity.Any())
	table.Add("sum", funcs.Sum).SetArity(arity.Any())

	// Conversion functions
	table.Add("toString", funcs.ToString).SetArity(arity.None())
	table.Add("convertsToString", funcs.ConvertsToString).SetArity(arity.None())
	table.Add("toBoolean", funcs.ToBoolean).SetArity(arity.None())
	table.Add("convertsToBoolean", funcs.ConvertsToBoolean).SetArity(arity.None())
	table.Add("toNumber", funcs.ToNumber).SetArity(arity.None())
	table.Add("convertsToNumber", funcs.ConvertsToNumber).SetArity(arity.None())
	table.Add("toInteger", funcs.ToInteger).SetArity(arity.None())
	table.Add("convertsToInteger", funcs.ConvertsToInteger).SetArity(arity.None())
	table.Add("toFloat", funcs.ToFloat).SetArity(arity.None())
	table.Add("convertsToFloat", funcs.ConvertsToFloat).SetArity(arity.None())
	table.Add("toSequence", funcs.ToSequence).SetArity(arity.None())

	// Reflection
	table.Add("reflect", funcs.Reflect).SetArity(arity.None())
	table.Add("isString", funcs.IsString).SetArity(arity.None())
	table.Add("isInteger", funcs.IsInteger).SetArity(arity.None())
	table.Add("isFloat", funcs.IsFloat).SetArity(arity.None())
	table.Add("isBoolean", funcs.IsBoolean).SetArity(arity.None())
	table.Add("isNull", funcs.IsNull).SetArity(arity.None())
	table.Add("isScalar", funcs.IsScalar).SetArity(arity.None())
	table.Add("isSequence", funcs.IsSequence).SetArity(arity.None())
	table.Add("isMapping", funcs.IsMapping).SetArity(arity.None())
	return table
})
