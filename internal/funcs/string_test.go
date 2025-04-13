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
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestLower(t *testing.T) {
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
		},
		{
			name:  "Input is string",
			input: []*yaml.Node{yamlconv.String("HELLO WORLD")},
			want:  []*yaml.Node{yamlconv.String("hello world")},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Lower(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Lower() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Lower() = %v, want %v", got, want)
			}
		})
	}
}

func TestUpper(t *testing.T) {
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
		},
		{
			name:  "Input is string",
			input: []*yaml.Node{yamlconv.String("hello world")},
			want:  []*yaml.Node{yamlconv.String("HELLO WORLD")},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Upper(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Upper() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Upper() = %v, want %v", got, want)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{},
		}, {
			name:  "Input matches prefix",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("hello")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Input does not match prefix",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Argument returns error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Argument is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.Bool(true)),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Argument is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Argument contains multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo"), yamlconv.String("bar")),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.StartsWith(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("StartsWith() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("StartsWith() = %v, want %v", got, want)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{},
		}, {
			name:  "Input matches suffix",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Input does not match suffix",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("hello")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Argument returns error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Argument is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.Bool(true)),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Argument is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Argument contains multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo"), yamlconv.String("bar")),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.EndsWith(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("EndsWith() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("EndsWith() = %v, want %v", got, want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  nil,
		}, {
			name:  "Input contains substring",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Input does not contain substring",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input is not a singleton",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Substring is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Substring is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Bool(true),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Substring is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Substring contains multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Contains(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Contains() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Contains() = %v, want %v", got, want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{},
		}, {
			name:  "Input contains substring",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("world")),
			want:  []*yaml.Node{yamlconv.Number(6)},
		}, {
			name:  "Input does not contain substring",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{yamlconv.Number(-1)},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input is not a singleton",
			input:   yamlconv.Strings("foo", "bar"),
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Substring is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Substring is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Bool(true),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Substring is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Substring is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IndexOf(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IndexOf() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("IndexOf() = %v, want %v", got, want)
			}
		})
	}
}

func TestSubstring(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		args    []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			args:  []invocation.Parameter{invocationtest.Int(0)},
			want:  []*yaml.Node{},
		}, {
			name:  "Only start index within string range",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args:  []invocation.Parameter{invocationtest.Int(6)},
			want:  []*yaml.Node{yamlconv.String("world")},
		}, {
			name:  "Only start index outside string range",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args:  []invocation.Parameter{invocationtest.Int(20)},
			want:  nil,
		}, {
			name:  "Start index and length within string range",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.Int(0),
				invocationtest.Int(5),
			},
			want: []*yaml.Node{yamlconv.String("hello")},
		}, {
			name:  "Start index and length, length exceeds string range",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.Int(6),
				invocationtest.Int(20),
			},
			want: []*yaml.Node{yamlconv.String("world")},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			args:    []invocation.Parameter{invocationtest.Int(0)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			args:    []invocation.Parameter{invocationtest.Int(0)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input contains multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			args:    []invocation.Parameter{invocationtest.Int(0)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Index returns error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Error(testErr)},
			wantErr: testErr,
		}, {
			name:    "Index is not an integer",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Index is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`))},
			wantErr: errs.ErrBadKind,
		}, {
			name:  "Index is not singleton",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.Numbers(0, 1, 2),
			},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Length is an error",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.Int(0),
				invocationtest.Error(testErr),
			},
			wantErr: testErr,
		}, {
			name:    "Length is not a number",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Int(0), invocationtest.String("foo")},
			wantErr: errs.ErrBadTag,
		}, {
			name:  "Length is not a scalar",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.Int(0),
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:  "Length is not singleton",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.Int(0),
				invocationtest.Ints(0, 1, 2),
			},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Substring(ctx, tc.args...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Substring() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Substring() = %v, want %v", got, want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		args    []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			args:  []invocation.Parameter{invocationtest.String("foo"), invocationtest.String("bar")},
			want:  []*yaml.Node{},
		}, {
			name:  "Input is string with match",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.String("hello"),
				invocationtest.String("foo"),
			},
			want: []*yaml.Node{yamlconv.String("foo world")},
		}, {
			name:  "Input is string without match",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.String("foo"),
				invocationtest.String("bar"),
			},
			want: []*yaml.Node{yamlconv.String("hello world")},
		}, {
			name:  "Input is not scalar node",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			args: []invocation.Parameter{
				invocationtest.String("hello"),
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrBadKind,
		}, {
			name:  "Input is non-string scalar node",
			input: []*yaml.Node{yamlconv.Bool(true)},
			args: []invocation.Parameter{
				invocationtest.String("hello"),
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrBadTag,
		}, {
			name:  "Input contains multiple elements",
			input: []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			args: []invocation.Parameter{
				invocationtest.String("hello"),
				invocationtest.String("foo"),
			},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Pattern is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Error(testErr)},
			wantErr: testErr,
		}, {
			name:    "Pattern is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Pattern is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`))},
			wantErr: errs.ErrBadKind,
		}, {
			name:  "Pattern is not singleton",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.Strings("foo", "bar"),
				invocationtest.String("bar"),
			},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Replacement is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.Error(testErr)},
			wantErr: testErr,
		}, {
			name:    "Replacement is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Replacement is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`))},
			wantErr: errs.ErrBadKind,
		}, {
			name:  "Replacement is not singleton",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.String("foo"),
				invocationtest.Strings("foo", "bar"),
			},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Replace(ctx, tc.args...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Replace() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Replace() = %v, want %v", got, want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.String("foo"),
			want:  []*yaml.Node{},
		}, {
			name:  "Input is string with match",
			input: yamlconv.Strings("hello world"),
			arg:   invocationtest.String(" "),
			want:  yamlconv.Strings("hello", "world"),
		}, {
			name:  "Input is string without match",
			input: yamlconv.Strings("hello world"),
			arg:   invocationtest.String("foo"),
			want:  []*yaml.Node{yamlconv.String("hello world")},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input is not singleton",
			input:   yamlconv.Strings("foo", "bar"),
			arg:     invocationtest.String("foo"),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Separator is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Separator is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Bool(true),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Separator is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Separator is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Split(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Split() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Split() = %v, want %v", got, want)
			}
		})
	}
}

