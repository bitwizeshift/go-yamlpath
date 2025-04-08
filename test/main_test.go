package test_test

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath"
	"rodusek.dev/pkg/yamlpath/internal/yamlcmp"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		Name:                "YAMLPath",
		Options:             &o,
		ScenarioInitializer: InitializeScenario,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^the yaml input:$`, SetYAMLInput)
	ctx.When("^the yamlpath `(.*)` is evaluated$", EvaluateYAMLPath)
	ctx.Then(`^the evaluation contains (\d) entries$`, TheEvaluationContainsNEntries)
	ctx.Then(`^the evaluation contains (1) entry$`, TheEvaluationContainsNEntries)
	ctx.Then(`^the evaluation result is empty$`, TheEvaluationIsEmpty)
	ctx.Then(`^an error is raised$`, AnErrorIsRaised)
	ctx.Then(`^the evaluation result is:$`, TheEvaluationResultIs)
}

type ctxKey int

const (
	ctxKeyInput ctxKey = iota
	ctxKetResult
	ctxKeyErrResult
)

func SetYAMLInput(ctx context.Context, input string) (context.Context, error) {
	decoder := yaml.NewDecoder(strings.NewReader(input))
	var node yaml.Node
	if err := decoder.Decode(&node); err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ctxKeyInput, &node), nil
}

func GetYAMLInput(ctx context.Context) *yaml.Node {
	node := ctx.Value(ctxKeyInput)
	if node == nil {
		return nil
	}
	return node.(*yaml.Node)
}

func EvaluateYAMLPath(ctx context.Context, path string) (context.Context, error) {
	node := GetYAMLInput(ctx)

	yp, err := yamlpath.Compile(path)
	if err != nil {
		return ctx, err
	}
	got, err := yp.Match(node)
	if err != nil {
		ctx = context.WithValue(ctx, ctxKeyErrResult, err)
	} else {
		ctx = context.WithValue(ctx, ctxKetResult, got)
	}
	return ctx, nil
}

func GetResult(ctx context.Context) (yamlpath.Collection, bool) {
	result := ctx.Value(ctxKetResult)
	if result == nil {
		return nil, false
	}
	return result.(yamlpath.Collection), true
}

func GetError(ctx context.Context) error {
	err := ctx.Value(ctxKeyErrResult)
	if err == nil {
		return nil
	}
	return err.(error)
}

func AnErrorIsRaised(ctx context.Context) error {
	err := GetError(ctx)
	if err == nil {
		return errors.New("expected an error, but none were raised")
	}
	return nil
}

func TheEvaluationIsEmpty(ctx context.Context) error {
	result, _ := GetResult(ctx)
	if len(result) != 0 {
		return fmt.Errorf("expected an empty result but got %d", len(result))
	}
	return nil
}

func TheEvaluationContainsNEntries(ctx context.Context, n int) error {
	result, _ := GetResult(ctx)
	if len(result) != n {
		return fmt.Errorf("unexpected %d results but got %d", n, len(result))
	}
	return nil
}

func TheEvaluationResultIs(ctx context.Context, content string) error {
	if err := GetError(ctx); err != nil {
		return fmt.Errorf("unexpected evaluation error: %v", err)
	}

	got, _ := GetResult(ctx)
	var want []*yaml.Node
	decoder := yaml.NewDecoder(strings.NewReader(content))
	for {
		var next yaml.Node
		if err := decoder.Decode(&next); err != nil {
			if !errors.Is(err, io.EOF) {
				return fmt.Errorf("error decoding expected yaml: %v", err)
			}
			break
		}
		want = append(want, &next)
	}
	want = yamlconv.FlattenDocuments(want...) // normalize documents

	if !yamlcmp.EqualRange(got, want) {
		return fmt.Errorf("unexpected got:\n%s\nwant:\n%s", nodesToString(got), nodesToString(want))
	}

	return nil
}

func nodesToString(nodes []*yaml.Node) string {
	var b strings.Builder
	encoder := yaml.NewEncoder(&b)
	for _, n := range nodes {
		if err := encoder.Encode(n); err != nil {
			return fmt.Sprintf("error encoding node: %v", err)
		}
	}
	return b.String()
}
