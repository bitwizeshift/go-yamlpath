package yamlpathtest_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
	"rodusek.dev/pkg/yamlpath/internal/yamltest"
	"rodusek.dev/pkg/yamlpath/yamlpathtest"
)

func TestCollection(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamlconv.String("a"),
		yamlconv.String("b"),
		yamlconv.String("c"),
	}
	sut := yamlpathtest.Collection(
		yamltest.MustParseNode(`"a"`),
		yamltest.MustParseNode(`"b"`),
		yamltest.MustParseNode(`"c"`),
	)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`{}`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Collection(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Collection(...).Match(...) = %v, want %v", got, want)
	}
}

func TestString(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamlconv.String("test")}
	sut := yamlpathtest.String("test")

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`"test"`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("String(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("String(...).Match(...) = %v, want %v", got, want)
	}
}

func TestStringSequence(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamltest.MustParseNode(`["a", "b", "c"]`),
	}
	sut := yamlpathtest.StringSequence("a", "b", "c")

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`["a", "b", "c"]`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("StringSequence(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("StringSequence(...).Match(...) = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamltest.MustParseNode(`true`)}
	sut := yamlpathtest.Bool(true)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`true`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Bool(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Bool(...).Match(...) = %v, want %v", got, want)
	}
}

func TestBoolSequence(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamltest.MustParseNode(`[true, false, true]`),
	}
	sut := yamlpathtest.BoolSequence(true, false, true)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`[true, false, true]`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("BoolSequence(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("BoolSequence(...).Match(...) = %v, want %v", got, want)
	}
}

func TestFloat(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamltest.MustParseNode(`1.23`)}
	sut := yamlpathtest.Float(1.23)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`1.23`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Float(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Float(...).Match(...) = %v, want %v", got, want)
	}
}

func TestFloatSequence(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamltest.MustParseNode(`[1.23, 4.56, 7.89]`),
	}
	sut := yamlpathtest.FloatSequence(1.23, 4.56, 7.89)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`[1.23, 4.56, 7.89]`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("FloatSequence(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("FloatSequence(...).Match(...) = %v, want %v", got, want)
	}
}

func TestObject(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamltest.MustParseNode(`{"key": "value"}`)}
	sut := yamlpathtest.Object(map[string]any{"key": "value"})

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`{"key": "value"}`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Object(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Object(...).Match(...) = %v, want %v", got, want)
	}
}

func TestObjectSequence(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamltest.MustParseNode(`[{"key": "value"}, {"key": "value2"}]`),
	}
	sut := yamlpathtest.ObjectSequence(
		map[string]any{"key": "value"},
		map[string]any{"key": "value2"},
	)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`[{"key": "value"}, {"key": "value2"}]`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("ObjectSequence(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("ObjectSequence(...).Match(...) = %v, want %v", got, want)
	}
}

func TestInt(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamltest.MustParseNode(`123`)}
	sut := yamlpathtest.Int(123)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`123`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Int(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Int(...).Match(...) = %v, want %v", got, want)
	}
}

func TestIntSequence(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{
		yamltest.MustParseNode(`[1, 2, 3]`),
	}
	sut := yamlpathtest.IntSequence(1, 2, 3)

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`[1, 2, 3]`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("IntSequence(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("IntSequence(...).Match(...) = %v, want %v", got, want)
	}
}

func TestNull(t *testing.T) {
	// Arrange
	want := yamlpath.Collection{yamlconv.Null()}
	sut := yamlpathtest.Null()

	// Act
	got, err := sut.Match(yamltest.MustParseNode(`null`))

	// Assert
	if got, want := err, (error)(nil); !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Fatalf("Null(...).Match(...) err = %v, want %v", got, want)
	}
	if !cmp.Equal(got, want, cmpopts.EquateEmpty()) {
		t.Errorf("Null(...).Match(...) = %v, want %v", got, want)
	}
}

func TestError(t *testing.T) {
	// Arrange
	testErr := errors.New("test error")

	// Act
	sut := yamlpathtest.Error(testErr)
	_, err := sut.Match(yamltest.MustParseNode(`{}`))

	if got, want := err, testErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
		t.Errorf("Error(...).Match(...) err = %v, want %v", got, want)
	}
}
