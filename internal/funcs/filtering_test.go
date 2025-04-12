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

func TestWhere(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Criteria parameter returns true",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Bool(true)).AddSuccess(yamlconv.Bool(false)),
			},
			input: []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			want:  []*yaml.Node{yamlconv.String("foo")},
		}, {
			name: "Criteria parameter returns false",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Bool(false)),
			},
			input: []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			want:  []*yaml.Node{},
		}, {
			name: "Criteria parameter returns error",
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

			got, err := funcs.Where(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Where() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Where() = %v, want %v", got, want)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		input   []*yaml.Node
		params  []invocation.Parameter
		want    []*yaml.Node
		wantErr error
	}{
		{
			name: "Projection parameter returns elements",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.String("hello")).AddSuccess(yamlconv.String("world")),
			},
			input: []*yaml.Node{yamlconv.String("foo"), yamlconv.String("bar")},
			want:  []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")},
		}, {
			name: "Projection parameter returns error",
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

			got, err := funcs.Transform(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Select() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Select() = %v, want %v", got, want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	testCases := []struct {
		name  string
		input []*yaml.Node
		want  []*yaml.Node
	}{
		{
			name:  "Keys of a mapping",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar", "baz": "qux"}`)},
			want:  []*yaml.Node{yamlconv.String("foo"), yamlconv.String("baz")},
		}, {
			name:  "Keys of a sequence",
			input: []*yaml.Node{yamltest.MustParseNode(`["foo", "bar", "baz", "qux"]`)},
			want:  []*yaml.Node{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Keys(ctx)

			if err != nil {
				t.Fatalf("Keys() error = %v", err)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Keys() = %v, want %v", got, want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name    string
		params  []invocation.Parameter
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection returns empty",
			input: []*yaml.Node{},
			want:  []*yaml.Node{},
		}, {
			name:  "Multiple elements, no parameters returns empty",
			input: []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")},
			want:  []*yaml.Node{},
		}, {
			name: "Parameter returns error",
			params: []invocation.Parameter{
				invocationtest.ErrorParameter(testErr),
			},
			input:   []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")},
			wantErr: testErr,
		}, {
			name: "Multiple elements, one parameter returns nothing",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(),
			},
			input: []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")},
		}, {
			name: "Multiple elements, one parameter returning single string",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.String("key")),
			},
			input: []*yaml.Node{yamltest.MustParseNode(`{"key": "value"}`)},
			want:  []*yaml.Node{yamlconv.String("value")},
		}, {
			name: "Multiple elements, one parameter returning single number",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Number(1)),
				invocationtest.SuccessParameter(yamlconv.Number(2)),
			},
			input: []*yaml.Node{yamltest.MustParseNode(`[1,2,3]`)},
			want:  []*yaml.Node{yamlconv.Number(2), yamlconv.Number(3)},
		}, {
			name: "Multiple elements, parameter returns non-scalar",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamltest.MustParseNode(`{"key": "value"}`)),
			},
			input:   []*yaml.Node{yamltest.MustParseNode(`{"key": "value"}`)},
			wantErr: errs.ErrBadKind,
		}, {
			name: "Multiple elements, parameter returns scalar with bad tag",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Bool(true)),
			},
			input:   []*yaml.Node{yamltest.MustParseNode(`{"key": "value"}`)},
			wantErr: errs.ErrBadTag,
		}, {
			name: "Multiple elements, parameter returns multiple values",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.String("key"), yamlconv.String("value")),
			},
			input:   []*yaml.Node{yamltest.MustParseNode(`{"key": "value"}`)},
			wantErr: errs.ErrNotSingleton,
		}, {
			name: "Multiple elements, parameter returns negative int",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Number(-1)),
			},
			input: []*yaml.Node{yamltest.MustParseNode(`["hello", "world"]`)},
			want:  []*yaml.Node{yamlconv.String("world")},
		}, {
			name: "Multiple elements, parameter returns index out of range",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.Number(2)),
			},
			input: []*yaml.Node{yamltest.MustParseNode(`["hello", "world"]`)},
			want:  []*yaml.Node{},
		}, {
			name: "invalid integer definition",
			params: []invocation.Parameter{
				invocationtest.SuccessParameter(yamlconv.NumberString("bad")),
			},
			input:   []*yaml.Node{yamltest.MustParseNode(`["hello", "world"]`)},
			wantErr: errs.ErrEval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Select(ctx, tc.params...)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Select() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !yamlcmp.EqualRange(got, want) {
				t.Errorf("Select() = %v, want %v", got, want)
			}
		})
	}
}
