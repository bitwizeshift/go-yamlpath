package yamlconv_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func toNumber[T constraints.Integer | constraints.Float](v T) func() *yaml.Node {
	return func() *yaml.Node {
		return yamlconv.Number(v)
	}
}

func toNode[T yamlconv.Primitive](v T) func() *yaml.Node {
	return func() *yaml.Node {
		return yamlconv.Node(v)
	}
}

func TestNumber(t *testing.T) {
	testCases := []struct {
		name string
		call func() *yaml.Node
		want *yaml.Node
	}{
		{
			name: "int",
			call: toNumber(42),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!int",
				Value: "42",
			},
		}, {
			name: "uint",
			call: toNumber(uint(42)),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!int",
				Value: "42",
			},
		}, {
			name: "float64",
			call: toNumber(42.3),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!float",
				Value: "42.3",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call()

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Number() = %v, want %v", got, want)
			}
		})
	}
}

func TestString(t *testing.T) {
	input := "foo"
	want := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!str",
		Value: input,
	}

	got := yamlconv.String("foo")

	if got, want := got, want; !cmp.Equal(got, want) {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	input := true
	want := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!bool",
		Value: "true",
	}

	got := yamlconv.Bool(input)

	if got, want := got, want; !cmp.Equal(got, want) {
		t.Errorf("Bool() = %v, want %v", got, want)
	}
}

func TestNode(t *testing.T) {
	testCases := []struct {
		name string
		call func() *yaml.Node
		want *yaml.Node
	}{
		{
			name: "int",
			call: toNode(42),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!int",
				Value: "42",
			},
		}, {
			name: "uint",
			call: toNode(uint(42)),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!int",
				Value: "42",
			},
		}, {
			name: "float64",
			call: toNode(42.3),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!float",
				Value: "42.3",
			},
		}, {
			name: "string",
			call: toNode("foo"),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!str",
				Value: "foo",
			},
		}, {
			name: "bool",
			call: toNode(true),
			want: &yaml.Node{
				Kind:  yaml.ScalarNode,
				Tag:   "!!bool",
				Value: "true",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call()

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Node() = %v, want %v", got, want)
			}
		})
	}
}

func TestInts(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []int{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []int{42},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				},
			},
		}, {
			name:  "multiple",
			input: []int{42, 43},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "43",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Ints(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Ints() = %v, want %v", got, want)
			}
		})
	}
}

func TestFloats(t *testing.T) {
	testCases := []struct {
		name  string
		input []float64
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []float64{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []float64{42.3},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!float",
					Value: "42.3",
				},
			},
		}, {
			name:  "multiple",
			input: []float64{42.3, 43.4},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!float",
					Value: "42.3",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!float",
					Value: "43.4",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Floats(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Floats() = %v, want %v", got, want)
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []int{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []int{42},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				},
			},
		}, {
			name:  "multiple",
			input: []int{42, 44},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "44",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Numbers(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Numbers() = %v, want %v", got, want)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []string{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []string{"foo"},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!str",
					Value: "foo",
				},
			},
		}, {
			name:  "multiple",
			input: []string{"foo", "bar"},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!str",
					Value: "foo",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!str",
					Value: "bar",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Strings(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Strings() = %v, want %v", got, want)
			}
		})
	}
}

func TestBools(t *testing.T) {
	testCases := []struct {
		name  string
		input []bool
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []bool{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []bool{true},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!bool",
					Value: "true",
				},
			},
		}, {
			name:  "multiple",
			input: []bool{true, false},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!bool",
					Value: "true",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!bool",
					Value: "false",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Bools(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Bools() = %v, want %v", got, want)
			}
		})
	}
}

func TestNodes(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []int{},
			want:  []*yaml.Node{},
		}, {
			name:  "single",
			input: []int{42},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				},
			},
		}, {
			name:  "multiple",
			input: []int{42, 50},
			want: []*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "42",
				}, {
					Kind:  yaml.ScalarNode,
					Tag:   "!!int",
					Value: "50",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Nodes(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Nodes() = %v, want %v", got, want)
			}
		})
	}
}

func TestSequence(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  *yaml.Node
	}{
		{
			name:  "empty",
			input: []*yaml.Node{},
			want: &yaml.Node{
				Kind:    yaml.SequenceNode,
				Tag:     "!!seq",
				Content: []*yaml.Node{},
			},
		}, {
			name:  "single",
			input: []*yaml.Node{yamlconv.Bool(false)},
			want: &yaml.Node{
				Kind:    yaml.SequenceNode,
				Tag:     "!!seq",
				Content: []*yaml.Node{yamlconv.Bool(false)},
			},
		}, {
			name:  "multiple",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(true)},
			want: &yaml.Node{
				Kind:    yaml.SequenceNode,
				Tag:     "!!seq",
				Content: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(true)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Sequence(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Sequence() = %v, want %v", got, want)
			}
		})
	}
}

