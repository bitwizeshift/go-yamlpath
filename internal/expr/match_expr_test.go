package expr_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestMatchExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		regex   string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty node evaluates to false",
			regex: ".*",
			expr:  exprtest.Return(),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:    "Subexpr evaluates error",
			regex:   ".*",
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:  "Subexpr returns multiple nodes",
			regex: ".*",
			expr:  exprtest.Return(yamlutil.String("hello"), yamlutil.String("world")),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Subexpr returns scalar non-string node",
			regex: ".*",
			expr:  exprtest.Return(yamlutil.Number("42")),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Subexpr returns non-scalar node",
			regex: ".*",
			expr:  exprtest.Return(yamltest.MustParseNode(`{"foo": "bar"}`)),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Subexpr returns scalar string node that matches",
			regex: ".*",
			expr:  exprtest.Return(yamlutil.String("hello")),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Subexpr returns scalar string node that does not match",
			regex: "world",
			expr:  exprtest.Return(yamlutil.String("hello")),
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.MatchExpr{
				Expr:  tc.expr,
				Regex: regexp.MustCompile(tc.regex),
			}

			got, err := sut.Eval(expr.NewContext(nil))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("MatchExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("MatchExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
