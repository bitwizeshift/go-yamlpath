package compile

import (
	"errors"
	"io"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/parser"
)

// Config provides compilation configuration to the [NewTree] function.
type Config struct {
	Table *invocation.Table
}

// NewTree converts a string FHIRPath expression into the proper Expression
// tree.
func NewTree(str string, cfg *Config) (expr.Expr, error) {
	return NewTreeFromReader(strings.NewReader(str), cfg)
}

// NewTreeFromReader converts a FHIRPath expression from an io.Reader into the
// proper Expression tree.
func NewTreeFromReader(r io.Reader, cfg *Config) (expr.Expr, error) {
	input := antlr.NewIoStream(r)

	lexerErrors := &ErrorListener{}
	lexer := parser.NewyamlpathLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parserErrors := &ErrorListener{}
	parser := parser.NewyamlpathParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	path := parser.Path()

	var errs []error
	errs = append(errs, lexerErrors.Errors...)
	errs = append(errs, parserErrors.Errors...)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	var visitor Visitor
	visitor.FuncTable = cfg.Table
	return visitor.VisitRoot(path)
}
