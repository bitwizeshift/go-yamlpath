package funcs_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestSingle(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "single value returns that value",
			input: []*yaml.Node{yamlutil.String("false")},
			want:  []*yaml.Node{yamlutil.String("false")},
		}, {
			name:    "multiple values returns singleton error",
			input:   []*yaml.Node{yamlutil.False, yamlutil.False},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Single(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Single() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Single() = %v, want %v", got, want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		param   invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Multiple elements, no parameter returns first parameter",
			input: []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			want:  []*yaml.Node{yamlutil.String("hello")},
		}, {
			name:    "Multiple elements, parameter returns error",
			input:   []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		}, {
			name:  "Multiple elements, parameter returns less than input",
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			param: invocationtest.SuccessParameter(yamlutil.Number("2")),
			want:  []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar")},
		}, {
			name:  "Multiple elemtns, parameter returns more than input",
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			param: invocationtest.SuccessParameter(yamlutil.Number("6")),
			want:  []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			var params []invocation.Parameter
			if tc.param != nil {
				params = append(params, tc.param)
			}
			got, err := funcs.First(ctx, params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("First() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("First() = %v, want %v", got, want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		param   invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Multiple elements, no parameter returns last parameter",
			input: []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			want:  []*yaml.Node{yamlutil.String("world")},
		}, {
			name:    "Multiple elements, parameter returns error",
			input:   []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		}, {
			name:  "Multiple elements, parameter returns less than input",
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			param: invocationtest.SuccessParameter(yamlutil.Number("2")),
			want:  []*yaml.Node{yamlutil.String("bar"), yamlutil.String("baz")},
		}, {
			name:  "Multiple elements, parameter returns more than input",
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			param: invocationtest.SuccessParameter(yamlutil.Number("6")),
			want:  []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			var params []invocation.Parameter
			if tc.param != nil {
				params = append(params, tc.param)
			}
			got, err := funcs.Last(ctx, params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Last() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Last() = %v, want %v", got, want)
			}
		})
	}
}

func TestSkip(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		param   invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			param: invocationtest.SuccessParameter(yamlutil.Number("2")),
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Multiple elements, skips n elements",
			param: invocationtest.SuccessParameter(yamlutil.Number("1")),
			input: []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			want:  []*yaml.Node{yamlutil.String("world")},
		}, {
			name:    "Multiple elements, parameter returns error",
			input:   []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		}, {
			name:  "Multiple elements, parameter returns less than input",
			param: invocationtest.SuccessParameter(yamlutil.Number("2")),
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			want:  []*yaml.Node{yamlutil.String("baz")},
		}, {
			name:  "Multiple elements, parameter returns more than input",
			param: invocationtest.SuccessParameter(yamlutil.Number("6")),
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			want:  []*yaml.Node{},
		}, {
			name:  "Multiple elements, negative number within range",
			param: invocationtest.SuccessParameter(yamlutil.Number("-2")),
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			want:  []*yaml.Node{yamlutil.String("foo")},
		}, {
			name:  "Multiple elements, parameter returns more than input",
			param: invocationtest.SuccessParameter(yamlutil.Number("-6")),
			input: []*yaml.Node{yamlutil.String("foo"), yamlutil.String("bar"), yamlutil.String("baz")},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Skip(ctx, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Skip() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Skip() = %v, want %v", got, want)
			}
		})
	}
}
