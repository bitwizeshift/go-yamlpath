package expr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestRootExpr(t *testing.T) {
	root := []*yaml.Node{yamlconv.String("hello")}
	current := []*yaml.Node{yamlconv.String("world")}

	testCases := []struct {
		name    string
		root    string
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "root expression '$'",
			root: "$",
			want: root,
		}, {
			name: "current expression '@'",
			root: "@",
			want: current,
		}, {
			name:    "unsupported root expression",
			root:    "/",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := &expr.RootExpr{Root: tc.root}

			ctx := expr.NewContext(root).NewContext(current)

			got, err := sut.Eval(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("RootExpr.Eval() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("RootExpr.Eval() = %v, want %v", got, want)
			}
		})
	}
}
