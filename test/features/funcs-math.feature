Feature: Funcs - Math

  YAMLPath provides support for various "math" functionalities that
  compute mathematical expressions and return the result as either integers or
  decimal values.

  Rule: abs() computes the absolute value of a number

    Returns the absolute value of the input.
    If the input collection is empty, the result is empty.
    If the input collection contains multiple items, the evaluation of the
    expression will end and signal an error to the calling environment.
    If the input collection contains a single non-numeric value, the evaluation
    of the expression will end and signal an error to the calling environment.

    Scenario: Absolute value of a positive number

      Given the yaml input:
        """
        num: 5
        """
      When the yamlpath `$.num.abs()` is evaluated
      Then the evaluation result is:
        """
        5
        """

    Scenario: Absolute value of a negative number

      Given the yaml input:
        """
        num: -3
        """
      When the yamlpath `$.num.abs()` is evaluated
      Then the evaluation result is:
        """
        3
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.abs()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.abs()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.abs()` is evaluated
      Then an error is raised

  Rule: ceil() computes the ceiling of a number

    Returns the first integer greater than or equal to the input.
    If the input collection is empty, the result is empty.
    If the input collection contains multiple items, the evaluation of the
    expression will end and signal an error to the calling environment.
    If the input collection contains a single non-numeric value, the evaluation
    of the expression will end and signal an error to the calling environment.

    Scenario: Ceiling of a positive decimal number

      Given the yaml input:
        """
        num: 5.3
        """
      When the yamlpath `$.num.ceil()` is evaluated
      Then the evaluation result is:
        """
        6
        """

    Scenario: Ceiling of a negative decimal number

      Given the yaml input:
        """
        num: -3.7
        """
      When the yamlpath `$.num.ceil()` is evaluated
      Then the evaluation result is:
        """
        -3
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.ceil()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1.1, 2.2, 3.3]
        """
      When the yamlpath `$.num.ceil()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.ceil()` is evaluated
      Then an error is raised

  Rule: floor() computes the floor of a number

    Returns the first integer less than or equal to the input.
    If the input collection is empty, the result is empty.
    If the input collection contains multiple items, the evaluation of the
    expression will end and signal an error to the calling environment.
    If the input collection contains a single non-numeric value, the evaluation
    of the expression will end and signal an error to the calling environment.

    Scenario: Floor of a positive decimal number

      Given the yaml input:
        """
        num: 5.7
        """
      When the yamlpath `$.num.floor()` is evaluated
      Then the evaluation result is:
        """
        5
        """

    Scenario: Floor of a negative decimal number

      Given the yaml input:
        """
        num: -3.3
        """
      When the yamlpath `$.num.floor()` is evaluated
      Then the evaluation result is:
        """
        -4
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.floor()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1.1, 2.2, 3.3]
        """
      When the yamlpath `$.num.floor()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.floor()` is evaluated
      Then an error is raised

  Rule: exp() computes the exponential of a number

    Returns e raised to the power of the input.
    If the input collection contains an Integer, it will be implicitly converted
    to a Decimal and the result will be a Decimal.
    If the input collection is empty, the result is empty.
    If the input collection contains multiple items, the evaluation of the
    expression will end and signal an error to the calling environment.
    If the input collection contains a single non-numeric value, the evaluation
    of the expression will end and signal an error to the calling environment.

    Scenario: Exponential of a positive number

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.exp()` is evaluated
      Then the evaluation result is:
        """
        7.38905
        """

    Scenario: Exponential of a negative number

      Given the yaml input:
        """
        num: -2
        """
      When the yamlpath `$.num.exp()` is evaluated
      Then the evaluation result is:
        """
        0.13533
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.exp()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.exp()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.exp()` is evaluated
      Then an error is raised

  Rule: ln() computes the natural logarithm of a number

    Returns the natural logarithm of the input (i.e. the logarithm base e).
    If the input collection is empty, the result is empty.
    If the input collection contains multiple items, the evaluation of the
    expression will end and signal an error to the calling environment.
    If the input collection contains a single non-numeric value, the evaluation
    of the expression will end and signal an error to the calling environment.

    Scenario: Natural logarithm of a positive number

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.ln()` is evaluated
      Then the evaluation result is:
        """
        0.69315
        """

    Scenario: Natural logarithm of a negative number

      Given the yaml input:
        """
        num: -2
        """
      When the yamlpath `$.num.ln()` is evaluated
      Then an error is raised
        """
        0.69315
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.ln()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.ln()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.ln()` is evaluated
      Then an error is raised

  Rule: pow() computes the power of a number

    Raises a number to the exponent power. If this function is used with
    Integers, the result is an Integer. If the function is used with Decimals,
    the result is a Decimal. If the function is used with a mixture of Integer
    and Decimal, the Integer is implicitly converted to a Decimal and the result
    is a Decimal.

    The evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If the input collection is not a singleton numeric value
    - If the exponent is not a singleton numeric value

    Scenario: Power of a positive number

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.pow(3)` is evaluated
      Then the evaluation result is:
        """
        8
        """

    Scenario: Power of a negative number

      Given the yaml input:
        """
        num: -2
        """
      When the yamlpath `$.num.pow(3)` is evaluated
      Then the evaluation result is:
        """
        -8
        """

    Scenario: Power of a decimal number
      Given the yaml input:
        """
        num: 2.5
        """
      When the yamlpath `$.num.pow(3)` is evaluated
      Then the evaluation result is:
        """
        15.625
        """

    Scenario: Exponent is decimal

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.pow(3.5)` is evaluated
      Then the evaluation result is:
        """
        11.313708504
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.pow(3)` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.pow(3)` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.pow(3)` is evaluated
      Then an error is raised

    Scenario: Exponent is a non-numeric value

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.pow("foo")` is evaluated
      Then an error is raised

  Rule: round() computes the rounded value of a number

    Rounds the decimal to the nearest whole number using a traditional round
    (i.e. 0.5 or higher will round to 1). If specified, the precision argument
    determines the decimal place at which the rounding will occur.
    If not specified, the rounding will default to 0 decimal places.

    If the input collection is empty, the result is empty.

    Evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If the input collection is not a singleton numeric value
    - If the precision argument is not a singleton numeric value

    Scenario: Round a positive decimal number

      Given the yaml input:
        """
        num: 2.5
        """
      When the yamlpath `$.num.round()` is evaluated
      Then the evaluation result is:
        """
        3
        """

    Scenario: Round a negative decimal number

      Given the yaml input:
        """
        num: -2.5
        """
      When the yamlpath `$.num.round()` is evaluated
      Then the evaluation result is:
        """
        -3
        """

    Scenario: Round a positive decimal number with precision

      Given the yaml input:
        """
        num: 2.555
        """
      When the yamlpath `$.num.round(2)` is evaluated
      Then the evaluation result is:
        """
        2.56
        """

    Scenario: Round a negative decimal number with precision

      Given the yaml input:
        """
        num: -2.555
        """
      When the yamlpath `$.num.round(2)` is evaluated
      Then the evaluation result is:
        """
        -2.56
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.round()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.round()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.round()` is evaluated
      Then an error is raised

    Scenario: Precision is a non-numeric value

      Given the yaml input:
        """
        num: 2
        """
      When the yamlpath `$.num.round("foo")` is evaluated
      Then an error is raised

  Rule: truncate() computes the truncated value of a number

    Returns the integer portion of the input.

    If the input collection is empty, the result is empty.

    Evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If the input collection is not a singleton numeric value

    Scenario: Truncate a positive decimal number

      Given the yaml input:
        """
        num: 2.5
        """
      When the yamlpath `$.num.truncate()` is evaluated
      Then the evaluation result is:
        """
        2
        """

    Scenario: Truncate a negative decimal number

      Given the yaml input:
        """
        num: -2.5
        """
      When the yamlpath `$.num.truncate()` is evaluated
      Then the evaluation result is:
        """
        -2
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: 0
        """
      When the yamlpath `$.bar.truncate()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains multiple items

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num.truncate()` is evaluated
      Then an error is raised

    Scenario: Collection contains a single non-numeric value

      Given the yaml input:
        """
        num: "foo"
        """
      When the yamlpath `$.num.truncate()` is evaluated
      Then an error is raised

  Rule: min() computes the minimum value of a collection

    Returns the minimum value of all the values within the input collection and
    its parameters.

    If the input collection and parameters are empty, the result is empty.

    Evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If any value in the input collection is not a number
    - If any of the parameters are not numbers

    Scenario: Minimum of values in the input collection

      Given the yaml input:
        """
        num: [1, 3, 5]
        """
      When the yamlpath `$.num[*].min()` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Minimum of values in the parameters

      Given the yaml input:
        """
        4
        """
      When the yamlpath `min(3, 5)` is evaluated
      Then the evaluation result is:
        """
        3
        """

    Scenario: Minimum of values in the input collection and parameters

      Given the yaml input:
        """
        num: [1, 3, 5]
        """
      When the yamlpath `$.num[*].min(2, 4)` is evaluated
      Then the evaluation result is:
        """
        1
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: []
        """
      When the yamlpath `$.num[*].min()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains non-numeric value

      Given the yaml input:
        """
        num: [1, "foo", 3]
        """
      When the yamlpath `$.num[*].min()` is evaluated
      Then an error is raised

    Scenario: Parameters contain non-numeric value

      Given the yaml input:
        """
        4
        """
      When the yamlpath `min(3, "foo")` is evaluated
      Then an error is raised

  Rule: max() computes the maximum value of a collection

    Returns the maximum value of all the values within the input collection and
    its parameters.

    If the input collection and parameters are empty, the result is empty.

    Evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If any value in the input collection is not a number
    - If any of the parameters are not numbers

    Scenario: Maximum of values in the input collection

      Given the yaml input:
        """
        num: [1, 3, 5]
        """
      When the yamlpath `$.num[*].max()` is evaluated
      Then the evaluation result is:
        """
        5
        """

    Scenario: Maximum of values in the parameters

      Given the yaml input:
        """
        4
        """
      When the yamlpath `max(3, 5)` is evaluated
      Then the evaluation result is:
        """
        5
        """

    Scenario: Maximum of values in the input collection and parameters

      Given the yaml input:
        """
        num: [1, 3, 5]
        """
      When the yamlpath `$.num[*].max(2, 4)` is evaluated
      Then the evaluation result is:
        """
        5
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: []
        """
      When the yamlpath `$.num[*].max()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains non-numeric value

      Given the yaml input:
        """
        num: [1, "foo", 3]
        """
      When the yamlpath `$.num[*].max()` is evaluated
      Then an error is raised

    Scenario: Parameters contain non-numeric value

      Given the yaml input:
        """
        4
        """
      When the yamlpath `max(3, "foo")` is evaluated
      Then an error is raised

  Rule: sum() computes the sum of values of a collection

    Returns the sum of all the values within the input collection and its
    parameters.

    If the input collection and parameters are empty, the result is empty.

    Evaluation of the expression will end and signal an error to the calling
    environment in the following conditions:

    - If any value in the input collection is not a number
    - If any of the parameters are not numbers

    Scenario: Sum of values in the input collection

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num[*].sum()` is evaluated
      Then the evaluation result is:
        """
        6
        """

    Scenario: Sum of values in the parameters

      Given the yaml input:
        """
        0
        """
      When the yamlpath `sum(1, 2)` is evaluated
      Then the evaluation result is:
        """
        3
        """

    Scenario: Sum of values in the input collection and parameters

      Given the yaml input:
        """
        num: [1, 2, 3]
        """
      When the yamlpath `$.num[*].sum(4, 5)` is evaluated
      Then the evaluation result is:
        """
        15
        """

    Scenario: Collection is empty

      Given the yaml input:
        """
        num: []
        """
      When the yamlpath `$.num[*].sum()` is evaluated
      Then the evaluation result is empty

    Scenario: Collection contains non-numeric value

      Given the yaml input:
        """
        num: [1, "foo", 3]
        """
      When the yamlpath `$.num[*].sum()` is evaluated
      Then an error is raised

    Scenario: Parameters contain non-numeric value

      Given the yaml input:
        """
        0
        """
      When the yamlpath `sum(1, "foo")` is evaluated
      Then an error is raised
