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
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
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
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlconv.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right expression returns error",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "operator '<' where left is less than right",
			compare: expr.CompareLess,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("world")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '<' where left is greater than right",
			compare: expr.CompareLess,
			left:    exprtest.Return(yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "operator '<=' where left is less than right",
			compare: expr.CompareLessEqual,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("world")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '<=' where left is equal to right",
			compare: expr.CompareLessEqual,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '<=' where left is greater than right",
			compare: expr.CompareLessEqual,
			left:    exprtest.Return(yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "operator '>' where left is greater than right",
			compare: expr.CompareGreater,
			left:    exprtest.Return(yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '>' where left is less than right",
			compare: expr.CompareGreater,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("world")),
			want:    []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "operator '>=' where left is greater than right",
			compare: expr.CompareGreaterEqual,
			left:    exprtest.Return(yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '>=' where left is equal to right",
			compare: expr.CompareGreaterEqual,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("hello")),
			want:    []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "operator '>=' where left is less than right",
			compare: expr.CompareGreaterEqual,
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("world")),
			want:    []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Left and right are incomparable",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
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

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("CompareExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("CompareExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
