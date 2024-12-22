package compile

import (
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/parser"
)

type Visitor struct{}

func (v *Visitor) VisitRoot(ctx parser.IYamlPathContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitRoot")
}

func (v *Visitor) VisitSelector(ctx *parser.SelectorContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitSelector")
}

func (v *Visitor) VisitDotSelector(ctx *parser.DotSelectorContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitDotSelector")
}

func (v *Visitor) VisitBracketSelector(ctx *parser.BracketSelectorContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitBracketSelector")
}

func (v *Visitor) VisitBracketExpression(ctx *parser.BracketExpressionContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitBracketExpression")
}

func (v *Visitor) VisitSlice(ctx *parser.SliceContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitSlice")
}

func (v *Visitor) VisitFilter(ctx *parser.FilterContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitFilter")
}

func (v *Visitor) VisitUnion(ctx *parser.UnionContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitUnion")
}

func (v *Visitor) VisitExpression(ctx *parser.ExpressionContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitExpression")
}

func (v *Visitor) VisitSubexpression(ctx *parser.SubexpressionContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitSubexpression")
}

func (v *Visitor) VisitValue(ctx *parser.ValueContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitValue")
}

func (v *Visitor) VisitQuotedName(ctx *parser.QuotedNameContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitQuotedName")
}

func (v *Visitor) VisitWildcard(ctx *parser.WildcardContext) (expr.Expression, error) {
	return nil, ErrNotImplemented("VisitWildcard")
}
