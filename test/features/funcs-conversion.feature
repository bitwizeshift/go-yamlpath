Feature: Funcs - Conversion

  YAMLPath shall provides conversion functionality between singleton collection
  objects and native YAML-type values.

  Rule: toString() converts a single value in the collection to a string

    Converts a single value in the collection to a string. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toString()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single string

      Given the yaml input:
        """
        name: "John"
        """
      When the yamlpath `$.name.toString()` is evaluated
      Then the evaluation result is:
        """
        "John"
        """

    Scenario: Collection contains single integer

      Given the yaml input:
        """
        age: 30
        """
      When the yamlpath `$.age.toString()` is evaluated
      Then the evaluation result is:
        """
        "30"
        """

    Scenario: Collection contains single float

      Given the yaml input:
        """
        height: 1.8
        """
      When the yamlpath `$.height.toString()` is evaluated
      Then the evaluation result is:
        """
        "1.8"
        """

    Scenario: Collection contains single boolean

      Given the yaml input:
        """
        is_active: true
        """
      When the yamlpath `$.is_active.toString()` is evaluated
      Then the evaluation result is:
        """
        "true"
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].toString()` is evaluated
      Then an error is raised

    Scenario: Collection contains single null

      Given the yaml input:
        """
        name: null
        """
      When the yamlpath `$.name.toString()` is evaluated
      Then the evaluation result is:
        """
        "null"
        """

  Rule: toBoolean() converts a single value in the collection to a boolean

    Converts a single value in the collection to a boolean. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toBoolean()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single string

      Given the yaml input:
        """
        is_active: "true"
        """
      When the yamlpath `$.is_active.toBoolean()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains single integer

      Given the yaml input:
        """
        is_active: 1
        """
      When the yamlpath `$.is_active.toBoolean()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains single float

      Given the yaml input:
        """
        is_active: 1.0
        """
      When the yamlpath `$.is_active.toBoolean()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains single boolean

      Given the yaml input:
        """
        is_active: true
        """
      When the yamlpath `$.is_active.toBoolean()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].toBoolean()` is evaluated
      Then an error is raised

  Rule: toNumber() converts a single value in the collection to a number

    Converts a single value in the collection to a number. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toNumber()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single string

      Given the yaml input:
        """
        age: "30"
        """
      When the yamlpath `$.age.toNumber()` is evaluated
      Then the evaluation result is:
        """
        30
        """

    Scenario: Collection contains single integer

      Given the yaml input:
        """
        age: 30
        """
      When the yamlpath `$.age.toNumber()` is evaluated
      Then the evaluation result is:
        """
        30
        """

    Scenario: Collection contains single float

      Given the yaml input:
        """
        height: 1.8
        """
      When the yamlpath `$.height.toNumber()` is evaluated
      Then the evaluation result is:
        """
        1.8
        """

    Scenario: Collection contains single boolean

      Given the yaml input:
        """
        is_active: true
        """
      When the yamlpath `$.is_active.toNumber()` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].toNumber()` is evaluated
      Then an error is raised

  Rule: toInteger() converts a single value in the collection to an integer

    Converts a single value in the collection to an integer. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toInteger()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single string

      Given the yaml input:
        """
        age: "30"
        """
      When the yamlpath `$.age.toInteger()` is evaluated
      Then the evaluation result is:
        """
        30
        """

    Scenario: Collection contains single integer

      Given the yaml input:
        """
        age: 30
        """
      When the yamlpath `$.age.toInteger()` is evaluated
      Then the evaluation result is:
        """
        30
        """

    Scenario: Collection contains single boolean

      Given the yaml input:
        """
        is_active: true
        """
      When the yamlpath `$.is_active.toInteger()` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].toInteger()` is evaluated
      Then an error is raised

  Rule: toFloat() converts a single value in the collection to a float

    Converts a single value in the collection to a float. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toFloat()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single string

      Given the yaml input:
        """
        height: "1.8"
        """
      When the yamlpath `$.height.toFloat()` is evaluated
      Then the evaluation result is:
        """
        1.8
        """

    Scenario: Collection contains single integer

      Given the yaml input:
        """
        height: 1
        """
      When the yamlpath `$.height.toFloat()` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection contains single float

      Given the yaml input:
        """
        height: 1.8
        """
      When the yamlpath `$.height.toFloat()` is evaluated
      Then the evaluation result is:
        """
        1.8
        """

    Scenario: Collection contains single boolean

      Given the yaml input:
        """
        height: true
        """
      When the yamlpath `$.height.toFloat()` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].toFloat()` is evaluated
      Then an error is raised

  Rule: toSequence() converts a single value in the collection to a sequence

    Converts a single value in the collection to a sequence. If the collection
    contains more than one value, an error is raised to the calling environment.
    If the collection is empty, this returns an empty collection.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toSequence()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single value

      Given the yaml input:
        """
        name: "John"
        """
      When the yamlpath `$.name.toSequence()` is evaluated
      Then the evaluation result is:
        """
        - "John"
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].name.toSequence()` is evaluated
      Then the evaluation result is:
        """
        - "John"
        - "Jane"
        """
