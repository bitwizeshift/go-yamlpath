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
			left:  exprtest.Return(),
			right: exprtest.Return(),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "left and right are the same",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamlconv.String("hello")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "left and right are different",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "left expression returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlconv.String("hello")),
			wantErr: testErr,
		}, {
			name:    "right expression returns error",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.EqualityExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("EqualityExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("EqualityExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
