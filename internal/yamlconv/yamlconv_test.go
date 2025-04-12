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
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
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

func TestInt(t *testing.T) {
	want := yamltest.MustParseNode(`42`)

	got := yamlconv.Int(42)

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("Int() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
	}
}

func TestFloat(t *testing.T) {
	want := yamltest.MustParseNode(`42.3`)

	got := yamlconv.Float(42.3)

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("Float() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
			want: yamltest.MustParseNode(`42`),
		}, {
			name: "uint",
			call: toNumber(uint(42)),
			want: yamltest.MustParseNode(`42`),
		}, {
			name: "float64",
			call: toNumber(42.3),
			want: yamltest.MustParseNode(`42.3`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call()

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Number() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
			}
		})
	}
}

func TestString(t *testing.T) {
	want := yamltest.MustParseNode(`"foo"`)

	got := yamlconv.String("foo")

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("String() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
	}
}

func TestBool(t *testing.T) {
	want := yamltest.MustParseNode(`true`)

	got := yamlconv.Bool(true)

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("Bool() = (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
	}
}

func TestNull(t *testing.T) {
	want := yamltest.MustParseNode(`null`)

	got := yamlconv.Null()

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("Null() = (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
			want: yamltest.MustParseNode(`42`),
		}, {
			name: "uint",
			call: toNode(uint(42)),
			want: yamltest.MustParseNode(`42`),
		}, {
			name: "float64",
			call: toNode(42.3),
			want: yamltest.MustParseNode(`42.3`),
		}, {
			name: "string",
			call: toNode("foo"),
			want: yamltest.MustParseNode(`"foo"`),
		}, {
			name: "bool",
			call: toNode(true),
			want: yamltest.MustParseNode(`true`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call()

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Node() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
			}
		})
	}
}

func TestIntString(t *testing.T) {
	want := yamltest.MustParseNode(`42`)

	got := yamlconv.IntString("42")

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("IntString() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
	}
}

func TestFloatString(t *testing.T) {
	want := yamltest.MustParseNode(`42.3`)

	got := yamlconv.FloatString("42.3")

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("FloatString() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
	}
}

func TestNumberString(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  *yaml.Node
	}{
		{
			name:  "int",
			input: "42",
			want:  yamltest.MustParseNode(`42`),
		}, {
			name:  "float",
			input: "42.3",
			want:  yamltest.MustParseNode(`42.3`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.NumberString(tc.input)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("NumberString() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
			}
		})
	}
}

func TestBoolString(t *testing.T) {
	want := yamltest.MustParseNode(`true`)

	got := yamlconv.BoolString("true")

	if got, want := got, want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
		t.Errorf("BoolString() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`42`),
			},
		}, {
			name:  "multiple",
			input: []int{42, 43},
			want: []*yaml.Node{
				yamltest.MustParseNode(`42`),
				yamltest.MustParseNode(`43`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Ints(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Ints() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`42.3`),
			},
		}, {
			name:  "multiple",
			input: []float64{42.3, 43.4},
			want: []*yaml.Node{
				yamltest.MustParseNode(`42.3`),
				yamltest.MustParseNode(`43.4`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Floats(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Floats() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`42`),
			},
		}, {
			name:  "multiple",
			input: []int{42, 44},
			want: []*yaml.Node{
				yamltest.MustParseNode(`42`),
				yamltest.MustParseNode(`44`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Numbers(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Numbers() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`"foo"`),
			},
		}, {
			name:  "multiple",
			input: []string{"foo", "bar"},
			want: []*yaml.Node{
				yamltest.MustParseNode(`"foo"`),
				yamltest.MustParseNode(`"bar"`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Strings(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Strings() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`true`),
			},
		}, {
			name:  "multiple",
			input: []bool{true, false},
			want: []*yaml.Node{
				yamltest.MustParseNode(`true`),
				yamltest.MustParseNode(`false`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Bools(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Bools() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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
				yamltest.MustParseNode(`42`),
			},
		}, {
			name:  "multiple",
			input: []int{42, 50},
			want: []*yaml.Node{
				yamltest.MustParseNode(`42`),
				yamltest.MustParseNode(`50`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlconv.Nodes(tc.input...)

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Nodes() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Sequence() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
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

			if got, want := got, tc.want; !cmp.Equal(got, want, yamltest.IgnoreMetaFields()) {
				t.Errorf("Document() = diff (-got,+want):\n%s", cmp.Diff(got, want, yamltest.IgnoreMetaFields()))
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    int64
		wantErr error
	}{
		{
			name:    "empty",
			input:   []*yaml.Node{},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "more than one",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "single value, valid int",
			input: yamlconv.Ints(42),
			want:  42,
		}, {
			name:    "single value, wrong kind",
			input:   []*yaml.Node{yamlconv.Sequence()},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "single value, wrong tag",
			input:   []*yaml.Node{yamlconv.String("foo")},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "single value, invalid int",
			input:   []*yaml.Node{yamlconv.IntString("foo")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseInt(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseInt() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseInt() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    float64
		wantErr error
	}{
		{
			name:    "empty",
			input:   []*yaml.Node{},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "more than one",
			input:   yamlconv.Floats(1.1, 2.2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "single value, valid float",
			input: yamlconv.Floats(42.3),
			want:  42.3,
		}, {
			name:    "single value, wrong kind",
			input:   []*yaml.Node{yamlconv.Sequence()},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "single value, wrong tag",
			input:   []*yaml.Node{yamlconv.String("foo")},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "single value, invalid float",
			input:   []*yaml.Node{yamlconv.FloatString("foo")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseFloat(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseFloat() error = %v, want %v", err, tc.wantErr)
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
		input   []*yaml.Node
		want    decimal.Decimal
		wantErr error
	}{
		{
			name:    "empty",
			input:   []*yaml.Node{},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "more than one",
			input:   yamlconv.Numbers(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "single value, valid decimal",
			input: yamlconv.Numbers(42),
			want:  decimal.New(42, 0),
		}, {
			name:    "single value, wrong kind",
			input:   []*yaml.Node{yamlconv.Sequence()},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "single value, wrong tag",
			input:   []*yaml.Node{yamlconv.String("foo")},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseDecimal(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseDecimal() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseDecimal() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    string
		wantErr error
	}{
		{
			name:    "empty",
			input:   []*yaml.Node{},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "more than one",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "single value, valid string",
			input: yamlconv.Strings("hello"),
			want:  "hello",
		}, {
			name:    "single value, wrong kind",
			input:   []*yaml.Node{yamlconv.Sequence()},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "single value, wrong tag",
			input:   []*yaml.Node{yamlconv.Int(42)},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseString(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseString() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    bool
		wantErr error
	}{
		{
			name:    "empty",
			input:   []*yaml.Node{},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "more than one",
			input:   yamlconv.Bools(true, false),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "single value, valid bool",
			input: yamlconv.Bools(true),
			want:  true,
		}, {
			name:    "single value, wrong kind",
			input:   []*yaml.Node{yamlconv.Sequence()},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "single value, wrong tag",
			input:   []*yaml.Node{yamlconv.Int(42)},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlconv.ParseBool(tc.input...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseBool() error = %v, want %v", err, tc.wantErr)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseBool() = %v, want %v", got, tc.want)
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
