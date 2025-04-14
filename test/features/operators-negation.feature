Feature: Negation operator

  YAMLPath supports a negation operator `!` (alternative: `not`) that may prefix
  any boolean expression to invert the result of the expression.

  The negation operator may only act on boolean expressions, and shall signal an
  error to the calling environment if the result returns more than one value, or
  is not a boolean value.

  Rule: The negation operator shall invert the result of the expression

    Scenario: Negation of boolean true value
      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `!(number == 42)` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Negation of boolean false value
      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `not (number != 42)` is evaluated
      Then the evaluation result is:
        """
        true
        """
