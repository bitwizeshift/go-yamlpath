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

func TestLiteralExpr(t *testing.T) {
	testCases := []struct {
		name    string
		value   *yaml.Node
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Input node does not affect output",
			value: yamlconv.String("Hello world"),
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamlconv.String("Hello world"),
			},
		}, {
			name:  "Sequence node returns all elements",
			value: yamlconv.String("Hello world"),
			want: []*yaml.Node{
				yamlconv.String("Hello world"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.LiteralExpr{Nodes: []*yaml.Node{tc.value}}
			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("LiteralExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("LiteralExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
