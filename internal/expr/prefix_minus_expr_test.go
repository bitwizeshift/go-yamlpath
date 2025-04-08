package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
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
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name: "subexpr returns empty",
			expr: exprtest.Return(),
		}, {
			name:    "subexpr returns multiple values",
			expr:    exprtest.Return(yamlconv.String("hello"), yamlconv.String("world")),
			wantErr: errs.ErrEval,
		}, {
			name:    "subexpr returns single non-scalar value",
			expr:    exprtest.Return(yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)),
			wantErr: errs.ErrEval,
		}, {
			name:    "subexpr returns scalar non-numeric value",
			expr:    exprtest.Return(yamlconv.String("hello")),
			wantErr: errs.ErrEval,
		}, {
			name: "subexpr returns positive scalar integer value",
			expr: exprtest.Return(yamlconv.Number(42)),
			want: []*yaml.Node{yamlconv.Number(-42)},
		}, {
			name: "subexpr returns positive scalar float value",
			expr: exprtest.Return(yamlconv.Number(42.5)),
			want: []*yaml.Node{yamlconv.Number(-42.5)},
		}, {
			name: "subexpr returns negative scalar integer value",
			expr: exprtest.Return(yamlconv.Number(-42)),
			want: []*yaml.Node{yamlconv.Number(42)},
		}, {
			name: "subexpr returns negative scalar float value",
			expr: exprtest.Return(yamlconv.Number(-42.5)),
			want: []*yaml.Node{yamlconv.Number(42.5)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.PrefixMinusExpr{Expr: tc.expr}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("PrefixPlusExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("PrefixPlusExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
