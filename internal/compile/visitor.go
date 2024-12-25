package compile

import (
	"strconv"

	antlr "github.com/antlr4-go/antlr/v4"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/parser"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Visitor is a visitor that walks the parse tree to generate an expression
type Visitor struct{}

// VisitRoot visits the root of the parse tree
func (v *Visitor) VisitRoot(ctx parser.IYamlPathContext) (expr.Expr, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitYAMLPath(ctx parser.IYamlPathContext) (expr.Expr, error) {
	return v.visitPath(ctx.Path())
}

func (v *Visitor) visitPath(ctx parser.IPathContext) (expr.Expr, error) {
	root, err := v.visitRoot(ctx.Root())
	if err != nil {
		return nil, err
	}

	var exprs expr.SequenceExpr
	exprs.Append(root)
	for _, ctx := range ctx.AllSelector() {
		expr, err := v.visitSelector(ctx)
		if err != nil {
			return nil, err
		}
		exprs.Append(expr)
	}
	return exprs, nil
}

func (v *Visitor) visitSelector(ctx parser.ISelectorContext) (expr.Expr, error) {
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

func (v *Visitor) visitDotSelector(ctx parser.IDotSelectorContext) (expr.Expr, error) {
	var result expr.Expr
	if node := ctx.NAME(); node != nil {
		result = &expr.FieldExpression{
			Name: node.GetText(),
		}
	} else if node := ctx.WILDCARD(); node != nil {
		result = &expr.WildcardExpr{}
	} else {
		return nil, ErrInternalf(ctx, "unhandled dot selector: %q", ctx.GetText())
	}
	return result, nil
}

func (v *Visitor) visitRecursiveSelector(ctx parser.IRecursiveSelectorContext) (expr.Expr, error) {
	var result expr.Expr = &expr.RecursiveDescentExpr{}
	if node := ctx.NAME(); node != nil {
		result = expr.SequenceExpr{result, &expr.FieldExpression{
			Name: node.GetText(),
		}}
	} else if node := ctx.WILDCARD(); node != nil {
		result = expr.SequenceExpr{result, &expr.WildcardExpr{}}
	}
	return result, nil
}

func (v *Visitor) visitBracketSelector(ctx parser.IBracketSelectorContext) (expr.Expr, error) {
	return v.visitBracketExpression(ctx.BracketExpression())
}

func (v *Visitor) visitBracketExpression(ctx parser.IBracketExpressionContext) (expr.Expr, error) {
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
		return &expr.WildcardExpr{}, nil
	}
	if number := ctx.NUMBER(); number != nil {
		i, err := strconv.ParseInt(number.GetText(), 10, 64)
		if err != nil {
			return nil, NewSemanticErrorf(number, "number: %s", number.GetText())
		}
		return &expr.IndexExpr{
			Index: i,
		}, nil
	}
	if slice := ctx.Slice(); slice != nil {
		s, err := expr.ParseSlice(slice.GetText())
		if err != nil {
			return nil, NewSemanticErrorf(slice, "slice: %s", slice.GetText())
		}
		return &expr.SliceExpr{
			Slice: s,
		}, nil
	}
	if union := ctx.UnionString(); union != nil {
		var u expr.Union
		for _, str := range union.AllQuotedName() {
			text := str.GetText()
			name, err := strconv.Unquote(text)
			if err != nil {
				return nil, NewSemanticErrorf(str, "string: %s", text)
			}
			u = append(u, name)
		}
		return &expr.UnionExpr{
			Union: u,
		}, nil
	}
	if union := ctx.UnionIndices(); union != nil {
		var u expr.Union
		for _, num := range union.AllNUMBER() {
			i, err := strconv.ParseInt(num.GetText(), 10, 64)
			if err != nil {
				return nil, NewSemanticErrorf(num, "number: %s", num.GetText())
			}
			u = append(u, i)
		}
		return &expr.UnionExpr{
			Union: u,
		}, nil
	}
	if filter := ctx.Filter(); filter != nil {
		return v.visitFilter(filter)
	}
	return nil, ErrInternalf(ctx, "unexpected bracket expression: %q", ctx.GetText())
}

func (v *Visitor) visitSlice(ctx *parser.SliceContext) (expr.Expr, error) {
	return v.visitNode(ctx)
}

func (v *Visitor) visitFilter(ctx parser.IFilterContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return &expr.FilterExpr{
		Expr: e,
	}, nil
}

func (v *Visitor) visitExpression(ctx parser.IExpressionContext) (expr.Expr, error) {
	if compare := ctx.CompareExpr(); compare != nil {
		return v.visitCompareExpr(compare)
	}
	if boolean := ctx.BooleanExpr(); boolean != nil {
		return v.visitBooleanExpr(boolean)
	}
	if containment := ctx.ContainmentExpr(); containment != nil {
		return v.visitContainmentExpr(containment)
	}
	if arithmetic := ctx.ArithmeticExpr(); arithmetic != nil {
		return v.visitArithmeticExpr(arithmetic)
	}
	if negation := ctx.NegationExpr(); negation != nil {
		return v.visitNegationExpr(negation)
	}
	if subexpr := ctx.Subexpression(); subexpr != nil {
		return v.visitSubexpression(subexpr)
	}
	return nil, ErrInternalf(ctx, "unexpected expression: %q", ctx.GetText())
}

func (v *Visitor) visitCompareExpr(ctx parser.ICompareExprContext) (expr.Expr, error) {
	return nil, ErrNotImplemented("compare")
}

func (v *Visitor) visitBooleanExpr(ctx parser.IBooleanExprContext) (expr.Expr, error) {
	return nil, ErrNotImplemented("boolean")
}

func (v *Visitor) visitContainmentExpr(ctx parser.IContainmentExprContext) (expr.Expr, error) {
	return nil, ErrNotImplemented("containment")
}

func (v *Visitor) visitArithmeticExpr(ctx parser.IArithmeticExprContext) (expr.Expr, error) {
	return nil, ErrNotImplemented("arithmetic")
}

func (v *Visitor) visitNegationExpr(ctx parser.INegationExprContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}

	return &expr.NegationExpr{
		Expr: e,
	}, nil
}

func (v *Visitor) visitSubexpression(ctx parser.ISubexpressionContext) (expr.Expr, error) {
	if path := ctx.Path(); path != nil {
		return v.visitPath(path)
	}
	if value := ctx.Value(); v != nil {
		return v.visitValue(value)
	}
	return nil, ErrInternalf(ctx, "unexpected subexpression: %q", ctx.GetText())
}

func (v *Visitor) visitValue(ctx parser.IValueContext) (expr.Expr, error) {
	if b := ctx.BOOLEAN(); b != nil {
		return &expr.ValueExpr{
			Node: yamlutil.Boolean(b.GetText()),
		}, nil
	}
	if n := ctx.NUMBER(); n != nil {
		return &expr.ValueExpr{
			Node: yamlutil.Number(n.GetText()),
		}, nil
	}
	if s := ctx.STRING(); s != nil {
		s, err := strconv.Unquote(s.GetText())
		if err != nil {
			return nil, NewSemanticErrorf(ctx, "string: %s", s)
		}
		return &expr.ValueExpr{
			Node: yamlutil.String(s),
		}, nil
	}
	if n := ctx.NULL(); n != nil {
		return &expr.ValueExpr{
			Node: yamlutil.Null,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unexpected value expression: %q", ctx.GetText())
}

func (v *Visitor) visitRoot(ctx parser.IRootContext) (expr.Expr, error) {
	root := ctx.GetText()
	switch root {
	case "@", "$":
		return &expr.RootExpr{
			Root: root,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unexpected root expression: %q", root)
}

func (v *Visitor) visitNode(node antlr.ParseTree) (expr.Expr, error) {
	switch ctx := node.(type) {
	case *parser.YamlPathContext:
		return v.visitYAMLPath(ctx)
	case parser.IRootContext:
		return v.visitRoot(ctx)
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
	case *parser.ExpressionContext:
		return v.visitExpression(ctx)
	case parser.ICompareExprContext:
		return v.visitCompareExpr(ctx)
	case parser.IBooleanExprContext:
		return v.visitBooleanExpr(ctx)
	case parser.IContainmentExprContext:
		return v.visitContainmentExpr(ctx)
	case parser.IArithmeticExprContext:
		return v.visitArithmeticExpr(ctx)
	case parser.INegationExprContext:
		return v.visitNegationExpr(ctx)
	case parser.ISubexpressionContext:
		return v.visitSubexpression(ctx)
	case parser.IValueContext:
		return v.visitValue(ctx)
	}
	return nil, ErrInternalf(node, "unexpected node type: %T", node)
}
