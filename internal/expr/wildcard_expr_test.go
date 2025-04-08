package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestWildcardExpr(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty node returns empty node",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name: "Mapping node returns all fields",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamlconv.String("Alice"),
				yamlconv.Number(30),
			},
		}, {
			name: "Sequence node returns empty node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`["Alice", "Bob"]`),
			},
			want: []*yaml.Node{},
		},
	}

	var sut expr.WildcardExpr
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("WildcardExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("WildcardExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
