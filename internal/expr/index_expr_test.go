package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestIndexExpr(t *testing.T) {
	testCases := []struct {
		name    string
		indices []int64
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:    "empty nodes select nothing",
			indices: []int64{},
			input:   []*yaml.Node{},
			want:    []*yaml.Node{},
		}, {
			name:    "negative index selects from end of the sequence",
			indices: []int64{-1},
			input:   []*yaml.Node{YAML(t, `["hello", "world"]`)},
			want:    []*yaml.Node{yamlutil.String("world")},
		}, {
			name:    "index out of range returns empty",
			indices: []int64{2},
			input:   []*yaml.Node{YAML(t, `["hello", "world"]`)},
			want:    []*yaml.Node{},
		}, {
			name:    "multiple indices select multiple nodes",
			indices: []int64{1, 2},
			input:   []*yaml.Node{YAML(t, `["foo", "bar", "baz", "buz"]`)},
			want:    []*yaml.Node{yamlutil.String("bar"), yamlutil.String("baz")},
		}, {
			name:    "Non-sequence nodes are ignored",
			indices: []int64{0},
			input:   []*yaml.Node{yamlutil.String("hello")},
			want:    []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.IndexExpr{
				Indices: tc.indices,
			}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IndexExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("IndexExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
