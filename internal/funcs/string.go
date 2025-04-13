package funcs

import (
	"strings"

	"gopkg.in/yaml.v3"
	"rodusek.dev/pkg/yamlpath/internal/errs"
	"rodusek.dev/pkg/yamlpath/internal/invocation"
	"rodusek.dev/pkg/yamlpath/internal/yamlconv"
)

// Lower converts a singleton string input into lowercase.
func Lower(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("lower()", current)
	}

	node := current[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("lower()", node, yaml.ScalarNode)
	}
	if node.Tag != "!!str" {
		return nil, errs.NewTagError("lower()", node, "!!str")
	}

	return []*yaml.Node{yamlconv.String(strings.ToLower(node.Value))}, nil
}

// Upper converts a singleton string input into uppercase.
func Upper(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}
	if len(current) != 1 {
		return nil, errs.NewSingletonError("upper()", current)
	}

	node := current[0]
	if node.Kind != yaml.ScalarNode {
		return nil, errs.NewKindError("upper()", node, yaml.ScalarNode)
	}
	if node.Tag != "!!str" {
		return nil, errs.NewTagError("upper()", node, "!!str")
	}
	return []*yaml.Node{yamlconv.String(strings.ToUpper(node.Value))}, nil
}

// StartsWith checks if the collection contains a single string that starts with
// substring from the argument.
//
// This will return true if the input collection starts with the prefix, and
// will return false if it does not. Otherwise, the behavior is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If either the input collection or the first argument are not string nodes,
//     an error is raised.
func StartsWith(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "startsWith(...)")
	}
	prefix, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "startsWith(prefix)")
	}

	if strings.HasPrefix(str, prefix) {
		return []*yaml.Node{yamlconv.Bool(true)}, nil
	}
	return []*yaml.Node{yamlconv.Bool(false)}, nil
}

// EndsWith checks if the collection contains a single string that ends with
// substring from the argument.
//
// This will return true if the input collection ends with the suffix, and
// will return false if it does not. Otherwise, the behavior is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If either the input collection or the first argument are not string nodes,
//     an error is raised.
func EndsWith(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "endsWith(...)")
	}
	suffix, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "endsWith(suffix)")
	}

	return []*yaml.Node{yamlconv.Bool(strings.HasSuffix(str, suffix))}, nil
}

// Contains checks if the collection contains a single string that contains
// substring from the argument.
//
// This will return true if the input collection contains the substring, and
// will return false if it does not. Otherwise, the behavior is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If either the input collection or the first argument are not string nodes,
//     an error is raised.
func Contains(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "contains(...)")
	}
	substr, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "contains(substring)")
	}

	return []*yaml.Node{yamlconv.Bool(strings.Contains(str, substr))}, nil
}

// IndexOf returns the index of the first occurrence of a substring in a
// string.
func IndexOf(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "indexOf(...)")
	}
	substr, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "indexOf(substring)")
	}

	index := strings.Index(str, substr)
	return []*yaml.Node{yamlconv.Number(index)}, nil
}

// Substring returns a substring of the input string.
//
// This will return a substring of the input string, starting at the index
// specified by the first argument, and ending at the index specified by the
// optional second argument. If the second argument is not provided, the
// substring will end at the end of the input string. Otherwise, the behavior
// is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If the input collection is not a string node, an error is raised.
//   - If the first argument is not a number, an error is raised.
//   - If the second argument is specified and is not a number, an error is
//     raised.
//   - If the first argument is greater than the length of the input string,
func Substring(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "substring(...)")
	}
	start, err := invocation.ParseInt(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "substring(start-index, ...)")
	}
	if start > int64(len(str)) {
		return nil, nil
	}

	length := int64(len(str))
	if len(params) > 1 {
		length, err = invocation.ParseInt(ctx, params[1])
		if err != nil {
			return nil, errs.IncludeSource(err, "substring(..., length)")
		}
	}
	length += start
	if length > int64(len(str)) {
		length = int64(len(str))
	}

	return []*yaml.Node{yamlconv.String(str[start:length])}, nil
}

// Replace replaces all occurrences of a substring in the input string.
//
// This will return a string with all occurrences of the substring replaced
// with the replacement string. Otherwise, the behavior is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
func Replace(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "replace(...)")
	}

	from, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "replace(pattern, ...)")
	}
	to, err := invocation.ParseString(ctx, params[1])
	if err != nil {
		return nil, errs.IncludeSource(err, "replace(..., substitution)")
	}

	return []*yaml.Node{yamlconv.String(strings.ReplaceAll(str, from, to))}, nil
}

// Split splits a string into a sequence of strings based on a separator.
func Split(ctx invocation.Context, params ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "split()")
	}
	sep, err := invocation.ParseString(ctx, params[0])
	if err != nil {
		return nil, errs.IncludeSource(err, "split(separator)")
	}

	result := strings.Split(str, sep)
	var nodes []*yaml.Node
	for _, s := range result {
		nodes = append(nodes, yamlconv.String(s))
	}
	return nodes, nil
}

// Length returns the length of the input string.
//
// This will return the length of the input string. Otherwise, the behavior
// is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If the input collection is not a string node, an error is raised.
func Length(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "length()")
	}

	return []*yaml.Node{yamlconv.Number(len(str))}, nil
}

// ToChars converts a string into a sequence of characters.
//
// This will return a sequence of characters from the input string. Otherwise,
// the behavior is as follows:
//
//   - If the input collection is empty, the output is also empty.
//   - If the input collection contains more than one element, an error is
//     raised to the calling environment.
//   - If the input collection is not a string node, an error is raised.
func ToChars(ctx invocation.Context, _ ...invocation.Parameter) ([]*yaml.Node, error) {
	current := ctx.Current()
	if len(current) == 0 {
		return nil, nil
	}

	str, err := yamlconv.ParseString(current...)
	if err != nil {
		return nil, errs.IncludeSource(err, "toChars()")
	}

	var result []*yaml.Node
	for _, c := range str {
		result = append(result, yamlconv.String(string(c)))
	}

	return result, nil
}
