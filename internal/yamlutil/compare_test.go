package yamlutil_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
			rhs:  []*yaml.Node{yamlutil.Boolean("true")},
			want: true,
		}, {
			name: "Right range is empty",
			lhs:  []*yaml.Node{yamlutil.Boolean("true")},
			want: false,
		}, {
			name:    "Ranges have different kinds",
			lhs:     []*yaml.Node{yamlutil.Number("42")},
			rhs:     []*yaml.Node{yamlutil.String("42")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.LessRange(tc.lhs, tc.rhs)

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
			rhs:  []*yaml.Node{yamlutil.Boolean("true")},
			want: -1,
		}, {
			name: "Right range is empty",
			lhs:  []*yaml.Node{yamlutil.Boolean("true")},
			want: 1,
		}, {
			name: "Ranges are different length",
			lhs:  []*yaml.Node{yamlutil.Number("42")},
			rhs:  []*yaml.Node{yamlutil.Number("42"), yamlutil.Number("43")},
			want: -1,
		}, {
			name: "Ranges have different values",
			lhs:  []*yaml.Node{yamlutil.Number("42"), yamlutil.Number("43")},
			rhs:  []*yaml.Node{yamlutil.Number("42"), yamlutil.Number("44")},
			want: -1,
		}, {
			name: "Ranges are equal",
			lhs:  []*yaml.Node{yamlutil.Number("42"), yamlutil.Number("43")},
			rhs:  []*yaml.Node{yamlutil.Number("42"), yamlutil.Number("43")},
		}, {
			name:    "Ranges have different kinds",
			lhs:     []*yaml.Node{yamlutil.Number("42")},
			rhs:     []*yaml.Node{yamlutil.String("42")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.CompareRange(tc.lhs, tc.rhs)

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
			rhs:  yamlutil.Boolean("true"),
			want: -1,
		}, {
			name: "Right node is nil",
			lhs:  yamlutil.Boolean("true"),
			want: 1,
		}, {
			name: "Both nodes are identical scalar strings",
			lhs:  yamlutil.String("hello"),
			rhs:  yamlutil.String("hello"),
			want: 0,
		}, {
			name: "Both nodes are different scalar strings",
			lhs:  yamlutil.String("hello"),
			rhs:  yamlutil.String("world"),
			want: -1,
		}, {
			name: "Both nodes are equivalent scalar numbers",
			lhs:  yamlutil.Number("42e0"),
			rhs:  yamlutil.Number("42"),
			want: 0,
		}, {
			name: "Both nodes are different scalar numbers",
			lhs:  yamlutil.Number("42"),
			rhs:  yamlutil.Number("43"),
			want: -1,
		}, {
			name: "Both nodes are the same scalar booleans",
			lhs:  yamlutil.Boolean("true"),
			rhs:  yamlutil.Boolean("true"),
			want: 0,
		}, {
			name: "Left node is false, right node is true",
			lhs:  yamlutil.Boolean("false"),
			rhs:  yamlutil.Boolean("true"),
			want: -1,
		}, {
			name: "Left node is true, right node is false",
			lhs:  yamlutil.Boolean("true"),
			rhs:  yamlutil.Boolean("false"),
			want: 1,
		}, {
			name: "Null nodes compare equal",
			lhs:  yamlutil.Null(),
			rhs:  yamlutil.Null(),
			want: 0,
		},
		// Sequence Values
		{
			name: "Sequences are different length",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2]`),
			want: 1,
		}, {
			name: "Sequences have different values",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2, 4]`),
			want: -1,
		}, {
			name: "Sequences are equal",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2, 3]`),
			want: 0,
		},
		// Errors
		{
			name: "Node kinds are different",
			lhs:  yamlutil.Boolean("true"),
			rhs: &yaml.Node{
				Kind: yaml.SequenceNode,
			},
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar values with different tags",
			lhs:     yamlutil.Number("1234"),
			rhs:     yamlutil.String("true"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar values with different non-numeric tags",
			lhs:     yamlutil.String("hello"),
			rhs:     yamlutil.Boolean("true"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar numeric values, left node is invalid",
			lhs:     yamlutil.Number("hello"),
			rhs:     yamlutil.Number("42"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar numeric values, right node is invalid",
			lhs:     yamlutil.Number("42"),
			rhs:     yamlutil.Number("world"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar boolean values, left node is invalid",
			lhs:     yamlutil.Boolean("hello"),
			rhs:     yamlutil.Boolean("true"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar boolean values, right node is invalid",
			lhs:     yamlutil.Boolean("true"),
			rhs:     yamlutil.Boolean("world"),
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
			lhs:     newYAML(t, `[1, 2, 3]`),
			rhs:     newYAML(t, `["1", "2", "3"]`),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.Compare(tc.lhs, tc.rhs)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Compare() = %v; want %v", got, tc.want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Compare() error = %v; want %v", err, tc.wantErr)
			}
		})
	}
}
