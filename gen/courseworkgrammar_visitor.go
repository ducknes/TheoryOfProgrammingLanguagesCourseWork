// Code generated from /Users/ilyaantonov/Downloads/TheoryOfProgrammingLanguagesCourseWork-test/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // courseWorkGrammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by courseWorkGrammarParser.
type courseWorkGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by courseWorkGrammarParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#progDeclare.
	VisitProgDeclare(ctx *ProgDeclareContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#varDeclare.
	VisitVarDeclare(ctx *VarDeclareContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#varList.
	VisitVarList(ctx *VarListContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#programDescription.
	VisitProgramDescription(ctx *ProgramDescriptionContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#subexpression.
	VisitSubexpression(ctx *SubexpressionContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#const.
	VisitConst(ctx *ConstContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#read.
	VisitRead(ctx *ReadContext) interface{}

	// Visit a parse tree produced by courseWorkGrammarParser#write.
	VisitWrite(ctx *WriteContext) interface{}
}
