package yamlcmp_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestLessRange(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs []*yaml.Node
		want     bool
		wantErr  error
	}{
		{
			name: "Empty ranges compare equal",
			want: false,
		}, {
			name: "Left range is empty",
			rhs:  []*yaml.Node{yamlconv.Bool(true)},
			want: true,
		}, {
			name: "Right range is empty",
			lhs:  []*yaml.Node{yamlconv.Bool(true)},
			want: false,
		}, {
			name:    "Ranges have different kinds",
			lhs:     []*yaml.Node{yamlconv.Number(42)},
			rhs:     []*yaml.Node{yamlconv.String("42")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlcmp.LessRange(tc.lhs, tc.rhs)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("LessRange() = %v; want %v", got, tc.want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("LessRange() error = %v; want %v", err, tc.wantErr)
			}
		})
	}
}

func TestCompareRange(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs []*yaml.Node
		want     int
		wantErr  error
	}{
		{
			name: "Empty ranges compare equal",
			want: 0,
		}, {
			name: "Left range is empty",
			rhs:  []*yaml.Node{yamlconv.Bool(true)},
			want: -1,
		}, {
			name: "Right range is empty",
			lhs:  []*yaml.Node{yamlconv.Bool(true)},
			want: 1,
		}, {
			name: "Ranges are different length",
			lhs:  []*yaml.Node{yamlconv.Number(42)},
			rhs:  []*yaml.Node{yamlconv.Number(42), yamlconv.Number(43)},
			want: -1,
		}, {
			name: "Ranges have different values",
			lhs:  []*yaml.Node{yamlconv.Number(42), yamlconv.Number(43)},
			rhs:  []*yaml.Node{yamlconv.Number(42), yamlconv.Number(44)},
			want: -1,
		}, {
			name: "Ranges are equal",
			lhs:  []*yaml.Node{yamlconv.Number(42), yamlconv.Number(43)},
			rhs:  []*yaml.Node{yamlconv.Number(42), yamlconv.Number(43)},
		}, {
			name:    "Ranges have different kinds",
			lhs:     []*yaml.Node{yamlconv.Number(42)},
			rhs:     []*yaml.Node{yamlconv.String("42")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlcmp.CompareRange(tc.lhs, tc.rhs)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("CompareRange() = %v; want %v", got, tc.want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("CompareRange() error = %v; want %v", err, tc.wantErr)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs *yaml.Node
		want     int
		wantErr  error
	}{
		// Scalar Values
		{
			name: "Nil nodes compare equal",
			want: 0,
		}, {
			name: "Left node is nil",
			rhs:  yamlconv.Bool(true),
			want: -1,
		}, {
			name: "Right node is nil",
			lhs:  yamlconv.Bool(true),
			want: 1,
		}, {
			name: "Both nodes are identical scalar strings",
			lhs:  yamlconv.String("hello"),
			rhs:  yamlconv.String("hello"),
			want: 0,
		}, {
			name: "Both nodes are different scalar strings",
			lhs:  yamlconv.String("hello"),
			rhs:  yamlconv.String("world"),
			want: -1,
		}, {
			name: "Both nodes are equivalent scalar numbers",
			lhs:  yamlconv.Number(42e0),
			rhs:  yamlconv.Number(42),
			want: 0,
		}, {
			name: "Both nodes are different scalar numbers",
			lhs:  yamlconv.Number(42),
			rhs:  yamlconv.Number(43),
			want: -1,
		}, {
			name: "Both nodes are the same scalar booleans",
			lhs:  yamlconv.Bool(true),
			rhs:  yamlconv.Bool(true),
			want: 0,
		}, {
			name: "Left node is false, right node is true",
			lhs:  yamlconv.Bool(false),
			rhs:  yamlconv.Bool(true),
			want: -1,
		}, {
			name: "Left node is true, right node is false",
			lhs:  yamlconv.Bool(true),
			rhs:  yamlconv.Bool(false),
			want: 1,
		}, {
			name: "Null nodes compare equal",
			lhs:  yamlconv.Null(),
			rhs:  yamlconv.Null(),
			want: 0,
		},
		// Sequence Values
		{
			name: "Sequences are different length",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2]`),
			want: 1,
		}, {
			name: "Sequences have different values",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2, 4]`),
			want: -1,
		}, {
			name: "Sequences are equal",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			want: 0,
		},
		// Errors
		{
			name: "Node kinds are different",
			lhs:  yamlconv.Bool(true),
			rhs: &yaml.Node{
				Kind: yaml.SequenceNode,
			},
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar values with different tags",
			lhs:     yamlconv.Number(1234),
			rhs:     yamlconv.String("true"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar values with different non-numeric tags",
			lhs:     yamlconv.String("hello"),
			rhs:     yamlconv.Bool(true),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar numeric values, left node is invalid",
			lhs:     yamlconv.RawNumber("hello"),
			rhs:     yamlconv.Number(42),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar numeric values, right node is invalid",
			lhs:     yamlconv.Number(42),
			rhs:     yamlconv.RawNumber("world"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar boolean values, left node is invalid",
			lhs:     yamlconv.RawBool("hello"),
			rhs:     yamlconv.Bool(true),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar boolean values, right node is invalid",
			lhs:     yamlconv.Bool(true),
			rhs:     yamlconv.RawBool("world"),
			wantErr: cmpopts.AnyError,
		}, {
			name: "Scalar nodes with equivalent but invalid tags",
			lhs: &yaml.Node{
				Kind: yaml.ScalarNode,
				Tag:  "!!invalid",
			},
			rhs: &yaml.Node{
				Kind: yaml.ScalarNode,
				Tag:  "!!invalid",
			},
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Sequences contain different node types",
			lhs:     yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:     yamltest.MustParseNode(`["1", "2", "3"]`),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlcmp.Compare(tc.lhs, tc.rhs)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Compare() = %v; want %v", got, tc.want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Compare() error = %v; want %v", err, tc.wantErr)
			}
		})
	}
}
