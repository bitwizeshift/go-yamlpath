package yamlcmp_test

import (
	"testing"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func nodes(nodes ...*yaml.Node) []*yaml.Node {
	return nodes
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs *yaml.Node
		want     bool
	}{
		{
			name: "Nil nodes are equal",
			want: true,
		}, {
			name: "Left node is nil",
			rhs:  yamlconv.Bool(true),
			want: false,
		}, {
			name: "Right node is nil",
			lhs:  yamlconv.Bool(true),
			want: false,
		}, {
			name: "Node kinds are different",
			lhs:  yamlconv.Bool(true),
			rhs: &yaml.Node{
				Kind: yaml.SequenceNode,
			},
		}, {
			name: "Both nodes are identical strings",
			lhs:  yamlconv.String("hello"),
			rhs:  yamlconv.String("hello"),
			want: true,
		}, {
			name: "Both nodes are equivalent numbers",
			lhs:  yamlconv.NumberString("42e0"),
			rhs:  yamlconv.Number(42),
			want: true,
		}, {
			name: "Both nodes are numbers with different precision",
			lhs:  yamlconv.Number(42),
			rhs:  yamlconv.NumberString("42.0"),
			want: false,
		}, {
			name: "Both nodes are different scalar kinds",
			lhs:  yamlconv.String("hello"),
			rhs:  yamlconv.Number(42),
			want: false,
		}, {
			name: "List nodes have different number of children",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2]`),
			want: false,
		}, {
			name: "List nodes same number of children but different values",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2, 4]`),
			want: false,
		}, {
			name: "List nodes same number of children and values",
			lhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			rhs:  yamltest.MustParseNode(`[1, 2, 3]`),
			want: true,
		}, {
			name: "Map nodes have different number of children",
			lhs:  yamltest.MustParseNode(`{"a": 1, "b": 2}`),
			rhs:  yamltest.MustParseNode(`{"a": 1}`),
			want: false,
		}, {
			name: "Map nodes same number of children but different values",
			lhs:  yamltest.MustParseNode(`{"a": 1, "b": 2}`),
			rhs:  yamltest.MustParseNode(`{"a": 1, "b": 3}`),
			want: false,
		}, {
			name: "Map nodes same number of children but different keys",
			lhs:  yamltest.MustParseNode(`{"a": 1, "b": 2}`),
			rhs:  yamltest.MustParseNode(`{"a": 1, "c": 2}`),
			want: false,
		}, {
			name: "Map nodes same number of children and values",
			lhs:  yamltest.MustParseNode(`{"a": 1, "b": 2}`),
			rhs:  yamltest.MustParseNode(`{"a": 1, "b": 2}`),
			want: true,
		}, {
			name: "Equivalent but invalid number fields return true",
			lhs:  yamlconv.NumberString("hello"),
			rhs:  yamlconv.NumberString("hello"),
			want: true,
		}, {
			name: "Different but invalid number fields return false",
			lhs:  yamlconv.NumberString("hello"),
			rhs:  yamlconv.NumberString("world"),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlcmp.Equal(tc.lhs, tc.rhs)

			if got != tc.want {
				t.Errorf("Equal() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestEqualRange(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs []*yaml.Node
		want     bool
	}{
		{
			name: "Nil nodes are equal",
			want: true,
		}, {
			name: "More nodes in lhs",
			rhs:  nodes(yamlconv.Bool(true)),
			want: false,
		}, {
			name: "More nodes in rhs",
			lhs:  nodes(yamlconv.Bool(true)),
			want: false,
		}, {
			name: "Same representation",
			lhs:  nodes(yamlconv.Bool(true), yamlconv.Bool(false)),
			rhs:  nodes(yamlconv.Bool(true), yamlconv.Bool(false)),
			want: true,
		}, {
			name: "Different representation",
			lhs:  nodes(yamlconv.Bool(true), yamlconv.Bool(false)),
			rhs:  nodes(yamlconv.Bool(true), yamlconv.Bool(true)),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlcmp.EqualRange(tc.lhs, tc.rhs)

			if got != tc.want {
				t.Errorf("EqualRange() = %v; want %v", got, tc.want)
			}
		})
	}
}