func TestDocument(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  *yaml.Node
	}{
		{
			name:  "empty",
			input: []*yaml.Node{},
			want: &yaml.Node{
				Kind:    yaml.DocumentNode,
				Content: []*yaml.Node{},
			},
		}, {
			name:  "single",
			input: []*yaml.Node{yamlconv.Bool(false)},
			want: &yaml.Node{
				Kind:    yaml.DocumentNode,
				Content: []*yaml.Node{yamlconv.Bool(false)},
			},
		}, {
			name:  "multiple",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(true)},
			want: &yaml.Node{
				Kind:    yaml.DocumentNode,
				Content: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(true)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Document(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Document() = %v, want %v", got, want)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	testCases := []struct {
		name    string
		input   *yaml.Node
		want    string
		wantErr error
	}{
		{
			name:    "wrong kind",
			input:   yamlconv.Sequence(),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "wrong tag",
			input:   yamlconv.Int(4),
			wantErr: errs.ErrBadTag,
		}, {
			name:  "valid string",
			input: yamlconv.String("foo"),
			want:  "foo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseString(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseString() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	testCases := []struct {
		name    string
		input   *yaml.Node
		want    int64
		wantErr error
	}{
		{
			name:    "wrong kind",
			input:   yamlconv.Sequence(),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "wrong tag",
			input:   yamlconv.String("foo"),
			wantErr: errs.ErrBadTag,
		}, {
			name:  "valid int",
			input: yamlconv.Int(42),
			want:  42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseInt(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseInt() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseInt() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	testCases := []struct {
		name    string
		input   *yaml.Node
		want    bool
		wantErr error
	}{
		{
			name:    "wrong kind",
			input:   yamlconv.Sequence(),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "wrong tag",
			input:   yamlconv.Int(42),
			wantErr: errs.ErrBadTag,
		}, {
			name:  "valid bool",
			input: yamlconv.Bool(true),
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseBool(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseBool() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseBool() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	testCases := []struct {
		name    string
		input   *yaml.Node
		want    float64
		wantErr error
	}{
		{
			name:    "wrong kind",
			input:   yamlconv.Sequence(),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "wrong tag",
			input:   yamlconv.String("foo"),
			wantErr: errs.ErrBadTag,
		}, {
			name:  "valid float",
			input: yamlconv.Float(42.3),
			want:  42.3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseFloat(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseFloat() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseFloat() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseDecimal(t *testing.T) {
	testCases := []struct {
		name    string
		input   *yaml.Node
		want    decimal.Decimal
		wantErr error
	}{
		{
			name:    "wrong kind",
			input:   yamlconv.Sequence(),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "wrong tag",
			input:   yamlconv.String("foo"),
			wantErr: errs.ErrBadTag,
		}, {
			name:  "integer",
			input: yamlconv.Int(42),
			want:  decimal.New(42, 0),
		}, {
			name:  "float",
			input: yamlconv.Float(42.3),
			want:  decimal.NewFromFloat(42.3),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseDecimal(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseDecimal() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseDecimal() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFlattenDocuments(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  []*yaml.Node
	}{
		{
			name:  "empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name: "single document",
			input: []*yaml.Node{
				yamlconv.Document(yamlconv.Int(42), yamlconv.String("foo")),
			},
			want: []*yaml.Node{
				yamlconv.Int(42),
				yamlconv.String("foo"),
			},
		}, {
			name: "multiple documents",
			input: []*yaml.Node{
				yamlconv.Document(yamlconv.Int(42), yamlconv.String("foo")),
				yamlconv.Document(yamlconv.Bool(true), yamlconv.Float(42.3)),
			},
			want: []*yaml.Node{
				yamlconv.Int(42),
				yamlconv.String("foo"),
				yamlconv.Bool(true),
				yamlconv.Float(42.3),
			},
		}, {
			name: "not a document",
			input: []*yaml.Node{
				yamlconv.Int(42),
			},
			want: []*yaml.Node{
				yamlconv.Int(42),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.FlattenDocuments(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("FlattenDocuments() (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestIsTruthy(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  bool
	}{
		{
			name:  "Empty node is falsy",
			input: yamlconv.Nodes[bool](),
			want:  false,
		}, {
			name:  "Single non-bool node is truthy",
			input: yamlconv.Nodes("hello"),
			want:  true,
		}, {
			name:  "Single true node is truthy",
			input: yamlconv.Nodes(true),
			want:  true,
		}, {
			name:  "Single false node is falsy",
			input: yamlconv.Nodes(false),
			want:  false,
		}, {
			name:  "Multiple nodes are truthy",
			input: yamlconv.Nodes(true, false),
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.IsTruthy(tc.input...)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
