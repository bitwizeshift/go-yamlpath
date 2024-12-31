package expr_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
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
			expr:    ExprReturnsError(testErr),
			wantErr: testErr,
		}, {
			name:    "Expr returns multiple nodes",
			expr:    ExprReturnsNodes(yamlutil.String("hello"), yamlutil.String("world")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Expr returns single non-scalar node",
			expr:    ExprReturnsNodes(YAML(t, `{"name": "Alice", "age": 30}`)),
			wantErr: expr.ErrEval,
		}, {
			name:  "Expr returns single scalar int node",
			expr:  ExprReturnsNodes(yamlutil.Number("1")),
			input: []*yaml.Node{YAML(t, `[1, 2, 3]`)},
			want:  []*yaml.Node{yamlutil.Number("2")},
		}, {
			name:  "Expr returns single scalar string node",
			expr:  ExprReturnsNodes(yamlutil.String("key")),
			input: []*yaml.Node{YAML(t, `{"key": "value"}`)},
			want:  []*yaml.Node{yamlutil.String("value")},
		}, {
			name:    "Expr returns invalid integer node",
			expr:    ExprReturnsNodes(yamlutil.Number("foo")),
			wantErr: expr.ErrEval,
		}, {
			name:    "Expr returns scalar boolean node",
			expr:    ExprReturnsNodes(yamlutil.True),
			wantErr: expr.ErrEval,
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
