Feature: Funcs - Subsetting

  YAMLPath provides support for various subsetting functionalities to retrieve
  smaller parts of a returned collection. This contrasts the builtin slice
  access functionality, which is used to retrieve a range of elements from a
  sequence instead of a collection.

  Rule: single() returns the only element of the collection

    Returns the only element of a collection, if a collection contains a single
    value. If the collection contains more than one value, it signals an error
    to the calling environment. If the collection is empty, it returns an
    empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].single()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains a single element

      Given the yaml input:
        """
        people:
          - name: "John"
        """
      When the yamlpath `$.people[*].single()` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].single()` is evaluated
      Then an error is raised

  Rule: first() returns the first n elements of the collection

    Returns the first n elements of a collection, where n is a positive integer.
    If the collection is empty, it returns an empty collection.
    If n is not specified, 1 is assumed.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].first()` is evaluated
      Then the evaluation result is empty

    Scenario: Count not specified

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].first()` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        """

    Scenario: Collection contains fewer elements than requested

      Given the yaml input:
        """
        people:
          - name: "John"
        """
      When the yamlpath `$.people[*].first(2)` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        """

    Scenario: Collection contains more elements than requested

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
          - name: "Bob"
        """
      When the yamlpath `$.people[*].first(2)` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        ---
        name: "Jane"
        """

  Rule: last() returns the last n elements of the collection

    Returns the last n elements of a collection, where n is a positive integer.
    If the collection is empty, it returns an empty collection.
    If n is not specified, 1 is assumed.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].last()` is evaluated
      Then the evaluation result is empty

    Scenario: Count not specified

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].last()` is evaluated
      Then the evaluation result is:
        """
        name: "Jane"
        """

    Scenario: Collection contains fewer elements than requested

      Given the yaml input:
        """
        people:
          - name: "John"
        """
      When the yamlpath `$.people[*].last(2)` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        """

    Scenario: Collection contains more elements than requested

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
          - name: "Bob"
        """
      When the yamlpath `$.people[*].last(2)` is evaluated
      Then the evaluation result is:
        """
        name: "Jane"
        ---
        name: "Bob"
        """

  Rule: skip() skips n elements of the collection

    Returns the collection with the first n elements removed, if n is positive.
    If n is negative, it skips the last n elements. If the collection is empty,
    it returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].skip(1)` is evaluated
      Then the evaluation result is empty

    Scenario: Positive n, within range of collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
          - name: "Bob"
        """
      When the yamlpath `$.people[*].skip(1)` is evaluated
      Then the evaluation result is:
        """
        name: "Jane"
        ---
        name: "Bob"
        """

    Scenario: Positive n, outside range of collection

      Given the yaml input:
        """
        people:
          - name: "John"
        """
      When the yamlpath `$.people[*].skip(2)` is evaluated
      Then the evaluation result is empty

    Scenario: Negative n, within range of collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
          - name: "Bob"
        """
      When the yamlpath `$.people[*].skip(-1)` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        ---
        name: "Jane"
        """

    Scenario: Negative n, outside range of collection

      Given the yaml input:
        """
        people:
          - name: "John"
        """
      When the yamlpath `$.people[*].skip(-2)` is evaluated
      Then the evaluation result is empty
