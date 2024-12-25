grammar yamlpath;

// Parser Rules
yamlPath          : path EOF ;
path              : root selector* ;
root              : '$' | '@' ;
selector          : dotSelector | recursiveSelector | bracketSelector ;
recursiveSelector : '..' (NAME | WILDCARD)? ;
dotSelector       : '.' (NAME | WILDCARD) ;
bracketSelector   : '[' bracketExpression ']' ;
bracketExpression
                  : quotedName
                  | NUMBER
                  | WILDCARD
                  | slice
                  | filter
                  | unionString
                  | unionIndices
                  ;
slice             : (NUMBER)? ':' (NUMBER)? (':' NUMBER)? ;
filter            : '?' '(' expression ')' ;
unionString       : (quotedName) (',' (quotedName))* ;
unionIndices      : (NUMBER) (',' (NUMBER))* ;
expression        : compareExpr
                  | booleanExpr
                  | arithmeticExpr
                  | containmentExpr
                  | negationExpr
                  | subexpression ;
compareExpr       : subexpression ('==' | '!=' | '<' | '>' | '<=' | '>=') subexpression ;
booleanExpr       : subexpression ('&&' | '||') subexpression ;
arithmeticExpr    : subexpression ('+' | '-' | '/' | '*') subexpression ;
containmentExpr   : subexpression ('in' | 'nin' | 'subsetof') subexpression ;
negationExpr      : '!' subexpression ;
subexpression     : value | path ;
value             : STRING | NUMBER | BOOLEAN | NULL ;

// Lexer Rules
NAME           : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER         : '-'? [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)? ;
STRING         : '"' (~["\\] | '\\' .)* '"' ;
BOOLEAN        : 'true' | 'false' ;
NULL           : 'null' ;
quotedName     : STRING ;
WILDCARD       : '*' ;
WS             : [ \t\r\n]+ -> skip ;
COMMENT        : '//' ~[\r\n]* -> skip ;
