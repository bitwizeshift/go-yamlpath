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
			left:  exprtest.Return(),
			right: exprtest.Return(),
			want:  []*yaml.Node{},
		}, {
			name:  "empty left expression returns empty expression",
			left:  exprtest.Return(),
			right: exprtest.Return(yamlconv.String("hello")),
			want:  []*yaml.Node{},
		}, {
			name:  "empty right expression returns empty expression",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(),
			want:  []*yaml.Node{},
		}, {
			name:    "left expression returns multiple nodes",
			left:    exprtest.Return(yamlconv.String("hello"), yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("hello")),
			wantErr: cmpopts.AnyError,
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
		}, {
			name:  "right returns list containing left element",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamltest.MustParseNode(`["hello", "world"]`)),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "right returns list not containing left element",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamltest.MustParseNode(`["world", "foo"]`)),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "right returns non-list node that matches",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamlconv.String("hello")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.InExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("InExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("InExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
