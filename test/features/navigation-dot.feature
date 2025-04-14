Feature: Dot path navigation

  The JSONPath specification defines how fields should be accessed.
  Fields are specified between dots, which are used to navigate the
  JSON (in this case, YAML) structure.

  Rule: Path selection with matching fields shall return matching elements

    Matching fields in path selection shall return the matching elements
    in the YAML structure.

    Scenario: Fields exists and specifies a single scalar value

      Given the yaml input:
        """
        foo:
          bar:
            baz: "hello"
        """
      When the yamlpath `$.foo.bar.baz` is evaluated
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
      When the yamlpath `$.foo.bar` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        """

    Scenario: Field name is complex, exists and specifies a value

      Given the yaml input:
        """
        "foo bar baz": 42
        """
      When the yamlpath `$."foo bar baz"` is evaluated
      Then the evaluation result is:
        """
        42
        """

    Scenario: Fields do not exist and return empty input

      Given the yaml input:
        """
        foo:
          bar: 42
          baz:
            buzz: "hello"
        """
      When the yamlpath `$.foo.baz.buzzz` is evaluated
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
      When the yamlpath `$.foo.*` is evaluated
      Then the evaluation result is:
        """
        baz: "hello"
        ---
        baz: "world"
        """

  Rule: Recursive descent shall return all fields

    Recursive descent shall match all fields in the YAML structure, flattening
    and returning all elements along the way.

    Scenario: Recursive descent matches all subfields

      Given the yaml input:
        """
        store:
          book:
            - category: fiction
              author: Author A
              title: Book A
            - category: non-fiction
              author: Author B
              title: Book B
          bicycle:
            color: red
            price: 19.95
        expensive: 50
        """
      When the yamlpath `$..` is evaluated
      Then the evaluation result is:
        """
        store:
          book:
            - category: fiction
              author: Author A
              title: Book A
            - category: non-fiction
              author: Author B
              title: Book B
          bicycle:
            color: red
            price: 19.95
        expensive: 50
        ---
        book:
          - category: fiction
            author: Author A
            title: Book A
          - category: non-fiction
            author: Author B
            title: Book B
        bicycle:
          color: red
          price: 19.95
        ---
          - category: fiction
            author: Author A
            title: Book A
          - category: non-fiction
            author: Author B
            title: Book B
        ---
        category: fiction
        author: Author A
        title: Book A
        ---
        "fiction"
        ---
        "Author A"
        ---
        "Book A"
        ---
        category: non-fiction
        author: Author B
        title: Book B
        ---
        "non-fiction"
        ---
        "Author B"
        ---
        "Book B"
        ---
        color: red
        price: 19.95
        ---
        "red"
        ---
        19.95
        ---
        50
        """

    Scenario: Recursive descent with trailing fields filter matching entries

      Given the yaml input:
        """
        store:
          book:
            - category: fiction
              author: Author A
              title: Book A
            - category: non-fiction
              author: Author B
              title: Book B
          bicycle:
            color: red
            price: 19.95
        expensive: 50
        """
      When the yamlpath `$..author` is evaluated
      Then the evaluation result is:
        """
        "Author A"
        ---
        "Author B"
        """

    Scenario: Recursive descent with wildcard matches all subfields

      Given the yaml input:
        """
        store:
          book:
            - category: fiction
              author: Author A
              title: Book A
            - category: non-fiction
              author: Author B
              title: Book B
          bicycle:
            color: red
            price: 19.95
        expensive: 50
        """
      When the yamlpath `$..*` is evaluated
      Then the evaluation result is:
        """
        book:
          - category: fiction
            author: Author A
            title: Book A
          - category: non-fiction
            author: Author B
            title: Book B
        bicycle:
          color: red
          price: 19.95
        ---
        50
        ---
        - category: fiction
          author: Author A
          title: Book A
        - category: non-fiction
          author: Author B
          title: Book B
        ---
        color: red
        price: 19.95
        ---
        fiction
        ---
        Author A
        ---
        Book A
        ---
        non-fiction
        ---
        Author B
        ---
        Book B
        ---
        red
        ---
        19.95
        """
