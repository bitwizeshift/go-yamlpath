package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
			expr: ExprReturnsNodes(),
			want: []*yaml.Node{yamlutil.True},
		}, {
			name:    "Subexpr evaluates error",
			expr:    ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name: "Subexpr returns scalar non-bool value",
			expr: ExprReturnsNodes(yamlutil.String("hello")),
			want: []*yaml.Node{yamlutil.False},
		}, {
			name: "Subexpr returns scalar bool value",
			expr: ExprReturnsNodes(yamlutil.True),
			want: []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.NegationExpr{Expr: tc.expr}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("NegationExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("NegationExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
