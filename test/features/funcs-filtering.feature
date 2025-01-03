Feature: Funcs - Filtering

  YAMLPath provides support for various "filter"-testing functions that
  evaluate the state of fields in the YAML document.

  Rule: where() filters the collection based on a criteria

    Returns a collection containing only those elements in the input collection
    for which the stated criteria expression evaluates to true. Elements for
    which the expression evaluates to false or empty are not included in
    the result.

    If the input collection is empty, the result is empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].where(@.name == "John")` is evaluated
      Then the evaluation result is empty

    Scenario: Criteria matches some elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].where(@.name == "John")` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        """

    Scenario: Criteria matches no elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].where(@.name == "Mary")` is evaluated
      Then the evaluation result is empty

  Rule: transform() transforms the collection to the projection

    Transforms the collection based on the projection expression provided
    If the input collection is empty, the result is empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].transform(@.name)` is evaluated
      Then the evaluation result is empty

    Scenario: Projection transforms the collection

      Given the yaml input:
        """
        people:
          - name: "John"
            age: 30
          - name: "Jane"
            age: 25
        """
      When the yamlpath `$.people[*].transform(@.name + " Doe")` is evaluated
      Then the evaluation result is:
        """
        "John Doe"
        ---
        "Jane Doe"
        """
