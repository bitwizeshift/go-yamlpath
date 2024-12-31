package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestFilterExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input evalutes to empty output",
			expr:  exprtest.Return(),
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Subexpr returns error",
			expr:    exprtest.Error(testErr),
			input:   []*yaml.Node{yamlutil.String("hello")},
			wantErr: testErr,
		}, {
			name: "Includs only filtered values",
			expr: exprtest.Func(func(ctx invocation.Context) ([]*yaml.Node, error) {
				nodes := ctx.Current()
				if nodes[0].Value == "hello" {
					return []*yaml.Node{yamlutil.True}, nil
				}
				return []*yaml.Node{yamlutil.False}, nil
			}),
			input: []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			want:  []*yaml.Node{yamlutil.String("hello")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.FilterExpr{
				Expr: tc.expr,
			}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("FilterExpr.Eval() error = %v, want %v", got, want)
			}

			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("FilterExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
