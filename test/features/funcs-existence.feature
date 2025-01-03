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

  Rule: count() evaluates the number of elements in a collection

    The `count()` function shall return the number of elements in the collection.
    Empty collections shall return 0.

    Scenario: Counting elements in a collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].count()` is evaluated
      Then the evaluation result is:
        """
        2
        """

    Scenario: Counting elements in an empty collection

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].count()` is evaluated
      Then the evaluation result is:
        """
        0
        """

  Rule: distinct() evaluates to unique elements in a collection

    The `distinct()` function shall return a collection with only the unique
    elements in the input collection. The order of the elements in the output
    collection is not guaranteed.

    Scenario: Distinct elements in a collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
          - name: "John"
        """
      When the yamlpath `$.people[*].distinct()` is evaluated
      Then the evaluation result is:
        """
        name: "John"
        ---
        name: "Jane"
        """

  Rule: isDistinct() evaluates to true if connection contains unique elements

    The `isDistinct()` function shall return true if the input collection contains
    only unique elements, and false otherwise.

    Scenario: Unique elements in a collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].isDistinct()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Non-unique elements in a collection

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "John"
        """
      When the yamlpath `$.people[*].isDistinct()` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: all() evaluates to true if all elements in a collection satisfy a criteria

    The `all()` function shall return true if all elements in the input collection
    satisfy the criteria, and false otherwise.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].all(@.name =~ /^J/)` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: All elements satisfy criteria

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].all(@.name =~ /^J/)` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Not all elements satisfy criteria

      Given the yaml input:
        """
        people:
          - name: "Mark"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].all(@.name =~ /^M/)` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: any() evaluates to true if any element in a collection satisfies a criteria

    The `any()` function shall return true if any element in the input collection
    satisfies the criteria, and false otherwise.

    Scenario: Collection is empty

      Given the yaml input:
        """
        people: []
        """
      When the yamlpath `$.people[*].any(@.name =~ /^J/)` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: All elements satisfies criteria

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].any(@.name =~ /^J/)` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: One element satisfies criteria

      Given the yaml input:
        """
        people:
          - name: "John"
          - name: "Mark"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].any(@.name =~ /^M/)` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: No element satisfies criteria

      Given the yaml input:
        """
        people:
          - name: "Mark"
          - name: "Jane"
        """
      When the yamlpath `$.people[*].any(@.name =~ /^G/)` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: allTrue() evaluates to true if all elements in a collection are boolean true

    The `allTrue()` function shall return true if all elements in the input
    collection are booleans with the value of true, and false otherwise.

    Scenario: Collection is empty

      Given the yaml input:
        """
        conditions: []
        """
      When the yamlpath `$.conditions[*].allTrue()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: All elements are true

      Given the yaml input:
        """
        conditions: [true, true, true]
        """
      When the yamlpath `$.conditions[*].allTrue()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Not all elements are true

      Given the yaml input:
        """
        conditions: [true, true, false]
        """
      When the yamlpath `$.conditions[*].allTrue()` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Not all elements are boolean

      Given the yaml input:
        """
        conditions: [true, true, 1]
        """
      When the yamlpath `$.conditions[*].allTrue()` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: anyTrue() evaluates to true if any element in a collection is boolean true

    The `anyTrue()` function shall return true if any element in the input
    collection is a boolean with the value of true, and false otherwise.

    Scenario: All elements are true

      Given the yaml input:
        """
        conditions: [true, true, true]
        """
      When the yamlpath `$.conditions[*].anyTrue()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: One element is true

      Given the yaml input:
        """
        conditions: [true, false, false]
        """
      When the yamlpath `$.conditions[*].anyTrue()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: No element is true

      Given the yaml input:
        """
        conditions: [false, false, false]
        """
      When the yamlpath `$.conditions[*].anyTrue()` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Not all elements are boolean

      Given the yaml input:
        """
        conditions: [false, 1, true]
        """
      When the yamlpath `$.conditions[*].anyTrue()` is evaluated
      Then the evaluation result is:
        """
        true
        """

  Rule: allFalse() evaluates to true if all elements in a collection are boolean false

    The `allFalse()` function shall return true if all elements in the input
    collection are booleans with the value of false, and false otherwise.

    Scenario: Collection is empty

      Given the yaml input:
        """
        conditions: []
        """
      When the yamlpath `$.conditions[*].allFalse()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: All elements are false

      Given the yaml input:
        """
        conditions: [false, false, false]
        """
      When the yamlpath `$.conditions[*].allFalse()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: Not all elements are false

      Given the yaml input:
        """
        conditions: [false, false, true]
        """
      When the yamlpath `$.conditions[*].allFalse()` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Not all elements are boolean

      Given the yaml input:
        """
        conditions: [false, false, 0]
        """
      When the yamlpath `$.conditions[*].allFalse()` is evaluated
      Then the evaluation result is:
        """
        false
        """

  Rule: anyFalse() evaluates to true if any element in a collection is boolean false

    The `anyFalse()` function shall return true if any element in the input
    collection is a boolean with the value of false, and false otherwise.

    Scenario: All elements are false

      Given the yaml input:
        """
        conditions: [false, false, false]
        """
      When the yamlpath `$.conditions[*].anyFalse()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: One element is false

      Given the yaml input:
        """
        conditions: [false, true, true]
        """
      When the yamlpath `$.conditions[*].anyFalse()` is evaluated
      Then the evaluation result is:
        """
        true
        """

    Scenario: No element is false

      Given the yaml input:
        """
        conditions: [true, true, true]
        """
      When the yamlpath `$.conditions[*].anyFalse()` is evaluated
      Then the evaluation result is:
        """
        false
        """

    Scenario: Not all elements are boolean

      Given the yaml input:
        """
        conditions: [true, 1, false]
        """
      When the yamlpath `$.conditions[*].anyFalse()` is evaluated
      Then the evaluation result is:
        """
        true
        """
