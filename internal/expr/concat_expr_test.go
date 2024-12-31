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
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{},
		}, {
			name:    "Left returns multiple elements",
			left:    ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			right:   ExprReturnsNodes(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right returns multiple elements",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.String("foo"), yamlutil.String("bar")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Left returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(yamlutil.String("foo")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    ExprReturnsNodes(yamlutil.String("foo")),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name:    "Left value is single but not scalar",
			left:    ExprReturnsNodes(YAML(t, `{"foo": "bar"}`)),
			right:   ExprReturnsNodes(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right value is single but not scalar",
			left:    ExprReturnsNodes(yamlutil.String("foo")),
			right:   ExprReturnsNodes(YAML(t, `{"foo": "bar"}`)),
			wantErr: expr.ErrEval,
		}, {
			name:  "Left value is scalar int, right is scalar int",
			left:  ExprReturnsNodes(yamlutil.Number("42")),
			right: ExprReturnsNodes(yamlutil.Number("42")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar int, right is scalar float",
			left:  ExprReturnsNodes(yamlutil.Number("42")),
			right: ExprReturnsNodes(yamlutil.Number("42.0")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar float, right is scalar int",
			left:  ExprReturnsNodes(yamlutil.Number("42.0")),
			right: ExprReturnsNodes(yamlutil.Number("42")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:  "Left value is scalar float, right is scalar float",
			left:  ExprReturnsNodes(yamlutil.Number("42.0")),
			right: ExprReturnsNodes(yamlutil.Number("42.0")),
			want:  []*yaml.Node{yamlutil.Number("84")},
		}, {
			name:    "Left and right are scalar ints, left has bad representation",
			left:    ExprReturnsNodes(yamlutil.Number("hello")),
			right:   ExprReturnsNodes(yamlutil.Number("42")),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Left and right are scalar ints, right has bad representation",
			left:    ExprReturnsNodes(yamlutil.Number("42")),
			right:   ExprReturnsNodes(yamlutil.Number("hello")),
			wantErr: cmpopts.AnyError,
		}, {
			name:  "left and right are scalar string types",
			left:  ExprReturnsNodes(yamlutil.String("hello")),
			right: ExprReturnsNodes(yamlutil.String("world")),
			want:  []*yaml.Node{yamlutil.String("helloworld")},
		}, {
			name:    "left and right are incompatible types",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsNodes(yamlutil.Number("42")),
			wantErr: expr.ErrEval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.ConcatExpr{
				Left:  tc.left,
				Right: tc.right,
			}

			got, err := sut.Eval(nil, nil)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConcatExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("ConcatExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
