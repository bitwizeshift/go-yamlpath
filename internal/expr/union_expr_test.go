package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestUnionExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		exprs   []expr.Expr
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			exprs: []expr.Expr{},
			want:  []*yaml.Node{},
		}, {
			name: "Single expression returns multiple values",
			exprs: []expr.Expr{
				exprtest.Return(yamlconv.String("Hello"), yamlconv.String("World")),
			},
			want: []*yaml.Node{
				yamlconv.String("Hello"),
				yamlconv.String("World"),
			},
		}, {
			name: "Multiple expressions return multiple values",
			exprs: []expr.Expr{
				exprtest.Return(yamlconv.String("Hello"), yamlconv.String("World")),
				exprtest.Return(yamlconv.String("Goodbye"), yamlconv.String("World")),
			},
			want: []*yaml.Node{
				yamlconv.String("Hello"),
				yamlconv.String("World"),
				yamlconv.String("Goodbye"),
				yamlconv.String("World"),
			},
		}, {
			name: "Expression returns error",
			exprs: []expr.Expr{
				exprtest.Error(testErr),
			},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := expr.UnionExpr(tc.exprs)
			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("UnionExpr.Eval() error = %v, want nil", err)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("UnionExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
