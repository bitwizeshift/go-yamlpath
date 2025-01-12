Feature: Funcs - Reflection

  YAMLPath provides support for a `reflect` functionality to introspect the
  results of each YAML node into corresponding types.

  Rule: reflect() returns introspection representation of inputs

    Returns a collection where each node from the input is transformed into an
    introspection representation node on the output, which encodes the `kind`,
    `type`, and `value` of the input node.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].reflect()` is evaluated
      Then the evaluation result is empty

    Scenario Outline: Input is scalar <tag> value

      Given the yaml input:
        """
        input: <value>
        """
      When the yamlpath `$.input.reflect()` is evaluated
      Then the evaluation result is:
        """
        kind: scalar
        tag: '<tag>'
        value: <value>
        """

      Examples:
        | tag     | value    |
        | !!str   | John Doe |
        | !!int   | 30       |
        | !!float | 3.14     |
        | !!bool  | true     |
        | !!null  | null     |

    Scenario: Input is a sequence

      Given the yaml input:
        """
        people:
          - John
          - Jane
        """
      When the yamlpath `$.people.reflect()` is evaluated
      Then the evaluation result is:
        """
        kind: sequence
        tag: '!!seq'
        entries:
        - kind: scalar
          tag: '!!str'
          value: 'John'
        - kind: scalar
          tag: '!!str'
          value: 'Jane'
        """

    Scenario: Input is a mapping

      Given the yaml input:
        """
        people:
          John: 30
          Jane: 25
        """
      When the yamlpath `$.people.reflect()` is evaluated
      Then the evaluation result is:
        """
        kind: mapping
        tag: '!!map'
        entries:
          - key: John
            value:
              kind: scalar
              tag: '!!int'
              value: 30
          - key: Jane
            value:
              kind: scalar
              tag: '!!int'
              value: 25
        """
