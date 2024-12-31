package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
			left:    exprtest.Return(yamlutil.String("hello"), yamlutil.String("world")),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right returns multiple elements",
			left:    exprtest.Return(yamlutil.String("hello")),
			right:   exprtest.Return(yamlutil.String("foo"), yamlutil.String("bar")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(yamlutil.String("foo")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Left value is single but not scalar",
			left:    exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right value is single but not scalar",
			left:    exprtest.Return(yamlutil.String("foo")),
			right:   exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: expr.ErrEval,
		}, {
			name:  "Left value is scalar int, right is scalar int",
			left:  exprtest.Return(yamlutil.Number("42")),
			right: exprtest.Return(yamlutil.Number("42")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar int, right is scalar float",
			left:  exprtest.Return(yamlutil.Number("42")),
			right: exprtest.Return(yamlutil.Number("42.0")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar float, right is scalar int",
			left:  exprtest.Return(yamlutil.Number("42.0")),
			right: exprtest.Return(yamlutil.Number("42")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar float, right is scalar float",
			left:  exprtest.Return(yamlutil.Number("42.0")),
			right: exprtest.Return(yamlutil.Number("42.0")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:    "Left and right are scalar ints, left has bad representation",
			left:    exprtest.Return(yamlutil.Number("hello")),
			right:   exprtest.Return(yamlutil.Number("42")),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Left and right are scalar ints, right has bad representation",
			left:    exprtest.Return(yamlutil.Number("42")),
			right:   exprtest.Return(yamlutil.Number("hello")),
			wantErr: cmpopts.AnyError,
		}, {
			name:  "left and right are scalar string types",
			left:  exprtest.Return(yamlutil.String("hello")),
			right: exprtest.Return(yamlutil.String("world")),
			want:  []*yaml.Node{yamlutil.String("helloworld")},
		}, {
			name:    "left and right are incompatible types",
			left:    exprtest.Return(yamlutil.String("hello")),
			right:   exprtest.Return(yamlutil.Number("42")),
			wantErr: expr.ErrEval,
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
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("ConcatExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
