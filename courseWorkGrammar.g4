grammar courseWorkGrammar;

//Вся программа 0
program: varDeclare progDeclare;

//Описание программы 1
progDeclare: BEGIN programDescription END;

//Объявление переменных 2
varDeclare: VAR varList COLON INT SEMICOLON ;

//Перечисление переменных 3
varList: id (COMMA id)*;

//Текст программы 4
programDescription: (assign+ | loop | read | write)+;

//Присвоение 5
assign: id ASSIGN expression SEMICOLON;

//Выражение (подвыражение или отрицательное подвыражение) 6
expression: unaryOperator subexpression | subexpression;

//Подвыражение (выражение в скобках, операнд, математическое выражение с подвыражениями) 7
subexpression: LEFT_BRACKET expression RIGHT_BRACKET | operand | subexpression binaryOperator subexpression;

//Унарные оператор (делает отрицательное число/переменную) 8
unaryOperator: SUB;

//Бинарный оператор (знак математического выражения) 9
binaryOperator: ADD | SUB | MUL;

//Операнд (переменная или число/цифра) 10
operand: id | const;

//Переменная 11
id: LITERAL*;

//Число/цифра 12
const: INTEGER*;

//Цикл for 13
loop: FOR assign TO expression DO programDescription END_FOR;

//Чтение данных 14
read: READ LEFT_BRACKET varList RIGHT_BRACKET ;

//Вывод данных 15
write: WRITE LEFT_BRACKET varList RIGHT_BRACKET;

//Регулярочка для переменных
LITERAL: [a-zA-Z];

//Регулярочка для константы
INTEGER: [0-9];

//Пропускаемые символы (мы их не учитываем в проге)
WS : [ \t\r\n]+ -> skip;

ADD: '+';

SUB: '-';

MUL: '*';

ASSIGN: '=';

LEFT_BRACKET: '(';

RIGHT_BRACKET: ')';

READ: 'READ';

WRITE: 'WRITE';

FOR: 'FOR';

TO: 'TO';

SEMICOLON : ';';

BEGIN: 'BEGIN';

END: 'END';

INT: 'INTEGER';

VAR: 'VAR';

COLON: ':';

COMMA: ',';

DO: 'DO';

END_FOR: 'END_FOR';