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

  Rule: startsWith() checks if the string starts with a prefix

    Shall return true if the string starts with the given prefix. If the value
    is not a string, or if there are more than one value in the collection, an
    error is raised to the calling environment. If the 'prefix' argument is not
    a string, an error is raised to the calling environment. If the collection
    is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].startsWith("john")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.startsWith("john")` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Collection contains string that does match

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.startsWith("jo")` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].startsWith("john")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.startsWith("john")` is evaluated
      Then an error is raised

    Scenario: Prefix argument is not a string

      The prefix argument must be a string. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.startsWith(123)` is evaluated
      Then an error is raised

  Rule: endsWith() checks if the string ends with a suffix

    Shall return true if the string ends with the given suffix. If the value is
    not a string, or if there are more than one value in the collection, an
    error is raised to the calling environment. If the 'suffix' argument is not
    a string, an error is raised to the calling environment. If the collection
    is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].endsWith("john")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.endsWith("john")` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Collection contains string that does match

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.endsWith("hn")` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].endsWith("john")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.endsWith("john")` is evaluated
      Then an error is raised

    Scenario: Suffix argument is not a string

      The suffix argument must be a string. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.endsWith(123)` is evaluated
      Then an error is raised

  Rule: contains() checks if the string contains a substring

    Shall return true if the string contains the given substring. If the value
    is not a string, or if there are more than one value in the collection, an
    error is raised to the calling environment. If the 'substring' argument is
    not a string, an error is raised to the calling environment. If the
    collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].contains("john")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.contains("john")` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Collection contains string that does match

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.contains("oh")` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].contains("john")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.contains("john")` is evaluated
      Then an error is raised

    Scenario: Substring argument is not a string

      The substring argument must be a string. If it is not, an error is raised
      to the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.contains(123)` is evaluated
      Then an error is raised

  Rule: indexOf() returns the index of the first occurrence of a substring

    Shall return the index of the first occurrence of the substring in the
    string. If the value is not a string, or if there are more than one value
    in the collection, an error is raised to the calling environment. If the
    'substring' argument is not a string, an error is raised to the calling
    environment. If the collection is empty, this returns empty. If the
    substring is not found, -1 is returned.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].indexOf("john")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: ""
        """
      When the yamlpath `$.name.indexOf("john")` is evaluated
      Then the evaluation result is:
        """
        -1
        """

    Scenario: Collection contains string that does match

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.indexOf("oh")` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].indexOf("john")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.indexOf("john")` is evaluated
      Then an error is raised

    Scenario: Substring argument is not a string

      The substring argument must be a string. If it is not, an error is raised
      to the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.indexOf(123)` is evaluated
      Then an error is raised

  Rule: substring() returns a substring of the string

    Shall return a substring of the string. If the value is not a string, or if
    there are more than one value in the collection, an error is raised to the
    calling environment. If the 'start' argument is not an integer, an error is
    raised to the calling environment. If the 'length' argument is not an
    integer, an error is raised to the calling environment. If the collection
    is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].substring(0, 2)` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string, index and length are within string

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.substring(1, 2)` is evaluated
      Then the evaluation result is:
        """
        oh
        """

    Scenario: Collection contains string, index is too long

      If the index argument excees the length of the string, an empty collection
      is returned.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.substring(10, 2)` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string, index is within string, length is too long

      If the length argument is too long, the substring will be truncated to the
      end of the string.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.substring(1, 10)` is evaluated
      Then the evaluation result is:
        """
        ohn
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].substring(0, 2)` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.substring(0, 2)` is evaluated
      Then an error is raised

    Scenario: Start argument is not an integer

      The start argument must be an integer. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.substring("0")` is evaluated
      Then an error is raised

    Scenario: Length argument is not an integer

      The length argument must be an integer. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.substring(0, "2")` is evaluated
      Then an error is raised

  Rule: replace() replaces all occurrences of a pattern with a replacement

    Shall return a string with all occurrences of the pattern replaced with the
    replacement. If the value is not a string, or if there are more than one
    value in the collection, an error is raised to the calling environment. If
    the 'pattern' argument is not a string, an error is raised to the calling
    environment. If the 'replacement' argument is not a string, an error is
    raised to the calling environment. If the collection is empty, this returns
    empty. If the pattern is not found, the original string is returned.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].replace("john", "jane")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: "hello world"
        """
      When the yamlpath `$.name.replace("john", "jane")` is evaluated
      Then the evaluation result is:
        """
        "hello world"
        """

    Scenario: Collection contains string that does match
      Given the yaml input:
        """
        name: "hello john"
        """
      When the yamlpath `$.name.replace("john", "jane")` is evaluated
      Then the evaluation result is:
        """
        hello jane
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].replace("john", "jane")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.replace("john", "jane")` is evaluated
      Then an error is raised

    Scenario: Pattern argument is not a string

      The pattern argument must be a string. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.replace(123, "jane")` is evaluated
      Then an error is raised

    Scenario: Replacement argument is not a string

      The replacement argument must be a string. If it is not, an error is
      raised to the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.replace("john", 123)` is evaluated
      Then an error is raised

  Rule: length() returns the length of the string

    Shall return the length of the string. If the value is not a string, or if
    there are more than one value in the collection, an error is raised to the
    calling environment. If the collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].length()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.length()` is evaluated
      Then the evaluation result is:
        """
        4
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].length()` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.length()` is evaluated
      Then an error is raised

  Rule: split() splits a string into a sequence of strings

    Shall return a sequence of strings, split by the given separator. If the
    value is not a string, or if there are more than one value in the collection,
    an error is raised to the calling environment. If the 'separator' argument
    is not a string, an error is raised to the calling environment. If the
    collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].split(",")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string

      Given the yaml input:
        """
        name: "john,jane"
        """
      When the yamlpath `$.name.split(",")` is evaluated
      Then the evaluation result is:
        """
        "john"
        ---
        "jane"
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john,jane"
          - name: "jane,john"
        """
      When the yamlpath `$.people[*].split(",")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.split(",")` is evaluated
      Then an error is raised

    Scenario: Separator argument is not a string

      The separator argument must be a string. If it is not, an error is raised
      to the calling environment.

      Given the yaml input:
        """
        name: "john,jane"
        """
      When the yamlpath `$.name.split(123)` is evaluated
      Then an error is raised

  Rule: toChars() returns a sequence of characters

    Shall return a sequence of characters, each character being a string of
    length 1. If the value is not a string, or if there are more than one value
    in the collection, an error is raised to the calling environment. If the
    collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].toChars()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.toChars()` is evaluated
      Then the evaluation result is:
        """
        "j"
        ---
        "o"
        ---
        "h"
        ---
        "n"
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].toChars()` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.toChars()` is evaluated
      Then an error is raised

  Rule: matches() checks if the string matches a regular expression

    Shall return true if the string matches the regular expression. If the
    value is not a string, or if there are more than one value in the collection,
    an error is raised to the calling environment. If the 'regex' argument is
    not a string, an error is raised to the calling environment. If the
    collection is empty, this returns empty.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].matches("john")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: "hello world"
        """
      When the yamlpath `$.name.matches("john")` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Collection contains string that does match

      Given the yaml input:
        """
        name: "hello john"
        """
      When the yamlpath `$.name.matches("^h.*lo")` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Collection contains more than one string

      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].matches("john")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value

      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.matches("john")` is evaluated
      Then an error is raised

    Scenario: Regex argument is not a string

      The regex argument must be a string. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.matches(123)` is evaluated
      Then an error is raised

  Rule: replaceMatches() replaces all occurrences of a pattern with a replacement

    Shall return a string with all occurrences of the pattern replaced with the
    replacement. If the value is not a string, or if there are more than one
    value in the collection, an error is raised to the calling environment. If
    the 'pattern' argument is not a string, an error is raised to the calling
    environment. If the 'replacement' argument is not a string, an error is
    raised to the calling environment. If the collection is empty, this returns
    empty. If the pattern is not found, the original string is returned.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].replaceMatches("john", "jane")` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains string that does not match

      Given the yaml input:
        """
        name: "hello world"
        """
      When the yamlpath `$.name.replaceMatches("john", "jane")` is evaluated
      Then the evaluation result is:
        """
        "hello world"
        """

    Scenario: Collection contains string that does match
      Given the yaml input:
        """
        name: "hello john"
        """
      When the yamlpath `$.name.replaceMatches("j.*n$", "jane")` is evaluated
      Then the evaluation result is:
        """
        hello jane
        """

    Scenario: Collection contains more than one string
      Collections containing more than one string element shall raise an error
      to the calling environment

      Given the yaml input:
        """
        people:
          - name: "john"
          - name: "jane"
        """
      When the yamlpath `$.people[*].replaceMatches("john", "jane")` is evaluated
      Then an error is raised

    Scenario: Collection contains non-string value
      Collections containing non-string values shall raise an error to the
      calling environment

      Given the yaml input:
        """
        name: 123
        """
      When the yamlpath `$.name.replaceMatches("john", "jane")` is evaluated
      Then an error is raised

    Scenario: Pattern argument is not a string
      The pattern argument must be a string. If it is not, an error is raised to
      the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.replaceMatches(123, "jane")` is evaluated
      Then an error is raised

    Scenario: Replacement argument is not a string
      The replacement argument must be a string. If it is not, an error is
      raised to the calling environment.

      Given the yaml input:
        """
        name: "john"
        """
      When the yamlpath `$.name.replaceMatches("john", 123)` is evaluated
      Then an error is raised
