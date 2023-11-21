package customVisitor

import (
	"TheoryOfProgrammingLanguagesCourseWork/errorhandler"
	genAntlr "TheoryOfProgrammingLanguagesCourseWork/gen"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

const (
	success = true
	fail    = false
)

type Visitor struct {
	genAntlr.BasecourseWorkGrammarVisitor
	errorHandler *errorhandler.CustomErrorHandler

	ids   map[string]int
	stack []int
}

func New(errorHandler *errorhandler.CustomErrorHandler) Visitor {
	return Visitor{
		errorHandler: errorHandler,
		ids:          make(map[string]int),
		stack:        make([]int, 0),
	}
}

func (v *Visitor) push(value int) {
	v.stack = append(v.stack, value)
}

func (v *Visitor) pop() int {
	if len(v.stack) == 0 {
		v.errorHandler.Error(errorhandler.CustomError{
			Type:    errorhandler.StackError,
			Message: "стек пуст\n",
		})
	}

	value := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	return value
}

func (v *Visitor) calc(binaryOperator string) {
	switch binaryOperator {
	case "+":
		v.push(v.pop() + v.pop())
	case "-":
		v.push(v.pop() - v.pop())
	case "*":
		v.push(v.pop() * v.pop())
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	vdCount := 0
	pdCount := 0

	for _, ctx := range tree.GetChildren() {
		switch ctx.(type) {
		case *genAntlr.VarDeclareContext:
			result := v.VisitVarDeclare(ctx.(*genAntlr.VarDeclareContext))
			if !result.(bool) {
				v.errorHandler.Error(errorhandler.CustomError{
					Line:     ctx.(*genAntlr.VarDeclareContext).GetStart().GetLine(),
					Position: ctx.(*genAntlr.VarDeclareContext).GetStart().GetColumn(),
					Type:     errorhandler.ProgramError,
					Message:  "ошибка при объявлении переменных\n",
				})
				return fail
			}
			vdCount += 1
		case *genAntlr.ProgDeclareContext:
			result := v.VisitProgDeclare(ctx.(*genAntlr.ProgDeclareContext))
			if !result.(bool) {
				v.errorHandler.Error(errorhandler.CustomError{
					Line:     ctx.(*genAntlr.ProgDeclareContext).GetStart().GetLine(),
					Position: ctx.(*genAntlr.ProgDeclareContext).GetStart().GetColumn(),
					Type:     errorhandler.ProgramError,
					Message:  "ошибка при объявлении программы\n",
				})
				return fail
			}
			pdCount += 1
		}
	}

	return success
}

func (v *Visitor) VisitChildren(_ antlr.RuleNode) interface{} {
	return nil
}

func (v *Visitor) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}

func (v *Visitor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (v *Visitor) VisitProgram(ctx *genAntlr.ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitProgDeclare(ctx *genAntlr.ProgDeclareContext) interface{} {
	return v.VisitProgramDescription(ctx.GetChild(1).(*genAntlr.ProgramDescriptionContext))
}

func (v *Visitor) VisitVarDeclare(ctx *genAntlr.VarDeclareContext) interface{} {
	for _, vdCtx := range ctx.GetChildren() {
		if _, ok := vdCtx.(*genAntlr.VarListContext); !ok {
			continue
		}
		result := v.VisitVarList(vdCtx.(*genAntlr.VarListContext))
		if !result.(bool) {
			v.errorHandler.Error(errorhandler.CustomError{
				Type:     errorhandler.ProgramError,
				Line:     vdCtx.(*genAntlr.VarListContext).GetStart().GetLine(),
				Position: vdCtx.(*genAntlr.VarListContext).GetStart().GetColumn(),
				Message:  "объявление переменных закончилось неудачей\n",
			})
			return fail
		}
	}
	return success
}

func (v *Visitor) VisitVarList(ctx *genAntlr.VarListContext) interface{} {
	idCtxCount := 0
	for _, vlCtx := range ctx.GetChildren() {
		if _, ok := vlCtx.(*genAntlr.IdContext); !ok {
			continue
		}
		v.ids[vlCtx.(*genAntlr.IdContext).GetText()] = 0
		idCtxCount += 1
	}

	if idCtxCount == 0 {
		return fail
	}
	return success
}

func (v *Visitor) VisitProgramDescription(ctx *genAntlr.ProgramDescriptionContext) interface{} {
	for _, vpCtx := range ctx.GetChildren() {
		switch vpCtx.(type) {
		case *genAntlr.AssignContext:
			v.VisitAssign(vpCtx.(*genAntlr.AssignContext))
		case *genAntlr.LoopContext:
			v.VisitLoop(vpCtx.(*genAntlr.LoopContext))
		case *genAntlr.ReadContext:
			v.VisitRead(vpCtx.(*genAntlr.ReadContext))
		case *genAntlr.WriteContext:
			v.VisitWrite(vpCtx.(*genAntlr.WriteContext))
		}
	}
	return success
}

func (v *Visitor) VisitAssign(ctx *genAntlr.AssignContext) interface{} {
	id := v.VisitId(ctx.GetChild(0).(*genAntlr.IdContext))
	if id == nil {
		v.errorHandler.Error(errorhandler.CustomError{
			Type:     errorhandler.ProgramError,
			Line:     ctx.GetStart().GetLine(),
			Position: ctx.GetStart().GetColumn(),
			Message:  "переменная не определена",
		})
	}
	v.VisitExpression(ctx.GetChild(2).(*genAntlr.ExpressionContext))
	v.ids[ctx.GetChild(0).(*genAntlr.IdContext).GetText()] = v.pop()
	return success
}

func (v *Visitor) VisitExpression(ctx *genAntlr.ExpressionContext) interface{} {
	if _, ok := ctx.GetChild(0).(*genAntlr.UnaryOperatorContext); ok {
		v.push(-v.VisitSubexpression(ctx.GetChild(1).(*genAntlr.SubexpressionContext)).(int))
	}
	v.push(v.VisitSubexpression(ctx.GetChild(0).(*genAntlr.SubexpressionContext)).(int))
	return success
}

func (v *Visitor) VisitSubexpression(ctx *genAntlr.SubexpressionContext) interface{} {
	var lastBinaryOperator string
	for _, seCtx := range ctx.GetChildren() {
		switch seCtx.(type) {
		case *genAntlr.ExpressionContext:
			v.VisitExpression(seCtx.(*genAntlr.ExpressionContext))
		case *genAntlr.OperandContext:
			v.VisitOperand(seCtx.(*genAntlr.OperandContext))
		case *genAntlr.SubexpressionContext:
			result := v.VisitSubexpression(seCtx.(*genAntlr.SubexpressionContext))
			v.push(result.(int))
			if len(v.stack) == 2 {
				v.calc(lastBinaryOperator)
			}
		case *genAntlr.BinaryOperatorContext:
			lastBinaryOperator = v.VisitBinaryOperator(seCtx.(*genAntlr.BinaryOperatorContext)).(string)
		}
	}
	return v.pop()
}

func (v *Visitor) VisitUnaryOperator(ctx *genAntlr.UnaryOperatorContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitBinaryOperator(ctx *genAntlr.BinaryOperatorContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitOperand(ctx *genAntlr.OperandContext) interface{} {
	oCtx := ctx.GetChild(0)
	switch oCtx.(type) {
	case *genAntlr.IdContext:
		num := v.VisitId(oCtx.(*genAntlr.IdContext))
		if num == nil {
			v.errorHandler.Error(errorhandler.CustomError{
				Type:     errorhandler.ProgramError,
				Line:     ctx.GetStart().GetLine(),
				Position: ctx.GetStart().GetColumn(),
				Message:  "переменная не определена",
			})
		}
		v.push(num.(int))
	case *genAntlr.ConstContext:
		num := v.VisitConst(oCtx.(*genAntlr.ConstContext))
		if num == nil {
			v.errorHandler.Error(errorhandler.CustomError{
				Type:     errorhandler.ProgramError,
				Line:     ctx.GetStart().GetLine(),
				Position: ctx.GetStart().GetColumn(),
				Message:  "встретилось не число",
			})
		}
		v.push(num.(int))
	}
	return success
}

func (v *Visitor) VisitId(ctx *genAntlr.IdContext) interface{} {
	id, ok := v.ids[ctx.GetText()]
	if !ok {
		return nil
	}
	return id
}

func (v *Visitor) VisitConst(ctx *genAntlr.ConstContext) interface{} {
	num, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		return nil
	}

	return num
}

func (v *Visitor) VisitLoop(ctx *genAntlr.LoopContext) interface{} {
	loopId := ctx.GetChild(1).(*genAntlr.AssignContext).Id().GetText()
	loopIdValueBefore := v.ids[loopId]
	v.VisitAssign(ctx.GetChild(1).(*genAntlr.AssignContext))
	loopIdValue := v.ids[loopId]
	v.VisitExpression(ctx.GetChild(3).(*genAntlr.ExpressionContext))
	loopToValue := v.pop()

	for i := loopIdValue; i < loopToValue; i++ {
		v.VisitProgramDescription(ctx.GetChild(5).(*genAntlr.ProgramDescriptionContext))
	}

	v.ids[loopId] = loopIdValueBefore
	return success
}

func (v *Visitor) VisitRead(ctx *genAntlr.ReadContext) interface{} {
	varListCtx := ctx.GetChild(2).(*genAntlr.VarListContext).GetChildren()
	readIds := make([]string, 0)
	for _, vlCtx := range varListCtx {
		if idCtx, ok := vlCtx.(*genAntlr.IdContext); ok {
			readIds = append(readIds, idCtx.GetText())
		}
	}

	for _, id := range readIds {
		var tempNum int
		fmt.Scanln(&tempNum)
		v.ids[id] = tempNum
	}
	return success
}

func (v *Visitor) VisitWrite(ctx *genAntlr.WriteContext) interface{} {
	varListCtx := ctx.GetChild(2).(*genAntlr.VarListContext).GetChildren()
	for _, vlCtx := range varListCtx {
		if idCtx, ok := vlCtx.(*genAntlr.IdContext); ok {
			fmt.Printf("%s = %d ", idCtx.GetText(), v.ids[idCtx.GetText()])
		}
	}
	fmt.Println()
	return success
}
