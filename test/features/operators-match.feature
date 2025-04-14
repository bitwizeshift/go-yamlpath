Feature: Regex match operator

  JSONPath defines a regex-based match comparison operator. This operator is
  used to match a string against a regular expression.

  The match operator is represented by the '=~' symbol, and is used to compare a
  singleton string against a regular expression. The match operator returns true
  if the string matches the regular expression, and false otherwise. If more
  than one string is provided, the match operator returns false.

  Scenario: Match case sensitive string
    Given the yaml input:
      """
      people:
        - name: "John Doe"
        - name: "Jane Doe"
      """
    When the yamlpath `$.people[*].where(@.name =~ /John Doe/)` is evaluated
    Then the evaluation result is:
      """
      name: "John Doe"
      """

  Scenario: Match case insensitive string
    Given the yaml input:
      """
      people:
        - name: "John Doe"
        - name: "Jane Doe"
      """
    When the yamlpath `$.people[*].where(@.name =~ /john doe/i)` is evaluated
    Then the evaluation result is:
      """
      name: "John Doe"
      """
