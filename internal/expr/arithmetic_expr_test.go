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
			right:   exprtest.Return(yamlconv.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    exprtest.Return(yamlconv.String("hello")),
			right:   exprtest.Error(testErr),
			wantErr: testErr,
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
			name:    "Left returns non-scalar value",
			left:    exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			right:   exprtest.Return(yamlconv.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns non-scalar value",
			left:    exprtest.Return(yamlconv.String("foo")),
			right:   exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left returns scalar non-numeric value",
			left:    exprtest.Return(yamlconv.String("foo")),
			right:   exprtest.Return(yamlconv.Number(42)),
			wantErr: errs.ErrEval,
		}, {
			name:    "Right returns scalar non-numeric value",
			left:    exprtest.Return(yamlconv.Number(42)),
			right:   exprtest.Return(yamlconv.String("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Left value is invalid number",
			left:    exprtest.Return(yamlconv.RawNumber("foo")),
			right:   exprtest.Return(yamlconv.Number(42)),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Right value is invalid number",
			left:    exprtest.Return(yamlconv.Number(42)),
			right:   exprtest.Return(yamlconv.RawNumber("foo")),
			wantErr: cmpopts.AnyError,
		}, {
			name:      "operator '-' where left and right are ints",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlconv.Number(42)),
			right:     exprtest.Return(yamlconv.Number(42)),
			want:      []*yaml.Node{yamlconv.Number(0)},
		}, {
			name:      "operator '-' where left and right are floats",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlconv.Number(42.69)),
			right:     exprtest.Return(yamlconv.Number(42.69)),
			want:      []*yaml.Node{yamlconv.Number(0.00)},
		}, {
			name:      "operator '-' where left is int and right is float",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlconv.Number(42)),
			right:     exprtest.Return(yamlconv.Number(42.69)),
			want:      []*yaml.Node{yamlconv.Number(-0.69)},
		}, {
			name:      "operator '-' where left is float and right is int",
			operation: expr.Subtraction,
			left:      exprtest.Return(yamlconv.Number(42.69)),
			right:     exprtest.Return(yamlconv.Number(42)),
			want:      []*yaml.Node{yamlconv.Number(0.69)},
		}, {
			name:      "operator '*' where left and right are ints",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlconv.Number(5)),
			right:     exprtest.Return(yamlconv.Number(5)),
			want:      []*yaml.Node{yamlconv.Number(25)},
		}, {
			name:      "operator '*' where left and right are floats",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlconv.Number(5.5)),
			right:     exprtest.Return(yamlconv.Number(5.5)),
			want:      []*yaml.Node{yamlconv.Number(30.25)},
		}, {
			name:      "operator '*' where left is int and right is float",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlconv.Number(5)),
			right:     exprtest.Return(yamlconv.Number(5.5)),
			want:      []*yaml.Node{yamlconv.Number(27.5)},
		}, {
			name:      "operator '*' where left is float and right is int",
			operation: expr.Multiplication,
			left:      exprtest.Return(yamlconv.Number(5.5)),
			right:     exprtest.Return(yamlconv.Number(5)),
			want:      []*yaml.Node{yamlconv.Number(27.5)},
		}, {
			name:      "operator '/' where left and right are ints",
			operation: expr.Division,
			left:      exprtest.Return(yamlconv.Number(10)),
			right:     exprtest.Return(yamlconv.Number(2)),
			want:      []*yaml.Node{yamlconv.Number(5)},
		}, {
			name:      "operator '/' where left and right are floats",
			operation: expr.Division,
			left:      exprtest.Return(yamlconv.Number(10.5)),
			right:     exprtest.Return(yamlconv.Number(2.5)),
			want:      []*yaml.Node{yamlconv.Number(4.2)},
		}, {
			name:      "operator '/' where left is int and right is float",
			operation: expr.Division,
			left:      exprtest.Return(yamlconv.Number(10)),
			right:     exprtest.Return(yamlconv.Number(2.5)),
			want:      []*yaml.Node{yamlconv.Number(4)},
		}, {
			name:      "operator '/' where left is float and right is int",
			operation: expr.Division,
			left:      exprtest.Return(yamlconv.Number(10.5)),
			right:     exprtest.Return(yamlconv.Number(2)),
			want:      []*yaml.Node{yamlconv.Number(5.25)},
		}, {
			name:      "operator '%' where left and right are ints",
			operation: expr.Modulus,
			left:      exprtest.Return(yamlconv.Number(10)),
			right:     exprtest.Return(yamlconv.Number(3)),
			want:      []*yaml.Node{yamlconv.Number(1)},
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ArithmeticExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
