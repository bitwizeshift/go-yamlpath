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
			left:  ExprReturnsNodes(),
			right: ExprReturnsNodes(),
			want:  []*yaml.Node{},
		}, {
			name:    "Left returns error",
			left:    ExprReturnsError(testErr),
			right:   ExprReturnsNodes(yamlutil.String("hello")),
			wantErr: testErr,
		}, {
			name:    "Right returns error",
			left:    ExprReturnsNodes(yamlutil.String("hello")),
			right:   ExprReturnsError(testErr),
			wantErr: testErr,
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
			name:    "Left returns non-scalar value",
			left:    ExprReturnsNodes(YAML(t, `{"foo": "bar"}`)),
			right:   ExprReturnsNodes(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right returns non-scalar value",
			left:    ExprReturnsNodes(yamlutil.String("foo")),
			right:   ExprReturnsNodes(YAML(t, `{"foo": "bar"}`)),
			wantErr: expr.ErrEval,
		}, {
			name:    "Left returns scalar non-numeric value",
			left:    ExprReturnsNodes(yamlutil.String("foo")),
			right:   ExprReturnsNodes(yamlutil.Number("42")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Right returns scalar non-numeric value",
			left:    ExprReturnsNodes(yamlutil.Number("42")),
			right:   ExprReturnsNodes(yamlutil.String("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Left value is invalid number",
			left:    ExprReturnsNodes(yamlutil.Number("foo")),
			right:   ExprReturnsNodes(yamlutil.Number("42")),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Right value is invalid number",
			left:    ExprReturnsNodes(yamlutil.Number("42")),
			right:   ExprReturnsNodes(yamlutil.Number("foo")),
			wantErr: cmpopts.AnyError,
		}, {
			name:      "operator '-' where left and right are ints",
			operation: expr.Subtraction,
			left:      ExprReturnsNodes(yamlutil.Number("42")),
			right:     ExprReturnsNodes(yamlutil.Number("42")),
			want:      []*yaml.Node{yamlutil.Number("0")},
		}, {
			name:      "operator '-' where left and right are floats",
			operation: expr.Subtraction,
			left:      ExprReturnsNodes(yamlutil.Number("42.69")),
			right:     ExprReturnsNodes(yamlutil.Number("42.69")),
			want:      []*yaml.Node{yamlutil.Number("0.00")},
		}, {
			name:      "operator '-' where left is int and right is float",
			operation: expr.Subtraction,
			left:      ExprReturnsNodes(yamlutil.Number("42")),
			right:     ExprReturnsNodes(yamlutil.Number("42.69")),
			want:      []*yaml.Node{yamlutil.Number("-0.69")},
		}, {
			name:      "operator '-' where left is float and right is int",
			operation: expr.Subtraction,
			left:      ExprReturnsNodes(yamlutil.Number("42.69")),
			right:     ExprReturnsNodes(yamlutil.Number("42")),
			want:      []*yaml.Node{yamlutil.Number("0.69")},
		}, {
			name:      "operator '*' where left and right are ints",
			operation: expr.Multiplication,
			left:      ExprReturnsNodes(yamlutil.Number("5")),
			right:     ExprReturnsNodes(yamlutil.Number("5")),
			want:      []*yaml.Node{yamlutil.Number("25")},
		}, {
			name:      "operator '*' where left and right are floats",
			operation: expr.Multiplication,
			left:      ExprReturnsNodes(yamlutil.Number("5.5")),
			right:     ExprReturnsNodes(yamlutil.Number("5.5")),
			want:      []*yaml.Node{yamlutil.Number("30.25")},
		}, {
			name:      "operator '*' where left is int and right is float",
			operation: expr.Multiplication,
			left:      ExprReturnsNodes(yamlutil.Number("5")),
			right:     ExprReturnsNodes(yamlutil.Number("5.5")),
			want:      []*yaml.Node{yamlutil.Number("27.5")},
		}, {
			name:      "operator '*' where left is float and right is int",
			operation: expr.Multiplication,
			left:      ExprReturnsNodes(yamlutil.Number("5.5")),
			right:     ExprReturnsNodes(yamlutil.Number("5")),
			want:      []*yaml.Node{yamlutil.Number("27.5")},
		}, {
			name:      "operator '/' where left and right are ints",
			operation: expr.Division,
			left:      ExprReturnsNodes(yamlutil.Number("10")),
			right:     ExprReturnsNodes(yamlutil.Number("2")),
			want:      []*yaml.Node{yamlutil.Number("5")},
		}, {
			name:      "operator '/' where left and right are floats",
			operation: expr.Division,
			left:      ExprReturnsNodes(yamlutil.Number("10.5")),
			right:     ExprReturnsNodes(yamlutil.Number("2.5")),
			want:      []*yaml.Node{yamlutil.Number("4.2")},
		}, {
			name:      "operator '/' where left is int and right is float",
			operation: expr.Division,
			left:      ExprReturnsNodes(yamlutil.Number("10")),
			right:     ExprReturnsNodes(yamlutil.Number("2.5")),
			want:      []*yaml.Node{yamlutil.Number("4")},
		}, {
			name:      "operator '/' where left is float and right is int",
			operation: expr.Division,
			left:      ExprReturnsNodes(yamlutil.Number("10.5")),
			right:     ExprReturnsNodes(yamlutil.Number("2")),
			want:      []*yaml.Node{yamlutil.Number("5.25")},
		}, {
			name:      "operator '%' where left and right are ints",
			operation: expr.Modulus,
			left:      ExprReturnsNodes(yamlutil.Number("10")),
			right:     ExprReturnsNodes(yamlutil.Number("3")),
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
