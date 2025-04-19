package compile

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/parser"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Visitor is a visitor that walks the parse tree to generate an expression
type Visitor struct {
	FuncTable *invocation.Table
	Constants map[string][]*yaml.Node
}

// VisitRoot visits the root of the parse tree
func (v *Visitor) VisitRoot(ctx parser.IPathContext) (expr.Expr, error) {
	return v.visitExpression(ctx.Expression())
}

//------------------------------------------------------------------------------
// Expressions
//------------------------------------------------------------------------------

func (v *Visitor) visitExpression(ctx parser.IExpressionContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.TermExpressionContext:
		return v.visitTermExpression(ctx)
	case *parser.FieldExpressionContext:
		return v.visitFieldExpression(ctx)
	case *parser.IndexExpressionContext:
		return v.visitIndexExpression(ctx)
	case *parser.AdditiveExpressionContext:
		return v.visitAdditiveExpression(ctx)
	case *parser.MultiplicativeExpressionContext:
		return v.visitMultiplicativeExpression(ctx)
	case *parser.InequalityExpressionContext:
		return v.visitInequalityExpression(ctx)
	case *parser.EqualityExpressionContext:
		return v.visitEqualityExpression(ctx)
	case *parser.MatchExpressionContext:
		return v.visitMatchExpression(ctx)
	case *parser.MembershipExpressionContext:
		return v.visitMembershipExpression(ctx)
	case *parser.AndExpressionContext:
		return v.visitAndExpression(ctx)
	case *parser.OrExpressionContext:
		return v.visitOrExpression(ctx)
	case *parser.NegationExpressionContext:
		return v.visitNegationExpression(ctx)
	case *parser.UnionExpressionContext:
		return v.visitUnionExpression(ctx)
	case *parser.ParenthesisExpressionContext:
		return v.visitExpression(ctx.Expression())
	case *parser.PolarityExpressionContext:
		return v.visitPolarityExpression(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected expression type: %T", ctx)
}

func (v *Visitor) visitTermExpression(ctx *parser.TermExpressionContext) (expr.Expr, error) {
	return v.visitTerm(ctx.Term())
}

func (v *Visitor) visitTerm(ctx parser.ITermContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.RootTermContext:
		return v.visitRootTerm(ctx)
	case *parser.ExternalConstantTermContext:
		return v.visitExternalConstantTerm(ctx)
	case *parser.LiteralTermContext:
		return v.visitLiteral(ctx.Literal())
	case *parser.InvocationTermContext:
		return v.visitInvocation(ctx.Invocation())
	}
	return nil, ErrInternalf(ctx, "unexpected term type: %T", ctx)
}

