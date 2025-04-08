package expr_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestSliceExpr(t *testing.T) {
	testCases := []struct {
		name    string
		slice   *expr.Slice
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty node returns empty node",
			slice: &expr.Slice{Start: 1, End: math.MaxInt},
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Input is sequence, slice is valid",
			slice: &expr.Slice{Start: 1, End: 3},
			input: []*yaml.Node{
				yamltest.MustParseNode(`["Alice", "Bob", "Charlie"]`),
			},
			want: []*yaml.Node{
				yamlconv.String("Bob"),
				yamlconv.String("Charlie"),
			},
		}, {
			name:  "Input is mapping",
			slice: &expr.Slice{Start: 1, End: 3},
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			},
			want: nil,
		}, {
			name:  "Input is sequence with larger step",
			slice: &expr.Slice{Start: 1, End: 4, Step: 2},
			input: []*yaml.Node{
				yamltest.MustParseNode(`["Alice", "Bob", "Charlie", "David"]`),
			},
			want: []*yaml.Node{
				yamlconv.String("Bob"),
				yamlconv.String("David"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.SliceExpr{
				Slice: tc.slice,
			}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("SliceExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("SliceExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
