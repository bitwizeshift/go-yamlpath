package invocationtest_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/invocation/invocationtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestSuccessParameter(t *testing.T) {
	input := []*yaml.Node{yamlconv.String("hello world")}

	sut := invocationtest.SuccessParameter(input...)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("SuccessParameter(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, input; !yamlcmp.EqualRange(got, want) {
		t.Errorf("SuccessParameter(...).GetArg() = %v, want %v", got, want)
	}
}

func TestErrorParameter(t *testing.T) {
	testErr := errors.New("test error")

	sut := invocationtest.ErrorParameter(testErr)

	got, err := sut.GetArg(nil)

	if got, want := err, testErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("ErrorParameter(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, ([]*yaml.Node)(nil); !cmp.Equal(got, want) {
		t.Errorf("ErrorParameter(...).GetArg() = %v, want %v", got, want)
	}
}
