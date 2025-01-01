package errs_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func TestKindError_Is(t *testing.T) {
	var err errs.KindError

	if got, want := &err, errs.ErrBadKind; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("KindError.Unwrap() error = %v, want %v", got, want)
	}
	if got, want := &err, errs.ErrEval; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("KindError.Unwrap() error = %v, want %v", got, want)
	}
}

func TestKindError_Error(t *testing.T) {
	testCases := []struct {
		name     string
		source   string
		expected []yaml.Kind
		node     *yaml.Node
		want     string
	}{
		{
			name:     "single kind",
			source:   "source",
			expected: []yaml.Kind{yaml.ScalarNode},
			node:     &yaml.Node{Kind: yaml.DocumentNode},
			want:     "yamlpath eval: source expected YAML node with kind 'scalar', but received 'document'",
		}, {
			name:     "two kinds",
			source:   "source",
			expected: []yaml.Kind{yaml.ScalarNode, yaml.MappingNode},
			node:     &yaml.Node{Kind: yaml.SequenceNode},
			want:     "yamlpath eval: source expected YAML node with kind 'scalar' or 'mapping', but received 'sequence'",
		}, {
			name:   "three kinds",
			source: "source",
			expected: []yaml.Kind{
				yaml.ScalarNode,
				yaml.MappingNode,
				yaml.SequenceNode,
				yaml.AliasNode,
			},
			node: &yaml.Node{Kind: yaml.DocumentNode},
			want: "yamlpath eval: source expected YAML node with kind 'scalar', 'mapping', 'sequence', or 'alias', but received 'document'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := errs.NewKindError(tc.source, tc.node, tc.expected...)

			got := sut.Error()

			if got, want := got, tc.want; got != want {
				t.Errorf(`KindError.Error() = "%v", want "%v"`, got, want)
			}
		})
	}
}

func TestTagError_Is(t *testing.T) {
	var err errs.TagError

	if got, want := &err, errs.ErrBadTag; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("TagError.Unwrap() error = %v, want %v", got, want)
	}
	if got, want := &err, errs.ErrEval; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("TagError.Unwrap() error = %v, want %v", got, want)
	}
}

func TestTagError_Error(t *testing.T) {
	testCases := []struct {
		name     string
		source   string
		expected []string
		node     *yaml.Node
		want     string
	}{
		{
			name:     "single tag",
			source:   "source",
			expected: []string{"!!str"},
			node:     &yaml.Node{Tag: "!!int"},
			want:     "yamlpath eval: source expected YAML node with tag '!!str', but received '!!int'",
		}, {
			name:     "two tags",
			source:   "source",
			expected: []string{"!!str", "!!int"},
			node:     &yaml.Node{Tag: "!!seq"},
			want:     "yamlpath eval: source expected YAML node with tag '!!str' or '!!int', but received '!!seq'",
		}, {
			name:     "three tags",
			source:   "source",
			expected: []string{"!!str", "!!int", "!!seq"},
			node:     &yaml.Node{Tag: "!!map"},
			want:     "yamlpath eval: source expected YAML node with tag '!!str', '!!int', or '!!seq', but received '!!map'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := errs.NewTagError(tc.source, tc.node, tc.expected...)

			got := sut.Error()

			if got, want := got, tc.want; got != want {
				t.Errorf(`TagError.Error() = "%v", want "%v"`, got, want)
			}
		})
	}
}

func TestNotSingletonError_Is(t *testing.T) {
	var err errs.NotSingletonError

	if got, want := &err, errs.ErrNotSingleton; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("NotSingletonError.Unwrap() error = %v, want %v", got, want)
	}
	if got, want := &err, errs.ErrEval; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("NotSingletonError.Unwrap() error = %v, want %v", got, want)
	}
}

func TestNotSingletonError_Error(t *testing.T) {
	sut := errs.NewSingletonError("source", []*yaml.Node{})

	want := "yamlpath eval: source expected YAML singleton collection, but received 0 nodes"
	got := sut.Error()

	if got != want {
		t.Errorf(`NotSingletonError.Error() = "%v", want "%v"`, got, want)
	}
}

func TestIncompatibleError_Is(t *testing.T) {
	sut := errs.NewIncompatibleError("source",
		yamlutil.String("hello"),
		yamlutil.Boolean("true"),
	)

	if got, want := sut, errs.ErrIncompatible; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("IncompatibleError.Unwrap() error = %v, want %v", got, want)
	}
	if got, want := sut, errs.ErrEval; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("IncompatibleError.Unwrap() error = %v, want %v", got, want)
	}
}

func TestIncompatibleError_Error(t *testing.T) {
	sut := errs.NewIncompatibleError("source",
		yamlutil.String("hello"),
		yamlutil.Boolean("true"),
	)

	want := "yamlpath eval: source has incompatible operands (scalar,!!str) and (scalar,!!bool)"
	got := sut.Error()

	if got != want {
		t.Errorf(`IncompatibleError.Error() = "%v", want "%v"`, got, want)
	}
}

func TestUnsupportedError_Is(t *testing.T) {
	sut := errs.NewUnsupportedErrorf("")

	if got, want := sut, errs.ErrUnsupported; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("UnsupportedError.Unwrap() error = %v, want %v", got, want)
	}
	if got, want := sut, errs.ErrEval; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("UnsupportedError.Unwrap() error = %v, want %v", got, want)
	}
}

func TestNewEvalError(t *testing.T) {
	testCases := []struct {
		name  string
		input error
		want  error
	}{
		{
			name:  "error is not eval error",
			input: errors.New("test error"),
			want:  errs.ErrEval,
		}, {
			name:  "error is eval error",
			input: errs.ErrEval,
			want:  errs.ErrEval,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := errs.NewEvalError(tc.input)

			if got, want := got, tc.want; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("NewEvalError() error = %v, want %v", got, want)
			}
		})
	}
}
