package compile

import (
	"strconv"

	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/parser"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

// Visitor is a visitor that walks the parse tree to generate an expression
type Visitor struct{}

// VisitRoot visits the root of the parse tree
func (v *Visitor) VisitRoot(ctx parser.IPathContext) (expr.Expr, error) {
	return v.visitExpression(ctx.Expression())
}

//------------------------------------------------------------------------------
// Expressions
//------------------------------------------------------------------------------

func (v *Visitor) visitExpression(ctx parser.IExpressionContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.RootExpressionContext:
		return v.visitRootExpression(ctx)
	case *parser.FieldExpressionContext:
		return v.visitFieldExpression(ctx)
	case *parser.RecursiveExpressionContext:
		return v.visitRecursiveExpression(ctx)
	case *parser.IndexExpressionContext:
		return v.visitIndexExpression(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected expression type: %T", ctx)
}

func (v *Visitor) visitRootExpression(ctx *parser.RootExpressionContext) (expr.Expr, error) {
	root := ctx.GetText()
	switch root {
	case "@", "$":
		return &expr.RootExpr{
			Root: root,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unexpected root expression: %q", root)
}

func (v *Visitor) visitFieldExpression(ctx *parser.FieldExpressionContext) (expr.Expr, error) {
	var result expr.SequenceExpr
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	result.Append(left)

	right, err := v.visitInvocation(ctx.Invocation())
	if err != nil {
		return nil, err
	}
	result.Append(right)
	return result, nil
}

func (v *Visitor) visitRecursiveExpression(ctx *parser.RecursiveExpressionContext) (expr.Expr, error) {
	var result expr.SequenceExpr
	result.Append(&expr.RecursiveDescentExpr{})
	if invocation := ctx.Invocation(); invocation != nil {
		right, err := v.visitInvocation(invocation)
		if err != nil {
			return nil, err
		}
		result.Append(right)
	}
	return result, nil
}

func (v *Visitor) visitIndexExpression(ctx *parser.IndexExpressionContext) (expr.Expr, error) {
	var result expr.SequenceExpr
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	result.Append(left)

	right, err := v.visitBracketParam(ctx.BracketParam())
	if err != nil {
		return nil, err
	}
	result.Append(right)
	return result, nil
}

//------------------------------------------------------------------------------
// Invocations
//------------------------------------------------------------------------------

func (v *Visitor) visitInvocation(ctx parser.IInvocationContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.MemberInvocationContext:
		return v.visitMemberInvocation(ctx)
	case *parser.WildcardInvocationContext:
		return v.visitWildcardInvocation(ctx)
	case *parser.FunctionInvocationContext:
		return v.visitFunctionInvocation(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected invocation type: %T", ctx)
}

func (v *Visitor) visitMemberInvocation(ctx *parser.MemberInvocationContext) (expr.Expr, error) {
	return &expr.FieldExpression{
		Names: []string{ctx.GetText()},
	}, nil
}

func (v *Visitor) visitWildcardInvocation(_ *parser.WildcardInvocationContext) (expr.Expr, error) {
	return &expr.WildcardExpr{}, nil
}

func (v *Visitor) visitFunctionInvocation(ctx *parser.FunctionInvocationContext) (expr.Expr, error) {
	return nil, ErrInternalf(ctx, "function invocation not yet implemented")
}

//------------------------------------------------------------------------------
// BracketParams
//------------------------------------------------------------------------------

func (v *Visitor) visitBracketParam(ctx parser.IBracketParamContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.BracketUnionNumberContext:
		return v.visitBracketUnionNumber(ctx)
	case *parser.BracketUnionStringContext:
		return v.visitBracketUnionString(ctx)
	case *parser.BracketWildcardContext:
		return v.visitBracketWildcard(ctx)
	case *parser.BracketSliceContext:
		return v.visitBracketSlice(ctx)
	case *parser.BracketFilterContext:
		return v.visitBracketFilter(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected bracket param type: %T", ctx)
}

func (v *Visitor) visitBracketUnionString(ctx *parser.BracketUnionStringContext) (expr.Expr, error) {
	var names []string
	for _, str := range ctx.AllSTRING() {
		name, err := strconv.Unquote(str.GetText())
		if err != nil {
			return nil, NewSemanticErrorf(str, "string: %s", str.GetText())
		}
		names = append(names, name)
	}
	return &expr.FieldExpression{
		Names: names,
	}, nil
}

func (v *Visitor) visitBracketUnionNumber(ctx *parser.BracketUnionNumberContext) (expr.Expr, error) {
	var indicies []int64
	for _, num := range ctx.AllNUMBER() {
		i, err := strconv.ParseInt(num.GetText(), 10, 64)
		if err != nil {
			return nil, NewSemanticErrorf(num, "number: %s", num.GetText())
		}
		indicies = append(indicies, i)
	}
	return &expr.IndexExpr{
		Indices: indicies,
	}, nil
}

func (v *Visitor) visitBracketWildcard(_ *parser.BracketWildcardContext) (expr.Expr, error) {
	return &expr.WildcardExpr{}, nil
}

func (v *Visitor) visitBracketSlice(ctx *parser.BracketSliceContext) (expr.Expr, error) {
	s, err := expr.ParseSlice(ctx.GetText())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "slice: %s", ctx.GetText())
	}
	return &expr.SliceExpr{
		Slice: s,
	}, nil
}

func (v *Visitor) visitBracketFilter(ctx *parser.BracketFilterContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}
	return &expr.FilterExpr{
		Expr: e,
	}, nil
}

//------------------------------------------------------------------------------
// Subexpressions
//------------------------------------------------------------------------------

func (v *Visitor) visitSubexpression(ctx parser.ISubexpressionContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.AdditiveSubexpressionContext:
		return v.visitAdditiveSubexpression(ctx)
	case *parser.MultiplicativeSubexpressionContext:
		return v.visitMultiplicativeSubexpression(ctx)
	case *parser.InequalitySubexpressionContext:
		return v.visitInequalitySubexpression(ctx)
	case *parser.EqualitySubexpressionContext:
		return v.visitEqualitySubexpression(ctx)
	case *parser.MembershipSubexpressionContext:
		return v.visitMembershipSubexpression(ctx)
	case *parser.AndSubexpressionContext:
		return v.visitAndSubexpression(ctx)
	case *parser.OrSubexpressionContext:
		return v.visitOrSubexpression(ctx)
	case *parser.NegationSubexpressionContext:
		return v.visitNegationSubexpression(ctx)
	case *parser.LiteralSubexpressionContext:
		return v.visitLiteral(ctx.Literal())
	case *parser.RootSubexpressionContext:
		return v.visitExpression(ctx.Expression())
	}
	return nil, ErrInternalf(ctx, "unexpected expression: %q", ctx.GetText())
}

func (v *Visitor) visitAdditiveSubexpression(ctx *parser.AdditiveSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitMultiplicativeSubexpression(ctx *parser.MultiplicativeSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitInequalitySubexpression(ctx *parser.InequalitySubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitEqualitySubexpression(ctx *parser.EqualitySubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitMembershipSubexpression(ctx *parser.MembershipSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitAndSubexpression(ctx *parser.AndSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}

	return &expr.BooleanAndExpr{
		Left:  lhs,
		Right: rhs,
	}, nil
}

func (v *Visitor) visitOrSubexpression(ctx *parser.OrSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}

	return &expr.BooleanOrExpr{
		Left:  lhs,
		Right: rhs,
	}, nil
}

type binaryExpressionContext interface {
	Subexpression(int) parser.ISubexpressionContext
}

func (v *Visitor) visitBinaryExpression(ctx binaryExpressionContext) (expr.Expr, expr.Expr, error) {
	lhs, err := v.visitSubexpression(ctx.Subexpression(0))
	if err != nil {
		return nil, nil, err
	}
	rhs, err := v.visitSubexpression(ctx.Subexpression(1))
	if err != nil {
		return nil, nil, err
	}

	return lhs, rhs, nil
}

func (v *Visitor) visitNegationSubexpression(ctx *parser.NegationSubexpressionContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}

	return &expr.NegationExpr{
		Expr: e,
	}, nil
}

//------------------------------------------------------------------------------
// Literals
//------------------------------------------------------------------------------

func (v *Visitor) visitLiteral(ctx parser.ILiteralContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.StringLiteralContext:
		return v.visitStringLiteral(ctx)
	case *parser.NumberLiteralContext:
		return v.visitNumberLiteral(ctx)
	case *parser.BooleanLiteralContext:
		return v.visitBooleanLiteral(ctx)
	case *parser.NullLiteralContext:
		return v.visitNullLiteral(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected literal subexpression: %q", ctx.GetText())
}

func (v *Visitor) visitStringLiteral(ctx *parser.StringLiteralContext) (expr.Expr, error) {
	s, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "string: %s", ctx.GetText())
	}
	return &expr.ValueExpr{
		Node: yamlutil.String(s),
	}, nil
}

func (v *Visitor) visitNumberLiteral(ctx *parser.NumberLiteralContext) (expr.Expr, error) {
	return &expr.ValueExpr{
		Node: yamlutil.Number(ctx.GetText()),
	}, nil
}

func (v *Visitor) visitBooleanLiteral(ctx *parser.BooleanLiteralContext) (expr.Expr, error) {
	return &expr.ValueExpr{
		Node: yamlutil.Boolean(ctx.GetText()),
	}, nil
}

func (v *Visitor) visitNullLiteral(_ *parser.NullLiteralContext) (expr.Expr, error) {
	return &expr.ValueExpr{
		Node: yamlutil.Null,
	}, nil
}
