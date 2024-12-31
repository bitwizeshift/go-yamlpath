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

func TestBinaryAndExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:    "Left returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    ExprReturnsNodes(yamlutil.True),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name:  "Short-circuits on left false",
			left:  ExprReturnsNodes(),
			right: ExprReturnsError(testErr),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Left and right evaluate to true",
			left:  ExprReturnsNodes(yamlutil.True),
			right: ExprReturnsNodes(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Left evaluates to true",
			left:  ExprReturnsNodes(yamlutil.True),
			right: ExprReturnsNodes(yamlutil.False),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Right evaluates to true",
			left:  ExprReturnsNodes(yamlutil.False),
			right: ExprReturnsNodes(yamlutil.True),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Left and right evaluate to false",
			left:  ExprReturnsNodes(yamlutil.False),
			right: ExprReturnsNodes(yamlutil.False),
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.BooleanAndExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("BooleanAndExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("BooleanAndExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
