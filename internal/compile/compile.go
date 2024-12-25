package compile

import (
	"errors"

	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/parser"
)

// NewTree converts a string FHIRPath expression into the proper Expression
// tree.
func NewTree(str string) (expr.Expr, error) {
	input := antlr.NewInputStream(str)

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
	path := parser.YamlPath()

	var errs []error
	errs = append(errs, lexerErrors.Errors...)
	errs = append(errs, parserErrors.Errors...)
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	var visitor Visitor
	return visitor.VisitRoot(path)
}
