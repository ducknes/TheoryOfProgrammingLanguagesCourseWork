package customVisitor

import (
	"TheoryOfProgrammingLanguagesCourseWork/errorhandler"
	genAntlr "TheoryOfProgrammingLanguagesCourseWork/gen"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
	"strings"
)

const (
	success = true
	fail    = false
)

type Exp struct {
	znar rune
	typ  int
}

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
	prog := ctx.GetText()
	sim := string(rune(prog[len(prog)-1]))
	if sim != ";" {
		v.errorHandler.Error(errorhandler.CustomError{
			Line:     ctx.GetStart().GetLine(),
			Position: ctx.GetStart().GetColumn(),
			Type:     errorhandler.ProgramError,
			Message:  "много",
		})
	}
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
	expRes := v.calculateExpression(ctx.GetChild(2).(*genAntlr.ExpressionContext).GetText())
	v.ids[ctx.GetChild(0).(*genAntlr.IdContext).GetText()] = expRes
	return success
}

func (v *Visitor) VisitExpression(ctx *genAntlr.ExpressionContext) interface{} {
	if _, ok := ctx.GetChild(0).(*genAntlr.UnaryOperatorContext); ok {
		v.push(-v.VisitSubexpression(ctx.GetChild(1).(*genAntlr.SubexpressionContext)).(int))
	} else {
		v.push(v.VisitSubexpression(ctx.GetChild(0).(*genAntlr.SubexpressionContext)).(int))
	}
	return success
}

func (v *Visitor) VisitSubexpression(ctx *genAntlr.SubexpressionContext) interface{} {
	var lastBinaryOperator string

	for _, seCtx := range ctx.GetChildren() {
		switch seCtx.(type) {
		case *genAntlr.ExpressionContext:
			v.VisitExpression(seCtx.(*genAntlr.ExpressionContext))
			if len(v.stack) == 3 {
				v.calc(lastBinaryOperator)
			}
		case *genAntlr.OperandContext:
			v.VisitOperand(seCtx.(*genAntlr.OperandContext))
		case *genAntlr.SubexpressionContext:
			result := v.VisitSubexpression(seCtx.(*genAntlr.SubexpressionContext))
			v.push(result.(int))
			if len(v.stack) >= 2 {
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
				Message:  fmt.Sprintf("переменная '%s' не определена", oCtx.(*genAntlr.IdContext).GetText()),
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
		v.ids[loopId] = i
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
		fmt.Printf("%s --> ", id)
		if _, err := fmt.Scanln(&tempNum); err != nil {
			v.errorHandler.Error(errorhandler.CustomError{
				Line:     ctx.GetStart().GetLine(),
				Position: ctx.GetStart().GetColumn(),
				Type:     errorhandler.ProgramError,
				Message:  "ошибка при чтении данных с клавиатуры",
			})
		}
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

func (v *Visitor) calculateExpression(expression string) int {
	var operators []Exp
	var values []int

	for key, value := range v.ids {
		expression = strings.ReplaceAll(expression, key, strconv.Itoa(value))
	}

	// Функция для применения оператора
	applyOperator := func() {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]

		right := values[len(values)-1]
		values = values[:len(values)-1]

		if operator.typ == 1 {
			values = append(values, -right)
			return
		}

		var left int
		if len(values) == 0 {
			left = 0
		} else {
			left = values[len(values)-1]
			values = values[:len(values)-1]
		}

		switch operator.znar {
		case '+':
			values = append(values, left+right)
		case '-':
			values = append(values, left-right)
		case '*':
			values = append(values, left*right)
		case '/':
			values = append(values, left/right)
		}
	}

	// Функция для получения приоритета оператора
	getPriority := func(operator rune) int {
		switch operator {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		}
		return 0
	}

	for i := 0; i < len(expression); i++ {
		switch {
		case expression[i] >= '0' && expression[i] <= '9':
			j := i
			for j < len(expression) && (expression[j] >= '0' && expression[j] <= '9' || expression[j] == '.') {
				j++
			}
			num, _ := strconv.Atoi(expression[i:j])
			values = append(values, num)
			i = j - 1
		case expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/':
			for len(operators) > 0 && getPriority(operators[len(operators)-1].znar) >= getPriority(rune(expression[i])) {
				applyOperator()
			}

			if i == 0 && expression[i] == '-' && ((expression[i+1] >= '0' && expression[i+1] <= '9') || expression[i+1] <= '(') {
				operators = append(operators, Exp{
					znar: rune(expression[i]),
					typ:  1,
				})
				continue
			}

			va := string(rune(expression[i-1]))

			if checkOp(va) && expression[i+1] >= '0' && expression[i+1] <= '9' {
				operators = append(operators, Exp{
					znar: rune(expression[i]),
					typ:  1,
				})
				continue
			}
			operators = append(operators, Exp{
				znar: rune(expression[i]),
				typ:  0,
			})
		case expression[i] == '(':
			operators = append(operators, Exp{
				znar: '(',
				typ:  0,
			})
		case expression[i] == ')':
			for operators[len(operators)-1].znar != '(' {
				applyOperator()
			}
			operators = operators[:len(operators)-1]
		}
	}

	for len(operators) > 0 {
		applyOperator()
	}

	return values[0]
}

func checkOp(op string) bool {
	ops := []string{"+", "-", "*", "("}
	for _, value := range ops {
		if value == op {
			return true
		}
	}
	return false
}
