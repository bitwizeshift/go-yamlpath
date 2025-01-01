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

func TestEmpty(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input evalutes to true",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Non-empty input evalutes to false",
			input: []*yaml.Node{yamlutil.String("hello")},
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Empty(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Empty() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Empty() = %v, want %v", got, want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input set returns false",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "Non-empty input set returns true",
			input: []*yaml.Node{yamlutil.String("hello")},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name: "Non-empty input set with params that returns values is true",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlutil.String("example")),
			},
			input: []*yaml.Node{yamlutil.String("hello")},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name: "Non-empty input set with params that returns no values is false",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(),
			},
			input: []*yaml.Node{yamlutil.String("hello")},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name: "Non-empty input set with params that returns error is error",
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

			got, err := funcs.Exists(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Exists() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Exists() = %v, want %v", got, want)
			}
		})
	}
}
