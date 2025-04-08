package exprtest_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/expr/exprtest"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

func TestEmpty_Eval(t *testing.T) {
	got, err := exprtest.Empty().Eval(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Empty.Eval() error = %v, want %v", got, want)
	}
	if got, want := len(got), 0; got != want {
		t.Errorf("Empty.Eval() = %v, want %v", got, want)
	}
}

func TestReturn_Eval(t *testing.T) {
	nodes := []*yaml.Node{yamlconv.Bool(true)}

	got, err := exprtest.Return(nodes...).Eval(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Return.Eval() error = %v, want %v", got, want)
	}
	if got, want := got, nodes; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Return.Eval() = %v, want %v", got, want)
	}
}

func TestError_Eval(t *testing.T) {
	testErr := errors.New("test error")

	got, err := exprtest.Error(testErr).Eval(nil)

	if got, want := err, testErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Error.Eval() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{}; !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Error.Eval() = %v, want %v", got, want)
	}
}
