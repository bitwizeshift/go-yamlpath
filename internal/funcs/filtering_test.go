package funcs_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestWhere(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Criteria parameter returns true",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlutil.True),
			},
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
			want:  []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
		}, {
			name: "Criteria parameter returns false",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlutil.False),
			},
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
			want:  []*yaml.Node{},
		}, {
			name: "Criteria parameter returns error",
			params: []invocation.Parameter{
				invocationtest.ErrorParameter(testErr),
			},
			input:   []*yaml.Node{yamlutil.String("hello")},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Where(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Where() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Where() = %v, want %v", got, want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Projection parameter returns elements",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlutil.String("hello")),
			},
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
			want:  []*yaml.Node{yamlutil.String("hello"), yamlutil.String("hello")},
		}, {
			name: "Projection parameter returns error",
			params: []invocation.Parameter{
				invocationtest.ErrorParameter(testErr),
			},
			input:   []*yaml.Node{yamlutil.String("hello")},
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Transform(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Select() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Select() = %v, want %v", got, want)
			}
		})
	}
}
