Feature: Bracket path navigation

  The JSONPath specification defines how fields should be accessed.
  Fields are specified between dots, which are used to navigate the
  JSON (in this case, YAML) structure.

  Rule: Brackets with strings shall return matching elements

    Matching fields in bracket expressions produce  shall return the matching elements
    in the YAML structure.

    Scenario: Fields exists and specifies a single scalar value

      Given the yaml input:
        """
        foo:
          bar:
            baz: "hello"
        """
      When the yamlpath `$["foo"]["bar"]["baz"]` is evaluated
      Then the evaluation result is:
        """
        "hello"
        """


    Scenario: Fields exists and specify a complex value

      Given the yaml input:
        """
        foo:
          bar:
            baz: "hello"
        """
      When the yamlpath `$["foo"]["bar"]` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        """

    Scenario: Fields do not exist and return empty input

      Given the yaml input:
        """
        foo:
          bar: 42
          baz:
            buzz: "hello"
        """
      When the yamlpath `$["foo"]["baz"]["buzzz"]` is evaluated
      Then the evaluation result is empty

  Rule: Wildcards shall match all fields

    Wildcards shall match all fields in the YAML structure.

    Scenario: Wildcard matches all subfields

      Given the yaml input:
        """
        foo:
          bar-1:
            baz: "hello"
          bar-2:
            baz: "world"
        """
      When the yamlpath `$["foo"][*]` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        ---
        baz: "world"
        """

  Rule: Index shall match the nth element

    Indexes shall match the nth element in the YAML structure.

    Scenario: Index matches the nth element

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
        """
      When the yamlpath `$["foo"][1]` is evaluated
      Then the evaluation result is:
        """
        bar: "world"
        """

    Scenario: Index does not match any element

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
        """
      When the yamlpath `$["foo"][2]` is evaluated
      Then the evaluation result is empty

    Scenario: Index is negative and matches the nth element from the end

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
        """
      When the yamlpath `$["foo"][-1]` is evaluated
      Then the evaluation result is:
        """
        bar: "world"
        """