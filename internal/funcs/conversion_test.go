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
			input:   []*yaml.Node{yamlconv.RawNumber("whatever")},
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
			input: []*yaml.Node{yamlconv.RawBool("whatever")},
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
