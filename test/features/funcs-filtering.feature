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

    Scenario: Criteria matches a true boolean element

      Given the yaml input:
        """
        foo:
          bar:
            enable: true
          baz:
            enable: false
        """
      When the yamlpath `$.foo.*.where(@.enable)` is evaluated
      Then the evaluation result is:
        """
        enable: true
        """

    Scenario: Criteria matches a false boolean element

      Given the yaml input:
        """
        foo:
          baz:
            enable: false
        """
      When the yamlpath `$.foo.baz.where(@.enable)` is evaluated
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

  Rule: keys() returns the keys of the map

    Returns the keys of the map as a collection. If the input is not a map,
    the result is empty.

    Scenario: Input is not a map

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people.keys()` is evaluated
      Then the evaluation result is empty

    Scenario: Map has keys

      Given the yaml input:
        """
        people:
          John: 30
          Jane: 25
        """
      When the yamlpath `$.people.keys()` is evaluated
      Then the evaluation result is:
        """
        "John"
        ---
        "Jane"
        """

    Scenario: Recursive-descent of map with keys

      Given the yaml input:
        """
        people:
          John:
            age: 30
          Jane:
            age: 25
        """
      When the yamlpath `$..keys()` is evaluated
      Then the evaluation result is:
        """
        "people"
        ---
        "John"
        ---
        "Jane"
        ---
        "age"
        ---
        "age"
        """

  Rule: select() returns matching keys and indices

    Returns the keys and indices of the map and list as a collection. If a node
    is not a map or list, the element is ignored for inclusion in the result.

    Scenario: Input is not a map or list

      Given the yaml input:
        """
        people: 30
        """
      When the yamlpath `$.people.select(0)` is evaluated
      Then the evaluation result is empty

    Scenario: Map has keys

      Given the yaml input:
        """
        people:
          alice: 30
          bob: 25
        """
      When the yamlpath `$.people.select("alice", "bob")` is evaluated
      Then the evaluation result is:
        """
        30
        ---
        25
        """

    Scenario: List has indices

      Given the yaml input:
        """
        people:
          - Alice
          - Bob
          - Charlie
        """
      When the yamlpath `$.people.select(0, 2)` is evaluated
      Then the evaluation result is:
        """
        "Alice"
        ---
        "Charlie"
        """
