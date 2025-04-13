grammar yamlpath;

/*****************************************************************************
  Parser rules
******************************************************************************/

path: expression EOF ;

expression
      : term                                              # termExpression
      | expression '..' (invocation)?                     # recursiveExpression
      | expression '.' invocation                         # fieldExpression
      | expression '[' indexParam ']'                     # indexExpression
      | '(' expression ')'                                # parenthesisExpression
      | ('!' | 'not') expression                          # negationExpression
      | ('+' | '-') expression                            # polarityExpression
      | expression ('*' | '/' | '%') expression           # multiplicativeExpression
      | expression ('+' | '-') expression                 # additiveExpression
      | expression ('|' expression)+                      # unionExpression
      | expression ('<=' | '<' | '>' | '>=') expression   # inequalityExpression
      | expression ('==' | '!=') expression               # equalityExpression
      | expression '=~' regex                             # matchExpression
      | expression ('in' | 'nin' | 'subsetof') expression # membershipExpression
      | expression ('&&' | 'and') expression              # andExpression
      | expression ('||' | 'or') expression               # orExpression
      ;

term
      : ('$' | '@')                                       # rootTerm
      | externalConstant                                  # externalConstantTerm
      | literal                                           # literalTerm
      | invocation                                        # invocationTerm
      ;

externalConstant
        : '%' identifier
        ;

indexParam
      : '*'                                               # wildcardIndex
      | (NUMBER)? ':' (NUMBER)? (':' NUMBER)?             # sliceIndex
      | expression                                        # expressionIndex
      ;

literal
      : STRING                                            # stringLiteral
      | NUMBER                                            # numberLiteral
      | ('true' | 'false')                                # booleanLiteral
      | 'null'                                            # nullLiteral
      | '[' (listEntries)? ']'                            # listLiteral
      | '{' (mapEntries)? '}'                             # mapLiteral
      ;

listEntries
      : (literal) (',' literal)*
      ;

mapEntries
      : STRING ':' literal (',' STRING ':' literal)*
      ;

invocation
      : identifier                                        # memberInvocation
      | '*'                                               # wildcardInvocation
      | identifier '(' paramList? ')'                     # functionInvocation
      ;

paramList
      : expression (',' expression)*
      ;

identifier
      : IDENTIFIER                                        # plainIdentifier
      | STRING                                            # quotedIdentifier
      ;

regex
      : REGEX ('i' | 'm' | 's')*?
      ;

/*****************************************************************************
  Lexer rules
******************************************************************************/

IDENTIFIER     : [a-zA-Z_][a-zA-Z0-9_-]*([a-zA-Z0-9_])?;
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
