package yamlpath_test

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlutil"
)

func newYAML(t *testing.T, s string) *yaml.Node {
	var n yaml.Node
	if err := yaml.Unmarshal([]byte(s), &n); err != nil {
		t.Fatal(err)
	}
	return yamlutil.Normalize(&n)[0]
}

func TestDecoder_Decode_DecodesSequenceAndReturnsEOF(t *testing.T) {
	type object struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}
	collection := yamlpath.Collection{
		newYAML(t, `{"name": "Alice", "age": 30}`),
		newYAML(t, `{"name": "Bob", "age": 25}`),
	}

	sut := collection.Decoder()

	var got object
	if err := sut.Decode(&got); err != nil {
		t.Fatalf("Decoder.Decoder() = error %v, want nil", err)
	}
	if want := (object{Name: "Alice", Age: 30}); !cmp.Equal(got, want) {
		t.Errorf("Decoder.Decoder() = %v, want %v", got, want)
	}

	if err := sut.Decode(&got); err != nil {
		t.Fatalf("Decoder.Decoder() = error %v, want nil", err)
	}
	if want := (object{Name: "Bob", Age: 25}); !cmp.Equal(got, want) {
		t.Errorf("Decoder.Decoder() = %v, want %v", got, want)
	}

	// End of sequence
	err := sut.Decode(&got)
	if got, want := err, io.EOF; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("Decoder.Decoder() = error %v, want %v", got, want)
	}
}

func TestDecoder_Decode_BadDecodingReturnsError(t *testing.T) {
	collection := yamlpath.Collection{
		yamlutil.String("hello"),
	}

	sut := collection.Decoder()

	var got int
	err := sut.Decode(&got)

	if got, want := err, cmpopts.AnyError; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Decoder.Decoder() = error %v, want %v", got, want)
	}
	if got, want := got, 0; !cmp.Equal(got, want) {
		t.Errorf("Decoder.Decoder() = %v, want %v", got, want)
	}
}
