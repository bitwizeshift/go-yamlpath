Feature: Funcs - Existence

  YAMLPath provides support for various "existence"-testing functions that
  evaluate the state of fields in the YAML document.

  Rule: empty() evaluates whether the collection is empty

    Returns true if the input collection is empty and false otherwise.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].empty()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection is non-empty

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].empty()` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: exists() evaluates the existence of fields

    The `exists()` function shall return true if the collection has any
    elements, and false otherwise. This is the opposite of empty(), and as such
    is a shorthand for `!(@.empty())`.

    If the input collection is empty, the result is false.

    The function can also take an optional criteria to be applied to the
    collection prior to the determination of the exists. In this case, the
    function is shorthand for where(criteria).exists().

    Scenario: No criteria on non-empty collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].exists()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: No criteria on empty collection

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].exists()` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Criteria on non-empty collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].exists(@.name == "John")` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Criteria on empty collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].exists(@.name == "Mary")` is evaluated
      Then the evaluation result is:
        """
        false
        """
