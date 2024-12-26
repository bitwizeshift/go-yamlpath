# YAMLPath Docs

> [!NOTE]
> This page is currently a stub, and will be expanded in the future.

This page contains documentation for the `yamlpath` package, which provides an
implementation of the [JSONPath] query language for YAML documents using the
[`gopkg.in/yaml.v3`] package.

[JSONPath]: https://goessner.net/articles/JsonPath/
[`gopkg.in/yaml.v3`]: https://gopkg.in/yaml.v3

## Table of Contents

* [Quickstart](#quickstart)
* [Motivation](#motivation)
* [Examples](../examples/)
* YAMLPath Language
  * [Features](./features.md)
  * [Extensions](./extensions.md)
* Implementation Details
  * [Grammar](../data/yamlpath.g4)
  * [Testing](../test/)

## Quickstart


## Motivation

Although there are many uses for query-languages on YAML data structures, the
primary motivation for this package was to make it easier to validate YAML
configuration and provide proper line-diagnostics for errors. The
[`gopkg.in/yaml.v3`] package makes it really easy to `Marshal` and `Unmarshal`
between YAML and Go types, but it loses source-information in the process,
making it difficult to provide accurate diagnostics back to the user.

This package helps solve that problem, since it communicates fully in
`*yaml.Node` objects -- which preserve source information. In addition to this,
`*yaml.Node` objects may be `Decode`d back into Go types, meaning there is no
loss in functionality by communicating in `*yaml.Node` objects.
