package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestConcatExpr(t *testing.T) {
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
			name:    "Left returns multiple elements",
			left:    exprtest.Return(yamlconv.String("hello"), yamlconv.String("world")),
			right:   exprtest.Return(yamlconv.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns multiple elements",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.String("foo"), yamlconv.String("bar")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlconv.String("foo")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(yamlconv.String("foo")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Left value is single but not scalar",
			left:    exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			right:   exprtest.Return(yamlconv.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right value is single but not scalar",
			left:    exprtest.Return(yamlconv.String("foo")),
			right:   exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrEval,
		}, {
			name:  "Left value is scalar int, right is scalar int",
			left:  exprtest.Return(yamlconv.Number(42)),
			right: exprtest.Return(yamlconv.Number(42)),
			want:  []*yaml.Node{yamlconv.Number(84)},
		}, {
			name:  "Left value is scalar int, right is scalar float",
			left:  exprtest.Return(yamlconv.Number(42)),
			right: exprtest.Return(yamlconv.Number(42.0)),
			want:  []*yaml.Node{yamlconv.Number(84)},
		}, {
			name:  "Left value is scalar float, right is scalar int",
			left:  exprtest.Return(yamlconv.Number(42.0)),
			right: exprtest.Return(yamlconv.Number(42)),
			want:  []*yaml.Node{yamlconv.Number(84)},
		}, {
			name:  "Left value is scalar float, right is scalar float",
			left:  exprtest.Return(yamlconv.Number(42.0)),
			right: exprtest.Return(yamlconv.Number(42.0)),
			want:  []*yaml.Node{yamlconv.Number(84)},
		}, {
			name:    "Left and right are scalar ints, left has bad representation",
			left:    exprtest.Return(yamlconv.NumberString("hello")),
			right:   exprtest.Return(yamlconv.Number(42)),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Left and right are scalar ints, right has bad representation",
			left:    exprtest.Return(yamlconv.Number(42)),
			right:   exprtest.Return(yamlconv.NumberString("hello")),
			wantErr: cmpopts.AnyError,
		}, {
			name:  "left and right are scalar string types",
			left:  exprtest.Return(yamlconv.String("hello")),
			right: exprtest.Return(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.String("helloworld")},
		}, {
			name:    "left and right are incompatible types",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Return(yamlconv.Number(42)),
			wantErr: errs.ErrEval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.ConcatExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConcatExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ConcatExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