func (v *Visitor) visitRootTerm(ctx *parser.RootTermContext) (expr.Expr, error) {
	root := ctx.GetText()
	switch root {
	case "@", "$":
		return &expr.RootExpr{
			Root: root,
		}, nil
	}
	return nil, ErrInternalf(ctx, "unexpected root term: %q", root)
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

func (v *Visitor) visitIndexExpression(ctx *parser.IndexExpressionContext) (expr.Expr, error) {
	var result expr.SequenceExpr
	left, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	result.Append(left)

	right, err := v.visitBracketParam(ctx.IndexParam())
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
	identifier, err := v.visitIdentifier(ctx.Identifier())
	if err != nil {
		return nil, err
	}
	return &expr.FieldExpr{
		Fields: []string{identifier},
	}, nil
}

func (v *Visitor) visitIdentifier(ctx parser.IIdentifierContext) (string, error) {
	switch ctx := ctx.(type) {
	case *parser.PlainIdentifierContext:
		return v.visitPlainIdentifier(ctx)
	case *parser.QuotedIdentifierContext:
		return v.visitQuotedIdentifier(ctx)
	}
	return "", ErrInternalf(ctx, "unexpected identifier type: %T", ctx)
}

func (v *Visitor) visitPlainIdentifier(ctx *parser.PlainIdentifierContext) (string, error) {
	return ctx.GetText(), nil
}

func (v *Visitor) visitQuotedIdentifier(ctx *parser.QuotedIdentifierContext) (string, error) {
	text := ctx.GetText()
	return strconv.Unquote(text)
}

func (v *Visitor) visitWildcardInvocation(_ *parser.WildcardInvocationContext) (expr.Expr, error) {
	return &expr.WildcardExpr{}, nil
}

func (v *Visitor) visitFunctionInvocation(ctx *parser.FunctionInvocationContext) (expr.Expr, error) {
	params, err := v.visitParamList(ctx.ParamList())
	if err != nil {
		return nil, err
	}
	identifier := ctx.Identifier().GetText()
	entry, ok := v.FuncTable.Lookup(identifier)
	if !ok {
		return nil, NewSemanticErrorf(ctx, "unknown function: %s", identifier)
	}
	if err := entry.TestArity(len(params)); err != nil {
		return nil, NewSemanticErrorf(ctx, "%w", err)
	}
	return &expr.FuncExpr{
		Func:       entry,
		Parameters: params,
	}, nil
}

type parameter struct{ expr.Expr }

func (p parameter) GetArg(ctx invocation.Context) ([]*yaml.Node, error) {
	return p.Eval(ctx)
}

var _ invocation.Parameter = (*parameter)(nil)

//------------------------------------------------------------------------------
// BracketParams
//------------------------------------------------------------------------------

func (v *Visitor) visitBracketParam(ctx parser.IIndexParamContext) (expr.Expr, error) {
	switch ctx := ctx.(type) {
	case *parser.WildcardIndexContext:
		return v.visitWildcardIndex(ctx)
	case *parser.SliceIndexContext:
		return v.visitSliceIndex(ctx)
	case *parser.ExpressionIndexContext:
		return v.visitExpressionIndex(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected bracket param type: %T", ctx)
}

func (v *Visitor) visitWildcardIndex(_ *parser.WildcardIndexContext) (expr.Expr, error) {
	return &expr.WildcardIndexExpr{}, nil
}

func (v *Visitor) visitSliceIndex(ctx *parser.SliceIndexContext) (expr.Expr, error) {
	s, err := expr.ParseSlice(ctx.GetText())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "slice: %s", ctx.GetText())
	}
	return &expr.SliceExpr{
		Slice: s,
	}, nil
}

func (v *Visitor) visitExpressionIndex(ctx *parser.ExpressionIndexContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return &expr.IndexExpr{
		Indices: e,
	}, nil
}

func (v *Visitor) visitPolarityExpression(ctx *parser.PolarityExpressionContext) (expr.Expr, error) {
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

func (v *Visitor) visitAdditiveExpression(ctx *parser.AdditiveExpressionContext) (expr.Expr, error) {
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
			Left:      lhs,
			Right:     rhs,
			Operation: expr.Subtraction,
		}, nil
	}

	return nil, ErrInternalf(ctx, "unhandled binary operator: %q", "")
}

func (v *Visitor) visitUnionExpression(ctx *parser.UnionExpressionContext) (expr.Expr, error) {
	exprs := ctx.AllExpression()
	union := make(expr.UnionExpr, 0, len(exprs))
	for _, sub := range exprs {
		e, err := v.visitExpression(sub)
		if err != nil {
			return nil, err
		}
		union = append(union, e)
	}
	return union, nil
}

func (v *Visitor) visitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	var operation expr.ArithmeticOp
	switch op {
	case "*":
		operation = expr.Multiplication
	case "/":
		operation = expr.Division
	case "%":
		operation = expr.Modulus
	default:
		return nil, ErrInternalf(ctx, "unhandled binary operator %q", op)
	}
	return &expr.ArithmeticExpr{
		Left:      lhs,
		Right:     rhs,
		Operation: operation,
	}, nil
}

func (v *Visitor) visitInequalityExpression(ctx *parser.InequalityExpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}
	op := v.getTreeText(ctx.GetChild(1))
	var compare expr.Comparator
	switch op {
	case "<":
		compare = expr.CompareLess
	case "<=":
		compare = expr.CompareLessEqual
	case ">":
		compare = expr.CompareGreater
	case ">=":
		compare = expr.CompareGreaterEqual
	default:
		return nil, ErrInternalf(ctx, "unhandled operator %q", op)
	}
	return &expr.CompareExpr{
		Left:    lhs,
		Right:   rhs,
		Compare: compare,
	}, nil
}

func (v *Visitor) visitEqualityExpression(ctx *parser.EqualityExpressionContext) (expr.Expr, error) {
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
		return &expr.NegationExpr{
			Expr: &expr.EqualityExpr{
				Left:  lhs,
				Right: rhs,
			},
		}, nil
	}

	return nil, ErrInternalf(ctx, "unknown binary operator %q", op)
}

