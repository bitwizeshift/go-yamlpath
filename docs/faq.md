# Frequently Asked Questions

i.e. questions I think people will ask, but likely haven't actually asked.

## Why diverge from JSONPath?

The [JSONPath] specification is quite fragmented. Although there is a formal
specification, there are many competing implementations that diverge heavily
from the specification and provide different feature-sets, syntaxes, and
functions.

Additionally, the syntax is questionable in terms of readability.
Take the JSONPath filtering API for example. In JSONPath, a filter expression
requires `?()` to be used inside of a bracket expression -- resulting in
paths like: `$.users[?(@.name == "bitwizeshift")].age`. As an average developer,
it's unclear what this syntax means. Contrast this with the YAMLPath equivalent
defined here: `users.where(name == "bitwizeshift").age`. Immediately, the
intention is clear: we are filtering the `users` array for a user with the
name `bitwizeshift` and then selecting the `age` field.

At its core, this library aims to have full feature parity with JSONPath, but
it is provided in a more readable, fluent, and _extendable_ way.

## Why not use YQ?

[yq] is a wonderful _command line_ tool for querying YAML documents -- but it
is not simple to use this as a library in Go. This library is not meant to
replace `yq` at all; rather, it's meant to provide a simplified feature-set
for traversal that can be easily embedded in any Go application.

[yq]: https://github.com/mikefarah/yq
