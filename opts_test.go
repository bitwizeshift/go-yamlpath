package yamlpath_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestWithFunction_Error(t *testing.T) {
	testCases := []struct {
		name    string
		fn      any
		wantErr error
	}{
		{
			name:    "not a function",
			fn:      1,
			wantErr: errs.ErrNotAFunction,
		}, {
			name:    "function has too few arguments",
			fn:      func() {},
			wantErr: errs.ErrFuncTooFewArguments,
		}, {
			name:    "function has bad first argument",
			fn:      func(error) {},
			wantErr: errs.ErrBadFirstArgument,
		}, {
			name:    "function has too few return types",
			fn:      func(yamlpath.Collection) {},
			wantErr: errs.ErrBadReturnSignature,
		}, {
			name: "function has too many return types",
			fn: func(yamlpath.Collection) (yamlpath.Collection, error, error) {
				return nil, nil, nil
			},
			wantErr: errs.ErrBadReturnSignature,
		}, {
			name: "function has return with wrong second type",
			fn: func(yamlpath.Collection) (yamlpath.Collection, int) {
				return nil, 0
			},
			wantErr: errs.ErrBadReturnSignature,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			opts := []yamlpath.Option{
				yamlpath.WithFunction("test", tc.fn),
			}

			// Act
			_, err := yamlpath.Compile("$.test()", opts...)

			// Assert
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("WithFunction(%q) err = %v, want %v", tc.name, got, want)
			}
		})
	}
}

