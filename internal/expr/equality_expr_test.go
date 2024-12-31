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

func TestEqualityExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:  "empty left and right expression returns true",
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "left and right are the same",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(yamlutil.String("hello")),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "left and right are different",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(yamlutil.String("world")),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:    "left expression returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "right expression returns error",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.EqualityExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(context.Background(), nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("EqualityExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("EqualityExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
