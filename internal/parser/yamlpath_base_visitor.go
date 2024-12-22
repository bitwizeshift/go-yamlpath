// Code generated from yamlpath.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // yamlpath

import "github.com/antlr4-go/antlr/v4"


type BaseyamlpathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseyamlpathVisitor) VisitYamlPath(ctx *YamlPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitDotSelector(ctx *DotSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitBracketSelector(ctx *BracketSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitBracketExpression(ctx *BracketExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitFilter(ctx *FilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitUnion(ctx *UnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitSubexpression(ctx *SubexpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitQuotedName(ctx *QuotedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseyamlpathVisitor) VisitWildcard(ctx *WildcardContext) interface{} {
	return v.VisitChildren(ctx)
}