func TestWithFunction_ValidFunction(t *testing.T) {
	simpleNode := yamltest.MustParseNode(`{"foo": "bar"}`)
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   *yaml.Node
		path    string
		fn      any
		want    yamlpath.Collection
		wantErr error
	}{
		{
			name: "simple function, no error",
			path: "$.test()",
			fn: func(c yamlpath.Collection) yamlpath.Collection {
				return c
			},
			want:    yamlpath.Collection{simpleNode},
			wantErr: nil,
		}, {
			name: "simple function with error, returns value",
			path: "$.test()",
			fn: func(c yamlpath.Collection) (yamlpath.Collection, error) {
				return c, nil
			},
			want: yamlpath.Collection{simpleNode},
		}, {
			name: "simple function with error, returns error",
			path: "$.test()",
			fn: func(yamlpath.Collection) (yamlpath.Collection, error) {
				return nil, testErr
			},
			want:    nil,
			wantErr: testErr,
		}, {
			name: "function with collection argument",
			path: "$.test(@)",
			fn: func(_ yamlpath.Collection, in yamlpath.Collection) yamlpath.Collection {
				return in
			},
			want: yamlpath.Collection{simpleNode},
		}, {
			name: "function with scalar bool return type",
			path: "$.test()",
			fn: func(yamlpath.Collection) bool {
				return true
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`true`),
			},
		}, {
			name: "function with scalar int return type",
			path: "$.test()",
			fn: func(yamlpath.Collection) int {
				return 1
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`1`),
			},
		}, {
			name: "function with scalar string return type",
			path: "$.test()",
			fn: func(yamlpath.Collection) string {
				return "test"
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`test`),
			},
		}, {
			name: "function with slice bool return type",
			path: "$.test()",
			fn: func(yamlpath.Collection) []bool {
				return []bool{true, false}
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`[true, false]`),
			},
		}, {
			name: "function with slice int return type",
			path: "$.test()",
			fn: func(yamlpath.Collection) []int {
				return []int{1, 2}
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`[1, 2]`),
			},
		}, {
			name: "function with int argument",
			path: "$.test(1)",
			fn: func(yamlpath.Collection, int) yamlpath.Collection {
				return yamlpath.Collection{yamltest.MustParseNode(`1`)}
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`1`),
			},
		}, {
			name: "function with bool argument",
			path: "$.test(true)",
			fn: func(yamlpath.Collection, bool) yamlpath.Collection {
				return yamlpath.Collection{yamltest.MustParseNode(`true`)}
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`true`),
			},
		}, {
			name: "function with variadic int parameters, no argument",
			path: "$.test()",
			fn: func(yamlpath.Collection, ...int) yamlpath.Collection {
				return yamlpath.Collection{}
			},
			want: yamlpath.Collection{},
		}, {
			name: "function with variadic int parameters, one argument",
			path: "$.test(1)",
			fn: func(_ yamlpath.Collection, v ...int) int {
				return v[0]
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`1`),
			},
		}, {
			name: "function with variadic int parameters, multiple arguments",
			path: "$.test(1, 2, 3)",
			fn: func(_ yamlpath.Collection, v ...int) int {
				return v[0] + v[1] + v[2]
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`6`),
			},
		}, {
			name: "function with slice parameter",
			path: "$.test([1, 2, 3])",
			fn: func(_ yamlpath.Collection, v []int) int {
				return v[0] + v[1] + v[2]
			},
			want: yamlpath.Collection{
				yamltest.MustParseNode(`6`),
			},
		}, {
			name: "function with int parameter, non-int input",
			path: "$.test(false)",
			fn: func(_ yamlpath.Collection, v int) int {
				return v
			},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			opts := []yamlpath.Option{
				yamlpath.WithFunction("test", tc.fn),
			}
			yp := yamlpath.MustCompile(tc.path, opts...)

			// Act
			got, err := yp.Match(simpleNode)

			// Assert
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("WithFunction() err = %v, want %v", got, want)
			}
			if diff := cmp.Diff(got, tc.want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("WithFunction() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestWithConstant(t *testing.T) {
	testCases := []struct {
		name string
		opt  yamlpath.Option
		want []*yaml.Node
	}{
		{
			name: "string constant",
			opt:  yamlpath.WithConstant("key", "value"),
			want: yamlconv.Strings("value"),
		}, {
			name: "int constant",
			opt:  yamlpath.WithConstant("key", 1),
			want: yamlconv.Ints(1),
		}, {
			name: "uint constant",
			opt:  yamlpath.WithConstant("key", uint(1)),
			want: yamlconv.Ints(1),
		}, {
			name: "float constant",
			opt:  yamlpath.WithConstant("key", 1.1),
			want: yamlconv.Floats(1.1),
		}, {
			name: "bool constant",
			opt:  yamlpath.WithConstant("key", true),
			want: yamlconv.Bools(true),
		}, {
			name: "node value",
			opt:  yamlpath.WithConstant("key", yamltest.MustParseNode(`{"foo": "bar"}`)),
			want: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": "bar"}`),
			},
		}, {
			name: "slice of nodes",
			opt:  yamlpath.WithConstant("key", []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar"}`)}),
			want: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": "bar"}`),
			},
		}, {
			name: "collection of nodes",
			opt:  yamlpath.WithConstant("key", yamlpath.Collection{yamltest.MustParseNode(`{"foo": "bar"}`)}),
			want: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": "bar"}`),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			yp := yamlpath.MustCompile("%key", tc.opt)
			want := yamlpath.Collection(tc.want)

			// Act
			got, err := yp.Match(nil)

			// Assert
			if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("WithExternalConstant() error = %v", err)
			}
			if diff := cmp.Diff(got, want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("WithExternalConstant() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestWithConstants_ConstantAlreadyDefined_ReturnsError(t *testing.T) {
	opts := []yamlpath.Option{
		yamlpath.WithConstant("key", "value"),
		yamlpath.WithConstant("key", "value2"),
	}

	yp, err := yamlpath.Compile("%key", opts...)

	if got, want := err, cmpopts.AnyError; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("WithExternalConstant() error = %v, want %v", got, want)
	}
	if got, want := yp, (*yamlpath.YAMLPath)(nil); !cmp.Equal(got, want) {
		t.Fatalf("WithExternalConstant() yp = %v, want %v", got, want)
	}
}
