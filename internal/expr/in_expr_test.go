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

func TestInExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:  "empty left and right expression returns empty expression",
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{},
		}, {
			name:  "empty left expression returns empty expression",
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(yamlutil.String("hello")),
			want:  []*yaml.Node{},
		}, {
			name:  "empty right expression returns empty expression",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{},
		}, {
			name:    "left expression returns multiple nodes",
			left:    ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: cmpopts.AnyError,
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
		}, {
			name:  "right returns list containing left element",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(YAML(t, `["hello", "world"]`)),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "right returns list not containing left element",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(YAML(t, `["world", "foo"]`)),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "right returns non-list node that matches",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(yamlutil.String("hello")),
			want:  []*yaml.Node{yamlutil.True},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.InExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(context.Background(), nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("InExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("InExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
