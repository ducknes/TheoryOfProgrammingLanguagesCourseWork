package listener

import "github.com/antlr4-go/antlr/v4"

const (
	courseWorkGrammarParserEOF           = antlr.TokenEOF
	courseWorkGrammarParserLITERAL       = 1
	courseWorkGrammarParserINTEGER       = 2
	courseWorkGrammarParserWS            = 3
	courseWorkGrammarParserADD           = 4
	courseWorkGrammarParserSUB           = 5
	courseWorkGrammarParserMUL           = 6
	courseWorkGrammarParserASSIGN        = 7
	courseWorkGrammarParserLEFT_BRACKET  = 8
	courseWorkGrammarParserRIGHT_BRACKET = 9
	courseWorkGrammarParserREAD          = 10
	courseWorkGrammarParserWRITE         = 11
	courseWorkGrammarParserFOR           = 12
	courseWorkGrammarParserTO            = 13
	courseWorkGrammarParserSEMICOLON     = 14
	courseWorkGrammarParserBEGIN         = 15
	courseWorkGrammarParserEND           = 16
	courseWorkGrammarParserINT           = 17
	courseWorkGrammarParserVAR           = 18
	courseWorkGrammarParserCOLON         = 19
	courseWorkGrammarParserCOMMA         = 20
	courseWorkGrammarParserDO            = 21
	courseWorkGrammarParserEND_FOR       = 22

	courseWorkGrammarParserRULE_program            = 0
	courseWorkGrammarParserRULE_progDeclare        = 1
	courseWorkGrammarParserRULE_varDeclare         = 2
	courseWorkGrammarParserRULE_varList            = 3
	courseWorkGrammarParserRULE_programDescription = 4
	courseWorkGrammarParserRULE_assign             = 5
	courseWorkGrammarParserRULE_expression         = 6
	courseWorkGrammarParserRULE_subexpression      = 7
	courseWorkGrammarParserRULE_unaryOperator      = 8
	courseWorkGrammarParserRULE_binaryOperator     = 9
	courseWorkGrammarParserRULE_operand            = 10
	courseWorkGrammarParserRULE_id                 = 11
	courseWorkGrammarParserRULE_const              = 12
	courseWorkGrammarParserRULE_loop               = 13
	courseWorkGrammarParserRULE_read               = 14
	courseWorkGrammarParserRULE_write              = 15
)
