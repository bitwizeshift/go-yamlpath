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

func TestFieldExpr(t *testing.T) {
	testCases := []struct {
		name    string
		fields  []string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:   "empty nodes select nothing",
			fields: []string{},
			input:  []*yaml.Node{},
			want:   []*yaml.Node{},
		}, {
			name:   "field matches one node",
			fields: []string{"foo"},
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": 1, "bar": 2}`),
			},
			want: []*yaml.Node{
				yamlconv.Number(1),
			},
		}, {
			name:   "field matches multiple nodes",
			fields: []string{"foo"},
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": 1, "bar": 2}`),
				yamltest.MustParseNode(`{"foo": 3, "bar": 4}`),
			},
			want: []*yaml.Node{
				yamlconv.Number(1),
				yamlconv.Number(3),
			},
		}, {
			name:   "multiple fields match node",
			fields: []string{"foo", "bar"},
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": 1, "bar": 2}`),
			},
			want: []*yaml.Node{
				yamlconv.Number(1),
				yamlconv.Number(2),
			},
		}, {
			name:   "Non-map is ignored",
			fields: []string{"foo"},
			input: []*yaml.Node{
				yamlconv.String("hello"),
			},
			want: []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.FieldExpr{
				Fields: tc.fields,
			}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("FieldExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("FieldExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
