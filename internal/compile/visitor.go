package compile

import (
	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/parser"
)

// Visitor is a visitor that walks the parse tree to generate an expression
type Visitor struct{}

// VisitRoot visits the root of the parse tree
func (v *Visitor) VisitRoot(ctx parser.IYamlPathContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitPath(ctx *parser.YamlPathContext) (expr.Expression, error) {
	var exprs expr.SequenceExpression
	for _, ctx := range ctx.AllSelector() {
		expr, err := v.visitSelector(ctx)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr)
	}
	return exprs, nil
}

func (v *Visitor) visitSelector(ctx parser.ISelectorContext) (expr.Expression, error) {
	if selector := ctx.DotSelector(); selector != nil {
		return v.visitDotSelector(selector.(*parser.DotSelectorContext))
	}
	if selector := ctx.BracketSelector(); selector != nil {
		return v.visitBracketSelector(selector.(*parser.BracketSelectorContext))
	}
	return nil, ErrInternalf(ctx, "unexpected selector type: %T", ctx)
}

func (v *Visitor) visitDotSelector(ctx *parser.DotSelectorContext) (expr.Expression, error) {
	expr := &expr.FieldExpression{
		Name: ctx.NAME().GetText(),
	}
	return expr, nil
}

func (v *Visitor) visitBracketSelector(ctx *parser.BracketSelectorContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitBracketExpression(ctx *parser.BracketExpressionContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitSlice(ctx *parser.SliceContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitFilter(ctx *parser.FilterContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitUnion(ctx *parser.UnionContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitExpression(ctx *parser.ExpressionContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitSubexpression(ctx *parser.SubexpressionContext) (expr.Expression, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitValue(ctx *parser.ValueContext) (expr.Expression, error) {
	return nil, nil
}

func (v *Visitor) visitQuotedName(ctx *parser.QuotedNameContext) (expr.Expression, error) {
	return nil, nil
}

func (v *Visitor) visitWildcard(ctx *parser.WildcardContext) (expr.Expression, error) {
	return nil, nil
}

func (v *Visitor) visitNode(node antlr.ParseTree) (expr.Expression, error) {
	switch ctx := node.(type) {
	case *parser.YamlPathContext:
		return v.visitPath(ctx)
	case *parser.SelectorContext:
		return v.visitSelector(ctx)
	case *parser.DotSelectorContext:
		return v.visitDotSelector(ctx)
	case *parser.BracketSelectorContext:
		return v.visitBracketSelector(ctx)
	case *parser.BracketExpressionContext:
		return v.visitBracketExpression(ctx)
	case *parser.SliceContext:
		return v.visitSlice(ctx)
	case *parser.FilterContext:
		return v.visitFilter(ctx)
	case *parser.UnionContext:
		return v.visitUnion(ctx)
	case *parser.ExpressionContext:
		return v.visitExpression(ctx)
	case *parser.SubexpressionContext:
		return v.visitSubexpression(ctx)
	case *parser.ValueContext:
		return v.visitValue(ctx)
	case *parser.QuotedNameContext:
		return v.visitQuotedName(ctx)
	case *parser.WildcardContext:
		return v.visitWildcard(ctx)
	}
	return nil, ErrInternalf(node, "unexpected node type: %T", node)
}
