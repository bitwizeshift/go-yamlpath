package expr_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestValueExpr(t *testing.T) {
	testCases := []struct {
		name    string
		value   *yaml.Node
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Input node does not affect output",
			value: yamlutil.String("Hello world"),
			input: []*yaml.Node{
				YAML(t, `{"name": "Alice", "age": 30}`),
			},
			want: []*yaml.Node{
				yamlutil.String("Hello world"),
			},
		}, {
			name:  "Sequence node returns all elements",
			value: yamlutil.String("Hello world"),
			want: []*yaml.Node{
				yamlutil.String("Hello world"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.ValueExpr{Node: tc.value}
			got, err := sut.Eval(context.Background(), tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ValueExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("ValueExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}