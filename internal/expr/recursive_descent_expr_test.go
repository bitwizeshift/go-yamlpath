package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestRecursiveDescentExpr(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Empty node returns empty node",
		}, {
			name: "Input is sequence",
			input: []*yaml.Node{
				yamltest.MustParseNode(`["Alice", "Bob", "Charlie"]`),
			},
			want: []*yaml.Node{
				yamltest.MustParseNode(`["Alice", "Bob", "Charlie"]`),
				yamlutil.String("Alice"),
				yamlutil.String("Bob"),
				yamlutil.String("Charlie"),
			},
		}, {
			name: "Input is mapping",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
				yamlutil.String("Alice"),
				yamlutil.Number("30"),
			},
		}, {
			name: "Input is document",
			input: []*yaml.Node{
				yamltest.MustParseDocument(`{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamltest.MustParseDocument(`{"name": "Alice", "age": 30}`),
				yamlutil.String("Alice"),
				yamlutil.Number("30"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.RecursiveDescentExpr{}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("RecursiveDescentExpr.Eval() error = %v; want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("RecursiveDescentExpr.Apply() = %v, want %v", got, want)
			}
		})
	}
}
