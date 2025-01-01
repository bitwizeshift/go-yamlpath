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
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestScriptExpr(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		expr    expr.Expr
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:    "Expr returns error",
			expr:    exprtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Expr returns multiple nodes",
			expr:    exprtest.Return(yamlutil.String("hello"), yamlutil.String("world")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Expr returns single non-scalar node",
			expr:    exprtest.Return(yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)),
			wantErr: errs.ErrEval,
		}, {
			name:  "Expr returns single scalar int node",
			expr:  exprtest.Return(yamlutil.Number("1")),
			input: []*yaml.Node{yamltest.MustParseNode(`[1, 2, 3]`)},
			want:  []*yaml.Node{yamlutil.Number("2")},
		}, {
			name:  "Expr returns single scalar string node",
			expr:  exprtest.Return(yamlutil.String("key")),
			input: []*yaml.Node{yamltest.MustParseNode(`{"key": "value"}`)},
			want:  []*yaml.Node{yamlutil.String("value")},
		}, {
			name:    "Expr returns invalid integer node",
			expr:    exprtest.Return(yamlutil.Number("foo")),
			wantErr: errs.ErrEval,
		}, {
			name:    "Expr returns scalar boolean node",
			expr:    exprtest.Return(yamlutil.True),
			wantErr: errs.ErrEval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.ScriptExpr{Expr: tc.expr}

			got, err := sut.Eval(expr.NewContext(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ScriptExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("ScriptExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
