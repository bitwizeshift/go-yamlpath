package yamlutil_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestIsTruthy(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  bool
	}{
		{
			name:  "Empty node is falsy",
			input: nodes(),
			want:  false,
		}, {
			name:  "Single non-bool node is truthy",
			input: nodes(yamlutil.String("hello")),
			want:  true,
		}, {
			name:  "Single true node is truthy",
			input: nodes(yamlutil.True),
			want:  true,
		}, {
			name:  "Single false node is falsy",
			input: nodes(yamlutil.False),
			want:  false,
		}, {
			name:  "Multiple nodes are truthy",
			input: nodes(yamlutil.True, yamlutil.False),
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := yamlutil.IsTruthy(tc.input...)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	testCases := []struct {
		name    string
		node    *yaml.Node
		want    bool
		wantErr error
	}{
		{
			name: "Scalar node with bool kind",
			node: yamlutil.True,
			want: true,
		}, {
			name:    "Scalar node with string kind",
			node:    yamlutil.String("true"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Scalar bool node with invalid bool value",
			node:    yamlutil.Boolean("hello"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.ToBool(tc.node)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ToBool() = %v; want %v", got, want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ToBool() error = %v; want %v", err, tc.wantErr)
			}
		})
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		name    string
		node    *yaml.Node
		want    string
		wantErr error
	}{
		{
			name: "Scalar node with string kind",
			node: yamlutil.String("hello"),
			want: "hello",
		}, {
			name:    "Scalar node with bool kind",
			node:    yamlutil.True,
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.ToString(tc.node)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ToString() = %v; want %v", got, want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ToString() error = %v; want %v", tc.wantErr, nil)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	testCases := []struct {
		name    string
		node    *yaml.Node
		want    int
		wantErr error
	}{
		{
			name: "Scalar node with int kind",
			node: yamlutil.Number("42"),
			want: 42,
		}, {
			name:    "Scalar node with string kind",
			node:    yamlutil.String("42"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.ToInt(tc.node)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ToInt() = %v; want %v", got, want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ToInt() error = %v; want %v", tc.wantErr, nil)
			}
		})
	}
}

func TestToDecimal(t *testing.T) {
	testCases := []struct {
		name    string
		node    *yaml.Node
		want    decimal.Decimal
		wantErr error
	}{
		{
			name: "Scalar node with int kind",
			node: yamlutil.Number("42"),
			want: decimal.New(42, 0),
		}, {
			name: "Scalar node with float kind",
			node: yamlutil.Number("42.5"),
			want: decimal.New(425, -1),
		}, {
			name:    "Scalar node with string kind",
			node:    yamlutil.String("42"),
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlutil.ToDecimal(tc.node)

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ToDecimal() = %v; want %v", got, want)
			}
			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ToDecimal() error = %v; want %v", tc.wantErr, nil)
			}
		})
	}
}
