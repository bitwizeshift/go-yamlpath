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

func TestInt(t *testing.T) {
	input := 42
	sut := invocationtest.Int(input)
	want := yamlconv.Int(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Int(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Int(...).GetArg() = %v, want %v", got, want)
	}
}

func TestString(t *testing.T) {
	input := "hello world"
	sut := invocationtest.String(input)
	want := yamlconv.String(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("String(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("String(...).GetArg() = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	input := true
	sut := invocationtest.Bool(input)
	want := yamlconv.Bool(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Bool(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Bool(...).GetArg() = %v, want %v", got, want)
	}
}

func TestFloat(t *testing.T) {
	input := 3.14
	sut := invocationtest.Float(input)
	want := yamlconv.Float(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Float(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Float(...).GetArg() = %v, want %v", got, want)
	}
}

func TestNumber(t *testing.T) {
	input := 42.0
	sut := invocationtest.Number(input)
	want := yamlconv.Number(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Number(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Number(...).GetArg() = %v, want %v", got, want)
	}
}

func TestNode(t *testing.T) {
	input := "hello world"
	sut := invocationtest.Node(input)
	want := yamlconv.Node(input)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Node(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !yamlcmp.EqualRange(got, want) {
		t.Errorf("Node(...).GetArg() = %v, want %v", got, want)
	}
}

func TestSequence(t *testing.T) {
	input := []*yaml.Node{yamlconv.String("hello"), yamlconv.String("world")}
	sut := invocationtest.Sequence(input...)
	want := yamlconv.Sequence(input...)

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Sequence(...).GetArg() error = %v, want %v", got, want)
	}
	if got, want := got, []*yaml.Node{want}; !cmp.Equal(got, want) {
		t.Errorf("Sequence(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestInts(t *testing.T) {
	input := []int{1, 2, 3}
	sut := invocationtest.Ints(input...)
	want := []*yaml.Node{
		yamlconv.Int(1),
		yamlconv.Int(2),
		yamlconv.Int(3),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Ints(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Ints(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestFloats(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3}
	sut := invocationtest.Floats(input...)
	want := []*yaml.Node{
		yamlconv.Float(1.1),
		yamlconv.Float(2.2),
		yamlconv.Float(3.3),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Floats(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Floats(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestNumbers(t *testing.T) {
	input := []float64{1.1, 2.2, 3.3}
	sut := invocationtest.Numbers(input...)
	want := []*yaml.Node{
		yamlconv.Number(1.1),
		yamlconv.Number(2.2),
		yamlconv.Number(3.3),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Numbers(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Numbers(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestStrings(t *testing.T) {
	input := []string{"hello", "world"}
	sut := invocationtest.Strings(input...)
	want := []*yaml.Node{
		yamlconv.String("hello"),
		yamlconv.String("world"),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Strings(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Strings(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestBools(t *testing.T) {
	input := []bool{true, false}
	sut := invocationtest.Bools(input...)
	want := []*yaml.Node{
		yamlconv.Bool(true),
		yamlconv.Bool(false),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Bools(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Bools(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}

func TestNodes(t *testing.T) {
	input := []string{"hello", "world"}
	sut := invocationtest.Nodes(input...)
	want := []*yaml.Node{
		yamlconv.Node("hello"),
		yamlconv.Node("world"),
	}

	got, err := sut.GetArg(nil)

	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Nodes(...).GetArg() error = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Nodes(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
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
		t.Errorf("ErrorParameter(...).GetArg() diff (-got,+want):\n%s", cmp.Diff(got, want))
	}
}
