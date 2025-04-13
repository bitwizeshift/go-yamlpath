package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestIndexExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		indices expr.Expr
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:    "empty nodes select nothing",
			indices: toIndices(),
			input:   []*yaml.Node{},
			want:    []*yaml.Node{},
		}, {
			name:    "negative index selects from end of the sequence",
			indices: toIndices(-1),
			input:   []*yaml.Node{yamltest.MustParseNode(`["hello", "world"]`)},
			want:    []*yaml.Node{yamlconv.String("world")},
		}, {
			name:    "index out of range returns empty",
			indices: toIndices(2),
			input:   []*yaml.Node{yamltest.MustParseNode(`["hello", "world"]`)},
			want:    []*yaml.Node{},
		}, {
			name:    "multiple indices select multiple nodes",
			indices: toIndices(1, 2),
			input:   []*yaml.Node{yamltest.MustParseNode(`["foo", "bar", "baz", "buz"]`)},
			want:    []*yaml.Node{yamlconv.String("bar"), yamlconv.String("baz")},
		}, {
			name:    "Non-sequence nodes are ignored",
			indices: toIndices(0),
			input:   []*yaml.Node{yamlconv.String("hello")},
			want:    []*yaml.Node{},
		}, {
			name: "non-integer indices nodes return an error",
			indices: exprtest.Return(
				yamlconv.String("foo"),
			),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "expression returns an error",
			indices: exprtest.Error(testErr),
			wantErr: testErr,
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IndexExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}

func toIndices(indices ...int64) expr.Expr {
	return exprtest.Return(yamlconv.Ints(indices...)...)
}
