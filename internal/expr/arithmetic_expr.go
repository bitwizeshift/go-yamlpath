package expr

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// ArithmeticOp is an arithmetic operation that takes two decimal.Decimal
// values.
type ArithmeticOp func(lhs, rhs decimal.Decimal) decimal.Decimal

var (
	// Addition is not modeled here, since the concat operator handles this.

	// Subtraction is the arithmetic subtraction operation for decimal objects.
	Subtraction ArithmeticOp = func(lhs, rhs decimal.Decimal) decimal.Decimal {
		return lhs.Sub(rhs)
	}
	// Multiplication is the arithmetic multiplication operation for decimal objects.
	Multiplication ArithmeticOp = func(lhs, rhs decimal.Decimal) decimal.Decimal {
		return lhs.Mul(rhs)
	}
	// Division is the arithmetic division operation for decimal objects.
	Division ArithmeticOp = func(lhs, rhs decimal.Decimal) decimal.Decimal {
		return lhs.Div(rhs)
	}
	// Modulus is the arithmetic modulus operation for decimal objects.
	Modulus ArithmeticOp = func(lhs, rhs decimal.Decimal) decimal.Decimal {
		return lhs.Mod(rhs)
	}
)

// ArithmeticExpr is a representation of the various arithmetic operations
// in the YAMLPath grammar, such as '-', '/', '*', and '%'. The only operator
// not included here is addition, since this is handled by the concat expression.
//
// These operations are only able to operate on numeric values, and will error
// for all other types.
type ArithmeticExpr struct {
	Left, Right Expr
	Operation   ArithmeticOp
}

// Eval evaluates the arithmetic expression against the given context.
func (e *ArithmeticExpr) Eval(ctx invocation.Context) ([]*yaml.Node, error) {
	left, right, err := e.eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(left) > 1 {
		return nil, errs.NewSingletonError("arithmetic op lhs", left)
	}
	if len(right) > 1 {
		return nil, errs.NewSingletonError("arithmetic op rhs", right)
	}
	if len(left) == 0 || len(right) == 0 {
		return nil, nil
	}
	lhs, rhs := left[0], right[0]
	if lhs.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("arithmetic op lhs", lhs, yaml.ScalarNode)
	}
	if rhs.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("arithmetic", rhs, yaml.ScalarNode)
	}
	if lhs.Tag != "!!int" && lhs.Tag != "!!float" {
		return nil, errs.NewTagError("arithmetic", lhs, "!!int", "!!float")
	}
	if rhs.Tag != "!!int" && rhs.Tag != "!!float" {
		return nil, errs.NewTagError("arithmetic", rhs, "!!int", "!!float")
	}

	lv, err := decimal.NewFromString(lhs.Value)
	if err != nil {
		return nil, err
	}
	rv, err := decimal.NewFromString(rhs.Value)
	if err != nil {
		return nil, err
	}
	result := e.Operation(lv, rv)

	return []*yaml.Node{yamlconv.NumberString(result.String())}, nil
}

func (e *ArithmeticExpr) eval(ctx invocation.Context) ([]*yaml.Node, []*yaml.Node, error) {
	left, err := e.Left.Eval(ctx)
	if err != nil {
		return nil, nil, err
	}
	right, err := e.Right.Eval(ctx)
	if err != nil {
		return nil, nil, err
	}
	return left, right, nil
}

var _ Expr = (*ArithmeticExpr)(nil)
