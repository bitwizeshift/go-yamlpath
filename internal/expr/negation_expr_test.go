package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestNegationExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Empty node evaluates to true",
			expr: exprtest.Return(),
			want: []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "Subexpr evaluates error",
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name: "Subexpr returns scalar non-bool value",
			expr: exprtest.Return(yamlconv.String("hello")),
			want: []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Subexpr returns scalar bool value",
			expr: exprtest.Return(yamlconv.Bool(true)),
			want: []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.NegationExpr{Expr: tc.expr}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("NegationExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("NegationExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
