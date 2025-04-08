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
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
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
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Non-empty input evalutes to false",
			input: []*yaml.Node{yamlconv.String("hello")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Empty(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Empty() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "Non-empty input set returns true",
			input: []*yaml.Node{yamlconv.String("hello")},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name: "Non-empty input set with params that returns values is true",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.String("example")),
			},
			input: []*yaml.Node{yamlconv.String("hello")},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name: "Non-empty input set with params that returns no values is false",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(),
			},
			input: []*yaml.Node{yamlconv.String("hello")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Non-empty input set with params that returns error is error",
			params: []invocation.Parameter{
				invocationtest.ErrorParameter(testErr),
			},
			input:   []*yaml.Node{yamlconv.String("hello")},
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Number(0)},
		}, {
			name:  "Non-empty collection returns count",
			input: []*yaml.Node{yamlconv.String("hello")},
			want:  []*yaml.Node{yamlconv.Number(1)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Count(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Count() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
				yamlconv.String("hello"),
				yamlconv.String("world"),
				yamlconv.String("hello")},
			want: []*yaml.Node{
				yamlconv.String("hello"),
				yamlconv.String("world")},
		}, {
			name: "Input contains no duplicates",
			input: []*yaml.Node{
				yamlconv.String("hello"),
				yamlconv.String("world"),
				yamlconv.String("goodbye"),
			},
			want: []*yaml.Node{
				yamlconv.String("hello"),
				yamlconv.String("world"),
				yamlconv.String("goodbye"),
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
				yamlconv.String("hello"),
				yamlconv.String("world"),
				yamlconv.String("hello")},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		}, {
			name: "Input contains no duplicates",
			input: []*yaml.Node{
				yamlconv.String("hello"),
				yamlconv.String("world"),
				yamlconv.String("goodbye"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(true),
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(true)},
		},
		{
			name:  "Param returns truthy value",
			input: []*yaml.Node{yamlconv.String("hello")},
			param: invocationtest.SuccessParameter(yamlconv.Bool(true)),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Param returns falsey value",
			input: []*yaml.Node{yamlconv.String("hello")},
			param: invocationtest.SuccessParameter(yamlconv.Bool(false)),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Param returns error",
			input:   []*yaml.Node{yamlconv.String("hello")},
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
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
		{
			name:  "Param returns truthy value",
			input: []*yaml.Node{yamlconv.String("hello")},
			param: invocationtest.SuccessParameter(yamlconv.Bool(true)),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Param returns falsey value",
			input: []*yaml.Node{yamlconv.String("hello")},
			param: invocationtest.SuccessParameter(yamlconv.Bool(false)),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Param returns error",
			input:   []*yaml.Node{yamlconv.String("hello")},
			param:   invocationtest.ErrorParameter(testErr),
			wantErr: testErr,
		}, {
			name:  "Param returns false then true",
			input: []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")},
			param: invocationtest.NewParameter().AddSuccess(yamlconv.Bool(false)).AddSuccess(yamlconv.Bool(true)),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Any(ctx, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Any() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "non-boolean input returns false",
			input: []*yaml.Node{yamlconv.Bool(true), yamlconv.String("true"), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "all true values return true",
			input: []*yaml.Node{yamlconv.Bool(true), yamlconv.Bool(true), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "any false values return false",
			input: []*yaml.Node{yamlconv.Bool(true), yamlconv.Bool(false), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AllTrue(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AllTrue() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "any true value returns true",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "all false or non-boolean returns false",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.String("true"), yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AnyTrue(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AnyTrue() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "non-boolean input returns false",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.String("false"), yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "all false values return true",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(false), yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "any true values return false",
			input: []*yaml.Node{yamlconv.Bool(false), yamlconv.Bool(false), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AllFalse(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AllFalse() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
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
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:  "any false value returns true",
			input: []*yaml.Node{yamlconv.Bool(true), yamlconv.String("true"), yamlconv.Bool(false)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "all true or non-boolean returns false",
			input: []*yaml.Node{yamlconv.Bool(true), yamlconv.String("true"), yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.AnyFalse(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("AnyFalse() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("AnyFalse() = %v, want %v", got, want)
			}
		})
	}
}
