Feature: Index navigation

  The YAMLPath specification defines that `[...]` shall be used for indexing,
  where the contents of the brackets are evaluated as different forms of
  indexing expressions.

  The behavior is as follows:

  - `[*]` matches all elements in a sequence
  - `[<expr>]` matches any numeric indices of a sequence node, where `<expr>`
    evaluates as a subexpression that returns number(s).
  - `[<slice>]` matches a slice of elements in a sequence node, where
    `<slice>` is defined as `<start>:[<end>][:<step>]`.

  Rule: Wildcards shall match all fields of a sequence

    Wildcards shall match all sequence fields in the collection.
    Mapping nodes shall be left untouched.

    Scenario: Wildcard matches all subfields

      Given the yaml input:
        """
        foo:
          - baz: "hello"
          - baz: "world"
        """
      When the yamlpath `$.foo[*]` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        ---
        baz: "world"
        """

  Rule: Expressions yielding indices shall match the corresponding sequence element

    Expression yielding indices shall match each of the corresponding elements
    from a sequence node. Non-sequence nodes shall be left untouched.
    If a negative index is evaluated, it shall be treated as an index from the
    end of the sequence.

    Scenario: Index matches the nth element

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
        """
      When the yamlpath `$.foo[1]` is evaluated
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
      When the yamlpath `$.foo[2]` is evaluated
      Then the evaluation result is empty

    Scenario: Index is negative and matches the nth element from the end

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
        """
      When the yamlpath `$.foo[-1]` is evaluated
      Then the evaluation result is:
        """
        bar: "world"
        """

    Scenario: Multiple indices are provided

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
          - bar: "goodbye"
        """
      When the yamlpath `$.foo[0 | 2]` is evaluated
      Then the evaluation result is:
        """
        bar: "hello"
        ---
        bar: "goodbye"
        """

  Rule: Slices shall match a range of elements

    Scenario: Slice specifies only start index

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
          - bar: "goodbye"
        """
      When the yamlpath `$.foo[1:]` is evaluated
      Then the evaluation result is:
        """
        bar: "world"
        ---
        bar: "goodbye"
        """

    Scenario: Slice specifies start and end index

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
          - bar: "goodbye"
        """
      When the yamlpath `$.foo[1:2]` is evaluated
      Then the evaluation result is:
        """
        bar: "world"
        """

    Scenario: Slice specifies end

      Given the yaml input:
        """
        foo:
          - bar: "hello"
          - bar: "world"
          - bar: "goodbye"
        """
      When the yamlpath `$.foo[:2]` is evaluated
      Then the evaluation result is:
        """
        bar: "hello"
        ---
        bar: "world"
        """

    Scenario: Slice specifies end and step index

      Given the yaml input:
        """
        foo:
          - bar: 1
          - bar: 2
          - bar: 3
          - bar: 4
        """
      When the yamlpath `$.foo[:3:2]` is evaluated
      Then the evaluation result is:
        """
        bar: 1
        ---
        bar: 3
        """

    Scenario: Slice specifies step

      Given the yaml input:
        """
        foo:
          - bar: 1
          - bar: 2
          - bar: 3
          - bar: 4
        """
      When the yamlpath `$.foo[::2]` is evaluated
      Then the evaluation result is:
        """
        bar: 1
        ---
        bar: 3
        """

    Scenario: Slice specifies start, end, and step index

      Given the yaml input:
        """
        foo:
          - bar: 1
          - bar: 2
          - bar: 3
          - bar: 4
        """
      When the yamlpath `$.foo[0:3:2]` is evaluated
      Then the evaluation result is:
        """
        bar: 1
        ---
        bar: 3
        """