func TestLength(t *testing.T) {
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
			name:  "Input is string",
			input: yamlconv.Strings("hello world"),
			want:  yamlconv.Ints(11),
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   yamlconv.Bools(true),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input contains multiple elements",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Length(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Length() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Length() = %v, want %v", got, want)
			}
		})
	}
}

func TestToChars(t *testing.T) {
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
			name:  "Input is string",
			input: yamlconv.Strings("hello world"),
			want:  yamlconv.Strings("h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"),
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   yamlconv.Bools(true),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input contains multiple elements",
			input:   yamlconv.Strings("foo", "bar"),
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ToChars(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ToChars() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ToChars() = %v, want %v", got, want)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		arg     invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{},
		}, {
			name:  "Input matches regex",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("^h.*o")),
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Input does not match regex",
			input: []*yaml.Node{yamlconv.String("hello world")},
			arg:   invocationtest.SuccessParameter(yamlconv.String("foo")),
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo")),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Pattern returns error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.Error(testErr),
			wantErr: testErr,
		}, {
			name:    "Pattern is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.Bool(true)),
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Pattern is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`)),
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Pattern is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("foo"), yamlconv.String("bar")),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Pattern is invalid regex",
			input:   []*yaml.Node{yamlconv.String("foo")},
			arg:     invocationtest.SuccessParameter(yamlconv.String("[a-z")),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Matches(ctx, tc.arg)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Matches() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Matches() = diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}

func TestReplaceMatches(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		args    []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty input",
			input: []*yaml.Node{},
			args:  []invocation.Parameter{invocationtest.String("foo"), invocationtest.String("bar")},
			want:  []*yaml.Node{},
		}, {
			name:  "Input is string with match",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.String("^h.*lo"),
				invocationtest.String("foo"),
			},
			want: []*yaml.Node{yamlconv.String("foo world")},
		}, {
			name:  "Input is string without match",
			input: []*yaml.Node{yamlconv.String("hello world")},
			args: []invocation.Parameter{
				invocationtest.String("foo"),
				invocationtest.String("bar"),
			},
			want: []*yaml.Node{yamlconv.String("hello world")},
		}, {
			name:    "Input is not scalar node",
			input:   []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)},
			args:    []invocation.Parameter{invocationtest.String("hello"), invocationtest.String("foo")},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Input is non-string scalar node",
			input:   []*yaml.Node{yamlconv.Bool(true)},
			args:    []invocation.Parameter{invocationtest.String("hello"), invocationtest.String("foo")},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Input contains multiple elements",
			input:   []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			args:    []invocation.Parameter{invocationtest.String("hello"), invocationtest.String("foo")},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:    "Pattern returns error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Error(testErr)},
			wantErr: testErr,
		}, {
			name:    "Pattern is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Pattern is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`))},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Pattern is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.Strings("foo", "bar"), invocationtest.String("bar")},
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Pattern is invalid regex",
			input: []*yaml.Node{yamlconv.String("foo")},
			args: []invocation.Parameter{
				invocationtest.String("[a-z"),
				invocationtest.String("bar"),
			},
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Replacement is an error",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.Error(testErr)},
			wantErr: testErr,
		}, {
			name:    "Replacement is not a string",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.Bool(true)},
			wantErr: errs.ErrBadTag,
		}, {
			name:    "Replacement is not a scalar",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.SuccessParameter(yamltest.MustParseNode(`{"foo": "bar"}`))},
			wantErr: errs.ErrBadKind,
		}, {
			name:    "Replacement is not singleton",
			input:   []*yaml.Node{yamlconv.String("foo")},
			args:    []invocation.Parameter{invocationtest.String("foo"), invocationtest.Strings("foo", "bar")},
			wantErr: errs.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.ReplaceMatches(ctx, tc.args...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ReplaceMatches() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("ReplaceMatches() diff (-got,+want):\n%s", cmp.Diff(got, want))
			}
		})
	}
}
