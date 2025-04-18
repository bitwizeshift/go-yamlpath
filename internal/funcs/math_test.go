package funcs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestAbs(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single negative integer",
			input: yamlconv.Ints(-42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(42.5),
			want:  yamlconv.Floats(42.5),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-42.5),
			want:  yamlconv.Floats(42.5),
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Abs(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Abs() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Abs() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single negative integer",
			input: yamlconv.Ints(-42),
			want:  yamlconv.Ints(-42),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(42.5),
			want:  yamlconv.Ints(43),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-42.5),
			want:  yamlconv.Ints(-42),
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Ceil(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Ceil() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Ceil() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single negative integer",
			input: yamlconv.Ints(-42),
			want:  yamlconv.Ints(-42),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(42.5),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-42.5),
			want:  yamlconv.Ints(-43),
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Floor(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Floor() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Floor() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestExp(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(5),
			want:  yamlconv.Floats(148.41315),
		}, {
			name:  "Single negative integer",
			input: yamlconv.Ints(-5),
			want:  yamlconv.Floats(0.00673),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(5.5),
			want:  yamlconv.Floats(244.69193),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-5.5),
			want:  yamlconv.Floats(0.00408),
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Exp(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Exp() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Exp() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestLn(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(5),
			want:  yamlconv.Floats(1.60944),
		}, {
			name:    "Single negative integer",
			input:   yamlconv.Ints(-5),
			wantErr: cmpopts.AnyError,
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(5.5),
			want:  yamlconv.Floats(1.70475),
		}, {
			name:    "Single negative decimal",
			input:   yamlconv.Floats(-5.5),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Ln(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Ln() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Ln() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestLog(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single positive integer",
			input: yamlconv.Ints(1000),
			want:  yamlconv.Ints(3),
		}, {
			name:    "Single negative integer",
			input:   yamlconv.Ints(-5),
			wantErr: cmpopts.AnyError,
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(5.5),
			want:  yamlconv.Floats(0.73),
		}, {
			name:    "Single negative decimal",
			input:   yamlconv.Floats(-5.5),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Log(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Log() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Log() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		param   invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			param: invocationtest.Int(0),
			want:  []*yaml.Node{},
		}, {
			name:  "Collection and parameter are positive integers",
			input: yamlconv.Ints(2),
			param: invocationtest.Int(3),
			want:  yamlconv.Ints(8),
		}, {
			name:  "Collection is a positive integer, and parameter is a negative integer",
			input: yamlconv.Ints(2),
			param: invocationtest.Int(-3),
			want:  yamlconv.Floats(0.125),
		}, {
			name:  "Collection is a positive decimal, and parameter is a positive integer",
			input: yamlconv.Floats(2.5),
			param: invocationtest.Int(3),
			want:  yamlconv.Floats(15.625),
		}, {
			name:  "Collection is a positive decimal, and parameter is a negative integer",
			input: yamlconv.Floats(2.5),
			param: invocationtest.Int(-3),
			want:  yamlconv.Floats(0.064),
		}, {
			name:    "Input is non-singleton input",
			input:   yamlconv.Ints(1, 2),
			param:   invocationtest.Int(1),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Input is incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			param:   invocationtest.Int(1),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is incorrect tag",
			input:   yamlconv.Strings("string"),
			param:   invocationtest.Int(1),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Parameter is non-singleton input",
			input:   yamlconv.Ints(1),
			param:   invocationtest.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Parameter is incorrect kind",
			input:   yamlconv.Ints(1),
			param:   invocationtest.SuccessParameter(yamltest.MustParseNode(`{}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Parameter is incorrect tag",
			input:   yamlconv.Ints(1),
			param:   invocationtest.String("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Pow(ctx, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Pow() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Pow() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestRound(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(42.5),
			want:  yamlconv.Ints(43),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-42.5),
			want:  yamlconv.Ints(-43),
		}, {
			name:  "Single positive decimal with precision",
			input: yamlconv.Floats(42.555),
			params: []invocation.Parameter{
				invocationtest.Int(2),
			},
			want: yamlconv.Floats(42.56),
		}, {
			name:  "Single negative decimal with precision",
			input: yamlconv.Floats(-42.555),
			params: []invocation.Parameter{
				invocationtest.Int(2),
			},
			want: yamlconv.Floats(-42.56),
		}, {
			name:    "Input is non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Input is incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Parameter is non-singleton input",
			input:   yamlconv.Ints(1),
			params:  []invocation.Parameter{invocationtest.Ints(1, 2)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Parameter is incorrect kind",
			input: yamlconv.Ints(1),
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{}`)),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Parameter is incorrect tag",
			input:   yamlconv.Ints(1),
			params:  []invocation.Parameter{invocationtest.String("string")},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Round(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Round() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Round() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Single integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single positive decimal",
			input: yamlconv.Floats(42.5),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Single negative decimal",
			input: yamlconv.Floats(-42.5),
			want:  yamlconv.Ints(-42),
		}, {
			name:    "Input is non-singleton input",
			input:   yamlconv.Ints(1, 2),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Input is incorrect kind",
			input:   []*yaml.Node{yamltest.MustParseNode(`{}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is incorrect tag",
			input:   yamlconv.Strings("string"),
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Truncate(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Truncate() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Truncate() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Collection contains single integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Collection contains multiple integers",
			input: yamlconv.Ints(1, 2, 3),
			want:  yamlconv.Ints(3),
		}, {
			name:   "Parameters contains single integer",
			params: []invocation.Parameter{invocationtest.Ints(42)},
			want:   yamlconv.Ints(42),
		}, {
			name:   "Parameters contains multiple integers",
			params: []invocation.Parameter{invocationtest.Ints(42), invocationtest.Ints(1)},
			want:   yamlconv.Ints(42),
		}, {
			name:  "Collection and parameters contain integers",
			input: yamlconv.Ints(1, 2, 3),
			params: []invocation.Parameter{
				invocationtest.Int(4),
				invocationtest.Int(5),
			},
			want: yamlconv.Ints(5),
		}, {
			name: "Collection contains non-scalar",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Collection contains non-numeric",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrBadTag,
		}, {
			name: "Parameters contain non-scalar",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{}`)),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name: "Parameters contain non-numeric",
			params: []invocation.Parameter{
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Max(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Max() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Max() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Collection contains single integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Collection contains multiple integers",
			input: yamlconv.Ints(1, 2, 3),
			want:  yamlconv.Ints(1),
		}, {
			name:   "Parameters contains single integer",
			params: []invocation.Parameter{invocationtest.Ints(42)},
			want:   yamlconv.Ints(42),
		}, {
			name:   "Parameters contains multiple integers",
			params: []invocation.Parameter{invocationtest.Ints(42), invocationtest.Ints(1)},
			want:   yamlconv.Ints(1),
		}, {
			name:  "Collection and parameters contain integers",
			input: yamlconv.Ints(1, 2, 3),
			params: []invocation.Parameter{
				invocationtest.Int(4),
				invocationtest.Int(5),
			},
			want: yamlconv.Ints(1),
		}, {
			name: "Collection contains non-scalar",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Collection contains non-numeric",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrBadTag,
		}, {
			name: "Parameters contain non-scalar",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{}`)),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name: "Parameters contain non-numeric",
			params: []invocation.Parameter{
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Min(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Min() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Min() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Collection contains single integer",
			input: yamlconv.Ints(42),
			want:  yamlconv.Ints(42),
		}, {
			name:  "Collection contains multiple integers",
			input: yamlconv.Ints(1, 2, 3),
			want:  yamlconv.Ints(6),
		}, {
			name:   "Parameters contains single integer",
			params: []invocation.Parameter{invocationtest.Ints(42)},
			want:   yamlconv.Ints(42),
		}, {
			name:   "Parameters contains multiple integers",
			params: []invocation.Parameter{invocationtest.Ints(42), invocationtest.Ints(1)},
			want:   yamlconv.Ints(43),
		}, {
			name:  "Collection and parameters contain integers",
			input: yamlconv.Ints(1, 2, 3),
			params: []invocation.Parameter{
				invocationtest.Int(4),
				invocationtest.Int(5),
			},
			want: yamlconv.Ints(15),
		}, {
			name: "Collection contains non-scalar",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{}`),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Collection contains non-numeric",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrBadTag,
		}, {
			name: "Parameters contain non-scalar",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{}`)),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name: "Parameters contain non-numeric",
			params: []invocation.Parameter{
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrBadTag,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Sum(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Sum() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Sum() = diff (-got,want):\n%s", cmp.Diff(got, want, cmpopts.EquateEmpty()))
			}
		})
	}
}
