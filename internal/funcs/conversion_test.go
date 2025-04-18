package funcs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestToBoolean(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "collection containing multiple values returns singleton error",
			input:   []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(true)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "boolean returns boolean",
			input: []*yaml.Node{yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "string can convert to boolean true",
			input: []*yaml.Node{yamlconv.String("true")},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "string can convert to boolean false",
			input: []*yaml.Node{yamlconv.String("false")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "string cannot convert to boolean",
			input: []*yaml.Node{yamlconv.String("whatever")},
			want:  []*yaml.Node{},
		}, {
			name:    "integer in bad representation",
			input:   []*yaml.Node{yamlconv.NumberString("whatever")},
			wantErr: cmpopts.AnyError,
		}, {
			name:  "integer can convert to boolean true",
			input: []*yaml.Node{yamlconv.Number(1)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "integer can convert to boolean false",
			input: []*yaml.Node{yamlconv.Number(0)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToBoolean(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToBoolean() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ToBoolean() = %v, want %v", got, want)
			}
		})
	}
}

func TestConvertsToBoolean(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name: "Non-scalar node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			want: yamlconv.Bools(false),
		}, {
			name:  "Integer with non 1/0 value",
			input: yamlconv.Ints(42),
			want:  yamlconv.Bools(false),
		}, {
			name:  "Integer with 1 value",
			input: yamlconv.Ints(1),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Integer with 0 value",
			input: yamlconv.Ints(0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float with non 1/0 value",
			input: yamlconv.Floats(42.0),
			want:  yamlconv.Bools(false),
		}, {
			name:  "Float with 1 value",
			input: yamlconv.Floats(1.0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float with 0 value",
			input: yamlconv.Floats(0.0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Boolean",
			input: yamlconv.Bools(true),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with 1 value",
			input: yamlconv.Strings("1"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with 0 value",
			input: yamlconv.Strings("0"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with non 1/0 value",
			input: yamlconv.Strings("42"),
			want:  yamlconv.Bools(false),
		}, {
			name:  "String with true value",
			input: yamlconv.Strings("true"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with false value",
			input: yamlconv.Strings("false"),
			want:  yamlconv.Bools(true),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ConvertsToBoolean(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConvertsToBoolean() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ConvertsToBoolean() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "collection containing multiple values returns singleton error",
			input:   []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(true)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "boolean false returns string 'false'",
			input: []*yaml.Node{yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.String("false")},
		}, {
			name:  "boolean false returns string 'true'",
			input: []*yaml.Node{yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.String("true")},
		}, {
			name:  "number returns numeric string",
			input: []*yaml.Node{yamlconv.Number(12345)},
			want:  []*yaml.Node{yamlconv.String("12345")},
		}, {
			name:  "complex objects return empty",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToString(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToString() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ToString() = %v, want %v", got, want)
			}
		})
	}
}

func TestConvertsToString(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Boolean",
			input: yamlconv.Bools(true),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float",
			input: yamlconv.Floats(42.0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String",
			input: yamlconv.Strings("42"),
			want:  yamlconv.Bools(true),
		}, {
			name: "Mapping node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			want: yamlconv.Bools(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ConvertsToString(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConvertsToString() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ConvertsToString() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestToNumber(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "collection containing multiple values returns singleton error",
			input:   []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(true)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "boolean false returns '0'",
			input: []*yaml.Node{yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Number(0)},
		}, {
			name:  "boolean false returns '1'",
			input: []*yaml.Node{yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Number(1)},
		}, {
			name:  "bad boolean value returns empty",
			input: []*yaml.Node{yamlconv.BoolString("whatever")},
			want:  []*yaml.Node{},
		}, {
			name:  "number returns input",
			input: []*yaml.Node{yamlconv.Number(12345)},
			want:  []*yaml.Node{yamlconv.Number(12345)},
		}, {
			name:  "complex objects return empty",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			want:  []*yaml.Node{},
		}, {
			name:  "string parses as number",
			input: []*yaml.Node{yamlconv.String("12345.67")},
			want:  []*yaml.Node{yamlconv.Number(12345.67)},
		}, {
			name:  "string does not parse as number returns empty",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToNumber(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToNumber() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ToNumber() = %v, want %v", got, want)
			}
		})
	}
}

func TestConvertsToNumber(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Boolean",
			input: yamlconv.Bools(true),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float",
			input: yamlconv.Floats(42.0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with numeric value",
			input: yamlconv.Strings("42"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with non-numeric value",
			input: yamlconv.Strings("foo"),
			want:  yamlconv.Bools(false),
		}, {
			name: "Mapping node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			want: yamlconv.Bools(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ConvertsToNumber(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConvertsToNumber() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ConvertsToNumber() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestToInteger(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "collection containing multiple values returns singleton error",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "boolean false returns '0'",
			input: yamlconv.Bools(false),
			want:  yamlconv.Ints(0),
		}, {
			name:  "boolean true returns '1'",
			input: yamlconv.Bools(true),
			want:  yamlconv.Ints(1),
		}, {
			name:  "integer returns input",
			input: yamlconv.Ints(12345),
			want:  yamlconv.Ints(12345),
		}, {
			name:  "float returns empty",
			input: yamlconv.Floats(12345.67),
			want:  []*yaml.Node{},
		}, {
			name:  "string parses as decimal returns empty",
			input: yamlconv.Strings("12345.67"),
			want:  []*yaml.Node{},
		}, {
			name:  "string parses as integer returns integer",
			input: yamlconv.Strings("12345"),
			want:  yamlconv.Ints(12345),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToInteger(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToInteger() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ToInteger() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestConvertsToInteger(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Boolean",
			input: yamlconv.Bools(true),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float",
			input: yamlconv.Floats(42.0),
			want:  yamlconv.Bools(false),
		}, {
			name:  "String with numeric value",
			input: yamlconv.Strings("42"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with non-numeric value",
			input: yamlconv.Strings("foo"),
			want:  yamlconv.Bools(false),
		}, {
			name: "Mapping node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			want: yamlconv.Bools(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ConvertsToInteger(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConvertsToInteger() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ConvertsToInteger() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "collection containing multiple values returns singleton error",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "boolean false returns '0'",
			input: yamlconv.Bools(false),
			want:  yamlconv.Floats(0.0),
		}, {
			name:  "boolean true returns '1'",
			input: yamlconv.Bools(true),
			want:  yamlconv.Floats(1.0),
		}, {
			name:  "integer returns input",
			input: yamlconv.Ints(12345),
			want:  yamlconv.Floats(12345.0),
		}, {
			name:  "float returns input",
			input: yamlconv.Floats(12345.67),
			want:  yamlconv.Floats(12345.67),
		}, {
			name:  "string parses as decimal returns float",
			input: yamlconv.Strings("12345.67"),
			want:  yamlconv.Floats(12345.67),
		}, {
			name:  "string parses as integer returns float",
			input: yamlconv.Strings("12345"),
			want:  yamlconv.Floats(12345.0),
		}, {
			name:  "invalid int returns empty",
			input: []*yaml.Node{yamlconv.IntString("bad")},
			want:  []*yaml.Node{},
		}, {
			name:  "string with invalid decimal returns empty",
			input: yamlconv.Strings("bad"),
			want:  []*yaml.Node{},
		}, {
			name:  "complex objects return empty",
			input: []*yaml.Node{yamltest.MustParseNode(`{}`)},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToFloat(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToFloat() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ToFloat() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestConvertsToFloat(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Boolean",
			input: yamlconv.Bools(true),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Bools(true),
		}, {
			name:  "Float",
			input: yamlconv.Floats(42.0),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with numeric value",
			input: yamlconv.Strings("42"),
			want:  yamlconv.Bools(true),
		}, {
			name:  "String with non-numeric value",
			input: yamlconv.Strings("foo"),
			want:  yamlconv.Bools(false),
		}, {
			name: "Mapping node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			want: yamlconv.Bools(false),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ConvertsToFloat(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ConvertsToFloat() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ConvertsToFloat() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestToSequence(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "collection containing multiple values returns list",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(true)},
			want: []*yaml.Node{{
				Kind: yaml.SequenceNode,
				Tag:  "!!seq",
				Content: []*yaml.Node{
					yamlconv.Bool(false),
					yamlconv.String("true"),
					yamlconv.Bool(true),
				},
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToSequence(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToSequence() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("ToSequence() = %v, want %v", got, want)
			}
		})
	}
}
