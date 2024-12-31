package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
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
				YAML(t, `{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamlutil.String("Alice"),
				yamlutil.Number("30"),
			},
		}, {
			name: "Sequence node returns all elements",
			input: []*yaml.Node{
				YAML(t, `["Alice", "Bob"]`),
			},
			want: []*yaml.Node{
				yamlutil.String("Alice"),
				yamlutil.String("Bob"),
			},
		},
	}

	var sut expr.WildcardExpr
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("WildcardExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("WildcardExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
