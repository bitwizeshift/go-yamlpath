package yamlutil_test

import (
	"testing"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func nodes(nodes ...*yaml.Node) []*yaml.Node {
	return nodes
}

func newYAML(t *testing.T, s string) *yaml.Node {
	var n yaml.Node
	if err := yaml.Unmarshal([]byte(s), &n); err != nil {
		t.Fatal(err)
	}
	return yamlutil.Normalize(&n)[0]
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
			rhs:  yamlutil.Boolean("true"),
			want: false,
		}, {
			name: "Right node is nil",
			lhs:  yamlutil.Boolean("true"),
			want: false,
		}, {
			name: "Node kinds are different",
			lhs:  yamlutil.Boolean("true"),
			rhs: &yaml.Node{
				Kind: yaml.SequenceNode,
			},
		}, {
			name: "Both nodes are identical strings",
			lhs:  yamlutil.String("hello"),
			rhs:  yamlutil.String("hello"),
			want: true,
		}, {
			name: "Both nodes are equivalent numbers",
			lhs:  yamlutil.Number("42e0"),
			rhs:  yamlutil.Number("42"),
			want: true,
		}, {
			name: "Both nodes are numbers with different precision",
			lhs:  yamlutil.Number("42"),
			rhs:  yamlutil.Number("42.0"),
			want: false,
		}, {
			name: "Both nodes are different scalar kinds",
			lhs:  yamlutil.String("hello"),
			rhs:  yamlutil.Number("42"),
			want: false,
		}, {
			name: "List nodes have different number of children",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2]`),
			want: false,
		}, {
			name: "List nodes same number of children but different values",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2, 4]`),
			want: false,
		}, {
			name: "List nodes same number of children and values",
			lhs:  newYAML(t, `[1, 2, 3]`),
			rhs:  newYAML(t, `[1, 2, 3]`),
			want: true,
		}, {
			name: "Map nodes have different number of children",
			lhs:  newYAML(t, `{"a": 1, "b": 2}`),
			rhs:  newYAML(t, `{"a": 1}`),
			want: false,
		}, {
			name: "Map nodes same number of children but different values",
			lhs:  newYAML(t, `{"a": 1, "b": 2}`),
			rhs:  newYAML(t, `{"a": 1, "b": 3}`),
			want: false,
		}, {
			name: "Map nodes same number of children but different keys",
			lhs:  newYAML(t, `{"a": 1, "b": 2}`),
			rhs:  newYAML(t, `{"a": 1, "c": 2}`),
			want: false,
		}, {
			name: "Map nodes same number of children and values",
			lhs:  newYAML(t, `{"a": 1, "b": 2}`),
			rhs:  newYAML(t, `{"a": 1, "b": 2}`),
			want: true,
		}, {
			name: "Equivalent but invalid number fields return true",
			lhs:  yamlutil.Number("hello"),
			rhs:  yamlutil.Number("hello"),
			want: true,
		}, {
			name: "Different but invalid number fields return false",
			lhs:  yamlutil.Number("hello"),
			rhs:  yamlutil.Number("world"),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlutil.Equal(tc.lhs, tc.rhs)

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
			rhs:  nodes(yamlutil.Boolean("true")),
			want: false,
		}, {
			name: "More nodes in rhs",
			lhs:  nodes(yamlutil.Boolean("true")),
			want: false,
		}, {
			name: "Same representation",
			lhs:  nodes(yamlutil.Boolean("true"), yamlutil.Boolean("false")),
			rhs:  nodes(yamlutil.Boolean("true"), yamlutil.Boolean("false")),
			want: true,
		}, {
			name: "Different representation",
			lhs:  nodes(yamlutil.Boolean("true"), yamlutil.Boolean("false")),
			rhs:  nodes(yamlutil.Boolean("true"), yamlutil.Boolean("true")),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlutil.EqualRange(tc.lhs, tc.rhs)

			if got != tc.want {
				t.Errorf("EqualRange() = %v; want %v", got, tc.want)
			}
		})
	}
}
