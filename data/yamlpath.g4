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
      : subexpression ('+' | '-') subexpression                 # additiveSubexpression
      | subexpression ('*' | '/') subexpression                 # multiplicativeSubexpression
      | subexpression ('<=' | '<' | '>' | '>=') subexpression   # inequalitySubexpression
      | subexpression ('==' | '!=') subexpression               # equalitySubexpression
      | subexpression ('in' | 'nin' | 'subsetof') subexpression # membershipSubexpression
      | subexpression ('&&' | 'and') subexpression              # andSubexpression
      | subexpression ('||' | 'or') subexpression               # orSubexpression
      | ('!' | 'not') subexpression                             # negationSubexpression
      | literal                                                 # literalSubexpression
      | '(' subexpression ')'                                   # parenthesisSubexpression
      | expression                                              # rootSubexpression
      ;

literal
      : STRING                                            # stringLiteral
      | NUMBER                                            # numberLiteral
      | ('true' | 'false')                                # booleanLiteral
      | 'null'                                            # nullLiteral
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

/*****************************************************************************
  Lexer rules
******************************************************************************/

IDENTIFIER     : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER         : '-'? [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)? ;
STRING         : '"' (ESC | .)*? '"' ;

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
