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

  Rule: Boolean "and" expressions evaluate combinations of filters

    Boolean "and" expressions shall evaluate both left and right sub-expressions
    and only return elements where both sub-expressions evaluate to true.

    Scenario: Left sub-expression evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
            allow: true
        """
      When the yamlpath `$.foo.bar[?(@.enable && @.allow)]` is evaluated
      Then the evaluation result is empty

    Scenario: Right sub-expression evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
            allow: false
        """
      When the yamlpath `$.foo.bar[?(@.enable && @.allow)]` is evaluated
      Then the evaluation result is empty

    Scenario: Both sub-expressions evaluate false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
            allow: false
        """
      When the yamlpath `$.foo.bar[?(@.enable && @.allow)]` is evaluated
      Then the evaluation result is empty

    Scenario: Both sub-expressions evaluate true

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
            allow: true
        """
      When the yamlpath `$.foo.bar[?(@.enable && @.allow)]` is evaluated
      Then the evaluation result is:
        """
        enable: true
        allow: true
        """

  Rule: Boolean "or" expressions evaluate combinations of filters

    Boolean "or" expressions shall evaluate both left and right sub-expressions
    and return elements where either sub-expression evaluates to true.

    Scenario: Left sub-expression evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
            allow: true
        """
      When the yamlpath `$.foo.bar[?(@.enable || @.allow)]` is evaluated
      Then the evaluation result is:
        """
        enable: false
        allow: true
        """

    Scenario: Right sub-expression evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
            allow: false
        """
      When the yamlpath `$.foo.bar[?(@.enable || @.allow)]` is evaluated
      Then the evaluation result is:
        """
        enable: true
        allow: false
        """

    Scenario: Both sub-expressions evaluate false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
            allow: false
        """
      When the yamlpath `$.foo.bar[?(@.enable || @.allow)]` is evaluated
      Then the evaluation result is empty

    Scenario: Both sub-expressions evaluate true

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
            allow: true
        """
      When the yamlpath `$.foo.bar[?(@.enable || @.allow)]` is evaluated
      Then the evaluation result is:
        """
        enable: true
        allow: true
        """

  Rule: Equality operator evaluates equality of path values

    The equality operator shall evaluate the equality of the left and right
    sub-expressions. If the returned objects recursively compare "equal",
    the expression shall evaluate to true and be included in the output set.

    Scenario: Equality evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
        """
      When the yamlpath `$.foo.bar[?(@.enable == true)]` is evaluated
      Then the evaluation result is empty

    Scenario: Equality evaluates true

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
        """
      When the yamlpath `$.foo.bar[?(@.enable == false)]` is evaluated
      Then the evaluation result is:
        """
        enable: false
        """

    Scenario: Complex equality evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
        """
      When the yamlpath `$.foo[?(@.bar == true)]` is evaluated
      Then the evaluation result is empty

    Scenario: Identity equality evaluates true

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
        """
      When the yamlpath `$.foo[?(@.bar == @.bar)]` is evaluated
      Then the evaluation result is:
        """
        bar:
          enable: true
        """

  Rule: Inequality operator evaluates inequality of path values

    The inequality operator shall evaluate the inequality of the left and right
    sub-expressions. If the returned objects recursively compare "not equal",
    the expression shall evaluate to true and be included in the output set.

    Scenario: Inequality evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
        """
      When the yamlpath `$.foo.bar[?(@.enable != false)]` is evaluated
      Then the evaluation result is empty

    Scenario: Inequality evaluates true

      Given the yaml input:
        """
        foo:
          bar:
            enable: false
        """
      When the yamlpath `$.foo.bar[?(@.enable != true)]` is evaluated
      Then the evaluation result is:
        """
        enable: false
        """

    Scenario: Complex inequality evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
        """
      When the yamlpath `$.foo.bar[?(@.enable != true)]` is evaluated
      Then the evaluation result is empty

    Scenario: Identity inequality evaluates false

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
        """
      When the yamlpath `$.foo[?(@.bar != @.bar)]` is evaluated
      Then the evaluation result is empty

  Rule: Less-than operator evaluates less-than of path values

    The less-than operator shall evaluate the less-than relationship of the left
    and right sub-expressions. If the returned objects recursively compare as
    "less-than", the expression shall evaluate to true and be included in the
    output set.

    Scenario: Scalar left value is greater than scaler right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value < 0)]` is evaluated
      Then the evaluation result is empty

    Scenario: Scalar left value is less than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value < 100)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """

  Rule: Less-than-or-equal operator evaluates less-than-or-equal of path values

    The less-than-or-equal operator shall evaluate the less-than-or-equal
    relationship of the left and right sub-expressions. If the returned objects
    recursively compare as "less-than-or-equal", the expression shall evaluate
    to true and be included in the output set.

    Scenario: Scalar left value is greater than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value <= 0)]` is evaluated
      Then the evaluation result is empty

    Scenario: Scalar left value is less than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value <= 100)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """

    Scenario: Scalar left value is equal to scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value <= 10)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """

  Rule: Greater-than operator evaluates greater-than of path values

    The greater-than operator shall evaluate the greater-than relationship of
    the left and right sub-expressions. If the returned objects recursively
    compare as "greater-than", the expression shall evaluate to true and be
    included in the output set.

    Scenario: Scalar left value is less than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value > 100)]` is evaluated
      Then the evaluation result is empty

    Scenario: Scalar left value is greater than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value > 0)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """

  Rule: Greater-than-or-equal operator evaluates greater-than-or-equal of path values

    The greater-than-or-equal operator shall evaluate the greater-than-or-equal
    relationship of the left and right sub-expressions. If the returned objects
    recursively compare as "greater-than-or-equal", the expression shall evaluate
    to true and be included in the output set.

    Scenario: Scalar left value is less than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value >= 100)]` is evaluated
      Then the evaluation result is empty

    Scenario: Scalar left value is greater than scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value >= 0)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """

    Scenario: Scalar left value is equal to scalar right value

      Given the yaml input:
        """
        foo:
          bar:
            value: 10
        """
      When the yamlpath `$.foo.bar[?(@.value >= 10)]` is evaluated
      Then the evaluation result is:
        """
        value: 10
        """
