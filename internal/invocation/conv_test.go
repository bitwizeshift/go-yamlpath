package invocation_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestParseString(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		context invocation.Context
		param   invocation.Parameter
		want    string
		wantErr error
	}{
		{
			name:    "Parameter contains string",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.String("Hello"),
			want:    "Hello",
		}, {
			name:    "Parameter contains non-string",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Int(0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := invocation.ParseString(tc.context, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseString() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		context invocation.Context
		param   invocation.Parameter
		want    int64
		wantErr error
	}{
		{
			name:    "Parameter contains int",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Int(1),
			want:    1,
		}, {
			name:    "Parameter contains non-int",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := invocation.ParseInt(tc.context, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseInt() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseInt() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		context invocation.Context
		param   invocation.Parameter
		want    bool
		wantErr error
	}{
		{
			name:    "Parameter contains bool",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Bool(true),
			want:    true,
		}, {
			name:    "Parameter contains non-bool",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := invocation.ParseBool(tc.context, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseBool() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseBool() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		context invocation.Context
		param   invocation.Parameter
		want    float64
		wantErr error
	}{
		{
			name:    "Parameter contains float",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Float(1.0),
			want:    1.0,
		}, {
			name:    "Parameter contains non-float",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := invocation.ParseFloat(tc.context, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseFloat() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseFloat() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseDecimal(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		context invocation.Context
		param   invocation.Parameter
		want    decimal.Decimal
		wantErr error
	}{
		{
			name:    "Parameter contains decimal",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Int(1),
			want:    decimal.New(1, 0),
		}, {
			name:    "Parameter contains non-decimal",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: expr.NewContext(yamlconv.Strings("test")),
			param:   invocationtest.Error(testErr),
			wantErr: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := invocation.ParseDecimal(tc.context, tc.param)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("ParseDecimal() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseDecimal() = %v, want %v", got, tc.want)
			}
		})
	}
}
