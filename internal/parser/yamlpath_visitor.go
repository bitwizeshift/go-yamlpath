// Code generated from yamlpath.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // yamlpath

import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by yamlpathParser.
type yamlpathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by yamlpathParser#yamlPath.
	VisitYamlPath(ctx *YamlPathContext) interface{}

	// Visit a parse tree produced by yamlpathParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by yamlpathParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by yamlpathParser#dotSelector.
	VisitDotSelector(ctx *DotSelectorContext) interface{}

	// Visit a parse tree produced by yamlpathParser#bracketSelector.
	VisitBracketSelector(ctx *BracketSelectorContext) interface{}

	// Visit a parse tree produced by yamlpathParser#bracketExpression.
	VisitBracketExpression(ctx *BracketExpressionContext) interface{}

	// Visit a parse tree produced by yamlpathParser#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by yamlpathParser#filter.
	VisitFilter(ctx *FilterContext) interface{}

	// Visit a parse tree produced by yamlpathParser#union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by yamlpathParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by yamlpathParser#subexpression.
	VisitSubexpression(ctx *SubexpressionContext) interface{}

	// Visit a parse tree produced by yamlpathParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by yamlpathParser#quotedName.
	VisitQuotedName(ctx *QuotedNameContext) interface{}

	// Visit a parse tree produced by yamlpathParser#wildcard.
	VisitWildcard(ctx *WildcardContext) interface{}

}