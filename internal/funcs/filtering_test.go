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
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
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
				invocationtest.SuccessParameter(yamlutil.True).AddSuccess(yamlutil.False),
			},
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
			want:  []*yaml.Node{yamlutil.String("foo")},
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

func TestTransform(t *testing.T) {
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
				invocationtest.SuccessParameter(yamlutil.String("hello")).AddSuccess(yamlutil.String("world")),
			},
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
			want:  []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
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

func TestKeys(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  []*yaml.Node
	}{
		{
			name:  "Keys of a mapping",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar", "baz": "qux"}`)},
			want:  []*yaml.Node{yamlutil.String("foo"), yamlutil.String("baz")},
		}, {
			name:  "Keys of a sequence",
			input: []*yaml.Node{yamltest.MustParseNode(`["foo", "bar", "baz", "qux"]`)},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Keys(ctx)

			if err != nil {
				t.Fatalf("Keys() error = %v", err)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Keys() = %v, want %v", got, want)
			}
		})
	}
}
