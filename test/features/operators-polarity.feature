Feature: Unary prefix polarity operators

  YAMLPath sub-expressions may provide unary prefix polarity operators of '+'
  and '-' to alter the sign of a numeric value. These operators may only act on
  numeric scalar/singular values, and shall signal an error to the calling
  environment if the result returns more than one value.

  Rule: The '+' operator shall return the numeric value unchanged

    Scenario: Integer value is unchanged

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `$.where(+$.number == 42)` is evaluated
      Then the evaluation result is:
        """
        number: 42
        """

  Rule: The '-' operator shall negate the numeric value of the field

    Scenario: Integer value is negated

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `$.where(-$.number == -42)` is evaluated
      Then the evaluation result is:
        """
        number: 42
        """

    Scenario: Float value is negated

      Given the yaml input:
        """
        number: 42.69
        """
      When the yamlpath `$.where(-$.number == -42.69)` is evaluated
      Then the evaluation result is:
        """
        number: 42.69
        """
