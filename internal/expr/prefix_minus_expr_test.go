package expr_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestPrefixMinusExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:    "subexpr returns error",
			expr:    ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name: "subexpr returns empty",
			expr: ExprReturnsNodes(),
		}, {
			name:    "subexpr returns multiple values",
			expr:    ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			wantErr: expr.ErrEval,
		}, {
			name:    "subexpr returns single non-scalar value",
			expr:    ExprReturnsNodes(YAML(t, `{"name": "Alice", "age": 30}`)),
			wantErr: expr.ErrEval,
		}, {
			name:    "subexpr returns scalar non-numeric value",
			expr:    ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: expr.ErrEval,
		}, {
			name: "subexpr returns positive scalar integer value",
			expr: ExprReturnsNodes(yamlutil.Number("42")),
			want: []*yaml.Node{yamlutil.Number("-42")},
		}, {
			name: "subexpr returns positive scalar float value",
			expr: ExprReturnsNodes(yamlutil.Number("42.5")),
			want: []*yaml.Node{yamlutil.Number("-42.5")},
		}, {
			name: "subexpr returns negative scalar integer value",
			expr: ExprReturnsNodes(yamlutil.Number("-42")),
			want: []*yaml.Node{yamlutil.Number("42")},
		}, {
			name: "subexpr returns negative scalar float value",
			expr: ExprReturnsNodes(yamlutil.Number("-42.5")),
			want: []*yaml.Node{yamlutil.Number("42.5")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.PrefixMinusExpr{Expr: tc.expr}

			got, err := sut.Eval(context.Background(), tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("PrefixPlusExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("PrefixPlusExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
