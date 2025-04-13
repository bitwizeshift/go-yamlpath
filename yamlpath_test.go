package yamlpath_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
)

func TestCompile(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantValue bool
		wantErr   error
	}{
		{
			name:      "Invalid path returns error",
			input:     "#",
			wantValue: false,
			wantErr:   yamlpath.ErrCompile,
		}, {
			name:      "Valid path returns YAMLPath",
			input:     "$.store.book[*].author",
			wantValue: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlpath.Compile(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("Compile() error = %v, want %v", got, want)
			}
			if tc.wantValue && got == nil {
				t.Errorf("Compile() = %v, want non-nil", got)
			}
		})
	}
}

func TestMustCompile(t *testing.T) {
	got := yamlpath.MustCompile("$.store.book[*].author")

	if got == nil {
		t.Fatalf("MustCompile() = %v, want non-nil", got)
	}
}

func TestMustCompile_PanicsIfInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustCompile() did not panic")
		}
	}()

	_ = yamlpath.MustCompile("#")
}

func zero[T any]() *T {
	var v T
	return &v
}

func TestYAMLPath_String(t *testing.T) {
	testCases := []struct {
		name string
		path *yamlpath.YAMLPath
		want string
	}{
		{
			name: "Empty path",
			path: zero[yamlpath.YAMLPath](),
			want: "",
		}, {
			name: "Nil path",
			path: nil,
			want: "",
		}, {
			name: "Valid path",
			path: yamlpath.MustCompile("$.store.book[*].author"),
			want: "$.store.book[*].author",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.path.String()

			if got, want := got, tc.want; got != want {
				t.Errorf("YAMLPath.String() = %v, want %v", got, want)
			}
		})
	}
}

func TestYAMLPath_Match(t *testing.T) {
	testCases := []struct {
		name    string
		path    *yamlpath.YAMLPath
		input   *yaml.Node
		want    yamlpath.Collection
		wantErr error
	}{
		{
			name: "Nil path returns empty collection",
		}, {
			name: "Nil input returns empty collection",
			path: yamlpath.MustCompile("$.store.book[*].author"),
		}, {
			name:  "Valid path returns matching collection",
			path:  yamlpath.MustCompile("$[*].name"),
			input: yamltest.MustParseNode(`[{"name": "Alice"}, {"name": "Bob"}]`),
			want: yamlpath.Collection{
				yamlconv.String("Alice"),
				yamlconv.String("Bob"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.path.Match(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("YAMLPath.Match() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("YAMLPath.Match() = %v, want %v", got, want)
			}
		})
	}
}

func TestYAMLPath_MustMatch(t *testing.T) {
	sut := yamlpath.MustCompile("$[*].name")
	input := yamltest.MustParseNode(`[{"name": "Alice"}, {"name": "Bob"}]`)
	got := sut.MustMatch(input)

	want := yamlpath.Collection{
		yamlconv.String("Alice"),
		yamlconv.String("Bob"),
	}

	if got, want := got, want; !cmp.Equal(got, want) {
		t.Errorf("YAMLPath.MustMatch() = %v, want %v", got, want)
	}
}

func TestYAMLPath_MustMatch_PanicsIfErrorOccurs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustMatch() did not panic")
		}
	}()

	sut := yamlpath.MustCompile("$[?(@.age > 32)]")
	input := yamltest.MustParseNode(`{"age": "foo"}`)
	_ = sut.MustMatch(input)
}

func TestYAMLPath_UnmarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "Valid path unmarshals successfully",
			input: "$.store.book[*].author",
		}, {
			name:    "Invalid path returns error",
			input:   "#",
			wantErr: yamlpath.ErrCompile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got yamlpath.YAMLPath
			err := got.UnmarshalText([]byte(tc.input))

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("YAMLPath.UnmarshalText() error = %v, want %v", got, want)
			}
		})
	}
}

func TestYAMLPath_MarshalText(t *testing.T) {
	testCases := []struct {
		name    string
		path    *yamlpath.YAMLPath
		want    string
		wantErr error
	}{
		{
			name: "Nil path returns empty string",
		}, {
			name: "Valid path returns string",
			path: yamlpath.MustCompile("$.store.book[*].author"),
			want: "$.store.book[*].author",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.path.MarshalText()

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Fatalf("YAMLPath.MarshalText() error = %v, want %v", got, want)
			}
			if got, want := string(got), tc.want; got != want {
				t.Errorf("YAMLPath.MarshalText() = %v, want %v", got, want)
			}
		})
	}
}

func TestYAMLPath_Equal(t *testing.T) {
	testCases := []struct {
		name     string
		lhs, rhs *yamlpath.YAMLPath
		want     bool
	}{
		{
			name: "Nil paths compare equal",
			want: true,
		}, {
			name: "Nil lhs path with zero rhs path",
			rhs:  zero[yamlpath.YAMLPath](),
			want: true,
		}, {
			name: "Nil rhs path with zero lhs path",
			lhs:  zero[yamlpath.YAMLPath](),
			want: true,
		}, {
			name: "Zero value paths compare equal",
			lhs:  zero[yamlpath.YAMLPath](),
			rhs:  zero[yamlpath.YAMLPath](),
			want: true,
		}, {
			name: "Different paths compare not equal",
			lhs:  yamlpath.MustCompile("$.store.book[*].author"),
			rhs:  yamlpath.MustCompile("$.store.book[*].title"),
			want: false,
		}, {
			name: "Equal paths compare equal",
			lhs:  yamlpath.MustCompile("$.store.book[*].author"),
			rhs:  yamlpath.MustCompile("$.store.book[*].author"),
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.lhs.Equal(tc.rhs)

			if got != tc.want {
				t.Errorf("YAMLPath.Equal() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		input   *yaml.Node
		want    yamlpath.Collection
		wantErr error
	}{
		{
			name:  "Valid path returns matching collection",
			path:  "name",
			input: yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			want: yamlpath.Collection{
				yamlconv.String("Alice"),
			},
		}, {
			name:    "Invalid path returns error",
			path:    "$$",
			input:   yamltest.MustParseNode(`{"name": "Alice", "age": 30}`),
			wantErr: yamlpath.ErrSyntax,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := yamlpath.Match(tc.path, tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("yamlpath.Match() error = %v, want %v", got, want)
			}
			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("yamlpath.Match() = %v, want %v", got, want)
			}
		})
	}
}

func TestMustMatch(t *testing.T) {
	want := yamlpath.Collection{
		yamlconv.String("Alice"),
	}
	input := yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)

	got := yamlpath.MustMatch("name", input)

	if got, want := got, want; !cmp.Equal(got, want) {
		t.Errorf("yamlpath.MustMatch() = %v, want %v", got, want)
	}
}

func TestMustMatch_PanicsIfErrorOccurs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustMatch() did not panic")
		}
	}()

	input := yamltest.MustParseNode(`{"name": "Alice", "age": 30}`)
	_ = yamlpath.MustMatch("$$", input)
}
