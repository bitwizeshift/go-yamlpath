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

func TestSubsetOfExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		input       []*yaml.Node
		left, right expr.Expr
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:  "Empty ranges compare equal",
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:    "Left returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    ExprReturnsNodes(),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name:  "Left is subset of right",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Left is not subset of right",
			left:  ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			right: ExprReturnsNodes(yamlutil.String("hello")),
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.SubsetOfExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("SubsetofExpr.Eval() error = %v, want %v", got, want)
			}

			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Fatalf("SubsetofExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
