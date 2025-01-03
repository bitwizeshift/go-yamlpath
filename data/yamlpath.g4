grammar yamlpath;

/*****************************************************************************
  Parser rules
******************************************************************************/

path: expression EOF ;

expression
      : ('$' | '@')                                       # rootExpression
      | expression '..' (invocation)?                     # recursiveExpression
      | expression '.' invocation                         # fieldExpression
      | expression '[' bracketParam ']'                   # indexExpression
      ;

bracketParam
      : STRING (',' (STRING))*                            # unionStringBracket
      | NUMBER (',' (NUMBER))*                            # unionNumberBracket
      | '*'                                               # wildcardBracket
      | (NUMBER)? ':' (NUMBER)? (':' NUMBER)?             # sliceBracket
      | '?' '(' subexpression ')'                         # filterBracket
      | '(' subexpression ')'                             # scriptBracket
      ;

subexpression
      : expression                                              # rootSubexpression
      | literal                                                 # literalSubexpression
      | aggregation                                             # aggregationSubexpression
      | '(' subexpression ')'                                   # parenthesisSubexpression
      | ('!' | 'not') subexpression                             # negationSubexpression
      | ('+' | '-') expression                                  # polaritySubexpression
      | subexpression ('*' | '/' | '%') subexpression           # multiplicativeSubexpression
      | subexpression ('+' | '-') subexpression                 # additiveSubexpression
      | subexpression ('<=' | '<' | '>' | '>=') subexpression   # inequalitySubexpression
      | subexpression ('==' | '!=') subexpression               # equalitySubexpression
      | subexpression '=~' regex                                # matchSubexpression
      | subexpression ('in' | 'nin' | 'subsetof') subexpression # membershipSubexpression
      | subexpression ('&&' | 'and') subexpression              # andSubexpression
      | subexpression ('||' | 'or') subexpression               # orSubexpression
      ;

literal
      : STRING                                            # stringLiteral
      | NUMBER                                            # numberLiteral
      | ('true' | 'false')                                # booleanLiteral
      | 'null'                                            # nullLiteral
      ;

aggregation
      : '[' (listEntries)? ']'                            # listAggregation
      | '{' (mapEntries)? '}'                             # mapAggregation
      | literal                                           # literalAggregation
      ;

listEntries
      : (literal) (',' literal)*
      ;

mapEntries
      : STRING ':' aggregation (',' STRING ':' aggregation)*
      ;

invocation
      : identifier                                        # memberInvocation
      | '*'                                               # wildcardInvocation
      | identifier '(' paramList? ')'                     # functionInvocation
      ;

paramList
      : subexpression (',' subexpression)*
      ;

identifier
      : IDENTIFIER
      ;

regex
      : REGEX ('i' | 'm' | 's')*?
      ;

/*****************************************************************************
  Lexer rules
******************************************************************************/

IDENTIFIER     : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER         : '-'? [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)? ;
STRING         : '"' (ESC | .)*? '"' ;
REGEX          : '/' (ESC | .)*? '/' ;

// Pipe whitespace to the HIDDEN channel to support retrieving source text through the parser.
WS             : [ \t\r\n]+ -> channel(HIDDEN) ;
COMMENT        : '//' ~[\r\n]* -> channel(HIDDEN) ;

// Fragments

fragment ESC
      : '\\' ([`'\\/fnrt] | UNICODE)
      ;

fragment UNICODE
        : 'u' HEX HEX HEX HEX
        ;

fragment HEX
        : [0-9a-fA-F]
        ;
