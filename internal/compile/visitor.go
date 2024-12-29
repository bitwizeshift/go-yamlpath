package compile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/shopspring/decimal"
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
	return &expr.FieldExpr{
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
	case *parser.UnionNumberBracketContext:
		return v.visitUnionNumberBracket(ctx)
	case *parser.UnionStringBracketContext:
		return v.visitUnionStringBracket(ctx)
	case *parser.WildcardBracketContext:
		return v.visitWildcardBracket(ctx)
	case *parser.SliceBracketContext:
		return v.visitSliceBracket(ctx)
	case *parser.FilterBracketContext:
		return v.visitFilterBracket(ctx)
	case *parser.ScriptBracketContext:
		return v.visitScriptBracket(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected bracket param type: %T", ctx)
}

func (v *Visitor) visitUnionStringBracket(ctx *parser.UnionStringBracketContext) (expr.Expr, error) {
	var names []string
	for _, str := range ctx.AllSTRING() {
		name, err := strconv.Unquote(str.GetText())
		if err != nil {
			return nil, NewSemanticErrorf(str, "string: %s", str.GetText())
		}
		names = append(names, name)
	}
	return &expr.FieldExpr{
		Names: names,
	}, nil
}

func (v *Visitor) visitUnionNumberBracket(ctx *parser.UnionNumberBracketContext) (expr.Expr, error) {
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

func (v *Visitor) visitWildcardBracket(_ *parser.WildcardBracketContext) (expr.Expr, error) {
	return &expr.WildcardExpr{}, nil
}

func (v *Visitor) visitSliceBracket(ctx *parser.SliceBracketContext) (expr.Expr, error) {
	s, err := expr.ParseSlice(ctx.GetText())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "slice: %s", ctx.GetText())
	}
	return &expr.SliceExpr{
		Slice: s,
	}, nil
}

func (v *Visitor) visitFilterBracket(ctx *parser.FilterBracketContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}
	return &expr.FilterExpr{
		Expr: e,
	}, nil
}

func (v *Visitor) visitScriptBracket(ctx *parser.ScriptBracketContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}
	return &expr.ScriptExpr{
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
	case *parser.MatchSubexpressionContext:
		return v.visitMatchSubexpression(ctx)
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
	case *parser.ParenthesisSubexpressionContext:
		return v.visitSubexpression(ctx.Subexpression())
	case *parser.RootSubexpressionContext:
		return v.visitExpression(ctx.Expression())
	case *parser.PolaritySubexpressionContext:
		return v.visitPolaritySubexpression(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected expression: %q", ctx.GetText())
}

func (v *Visitor) visitPolaritySubexpression(ctx *parser.PolaritySubexpressionContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(0))
	switch op {
	case "+":
		return &expr.PrefixPlusExpr{
			Expr: e,
		}, nil
	case "-":
		return &expr.PrefixMinusExpr{
			Expr: e,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unknown polarity operator: %q", op)
}

func (v *Visitor) visitAdditiveSubexpression(ctx *parser.AdditiveSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "+":
		return &expr.ConcatExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	case "-":
		return &expr.ArithmeticExpr{
			Left:  lhs,
			Right: rhs,
			Operation: func(lhs, rhs decimal.Decimal) decimal.Decimal {
				return lhs.Sub(rhs)
			},
		}, nil
	}

	return nil, ErrInternalf(ctx, "unknown binary operator: %q", "")
}

func (v *Visitor) visitMultiplicativeSubexpression(ctx *parser.MultiplicativeSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "*":
		return &expr.ArithmeticExpr{
			Left:  lhs,
			Right: rhs,
			Operation: func(lhs, rhs decimal.Decimal) decimal.Decimal {
				return lhs.Mul(rhs)
			},
		}, nil
	case "/":
		return &expr.ArithmeticExpr{
			Left:  lhs,
			Right: rhs,
			Operation: func(lhs, rhs decimal.Decimal) decimal.Decimal {
				return lhs.Div(rhs)
			},
		}, nil
	case "%":
		return &expr.ArithmeticExpr{
			Left:  lhs,
			Right: rhs,
			Operation: func(lhs, rhs decimal.Decimal) decimal.Decimal {
				return lhs.Mod(rhs)
			},
		}, nil
	}
	_, _ = lhs, rhs

	return nil, ErrInternalf(ctx, "unknown binary operator %q", op)
}

func (v *Visitor) visitInequalitySubexpression(ctx *parser.InequalitySubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "<":
		return &expr.LessExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	case "<=":
		return &expr.NegationExpr{
			Expr: &expr.LessExpr{
				Left:  rhs,
				Right: lhs,
			},
		}, nil
	case ">":
		return &expr.LessExpr{
			Left:  rhs,
			Right: lhs,
		}, nil
	case ">=":
		return &expr.NegationExpr{
			Expr: &expr.LessExpr{
				Left:  lhs,
				Right: rhs,
			},
		}, nil
	}

	return nil, ErrInternalf(ctx, "unknown binary operator %q", op)
}

func (v *Visitor) visitEqualitySubexpression(ctx *parser.EqualitySubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "==":
		return &expr.EqualityExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	case "!=":
		return &expr.InequalityExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	}

	return nil, ErrInternalf(ctx, "unknown binary operator %q", op)
}

func (v *Visitor) visitMatchSubexpression(ctx *parser.MatchSubexpressionContext) (expr.Expr, error) {
	e, err := v.visitSubexpression(ctx.Subexpression())
	if err != nil {
		return nil, err
	}
	r, err := v.visitRegex(ctx.Regex())
	if err != nil {
		return nil, err
	}

	return &expr.MatchExpr{
		Expr:  e,
		Regex: r,
	}, nil
}

func (v *Visitor) visitMembershipSubexpression(ctx *parser.MembershipSubexpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	switch op {
	case "in":
		return &expr.InExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	case "nin":
		return &expr.NegationExpr{
			Expr: &expr.InExpr{
				Left:  lhs,
				Right: rhs,
			},
		}, nil
	case "subsetof":
		return &expr.SubsetOfExpr{
			Left:  lhs,
			Right: rhs,
		}, nil
	}

	return nil, ErrInternalf(ctx, "unknown binary operator %q", op)
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

func (v *Visitor) getTreeText(tree antlr.Tree) string {
	return tree.(interface{ GetText() string }).GetText()
}

func (v *Visitor) visitParamList(ctx parser.IParamListContext) ([]expr.Expr, error) {
	var params []expr.Expr
	for _, sub := range ctx.AllSubexpression() {
		e, err := v.visitSubexpression(sub)
		if err != nil {
			return nil, err
		}
		params = append(params, e)
	}
	return params, nil
}

func (v *Visitor) visitRegex(ctx parser.IRegexContext) (*regexp.Regexp, error) {
	r := ctx.REGEX().GetText()
	r = r[1 : len(r)-1] // trim off outside '/' characters

	flags := ""
	if len(ctx.GetChildren()) == 2 {
		flags = v.getTreeText(ctx.GetChild(1))
	}

	var sb strings.Builder
	for _, flag := range flags {
		fmt.Fprintf(&sb, "(?%s)", string(flag))
	}
	sb.WriteString(r)

	regex, err := regexp.Compile(sb.String())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "regex: %s", r)
	}

	return regex, nil
}
