# YAMLPath

[![Build](https://github.com/bitwizeshift/go-yamlpath/actions/workflows/build-and-test.yaml/badge.svg)](https://github.com/bitwizeshift/go-yamlpath/actions/workflows/build-and-test.yaml)
[![GoDoc](https://godoc.org/rodusek.dev/pkg/yamlpath?status.svg)](https://godoc.org/rodusek.dev/pkg/yamlpath)
[![Go Report Card](https://goreportcard.com/badge/rodusek.dev/pkg/yamlpath)](https://goreportcard.com/report/rodusek.dev/pkg/yamlpath)

A fluent and powerful YAML query language that can be embedded into any Go
library or application.

This language takes heavy inspiration from the [JSONPath] query language, but
features changes to the specification that simplify traversal and make the
language more fluent and readable.

This is built on top of the [`gopkg.in/yaml.v3`] library in Go by querying
[`*yaml.Node`] objects directly. This allows for preserving source-location
information -- which is powerful for providing diagnostics and error messages
when validating YAML documents.

[JSONPath]: https://goessner.net/articles/JsonPath/
[`gopkg.in/yaml.v3`]: https://gopkg.in/yaml.v3

## Quickstart

* [📦 Installation](#installation)
* [🧾 Features](#features)
* [🚂 Example Use](#example-use)
* [📚 Documentation](./docs/index.md)
  * [📄 Grammar](./data/yamlpath.g4)
  * [🔗 Extensions](./docs/extensions.md)
  * [❓ FAQ](./docs/faq.md)
* ⚖️ [MIT](./LICENSE-MIT) or [APACHE-2](./LICENSE-APACHE)

## Installation

To install the package, use the following `go get` command:

```sh
go get rodusek.dev/pkg/yamlpath
```

## Features

This library has full feature-parity with JSONPath, although not all features
are provided in the same way. The follow features are supported:

* ✅ Root and Current node selection with `$` and `@`[^1]
* ✅ Child node selection with `.<name>`
* ✅ Recursive descent with `..`
* ✅ Array index selection with `[<number>]`[^2]
* ✅ Array slice selection with `[<start>:<end>:<step>]`
* ✅ Union of multiple selections with `<path> | <path>`
* ✅ Filtering with the `where` function, e.g. `$.people.where(name == "bitwizeshift")`
* ✅ Comparison operators with `==`, `!=`, `<`, `>`, `<=`, `>=`
* ✅ Containment operators with `in`, `nin`, `subsetof` (e.g. `age in [1, 2, 3]`)
* ✅ Regular expression operator with `=~` (e.g. `name =~ /^b.*shift$/i`)
* ✅ Logical expression operator with `&&`, `||`
* ✅ Arithmetic operators with `+`, `-`, `*`, `/`, `%`
* ✅ String concatenation with `key + "string"`
* ✅ Function support (including custom user-defined functions!)
* ✅ Dynamic subexpressions; any expression can be used as inputs to functions[^3]
* 🚧 Custom environment variables with `%<var>`, specified by the calling
     environment at runtime.

[^1]: These are optional in YAMLPath definitions. The path is always assumed to
      be the "current" context path if unspecified; but can be provided for
      disambiguation.
[^2]: In YAMLPath, only indices are selected with the index operator. Fields are
      selected with the `.` operator. To select fields with a string value, the
      `select` function may be used (e.g. `$.select("some key")`).
[^3]: This, along with custom environment variables, provide feature-parity with
      JSONPath "script" execution without needing to provide a way to call out
      to JavaScript or other languages. This allows the calling language (Go) to
      dynamically determine data and feed it back into YAMLPath.

## Example Use

> [!NOTE]
> For more examples, see the [examples](./_examples) directory.

This illustrates a simple example using this library to validate configurations,
with output diagnostics appearing in GitHub annotation format. This validates
some criteria on a `book` yaml object:

```go
path := yamlpath.MustCompile("store.book.where(price < 10.00)")

filepath := "bookstore.yaml"   // Some filepath to the YAML file
file, err := os.Open(filepath) // some file handle to the YAML file
if err != nil {
  log.Fatal(err)
}
defer file.Close()

var node yaml.Node
if err := yaml.NewDecoder(file).Decode(&node); err != nil {
  log.Fatal(err)
}

result, err := path.Match(&node)
if err != nil {
  log.Fatal(err)
}

for _, node := range result {
  if err := validateBook(result); err != nil {
    // This is in GitHub annotation format
    fmt.Println("::error file=%s,line=%d,title=Book validation failed::%v",
      filepath,
      result.Line,
      err,
    )
  }
}
```
