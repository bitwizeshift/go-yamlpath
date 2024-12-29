# YAMLPath

[![GoDoc](https://godoc.org/rodusek.dev/pkg/yamlpath?status.svg)](https://godoc.org/rodusek.dev/pkg/yamlpath)
[![Go Report Card](https://goreportcard.com/badge/rodusek.dev/pkg/yamlpath)](https://goreportcard.com/report/rodusek.dev/pkg/yamlpath)

An implementation of the [JSONPath] query language, built for YAML and the
[`gopkg.in/yaml.v3`] package.

This package makes it easy to query [`*yaml.Node`] objects to do complex queries
and operations, such as validation, while still preserving source-location
information -- which is useful for providing diagnostics.

[JSONPath]: https://goessner.net/articles/JsonPath/
[`gopkg.in/yaml.v3`]: https://gopkg.in/yaml.v3

## Quickstart

* [📦 Installation](#installation)
* [🧾 Features](#features)
* [🚂 Example Use](#example-use)
* [📚 Documentation](./docs/index.md)
  * [📄 Grammar](./data/yamlpath.g4)
  * [🔗 Extensions](./docs/extensions.md)
* ⚖️ [MIT](./LICENSE-MIT) or [APACHE-2](./LICENSE-APACHE)

## Installation

To install the package, use the following `go get` command:

```sh
go get rodusek.dev/pkg/yamlpath
```

## Features

The following features of JSONPath are supported and thoroughly tested.

* [x] Root and Current node selection (`$` and `@`)
* [x] Child node selection (`.<name>`)
* [x] Recursive descent (`..`)
* [x] Array index selection (`[<number>]`)
* [x] Array slice selection (`[<start>:<end>:<step>]`)
* [x] Union selection (`<path>,<path>`)
* [ ] Filter selection (`[?(<sub-expression>)]`) (**partial**)
  * [x] Existence and field filters (`[?(@.key)]`)
  * [x] Comparison filters `==`, `!=`, `<`, `>`, `<=`, `>=` (e.g. `[?(@.key == "value")]`)
  * [x] Containment filters `in`, `nin`, `subsetof` (e.g. `[?(@.entry in $.keys]`)
  * [x] Regular expression filters (e.g. `[?(@.key =~ /pattern/)]`)
  * [x] Logical expressions `&&`, `||` (`[?(@.key && @.key2)]`)
  * [ ] Function filters (`[?(@.key.function() == "value")]`)
* [ ] Script selection (`[<expression>]`) (**partial**)
  * [x] Field access (`[(@.key)]`)
  * [x] Arithmetic expressions (`[(@.key + 1)]`)
  * [x] String concatenation (`[(@.key + "string")]`)
  * [ ] Function calls (`[(@.key.function())`)

## Example Use

This illustrates a simple example using this library to validate configurations,
with output diagnostics appearing in GitHub annotation format. This validates
some criteria on a `book` yaml object:

```go
path := yamlpath.MustCompile("$.store.book[?(@.price < 10.00)]")

filepath := ... // Some filepath to the YAML file
file := ...     // some file handle to the YAML file
var node yaml.Node
if err := yaml.NewDecoder(file).Decode(&node); err != nil {
  log.Fatal(err)
}

result, err := path.Eval(data)
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
