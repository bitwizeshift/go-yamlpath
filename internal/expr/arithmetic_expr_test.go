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
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestArithmeticExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name        string
		operation   expr.ArithmeticOp
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
			name:    "Left returns error",
			left:    exprtest.Error(testErr),
			right:   exprtest.Return(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(yamlutil.String("hello")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Left returns multiple elements",
			left:    exprtest.Return(yamlutil.String("hello"), yamlutil.String("world")),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns multiple elements",
			left:    exprtest.Return(yamlutil.String("hello")),
			right:   exprtest.Return(yamlutil.String("foo"), yamlutil.String("bar")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left returns non-scalar value",
			left:    exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns non-scalar value",
			left:    exprtest.Return(yamlutil.String("foo")),
			right:   exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left returns scalar non-numeric value",
			left:    exprtest.Return(yamlutil.String("foo")),
			right:   exprtest.Return(yamlutil.Number("42")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns scalar non-numeric value",
			left:    exprtest.Return(yamlutil.Number("42")),
			right:   exprtest.Return(yamlutil.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left value is invalid number",
			left:    exprtest.Return(yamlutil.Number("foo")),
			right:   exprtest.Return(yamlutil.Number("42")),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Right value is invalid number",
			left:    exprtest.Return(yamlutil.Number("42")),
			right:   exprtest.Return(yamlutil.Number("foo")),
			wantErr: cmpopts.AnyError,
		}, {
			name:      "operator '-' where left and right are ints",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlutil.Number("42")),
			right:     exprtest.Return(yamlutil.Number("42")),
			want:      []*yaml.Node{yamlutil.Number("0")},
		}, {
			name:      "operator '-' where left and right are floats",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlutil.Number("42.69")),
			right:     exprtest.Return(yamlutil.Number("42.69")),
			want:      []*yaml.Node{yamlutil.Number("0.00")},
		}, {
			name:      "operator '-' where left is int and right is float",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlutil.Number("42")),
			right:     exprtest.Return(yamlutil.Number("42.69")),
			want:      []*yaml.Node{yamlutil.Number("-0.69")},
		}, {
			name:      "operator '-' where left is float and right is int",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlutil.Number("42.69")),
			right:     exprtest.Return(yamlutil.Number("42")),
			want:      []*yaml.Node{yamlutil.Number("0.69")},
		}, {
			name:      "operator '*' where left and right are ints",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlutil.Number("5")),
			right:     exprtest.Return(yamlutil.Number("5")),
			want:      []*yaml.Node{yamlutil.Number("25")},
		}, {
			name:      "operator '*' where left and right are floats",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlutil.Number("5.5")),
			right:     exprtest.Return(yamlutil.Number("5.5")),
			want:      []*yaml.Node{yamlutil.Number("30.25")},
		}, {
			name:      "operator '*' where left is int and right is float",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlutil.Number("5")),
			right:     exprtest.Return(yamlutil.Number("5.5")),
			want:      []*yaml.Node{yamlutil.Number("27.5")},
		}, {
			name:      "operator '*' where left is float and right is int",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlutil.Number("5.5")),
			right:     exprtest.Return(yamlutil.Number("5")),
			want:      []*yaml.Node{yamlutil.Number("27.5")},
		}, {
			name:      "operator '/' where left and right are ints",
			operation: expr.Division,
			left:      exprtest.Return(yamlutil.Number("10")),
			right:     exprtest.Return(yamlutil.Number("2")),
			want:      []*yaml.Node{yamlutil.Number("5")},
		}, {
			name:      "operator '/' where left and right are floats",
			operation: expr.Division,
			left:      exprtest.Return(yamlutil.Number("10.5")),
			right:     exprtest.Return(yamlutil.Number("2.5")),
			want:      []*yaml.Node{yamlutil.Number("4.2")},
		}, {
			name:      "operator '/' where left is int and right is float",
			operation: expr.Division,
			left:      exprtest.Return(yamlutil.Number("10")),
			right:     exprtest.Return(yamlutil.Number("2.5")),
			want:      []*yaml.Node{yamlutil.Number("4")},
		}, {
			name:      "operator '/' where left is float and right is int",
			operation: expr.Division,
			left:      exprtest.Return(yamlutil.Number("10.5")),
			right:     exprtest.Return(yamlutil.Number("2")),
			want:      []*yaml.Node{yamlutil.Number("5.25")},
		}, {
			name:      "operator '%' where left and right are ints",
			operation: expr.Modulus,
			left:      exprtest.Return(yamlutil.Number("10")),
			right:     exprtest.Return(yamlutil.Number("3")),
			want:      []*yaml.Node{yamlutil.Number("1")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			operation := tc.operation
			if operation == nil {
				operation = expr.Subtraction
			}
			sut := &expr.ArithmeticExpr{
				Left:      tc.left,
				Right:     tc.right,
				Operation: operation,
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ArithmeticExpr.Eval() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("ArithmeticExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
