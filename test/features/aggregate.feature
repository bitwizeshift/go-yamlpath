Feature: Aggregations within YAMLPath expressions

  YAMLPath shall accept, as a sub-expression, construction of ephemeral
  aggregate types such as sequences or maps as an extension to support
  better queries.

  Rule: List aggregations produce ephemeral lists

    Scenario: A list aggregation produces an ephemeral list
      Given the yaml input:
        """
        - name: "Bjoern"
          gender: "M"
        - name: "Alice"
          gender: "N/A"
        """
      When the yamlpath `$[*][?(@.gender in ["M", "F"])]` is evaluated
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
      When the yamlpath `$[?($ == { "name": "Bjoern", "age": 42 })]` is evaluated
      Then the evaluation result is:
        """
        name: "Bjoern"
        age: 42
        """
