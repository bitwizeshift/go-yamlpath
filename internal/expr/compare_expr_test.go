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

func TestCompareExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		compare     expr.Comparator
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:    "Left expression returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right expression returns error",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name:    "operator '<' where left is less than right",
			compare: expr.CompareLess,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("world")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '<' where left is greater than right",
			compare: expr.CompareLess,
			left:    ExprReturnsNodes(yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.False},
		}, {
			name:    "operator '<=' where left is less than right",
			compare: expr.CompareLessEqual,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("world")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '<=' where left is equal to right",
			compare: expr.CompareLessEqual,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '<=' where left is greater than right",
			compare: expr.CompareLessEqual,
			left:    ExprReturnsNodes(yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.False},
		}, {
			name:    "operator '>' where left is greater than right",
			compare: expr.CompareGreater,
			left:    ExprReturnsNodes(yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '>' where left is less than right",
			compare: expr.CompareGreater,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("world")),
			want:    []*yaml.Node{yamlutil.False},
		}, {
			name:    "operator '>=' where left is greater than right",
			compare: expr.CompareGreaterEqual,
			left:    ExprReturnsNodes(yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '>=' where left is equal to right",
			compare: expr.CompareGreaterEqual,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			want:    []*yaml.Node{yamlutil.True},
		}, {
			name:    "operator '>=' where left is less than right",
			compare: expr.CompareGreaterEqual,
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("world")),
			want:    []*yaml.Node{yamlutil.False},
		}, {
			name:    "Left and right are incomparable",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(YAML(t, `{"foo": "bar"}`)),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			compare := tc.compare
			if compare == nil {
				compare = expr.CompareLess
			}
			sut := &expr.CompareExpr{
				Left:    tc.left,
				Right:   tc.right,
				Compare: compare,
			}

			got, err := sut.Eval(nil, nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("CompareExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("CompareExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
