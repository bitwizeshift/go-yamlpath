grammar yamlpath;

// Parser Rules
yamlPath       : root selector* ;
root           : '$' ;
selector       : dotSelector | bracketSelector ;
dotSelector    : '.' (NAME | '*' | '..') ;
bracketSelector: '[' (bracketExpression | wildcard) ']' ;
bracketExpression
               : quotedName
               | NUMBER
               | slice
               | filter
               | union
               ;
slice          : (NUMBER)? ':' (NUMBER)? (':' NUMBER)? ;
filter         : '?' '(' expression ')' ;
union          : (quotedName | NUMBER) (',' (quotedName | NUMBER))* ;
expression     : subexpression (('==' | '!=' | '<' | '>' | '<=' | '>=') subexpression)? ;
subexpression  : value | NAME | yamlPath ;
value          : STRING | NUMBER | BOOLEAN | NULL ;

// Lexer Rules
NAME           : [a-zA-Z_][a-zA-Z0-9_]* ;
NUMBER         : '-'? [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)? ;
STRING         : '"' (~["\\] | '\\' .)* '"' ;
BOOLEAN        : 'true' | 'false' ;
NULL           : 'null' ;
quotedName     : STRING ;
wildcard       : '*' ;
WS             : [ \t\r\n]+ -> skip ;
COMMENT        : '//' ~[\r\n]* -> skip ;
