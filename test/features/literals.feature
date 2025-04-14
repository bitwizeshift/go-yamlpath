Feature: Literals in YAMLPath expressions

  YAMLPath shall accept literals in expressions, including:

  - String literals
  - Integer literals
  - Float literals
  - Null literals
  - Boolean literals
  - Sequence literals
  - Mapping literals

  Rule: List aggregations produce ephemeral lists

    Scenario: A list aggregation produces an ephemeral list
      Given the yaml input:
        """
        - name: "Bjoern"
          gender: "M"
        - name: "Alice"
          gender: "N/A"
        """
      When the yamlpath `$[*].where(@.gender in ["M", "F"])` is evaluated
      Then the evaluation result is:
        """
        name: "Bjoern"
        gender: "M"
        """

  Rule: Map aggregations produce ephemeral maps

    Scenario: A map aggregation produces an ephemeral map
      Given the yaml input:
        """
        name: "Bjoern"
        age: 42
        """
      When the yamlpath `$.where($ == { "name": "Bjoern", "age": 42 })` is evaluated
      Then the evaluation result is:
        """
        name: "Bjoern"
        age: 42
        """

  Rule: Scalar literals produce scalar values

    Scenario: String literal
      Given the yaml input:
        """
        value: "hello"
        """
      When the yamlpath `value == "hello"` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Integer literal
      Given the yaml input:
        """
        value: 42
        """
      When the yamlpath `value == 42` is evaluated
      Then the evaluation result is:
        """
        true
        """
    Scenario: Float literal
      Given the yaml input:
        """
        value: 42.0
        """
      When the yamlpath `value == 42.0` is evaluated
      Then the evaluation result is:
        """
        true
        """
    Scenario: Null literal
      Given the yaml input:
        """
        value: null
        """
      When the yamlpath `value == null` is evaluated
      Then the evaluation result is:
        """
        true
        """
    Scenario: Boolean literal
      Given the yaml input:
        """
        value: true
        """
      When the yamlpath `value == true` is evaluated
      Then the evaluation result is:
        """
        true
        """