func (v *Visitor) visitMatchExpression(ctx *parser.MatchExpressionContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
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

func (v *Visitor) visitMembershipExpression(ctx *parser.MembershipExpressionContext) (expr.Expr, error) {
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

func (v *Visitor) visitAndExpression(ctx *parser.AndExpressionContext) (expr.Expr, error) {
	lhs, rhs, err := v.visitBinaryExpression(ctx)
	if err != nil {
		return nil, err
	}

	return &expr.BooleanAndExpr{
		Left:  lhs,
		Right: rhs,
	}, nil
}

func (v *Visitor) visitOrExpression(ctx *parser.OrExpressionContext) (expr.Expr, error) {
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
	Expression(int) parser.IExpressionContext
}

func (v *Visitor) visitBinaryExpression(ctx binaryExpressionContext) (expr.Expr, expr.Expr, error) {
	lhs, err := v.visitExpression(ctx.Expression(0))
	if err != nil {
		return nil, nil, err
	}
	rhs, err := v.visitExpression(ctx.Expression(1))
	if err != nil {
		return nil, nil, err
	}

	return lhs, rhs, nil
}

func (v *Visitor) visitNegationExpression(ctx *parser.NegationExpressionContext) (expr.Expr, error) {
	e, err := v.visitExpression(ctx.Expression())
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
	case *parser.ListLiteralContext:
		return v.visitListLiteral(ctx)
	case *parser.MapLiteralContext:
		return v.visitMapLiteral(ctx)
	}
	return nil, ErrInternalf(ctx, "unexpected literal expression: %q", ctx.GetText())
}

func (v *Visitor) visitStringLiteral(ctx *parser.StringLiteralContext) (expr.Expr, error) {
	s, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		return nil, NewSemanticErrorf(ctx, "string: %s", ctx.GetText())
	}
	return &expr.LiteralExpr{
		Nodes: []*yaml.Node{yamlconv.String(s)},
	}, nil
}

func (v *Visitor) visitNumberLiteral(ctx *parser.NumberLiteralContext) (expr.Expr, error) {
	return &expr.LiteralExpr{
		Nodes: []*yaml.Node{yamlconv.NumberString(ctx.GetText())},
	}, nil
}

func (v *Visitor) visitBooleanLiteral(ctx *parser.BooleanLiteralContext) (expr.Expr, error) {
	return &expr.LiteralExpr{
		Nodes: []*yaml.Node{yamlconv.BoolString(ctx.GetText())},
	}, nil
}

func (v *Visitor) visitNullLiteral(_ *parser.NullLiteralContext) (expr.Expr, error) {
	return &expr.LiteralExpr{
		Nodes: []*yaml.Node{yamlconv.Null()},
	}, nil
}

func (v *Visitor) visitListLiteral(ctx *parser.ListLiteralContext) (expr.Expr, error) {
	nodes, err := v.visitYAMLText(ctx.GetText())
	if err != nil {
		return nil, err
	}
	return &expr.LiteralExpr{
		Nodes: nodes,
	}, nil
}

func (v *Visitor) visitMapLiteral(ctx *parser.MapLiteralContext) (expr.Expr, error) {
	nodes, err := v.visitYAMLText(ctx.GetText())
	if err != nil {
		return nil, err
	}
	return &expr.LiteralExpr{
		Nodes: nodes,
	}, nil
}

func (v *Visitor) visitYAMLText(str string) ([]*yaml.Node, error) {
	decoder := yaml.NewDecoder(strings.NewReader(str))

	var node yaml.Node
	if err := decoder.Decode(&node); err != nil {
		return nil, err
	}
	return yamlconv.FlattenDocuments(v.setDefaults(&node)), nil
}

func (v *Visitor) setDefaults(node *yaml.Node) *yaml.Node {
	node.Line = 0
	node.Column = 0
	node.Style = 0
	for _, node := range node.Content {
		v.setDefaults(node)
	}
	return node
}

func (v *Visitor) getTreeText(tree antlr.Tree) string {
	return tree.(interface{ GetText() string }).GetText()
}

func (v *Visitor) visitParamList(ctx parser.IParamListContext) ([]invocation.Parameter, error) {
	if ctx == nil {
		return nil, nil
	}
	var params []invocation.Parameter
	for _, sub := range ctx.AllExpression() {
		e, err := v.visitExpression(sub)
		if err != nil {
			return nil, err
		}
		params = append(params, parameter{e})
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

func (v *Visitor) visitExternalConstantTerm(ctx *parser.ExternalConstantTermContext) (expr.Expr, error) {
	str, err := v.visitIdentifier(ctx.ExternalConstant().Identifier())
	if err != nil {
		return nil, err
	}
	if v.Constants == nil {
		return nil, NewSemanticErrorf(ctx, "external constant %q not defined", str)
	}
	constant, ok := v.Constants[str]
	if !ok {
		return nil, NewSemanticErrorf(ctx, "external constant %q not defined", str)
	}
	expr := &expr.LiteralExpr{
		Nodes: constant,
	}
	return expr, nil
}
