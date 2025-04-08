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
