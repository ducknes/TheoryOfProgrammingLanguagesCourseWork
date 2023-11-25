// Code generated from /Users/ilyaantonov/Downloads/TheoryOfProgrammingLanguagesCourseWork-test/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // courseWorkGrammar

import "github.com/antlr4-go/antlr/v4"

// courseWorkGrammarListener is a complete listener for a parse tree produced by courseWorkGrammarParser.
type courseWorkGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgDeclare is called when entering the progDeclare production.
	EnterProgDeclare(c *ProgDeclareContext)

	// EnterVarDeclare is called when entering the varDeclare production.
	EnterVarDeclare(c *VarDeclareContext)

	// EnterVarList is called when entering the varList production.
	EnterVarList(c *VarListContext)

	// EnterProgramDescription is called when entering the programDescription production.
	EnterProgramDescription(c *ProgramDescriptionContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubexpression is called when entering the subexpression production.
	EnterSubexpression(c *SubexpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterConst is called when entering the const production.
	EnterConst(c *ConstContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterWrite is called when entering the write production.
	EnterWrite(c *WriteContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgDeclare is called when exiting the progDeclare production.
	ExitProgDeclare(c *ProgDeclareContext)

	// ExitVarDeclare is called when exiting the varDeclare production.
	ExitVarDeclare(c *VarDeclareContext)

	// ExitVarList is called when exiting the varList production.
	ExitVarList(c *VarListContext)

	// ExitProgramDescription is called when exiting the programDescription production.
	ExitProgramDescription(c *ProgramDescriptionContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubexpression is called when exiting the subexpression production.
	ExitSubexpression(c *SubexpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitConst is called when exiting the const production.
	ExitConst(c *ConstContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitWrite is called when exiting the write production.
	ExitWrite(c *WriteContext)
}
