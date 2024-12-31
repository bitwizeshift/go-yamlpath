package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestPrefixPlusExpr(t *testing.T) {
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
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name: "subexpr returns empty",
			expr: exprtest.Return(),
		}, {
			name:    "subexpr returns multiple values",
			expr:    exprtest.Return(yamlutil.String("hello"), yamlutil.String("world")),
			wantErr: expr.ErrEval,
		}, {
			name:    "subexpr returns single non-scalar value",
			expr:    exprtest.Return(yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)),
			wantErr: expr.ErrEval,
		}, {
			name:    "subexpr returns scalar non-numeric value",
			expr:    exprtest.Return(yamlutil.String("hello")),
			wantErr: expr.ErrEval,
		}, {
			name: "subexpr returns scalar numeric value",
			expr: exprtest.Return(yamlutil.Number("42")),
			want: []*yaml.Node{yamlutil.Number("42")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.PrefixPlusExpr{Expr: tc.expr}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("PrefixPlusExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("PrefixPlusExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
