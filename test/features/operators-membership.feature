Feature: Membership Operators

  YAMLPath expressions may query for membership of a value in a set of values.
  The membership operators are `in`, `nin`, and `subsetof`.

  Rule: The `in` operator shall query if the left operand exists in the right

    The `in` operator shall return `true` if the left operand exists as an
    element in a sequence defined in the right operand.

    Scenario: Left operand exists within right operand

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `number in [42, 54]` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Left operand does not exist within right operand

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `number in [54, 69]` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: The `nin` operator shall query if the left operand does not exist in the right

    The `nin` operator shall return `true` if the left operand does not exist
    as an element in a sequence defined in the right operand.

    Scenario: Left operand does not exist within right operand

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `number nin [54, 69]` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Left operand exists within right operand

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `number nin [42, 54]` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: The `subsetof` operator shall query if the left operand is a subset of the right

    The `subsetof` operator shall return `true` if the left operand is a
    subset of the right operand.

    Scenario: Left operand is a subset of right operand

      Given the yaml input:
        """
        number: 42
        """
      When the yamlpath `number subsetof [42, 54]` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Left operand is not a subset of right operand
      Given the yaml input:
        """
        number: [54, 100]
        """
      When the yamlpath `number subsetof [54, 69]` is evaluated
      Then the evaluation result is:
        """
        false
        """
