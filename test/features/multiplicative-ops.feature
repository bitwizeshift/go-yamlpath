Feature: Multiplicative Operators

  YAMLPath sub-expressions may operate with multiplicative operators for both
  scalar and singular numeric values. These operators will return a numeric
  value, and shall signal an error to the calling environment if the result
  returns more than one value, or is not a numeric value.

  Rule: Multiplication operator computes product of two numeric values

    Scenario Outline: Left <left_type> multiplied by right <right_type>

      Given the yaml input:
        """
        number: <left>
        """
      When the yamlpath `$.where($.number * <right> == <result>)` is evaluated
      Then the evaluation result is:
        """
        number: <left>
        """

      Examples:
        | left_type | right_type | left | right | result |
        | integer   | integer    | 6    | 7     | 42     |
        | integer   | float      | 5    | 7.5   | 37.5   |
        | float     | integer    | 7.5  | 5     | 37.5   |
        | float     | float      | 1.3  | 3.5   | 4.55   |

  Rule: Division operator computes quotient of two numeric values

    Scenario Outline: Left <left_type> divided by right <right_type>

      Given the yaml input:
        """
        number: <left>
        """
      When the yamlpath `$.where($.number / <right> == <result>)` is evaluated
      Then the evaluation result is:
        """
        number: <left>
        """

      Examples:
        | left_type | right_type | left | right | result |
        | integer   | integer    | 42   | 7     | 6      |
        | integer   | float      | 37.5 | 7.5   | 5      |
        | float     | integer    | 37.5 | 5     | 7.5    |
        | float     | float      | 4.55 | 3.5   | 1.3    |
