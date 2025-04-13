package invocation

import (
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// ParseString parses the given parameter as a string.
func ParseString(ctx Context, param Parameter) (string, error) {
	return parseScalar(yamlconv.ParseString, ctx, param)
}

// ParseInt parses the given parameter as an integer.
func ParseInt(ctx Context, param Parameter) (int64, error) {
	return parseScalar(yamlconv.ParseInt, ctx, param)
}

// ParseFloat parses the given parameter as a float.
func ParseFloat(ctx Context, param Parameter) (float64, error) {
	return parseScalar(yamlconv.ParseFloat, ctx, param)
}

// ParseDecimal parses the given parameter as a decimal.
func ParseDecimal(ctx Context, param Parameter) (decimal.Decimal, error) {
	return parseScalar(yamlconv.ParseDecimal, ctx, param)
}

// ParseBool parses the given parameter as a boolean.
func ParseBool(ctx Context, param Parameter) (bool, error) {
	return parseScalar(yamlconv.ParseBool, ctx, param)
}

func parseScalar[T any](fn func(...*yaml.Node) (T, error), ctx Context, param Parameter) (T, error) {
	nodes, err := param.GetArg(ctx)
	if err != nil {
		var zero T
		return zero, err
	}
	return fn(nodes...)
}
