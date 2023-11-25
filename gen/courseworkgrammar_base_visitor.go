// Code generated from /Users/ilyaantonov/Downloads/TheoryOfProgrammingLanguagesCourseWork-test/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // courseWorkGrammar

import "github.com/antlr4-go/antlr/v4"

type BasecourseWorkGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecourseWorkGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitProgDeclare(ctx *ProgDeclareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitVarDeclare(ctx *VarDeclareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitVarList(ctx *VarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitProgramDescription(ctx *ProgramDescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitSubexpression(ctx *SubexpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitConst(ctx *ConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitRead(ctx *ReadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecourseWorkGrammarVisitor) VisitWrite(ctx *WriteContext) interface{} {
	return v.VisitChildren(ctx)
}
