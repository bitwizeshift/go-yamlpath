package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestBinaryOrExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		left, right expr.Expr
		want        []*yaml.Node
		wantErr     error
	}{
		{
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:  "Short-circuits on left true",
			left:  exprtest.Return(yamlutil.True),
			right: exprtest.Error(testErr),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Left and right evaluate to true",
			left:  exprtest.Return(yamlutil.True),
			right: exprtest.Return(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Left evaluates to true",
			left:  exprtest.Return(yamlutil.True),
			right: exprtest.Return(yamlutil.False),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Right evaluates to true",
			left:  exprtest.Return(yamlutil.False),
			right: exprtest.Return(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Left and right evaluate to false",
			left:  exprtest.Return(yamlutil.False),
			right: exprtest.Return(yamlutil.False),
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.BooleanOrExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("BooleanOrExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("BooleanOrExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
