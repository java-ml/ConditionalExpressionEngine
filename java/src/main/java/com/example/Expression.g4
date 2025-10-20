grammar Expression;

// 语法规则
expression
    : logicalOrExpression
    ;

logicalOrExpression
    : logicalAndExpression ('||' logicalAndExpression)*
    ;

logicalAndExpression
    : equalityExpression ('&&' equalityExpression)*
    ;

equalityExpression
    : relationalExpression (('==' | '!=') relationalExpression)*
    ;

relationalExpression
    : additiveExpression (('>' | '>=' | '<' | '<=') additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression (('+' | '-') multiplicativeExpression)*
    ;

multiplicativeExpression
    : unaryExpression (('*' | '/' | '%') unaryExpression)*
    ;

unaryExpression
    : '!' unaryExpression
    | '-' unaryExpression
    | primaryExpression
    ;

primaryExpression
    : '(' expression ')'
    | functionCall
    | variable
    | number
    | BOOLEAN
    | STRING
    ;

functionCall
    : IDENTIFIER '(' argumentList? ')'
    ;

argumentList
    : expression (',' expression)*
    ;

variable
    : IDENTIFIER
    ;

number
    : INTEGER
    | FLOAT
    ;

// 词法规则
BOOLEAN: 'true' | 'false';

// 支持转义字符的字符串
STRING: '"' (ESC | ~["\\])* '"';
fragment ESC: '\\' (["\\/bfnrt] | UNICODE);
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];

// 标识符支持转义反引号
IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]* | '`' (ESC_IDENT | ~[`\\])+ '`';
fragment ESC_IDENT: '\\' ([`\\] | UNICODE);

INTEGER: [0-9]+;
FLOAT: [0-9]+ '.' [0-9]* | '.' [0-9]+;

WS: [ \t\r\n]+ -> skip;

// 操作符
PLUS: '+';
MINUS: '-';
MULTIPLY: '*';
DIVIDE: '/';
MODULO: '%';
NOT: '!';
AND: '&&';
OR: '||';
EQ: '==';
NEQ: '!=';
GT: '>';
GE: '>=';
LT: '<';
LE: '<=';
LPAREN: '(';
RPAREN: ')';
COMMA: ',';
BACKTICK: '`';