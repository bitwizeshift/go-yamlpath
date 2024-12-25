package compile

import (
	"strconv"

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
		return v.visitDotSelector(selector)
	}
	if selector := ctx.BracketSelector(); selector != nil {
		return v.visitBracketSelector(selector)
	}
	if selector := ctx.RecursiveSelector(); selector != nil {
		return v.visitRecursiveSelector(selector)
	}
	return nil, ErrInternalf(ctx, "unexpected selector type: %T", ctx)
}

func (v *Visitor) visitDotSelector(ctx parser.IDotSelectorContext) (expr.Expression, error) {
	var result expr.Expression
	if node := ctx.NAME(); node != nil {
		result = &expr.FieldExpression{
			Name: node.GetText(),
		}
	} else if node := ctx.WILDCARD(); node != nil {
		result = &expr.WildcardExpression{}
	} else {
		return nil, ErrInternalf(ctx, "unhandled dot selector: %q", ctx.GetText())
	}
	return result, nil
}

func (v *Visitor) visitRecursiveSelector(ctx parser.IRecursiveSelectorContext) (expr.Expression, error) {
	var result expr.Expression = &expr.RecursiveDescentExpression{}
	if node := ctx.NAME(); node != nil {
		result = expr.SequenceExpression{result, &expr.FieldExpression{
			Name: node.GetText(),
		}}
	} else if node := ctx.WILDCARD(); node != nil {
		result = expr.SequenceExpression{result, &expr.WildcardExpression{}}
	}
	return result, nil
}

func (v *Visitor) visitBracketSelector(ctx parser.IBracketSelectorContext) (expr.Expression, error) {
	return v.visitBracketExpression(ctx.BracketExpression())
}

func (v *Visitor) visitBracketExpression(ctx parser.IBracketExpressionContext) (expr.Expression, error) {
	if qn := ctx.QuotedName(); qn != nil {
		text := qn.GetText()
		name, err := strconv.Unquote(text)
		if err != nil {
			return nil, NewSemanticErrorf(qn, "string: %s", text)
		}

		return &expr.FieldExpression{
			Name: name,
		}, nil
	}
	if wildcard := ctx.WILDCARD(); wildcard != nil {
		return &expr.WildcardExpression{}, nil
	}
	if number := ctx.NUMBER(); number != nil {
		i, err := strconv.ParseInt(number.GetText(), 10, 64)
		if err != nil {
			return nil, NewSemanticErrorf(number, "number: %s", number.GetText())
		}
		return &expr.IndexExpression{
			Index: i,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unexpected bracket expression: %q", ctx.GetText())
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

func (v *Visitor) visitNode(node antlr.ParseTree) (expr.Expression, error) {
	switch ctx := node.(type) {
	case *parser.YamlPathContext:
		return v.visitPath(ctx)
	case parser.ISelectorContext:
		return v.visitSelector(ctx)
	case parser.IDotSelectorContext:
		return v.visitDotSelector(ctx)
	case parser.IRecursiveSelectorContext:
		return v.visitRecursiveSelector(ctx)
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
	}
	return nil, ErrInternalf(node, "unexpected node type: %T", node)
}
