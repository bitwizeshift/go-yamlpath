Feature: Arithmetic Operators

  YAMLPath expressions may operate with arithmetic operators for both
  scalar and singular numeric values. These operators will return a numeric
  value, and shall signal an error to the calling environment if the result
  returns more than one value, or is not a numeric value.

  Rule: Addition operator computes sum of two numeric values

    Scenario Outline: Left <left_type> added to right <right_type>

      Given the yaml input:
        """
        number: <left>
        """
      When the yamlpath `$.where($.number + <right> == <result>)` is evaluated
      Then the evaluation result is:
        """
        number: <left>
        """

      Examples:
        | left_type | right_type | left | right | result |
        | integer   | integer    | 6    | 7     | 13     |
        | integer   | float      | 5    | 7.5   | 12.5   |
        | float     | integer    | 7.5  | 5     | 12.5   |
        | float     | float      | 1.3  | 3.5   | 4.8    |

  Rule: Subtraction operator computes difference of two numeric values

    Scenario Outline: Left <left_type> subtracted from right <right_type>

      Given the yaml input:
        """
        number: <left>
        """
      When the yamlpath `$.where($.number - <right> == <result>)` is evaluated
      Then the evaluation result is:
        """
        number: <left>
        """

      Examples:
        | left_type | right_type | left | right | result |
        | integer   | integer    | 13   | 7     | 6      |
        | integer   | float      | 12.5 | 7.5   | 5      |
        | float     | integer    | 12.5 | 5     | 7.5    |
        | float     | float      | 4.8  | 3.5   | 1.3    |

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

  Rule: Modulus operator computes remainder of two numeric values

    Scenario Outline: Left <left_type> modulus right <right_type>

      Given the yaml input:
        """
        number: <left>
        """
      When the yamlpath `$.where($.number % <right> == <result>)` is evaluated
      Then the evaluation result is:
        """
        number: <left>
        """

      Examples:
        | left_type | right_type | left | right | result |
        | integer   | integer    | 42   | 7     | 0      |
        | integer   | integer    | 37   | 5     | 2      |
        | integer   | float      | 37.5 | 7.5   | 0      |
        | float     | integer    | 37.5 | 5     | 2.5    |
        | float     | float      | 4.55 | 3.5   | 1.05   |
