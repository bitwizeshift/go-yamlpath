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
			left:  exprtest.Return(),
			right: exprtest.Return(),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:  "Left is subset of right",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamlconv.String("hello"), yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Left is not subset of right",
			left:  exprtest.Return(yamlconv.String("hello"), yamlconv.String("world")),
			right: exprtest.Return(yamlconv.String("hello")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
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

			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Fatalf("SubsetofExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
