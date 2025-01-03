Feature: Funcs - String

  YAMLPath shall provides string-manipulation functions to transform the
  string values in the YAML document.

  Rule: upper() converts the string to uppercase

    Shall return the string in uppercase. If the value is not a string, or if
    there are more than one value in the collection, an error is raised to the
    calling environment.
    If the collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].upper()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single empty string

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.upper()` is evaluated
      Then the evaluation result is:
        """
        ""
        """

    Scenario: Collection contains single non-empty string

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.upper()` is evaluated
      Then the evaluation result is:
        """
        JOHN
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].upper()` is evaluated
      Then an error is raised

    Scenario: Collection contains single non-string value

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.upper()` is evaluated
      Then an error is raised

  Rule: lower() converts the string to lowercase

    Shall return the string in lowercase. If the value is not a string, or if
    there are more than one value in the collection, an error is raised to the
    calling environment.
    If the collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].lower()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains single empty string

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.lower()` is evaluated
      Then the evaluation result is:
        """
        ""
        """

    Scenario: Collection contains single non-empty string

      Given the yaml input:
        """
        name: "JOHN"
        """
      When the yamlpath `$.name.lower()` is evaluated
      Then the evaluation result is:
        """
        john
        """

    Scenario: Collection contains multiple elements

      Given the yaml input:
        """
        people:
          - name: "JOHN"
          - name: "JANE"
        """
      When the yamlpath `$.people[*].lower()` is evaluated
      Then an error is raised

    Scenario: Collection contains single non-string value

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.lower()` is evaluated
      Then an error is raised
