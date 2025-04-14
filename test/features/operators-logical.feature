Feature: Logical boolean operators

  YAMLPath expressions support logical "and" and "or" operators. These
  operators are used to combine multiple boolean expressions into a single
  boolean expression. The logical operators are represented by the symbols
  '&&' and '||', respectively, with alternatives of 'and' or 'or' respectively.

  Rule: And operator returns true if both operands are truthy

    Scenario Outline: Left <left> and right <right>

      Given the yaml input:
        """
        left: <left>
        right: <right>
        """
      When the yamlpath `left && right` is evaluated
      Then the evaluation result is:
        """
        <result>
        """

      Examples:
        | left  | right | result |
        | true  | true  | true   |
        | true  | false | false  |
        | false | true  | false  |
        | false | false | false  |

  Rule: Or operator returns true if either operand is truthy

    Scenario Outline: Left <left> or right <right>

      Given the yaml input:
        """
        left: <left>
        right: <right>
        """
      When the yamlpath `left || right` is evaluated
      Then the evaluation result is:
        """
        <result>
        """

      Examples:
        | left  | right | result |
        | true  | true  | true   |
        | true  | false | true   |
        | false | true  | true   |
        | false | false | false  |
