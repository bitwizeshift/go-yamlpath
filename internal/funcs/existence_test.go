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

func TestCount(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input set returns 0",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.Number("0")},
		}, {
			name:  "Non-empty collection returns count",
			input: []*yaml.Node{yamlutil.String("hello")},
			want:  []*yaml.Node{yamlutil.Number("1")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Count(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Count() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Count() = %v, want %v", got, want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input set returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name: "Input contains duplicates",
			input: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world"),
				yamlutil.String("hello")},
			want: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world")},
		}, {
			name: "Input contains no duplicates",
			input: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world"),
				yamlutil.String("goodbye"),
			},
			want: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world"),
				yamlutil.String("goodbye"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Distinct(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Distinct() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Distinct() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsDistinct(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input set returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name: "Input contains duplicates",
			input: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world"),
				yamlutil.String("hello")},
			want: []*yaml.Node{
				yamlutil.Boolean("false"),
			},
		}, {
			name: "Input contains no duplicates",
			input: []*yaml.Node{
				yamlutil.String("hello"),
				yamlutil.String("world"),
				yamlutil.String("goodbye"),
			},
			want: []*yaml.Node{
				yamlutil.Boolean("true"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsDistinct(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Distinct() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Distinct() = %v, want %v", got, want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		param   invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns true",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.True},
		},
		{
			name:  "Param returns truthy value",
			input: []*yaml.Node{yamlutil.String("hello")},
			param: invocationtest.SuccessParameter(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Param returns falsey value",
			input: []*yaml.Node{yamlutil.String("hello")},
			param: invocationtest.SuccessParameter(yamlutil.False),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:    "Param returns error",
			input:   []*yaml.Node{yamlutil.String("hello")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.All(ctx, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("All() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("All() = %v, want %v", got, want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		param   invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns false",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.False},
		},
		{
			name:  "Param returns truthy value",
			input: []*yaml.Node{yamlutil.String("hello")},
			param: invocationtest.SuccessParameter(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "Param returns falsey value",
			input: []*yaml.Node{yamlutil.String("hello")},
			param: invocationtest.SuccessParameter(yamlutil.False),
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:    "Param returns error",
			input:   []*yaml.Node{yamlutil.String("hello")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		}, {
			name:  "Param returns false then true",
			input: []*yaml.Node{yamlutil.String("hello"), yamlutil.String("world")},
			param: invocationtest.NewParameter().AddSuccess(yamlutil.False).AddSuccess(yamlutil.True),
			want:  []*yaml.Node{yamlutil.True},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Any(ctx, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Any() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("Any() = %v, want %v", got, want)
			}
		})
	}
}

func TestAllTrue(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns true",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "non-boolean input returns false",
			input: []*yaml.Node{yamlutil.True, yamlutil.String("true"), yamlutil.True},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "all true values return true",
			input: []*yaml.Node{yamlutil.True, yamlutil.True, yamlutil.True},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "any false values return false",
			input: []*yaml.Node{yamlutil.True, yamlutil.False, yamlutil.True},
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AllTrue(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AllTrue() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("AllTrue() = %v, want %v", got, want)
			}
		})
	}
}

func TestAnyTrue(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns false",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "any true value returns true",
			input: []*yaml.Node{yamlutil.False, yamlutil.String("true"), yamlutil.True},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "all false or non-boolean returns false",
			input: []*yaml.Node{yamlutil.False, yamlutil.String("true"), yamlutil.False},
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AnyTrue(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AnyTrue() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("AnyTrue() = %v, want %v", got, want)
			}
		})
	}
}

func TestAllFalse(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns true",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "non-boolean input returns false",
			input: []*yaml.Node{yamlutil.False, yamlutil.String("false"), yamlutil.False},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "all false values return true",
			input: []*yaml.Node{yamlutil.False, yamlutil.False, yamlutil.False},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "any true values return false",
			input: []*yaml.Node{yamlutil.False, yamlutil.False, yamlutil.True},
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AllFalse(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AllFalse() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("AllFalse() = %v, want %v", got, want)
			}
		})
	}
}

func TestAnyFalse(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns false",
			input: []*yaml.Node{},
			want:  []*yaml.Node{yamlutil.False},
		}, {
			name:  "any false value returns true",
			input: []*yaml.Node{yamlutil.True, yamlutil.String("true"), yamlutil.False},
			want:  []*yaml.Node{yamlutil.True},
		}, {
			name:  "all true or non-boolean returns false",
			input: []*yaml.Node{yamlutil.True, yamlutil.String("true"), yamlutil.True},
			want:  []*yaml.Node{yamlutil.False},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AnyFalse(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AnyFalse() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlutil.EqualRange(got, want) {
				t.Errorf("AnyFalse() = %v, want %v", got, want)
			}
		})
	}
}
