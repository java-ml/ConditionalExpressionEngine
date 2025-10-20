grammar ConditionExpr;

// 语法解析器规则
expression
   : logicalExpression
   ;

logicalExpression
   : logicalExpression AND logicalExpression  # AndExpression
   | logicalExpression OR logicalExpression   # OrExpression
   | NOT logicalExpression                    # NotExpression
   | comparisonExpression                     # ComparisonExpr
   | '(' logicalExpression ')'                # ParenLogical
   ;

comparisonExpression
   : additiveExpression comparisonOperator additiveExpression # Comparison
   | additiveExpression                                       # SingleValue
   ;

comparisonOperator
   : GT | GE | LT | LE | EQ | NE
   ;

additiveExpression
   : multiplicativeExpression ( additiveOperator multiplicativeExpression )*
   ;

additiveOperator
   : ADD | SUB
   ;

multiplicativeExpression
   : unaryExpression ( multiplicativeOperator unaryExpression )*
   ;

multiplicativeOperator
   : MUL | DIV
   ;

unaryExpression
   : ( ADD | SUB ) unaryExpression  # UnarySigned
   | primaryExpression              # UnaryPrimary
   ;

primaryExpression
   : functionCall                   # FunctionPrimary
   | IDENTIFIER                     # IdentifierPrimary
   | literal                        # LiteralPrimary
   | '(' expression ')'             # ParenExpression
   ;

functionCall
   : IDENTIFIER '(' ( expression ( ',' expression )* )? ')'
   ;

literal
   : NUMBER
   | STRING
   | BOOL
   ;

// 词法规则 - 重新排序确保正确优先级
BOOL       : 'true' | 'false' ;
AND        : '&&' | 'AND' | 'and' ;
OR         : '||' | 'OR' | 'or' ;
NOT        : '!' | 'NOT' | 'not' ;

GT         : '>' ;
GE         : '>=' ;
LT         : '<' ;
LE         : '<=' ;
EQ         : '==' | '=' ;
NE         : '!=' | '<>' ;

ADD        : '+' ;
SUB        : '-' ;
MUL        : '*' ;
DIV        : '/' ;

LPAREN     : '(' ;
RPAREN     : ')' ;
COMMA      : ',' ;

// 简化的字符串规则 - 先确保基本功能
STRING     : '"' ( ~["\\] | '\\' ["\\] )* '"'
           | '\'' ( ~['\\] | '\\' ['\\] )* '\'' ;

IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
NUMBER     : [0-9]+ ( '.' [0-9]* )? ( [eE] [+\-]? [0-9]+ )? ;

WS         : [ \t\r\n]+ -> skip ;
COMMENT    : '//' ~[\r\n]* -> skip ;