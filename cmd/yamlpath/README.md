# YAMLPath CLI

This is a small CLI tool that uses YAMLPath to query YAML files from the command
line. This makes it easy to test and evaluate YAMLPath expressions without
needing to write a Go program, and can be used in shell scripts or other
automation.

> [!NOTE]
> This is _not_ meant to replace or in any way compete with [`yq`], which is a
> fantastic tool for working with YAML files. YAMLPath and `yq` have different
> syntaxes and use-cases, so you should use the tool that best fits your needs.

## Installation

```bash
go install rodusek.dev/pkg/yamlpath/cmd/yamlpath@latest
```

## Usage

The general usage of the `yamlpath` utility is simply
`yamlpath <expression> -i <file>`, where `<expression>` is a YAMLPath expression
and `<file>` is the path to the YAML file to query.

By default, output is printed to stdout. You can also use the `-o` flag to write
the output to a file.

**Example:**

```bash
yamlpath '.foo.bar' -i file.yaml
```

[`yq`]: https://github.com/mikefarah/yq
