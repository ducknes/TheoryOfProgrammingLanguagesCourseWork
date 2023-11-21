// Code generated from /Users/ilyaantonov/Downloads/3 курс/ТЯП/cwp/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // courseWorkGrammar

import "github.com/antlr4-go/antlr/v4"

// BasecourseWorkGrammarListener is a complete listener for a parse tree produced by courseWorkGrammarParser.
type BasecourseWorkGrammarListener struct{}

var _ courseWorkGrammarListener = &BasecourseWorkGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecourseWorkGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecourseWorkGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecourseWorkGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecourseWorkGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasecourseWorkGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasecourseWorkGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgDeclare is called when production progDeclare is entered.
func (s *BasecourseWorkGrammarListener) EnterProgDeclare(ctx *ProgDeclareContext) {}

// ExitProgDeclare is called when production progDeclare is exited.
func (s *BasecourseWorkGrammarListener) ExitProgDeclare(ctx *ProgDeclareContext) {}

// EnterVarDeclare is called when production varDeclare is entered.
func (s *BasecourseWorkGrammarListener) EnterVarDeclare(ctx *VarDeclareContext) {}

// ExitVarDeclare is called when production varDeclare is exited.
func (s *BasecourseWorkGrammarListener) ExitVarDeclare(ctx *VarDeclareContext) {}

// EnterVarList is called when production varList is entered.
func (s *BasecourseWorkGrammarListener) EnterVarList(ctx *VarListContext) {}

// ExitVarList is called when production varList is exited.
func (s *BasecourseWorkGrammarListener) ExitVarList(ctx *VarListContext) {}

// EnterProgramDescription is called when production programDescription is entered.
func (s *BasecourseWorkGrammarListener) EnterProgramDescription(ctx *ProgramDescriptionContext) {}

// ExitProgramDescription is called when production programDescription is exited.
func (s *BasecourseWorkGrammarListener) ExitProgramDescription(ctx *ProgramDescriptionContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasecourseWorkGrammarListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasecourseWorkGrammarListener) ExitAssign(ctx *AssignContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasecourseWorkGrammarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasecourseWorkGrammarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubexpression is called when production subexpression is entered.
func (s *BasecourseWorkGrammarListener) EnterSubexpression(ctx *SubexpressionContext) {}

// ExitSubexpression is called when production subexpression is exited.
func (s *BasecourseWorkGrammarListener) ExitSubexpression(ctx *SubexpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BasecourseWorkGrammarListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BasecourseWorkGrammarListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BasecourseWorkGrammarListener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BasecourseWorkGrammarListener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterOperand is called when production operand is entered.
func (s *BasecourseWorkGrammarListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BasecourseWorkGrammarListener) ExitOperand(ctx *OperandContext) {}

// EnterId is called when production id is entered.
func (s *BasecourseWorkGrammarListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BasecourseWorkGrammarListener) ExitId(ctx *IdContext) {}

// EnterConst is called when production const is entered.
func (s *BasecourseWorkGrammarListener) EnterConst(ctx *ConstContext) {}

// ExitConst is called when production const is exited.
func (s *BasecourseWorkGrammarListener) ExitConst(ctx *ConstContext) {}

// EnterLoop is called when production loop is entered.
func (s *BasecourseWorkGrammarListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BasecourseWorkGrammarListener) ExitLoop(ctx *LoopContext) {}

// EnterRead is called when production read is entered.
func (s *BasecourseWorkGrammarListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BasecourseWorkGrammarListener) ExitRead(ctx *ReadContext) {}

// EnterWrite is called when production write is entered.
func (s *BasecourseWorkGrammarListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *BasecourseWorkGrammarListener) ExitWrite(ctx *WriteContext) {}
