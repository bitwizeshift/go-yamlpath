Feature: Filtering

  The JSONPath specification supports filtering via `[?(...)]` where the `...`
  refers to some form of expression that can be evaluated to a boolean value.

  Implementations shall treat the expression as a boolean expression and return
  any elements for which the value is true. If the expression yields more than
  one value, or a value that is not a boolean, the implementation shall treat
  the result as `true` and return the element -- effectively using this as an
  existence check.

  Rule: Filter expressions with path elements returns filtered values

    Filter expressions that return path elements shall return the filtered
    values if the expression is "truthy".

    Scenario: Expression matches a true boolean element

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
          baz:
            enable: false
        """
      When the yamlpath `$.foo.*[?(@.enable)]` is evaluated
      Then the evaluation result is:
        """
        enable: true
        """

    Scenario: Expression matches a false boolean element

      Given the yaml input:
        """
        foo:
          baz:
            enable: false
        """
      When the yamlpath `$.foo.baz[?(@.enable)]` is evaluated
      Then the evaluation result is empty

    Scenario: Expression matches a non-boolean element

      Given the yaml input:
        """
        foo:
          bar:
            baz: 42
        """
      When the yamlpath `$.foo.bar[?(@.baz)]` is evaluated
      Then the evaluation result is:
        """
        baz: 42
        """

  Rule: Negation expressions negate the filter

    Negation expressions shall negate the filter expression.

    Scenario: Negation expression matches a false boolean element

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
        """
      When the yamlpath `$.foo.bar[?(!@.enable)]` is evaluated
      Then the evaluation result is:
        """
        enable: false
        """

    Scenario: Negation expression doesn't match a path element

      Given the yaml input:
        """
        foo:
          bar:
            baz: "hello"
        """
      When the yamlpath `$.foo.bar[?(!@.disable)]` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        """
