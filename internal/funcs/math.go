package funcs

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Abs computes the absolute value of the given number.
func Abs(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "abs()")
	}

	d = d.Abs()
	return []*yaml.Node{yamlconv.NumberString(d.String())}, nil
}

// Ceil computes the ceiling of the given number.
func Ceil(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "ceiling()")
	}

	d = d.Ceil()
	return []*yaml.Node{yamlconv.Int(d.IntPart())}, nil
}

// Floor computes the floor of the given number.
func Floor(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "floor()")
	}
	d = d.Floor()

	return []*yaml.Node{yamlconv.Int(d.IntPart())}, nil
}

// Exp calculates the natural exponent of decimal number.
func Exp(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "exp()")
	}

	digits := max(uint32(d.NumDigits()), 5)
	d, err = d.ExpHullAbrham(max(digits, 20))
	if err != nil {
		return nil, err
	}
	d = d.Truncate(int32(digits))

	return []*yaml.Node{yamlconv.NumberString(d.String())}, nil
}

// Ln calculates the natural logarithm of decimal number.
func Ln(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "ln()")
	}

	d, err = d.Ln(max(int32(d.NumDigits()), 5))
	if err != nil {
		return nil, err
	}

	return []*yaml.Node{yamlconv.NumberString(d.String())}, nil
}

// Log computes the natural logarithm of the given number.
func Log(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "log()")
	}

	digits := min(int32(d.NumDigits()), 5)
	d, err = d.Ln(digits)
	if err != nil {
		return nil, err
	}
	ln10 := decimal.NewFromFloat(2.302585092994046)
	d = d.Div(ln10).Truncate(digits)

	return []*yaml.Node{yamlconv.NumberString(d.String())}, nil
}

// Pow computes the power of the first argument raised to the second argument.
func Pow(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	base, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, err
	}
	exp, err := invocation.ParseDecimal(ctx, params[0])
	if err != nil {
		return nil, err
	}

	base = base.Pow(exp)

	return []*yaml.Node{yamlconv.NumberString(base.String())}, nil
}

// Round rounds the given number to the nearest integer.
func Round(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "round()")
	}

	precision := int64(0)
	if len(params) == 1 {
		precision, err = invocation.ParseInt(ctx, params[0])
		if err != nil {
			return nil, errs.IncludeSource(err, "round()")
		}
	}

	d = d.Round(int32(precision))

	return []*yaml.Node{yamlconv.NumberString(d.String())}, nil
}

// Truncate truncates the given number to the nearest integer.
func Truncate(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	d, err := yamlconv.ParseDecimal(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "truncate()")
	}

	i := d.Truncate(0).IntPart()

	return yamlconv.Ints(i), nil
}

// Min returns the minimum value from the current context and the provided
// parameters.
func Min(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	return minmax(ctx, params, func(lhs, rhs decimal.Decimal) bool {
		return lhs.LessThan(rhs)
	})
}

// Max returns the maximum value from the current context and the provided
// parameters.
func Max(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	return minmax(ctx, params, func(lhs, rhs decimal.Decimal) bool {
		return lhs.GreaterThan(rhs)
	})
}

func minmax(ctx invocation.Context, params []invocation.Parameter, compare func(lhs, rhs decimal.Decimal) bool) ([]*yaml.Node, error) {
	current := ctx.Current()

	var result *decimal.Decimal
	for _, node := range current {
		d, err := yamlconv.ParseDecimal(node)
		if err != nil {
			return nil, err
		}
		if result == nil || compare(d, *result) {
			result = &d
		}
	}

	for _, param := range params {
		d, err := invocation.ParseDecimal(ctx, param)
		if err != nil {
			return nil, err
		}
		if result == nil || compare(d, *result) {
			result = &d
		}
	}
	if result == nil {
		return nil, nil
	}

	return []*yaml.Node{yamlconv.NumberString(result.String())}, nil
}

// Sum computes the sum of the current context and the provided parameters.
func Sum(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()

	var total *decimal.Decimal
	for _, node := range current {
		d, err := yamlconv.ParseDecimal(node)
		if err != nil {
			return nil, err
		}
		if total == nil {
			z := decimal.Zero
			total = &z
		}
		*total = total.Add(d)
	}

	for _, param := range params {
		d, err := invocation.ParseDecimal(ctx, param)
		if err != nil {
			return nil, err
		}
		if total == nil {
			z := decimal.Zero
			total = &z
		}
		*total = total.Add(d)
	}

	if total == nil {
		return nil, nil
	}

	return []*yaml.Node{yamlconv.NumberString(total.String())}, nil
}
