Feature: Script evaluation

  JSONPath supports a mechanism of "script" evaluation, which calls out to the
  host language to evaluate a script. Since this is not JS, we will not have
  full javascript-support; but instead, YAMLPath will support the full
  evaluation functionality for this instead, so that computed fields can be
  programmatically decided based on other path values.

  Scenario: Simple script evaluation
    Given the yaml input:
      """
      foo: "bar"
      bar: 2
      """
    When the yamlpath `$[(@.foo)]` is evaluated
    Then the evaluation result is:
      """
      2
      """
