package funcs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/expr"
	"rodusek.dev/pkg/yamlpath/internal/funcs"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

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

func TestChildren(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: nil,
			want:  []*yaml.Node{},
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Strings("hello", "world"),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Singleton non-mapping node",
			input: []*yaml.Node{yamlconv.Sequence(yamlconv.Strings("hello", "world")...)},
			want:  nil,
		}, {
			name:  "Singleton mapping node",
			input: []*yaml.Node{yamltest.MustParseNode(`{"foo": "bar", "baz": "qux"}`)},
			want:  []*yaml.Node{yamlconv.String("bar"), yamlconv.String("qux")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Children(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Children() error = %v, want %v", got, want)
			}
			opts := []cmp.Option{cmpopts.EquateEmpty(), yamltest.IgnoreMetaFields()}
			if got, want := got, tc.want; !cmp.Equal(got, want, opts...) {
				t.Errorf("Children() = diff (-got, +want):\n%s", cmp.Diff(got, want, opts...))
			}
		})
	}
}

func TestDescendants(t *testing.T) {
	testCases := []struct {
		name    string
		input   []*yaml.Node
		want    []*yaml.Node
		wantErr error
	}{
		{
			name:  "Empty collection",
			input: nil,
			want:  nil,
		}, {
			name:    "Non-singleton collection",
			input:   yamlconv.Strings("hello", "world"),
			wantErr: errs.ErrNotSingleton,
		}, {
			name:  "Scalar value",
			input: yamlconv.Strings("hello"),
			want:  nil,
		}, {
			name:  "Sequence node",
			input: []*yaml.Node{yamlconv.Sequence(yamlconv.Strings("hello", "world")...)},
			want: []*yaml.Node{
				yamlconv.String("hello"),
				yamlconv.String("world"),
			},
		}, {
			name: "Mapping node",
			input: []*yaml.Node{
				yamltest.MustParseNode(`{"foo": 42, "bar": {"baz": "qux"}, "seq": ["val"]}`),
			},
			want: []*yaml.Node{
				yamlconv.Int(42),
				yamltest.MustParseNode(`{"baz": "qux"}`),
				yamlconv.String("qux"),
				yamlconv.Sequence(yamlconv.Strings("val")...),
				yamlconv.String("val"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := expr.NewContext(tc.input)

			got, err := funcs.Descendants(ctx)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("Descendants() error = %v, want %v", got, want)
			}
			opts := []cmp.Option{cmpopts.EquateEmpty(), yamltest.IgnoreMetaFields()}
			if got, want := got, tc.want; !cmp.Equal(got, want, opts...) {
				t.Errorf("Descendants() = diff (-got, +want):\n%s", cmp.Diff(got, want, opts...))
			}
		})
	}
}
