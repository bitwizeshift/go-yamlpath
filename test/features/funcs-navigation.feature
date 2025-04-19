Feature: Funcs - Navigation

  YAMLPath provides several groups of functions that can be used for navigating
  subfields of YAML documents.

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

  Rule: children() returns the children of a map

    Returns the children of all direct children of a map node.
    If the input is empty, or contains one element which is not a map, the
    result is empty.
    If the input is not a singleton, an error is raised to the calling
    environment.

    Scenario: Input is not a map

      Given the yaml input:
        """
        42
        """
      When the yamlpath `$.children()` is evaluated
      Then the evaluation result is empty

    Scenario: Input contains a singleton map node

      Given the yaml input:
        """
        people:
          John: 30
          Jane: 25
        """
      When the yamlpath `$.people.children()` is evaluated
      Then the evaluation result is:
        """
        30
        ---
        25
        """

    Scenario: Input contains a singleton sequence node

      Given the yaml input:
        """
        people:
          - John: 30
          - Jane: 25
        """
      When the yamlpath `$.people.children()` is evaluated
      Then the evaluation result is empty

    Scenario: Input is not singleton

      Given the yaml input:
        """
        people:
          John: 30
          Jane: 25
        """
      When the yamlpath `$.people.*.children()` is evaluated
      Then an error is raised

  Rule: descendants() returns the descendants of a map

    Returns the descendants of all recursive
    If the input is empty, or contains one element which is not a map, the
    result is empty.
    If the input is not a singleton, an error is raised to the calling
    environment.

    Scenario: Input is not a map

      Given the yaml input:
        """
        42
        """
      When the yamlpath `$.descendants()` is evaluated
      Then the evaluation result is empty

    Scenario: Input contains a singleton map node

      Given the yaml input:
        """
        people:
          John:
            age: 30
          Jane:
            age: 25
        """
      When the yamlpath `$.descendants()` is evaluated
      Then the evaluation result is:
        """
        John:
          age: 30
        Jane:
          age: 25
        ---
        age: 30
        ---
        30
        ---
        age: 25
        ---
        25
        """

    Scenario: Input is not singleton

      Given the yaml input:
        """
        people:
          John:
            age: 30
          Jane:
            age: 25
        """
      When the yamlpath `$.people.*.descendants()` is evaluated
      Then an error is raised
