package funcs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestIsString(t *testing.T) {
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
			name:  "Single string input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-string input",
			input: []*yaml.Node{yamlconv.Number(42)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.String("foo"),
				yamlconv.Number(42),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsString(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsString() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsString() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsInteger(t *testing.T) {
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
			name:  "Single integer input",
			input: []*yaml.Node{yamlconv.Number(42)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-integer input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.Number(42),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsInteger(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsInteger() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsInteger() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
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
			name:  "Single float input",
			input: []*yaml.Node{yamlconv.Number(42.0)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-float input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.Number(42.0),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsFloat(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsFloat() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsFloat() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
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
			name:  "Single integer input",
			input: []*yaml.Node{yamlconv.Int(42)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single float input",
			input: []*yaml.Node{yamlconv.Float(42.0)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-number input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.Number(42),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsNumber(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsNumber() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsNumber() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsBoolean(t *testing.T) {
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
			name:  "Single boolean input",
			input: []*yaml.Node{yamlconv.Bool(true)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-boolean input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.Bool(true),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsBoolean(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsBoolean() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsBoolean() = %v, want %v", got, want)
			}
		})
	}
}
func TestIsNull(t *testing.T) {
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
			name:  "Single null input",
			input: []*yaml.Node{yamlconv.Null()},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-null input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamlconv.Null(),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsNull(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsNull() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsNull() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsScalar(t *testing.T) {
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
			name:  "Single scalar input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-scalar input",
			input: []*yaml.Node{yamltest.MustParseNode(`[1,2,3]`)},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs",
			input: []*yaml.Node{
				yamlconv.String("foo"),
				yamlconv.String("bar"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsScalar(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsScalar() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsScalar() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsSequence(t *testing.T) {
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
			name:  "Single sequence input",
			input: []*yaml.Node{yamltest.MustParseNode(`[1,2,3]`)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-sequence input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamltest.MustParseNode(`[1,2,3]`),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsSequence(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsSequence() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsSequence() = %v, want %v", got, want)
			}
		})
	}
}

func TestIsMapping(t *testing.T) {
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
			name:  "Single mapping input",
			input: []*yaml.Node{yamltest.MustParseNode(`{foo: bar}`)},
			want:  []*yaml.Node{yamlconv.Bool(true)},
		}, {
			name:  "Single non-mapping input",
			input: []*yaml.Node{yamlconv.String("foo")},
			want:  []*yaml.Node{yamlconv.Bool(false)},
		}, {
			name: "Multiple inputs with mixed types",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{foo: bar}`),
				yamlconv.String("foo"),
			},
			want: []*yaml.Node{
				yamlconv.Bool(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.IsMapping(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("IsMapping() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("IsMapping() = %v, want %v", got, want)
			}
		})
	}
}
