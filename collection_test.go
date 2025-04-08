package yamlpath_test

import (
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestCollection_Nodes(t *testing.T) {
	collection := yamlpath.Collection{yamlconv.Bool(true)}

	want := ([]*yaml.Node)(collection)
	got := collection.Nodes()

	if unsafe.Pointer(&got[0]) != unsafe.Pointer(&want[0]) {
		t.Errorf("Collection.Nodes() = %v, want %v", got, want)
	}
}

func TestCollection_IsTruthy(t *testing.T) {
	testCases := []struct {
		name  string
		input yamlpath.Collection
		want  bool
	}{
		{
			name:  "Empty collection is falsey",
			input: yamlpath.Collection{},
			want:  false,
		}, {
			name:  "Single non-bool node is truthy",
			input: yamlpath.Collection{yamlconv.String("hello")},
			want:  true,
		}, {
			name:  "Single true node is truthy",
			input: yamlpath.Collection{yamlconv.Bool(true)},
			want:  true,
		}, {
			name:  "Single false node is falsey",
			input: yamlpath.Collection{yamlconv.Bool(false)},
			want:  false,
		}, {
			name:  "Multiple nodes are truthy",
			input: yamlpath.Collection{yamlconv.Bool(true), yamlconv.Bool(false)},
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.IsTruthy()

			if got != tc.want {
				t.Errorf("Collection.IsTruthy() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_IsEmpty(t *testing.T) {
	testCases := []struct {
		name  string
		input yamlpath.Collection
		want  bool
	}{
		{
			name:  "Empty collection is empty",
			input: yamlpath.Collection{},
			want:  true,
		}, {
			name:  "Non-empty collection is not empty",
			input: yamlpath.Collection{yamlconv.String("hello")},
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.input.IsEmpty()

			if got != tc.want {
				t.Errorf("Collection.IsEmpty() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Singleton(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    *yaml.Node
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single node returns node",
			input: yamlpath.Collection{yamlconv.String("hello")},
			want:  yamlconv.String("hello"),
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.String("hello"), yamlconv.String("world")},
			wantErr: yamlpath.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Singleton()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Collection.Singleton() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !yamlcmp.Equal(got, want) {
				t.Errorf("Collection.Singleton() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_SingletonString(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    string
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single node returns string",
			input: yamlpath.Collection{yamlconv.String("hello")},
			want:  "hello",
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.String("hello"), yamlconv.String("world")},
			wantErr: yamlpath.ErrNotSingleton,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.SingletonString()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonString() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_SingletonBool(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    bool
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single true node returns true",
			input: yamlpath.Collection{yamlconv.Bool(true)},
			want:  true,
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.Bool(true), yamlconv.Bool(false)},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:    "Single non-bool node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.SingletonBool()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonBool() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonBool() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_SingletonInt(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    int
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single node returns int",
			input: yamlpath.Collection{yamlconv.Number(42)},
			want:  42,
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.Number(42), yamlconv.Number(42)},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:    "Single non-int node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.SingletonInt()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonInt() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonInt() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_SingletonFloat64(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    float64
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single node returns float64",
			input: yamlpath.Collection{yamlconv.Number(42)},
			want:  42,
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.Number(42), yamlconv.Number(42)},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:    "Single non-float64 node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.SingletonFloat64()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonFloat64() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonFloat64() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_SingletonDecimal(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    decimal.Decimal
		wantErr error
	}{
		{
			name:    "Empty collection returns error",
			input:   yamlpath.Collection{},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:  "Single node returns decimal",
			input: yamlpath.Collection{yamlconv.Number(42.69)},
			want:  decimal.New(4269, -2),
		}, {
			name:    "Multiple nodes returns error",
			input:   yamlpath.Collection{yamlconv.Number(42), yamlconv.Number(42)},
			wantErr: yamlpath.ErrNotSingleton,
		}, {
			name:    "Single non-decimal node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.SingletonDecimal()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.SingletonDecimal() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("Collection.SingletonDecimal() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Strings(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    []string
		wantErr error
	}{
		{
			name:    "Empty collection returns empty slice",
			input:   yamlpath.Collection{},
			want:    []string{},
			wantErr: nil,
		}, {
			name:    "Single node returns slice",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			want:    []string{"hello"},
			wantErr: nil,
		}, {
			name:    "Multiple nodes returns slice",
			input:   yamlpath.Collection{yamlconv.String("hello"), yamlconv.String("world")},
			want:    []string{"hello", "world"},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Strings()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Strings() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Collection.Strings() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Bools(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    []bool
		wantErr error
	}{
		{
			name:    "Empty collection returns empty slice",
			input:   yamlpath.Collection{},
			want:    []bool{},
			wantErr: nil,
		}, {
			name:    "Single node returns slice",
			input:   yamlpath.Collection{yamlconv.Bool(true)},
			want:    []bool{true},
			wantErr: nil,
		}, {
			name:    "Multiple nodes returns slice",
			input:   yamlpath.Collection{yamlconv.Bool(true), yamlconv.Bool(false)},
			want:    []bool{true, false},
			wantErr: nil,
		}, {
			name:    "Single non-bool node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Bools()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Bools() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Collection.Bools() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Ints(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    []int
		wantErr error
	}{
		{
			name:    "Empty collection returns empty slice",
			input:   yamlpath.Collection{},
			want:    []int{},
			wantErr: nil,
		}, {
			name:    "Single node returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42)},
			want:    []int{42},
			wantErr: nil,
		}, {
			name:    "Multiple nodes returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42), yamlconv.Number(69)},
			want:    []int{42, 69},
			wantErr: nil,
		}, {
			name:    "Single non-int node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Ints()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Ints() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Collection.Ints() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Float64s(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    []float64
		wantErr error
	}{
		{
			name:    "Empty collection returns empty slice",
			input:   yamlpath.Collection{},
			want:    []float64{},
			wantErr: nil,
		}, {
			name:    "Single node returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42)},
			want:    []float64{42},
			wantErr: nil,
		}, {
			name:    "Multiple nodes returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42), yamlconv.Number(69)},
			want:    []float64{42, 69},
			wantErr: nil,
		}, {
			name:    "Single non-float64 node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Float64s()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Float64s() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Collection.Float64s() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCollection_Decimals(t *testing.T) {
	testCases := []struct {
		name    string
		input   yamlpath.Collection
		want    []decimal.Decimal
		wantErr error
	}{
		{
			name:    "Empty collection returns empty slice",
			input:   yamlpath.Collection{},
			want:    []decimal.Decimal{},
			wantErr: nil,
		}, {
			name:    "Single node returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42.69)},
			want:    []decimal.Decimal{decimal.New(4269, -2)},
			wantErr: nil,
		}, {
			name:    "Multiple nodes returns slice",
			input:   yamlpath.Collection{yamlconv.Number(42.69), yamlconv.Number(69.42)},
			want:    []decimal.Decimal{decimal.New(4269, -2), decimal.New(6942, -2)},
			wantErr: nil,
		}, {
			name:    "Single non-decimal node returns error",
			input:   yamlpath.Collection{yamlconv.String("hello")},
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.input.Decimals()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Collection.Decimals() error = %v, want %v", err, tc.wantErr)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
				t.Errorf("Collection.Decimals() = %v, want %v", got, tc.want)
			}
		})
	}
}
