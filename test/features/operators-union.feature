Feature: Union operator

  YAMLPath expressions may invoke the union operator `|` to combine
  multiple collections into a single collection. The union operator shall return
  a collection of all elements in the left operand, followed by all elements in
  the right operand. The union operator shall not remove duplicate values from
  the resulting sequence.

  Scenario: Two expressions are combined

    Given the yaml input:
      """
      first: "hello"
      second: "world"
      """
    When the yamlpath `first | second` is evaluated
    Then the evaluation result is:
      """
      hello
      ---
      world
      """
