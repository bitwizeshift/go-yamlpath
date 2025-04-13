package invocation_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shopspring/decimal"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

type testContext struct {
	err   error
	nodes []*yaml.Node
}

func (c *testContext) Root() []*yaml.Node {
	return c.nodes
}
func (c *testContext) Current() []*yaml.Node {
	return c.nodes
}
func (c *testContext) NewContext([]*yaml.Node) invocation.Context {
	return &testContext{
		nodes: c.nodes,
		err:   c.err,
	}
}
func (c *testContext) WithContext(context.Context) invocation.Context {
	return nil
}
func (c *testContext) Context() context.Context {
	return nil
}
func goodContext(nodes ...*yaml.Node) invocation.Context {
	return &testContext{
		nodes: nodes,
	}
}

func errContext(err error) invocation.Context {
	return &testContext{
		err: err,
	}
}

var _ invocation.Context = (*testContext)(nil)

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
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.String("Hello"),
			want:    "Hello",
		}, {
			name:    "Parameter contains non-string",
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.Int(0),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: goodContext(yamlconv.String("test")),
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
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.Int(1),
			want:    1,
		}, {
			name:    "Parameter contains non-int",
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: goodContext(yamlconv.String("test")),
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
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.Bool(true),
			want:    true,
		}, {
			name:    "Parameter contains non-bool",
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: goodContext(yamlconv.String("test")),
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
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.Float(1.0),
			want:    1.0,
		}, {
			name:    "Parameter contains non-float",
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: goodContext(yamlconv.String("test")),
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
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.Int(1),
			want:    decimal.New(1, 0),
		}, {
			name:    "Parameter contains non-decimal",
			context: goodContext(yamlconv.String("test")),
			param:   invocationtest.String("foo"),
			wantErr: cmpopts.AnyError,
		}, {
			name:    "Parameter contains error",
			context: goodContext(yamlconv.String("test")),
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
