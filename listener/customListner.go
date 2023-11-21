package listener

import (
	"TheoryOfProgrammingLanguagesCourseWork/errorhandler"
	genAntlr "TheoryOfProgrammingLanguagesCourseWork/gen"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type Listener struct {
	program string
	genAntlr.BasecourseWorkGrammarListener
	errorHandler *errorhandler.CustomErrorHandler
	ctx          antlr.BaseInterpreterRuleContext

	ids   map[string]int
	stack []int
}

func New(errorHandler *errorhandler.CustomErrorHandler, program string) *Listener {
	return &Listener{
		errorHandler: errorHandler,
		program:      program,
		ids:          make(map[string]int),
	}
}

func (l *Listener) calc(symbol string) int {
	switch symbol {
	case "+":
		return l.pop() + l.pop()
	case "-":
		return l.pop() - l.pop()
	case "*":
		return l.pop() * l.pop()
	default:
		return 0
	}
}

func (l *Listener) push(value int) {
	l.stack = append(l.stack, value)
}

func (l *Listener) pop() int {
	if len(l.stack) == 0 {
		l.errorHandler.Error(errorhandler.CustomError{
			Type:    errorhandler.StackError,
			Message: "стек пуст\n",
		})
	}

	value := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return value
}

// EnterEveryRule is called when any rule is entered.
func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (l *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (l *Listener) EnterProgram(ctx *genAntlr.ProgramContext) {}

// ExitProgram is called when production program is exited.
func (l *Listener) ExitProgram(ctx *genAntlr.ProgramContext) {
	for key, value := range l.ids {
		fmt.Printf("%s = %d\n", key, value)
	}
}

// EnterProgDeclare is called when production progDeclare is entered.
func (l *Listener) EnterProgDeclare(ctx *genAntlr.ProgDeclareContext) {}

// ExitProgDeclare is called when production progDeclare is exited.
func (l *Listener) ExitProgDeclare(ctx *genAntlr.ProgDeclareContext) {}

// EnterVarDeclare is called when production varDeclare is entered.
func (l *Listener) EnterVarDeclare(ctx *genAntlr.VarDeclareContext) {}

// ExitVarDeclare is called when production varDeclare is exited.
func (l *Listener) ExitVarDeclare(ctx *genAntlr.VarDeclareContext) {}

// EnterVarList is called when production varList is entered.
func (l *Listener) EnterVarList(ctx *genAntlr.VarListContext) {
	idList := ctx.AllId()
	if len(idList) <= 1 {
		l.errorHandler.Error(errorhandler.CustomError{
			Line:    ctx.GetStart().GetLine(),
			Message: "нет ни одной переменной для объявления",
		})
	}
	for _, id := range idList {
		if len(id.GetText()) > 9 {
			l.errorHandler.Error(errorhandler.CustomError{
				Line:    ctx.GetStart().GetLine(),
				Message: fmt.Sprintf("длина переменной '%s' (%d) больше допустимой (9)", id.GetText(), len(id.GetText())),
			})
		}
		l.ids[id.GetText()] = 0
	}
}

// ExitVarList is called when production varList is exited.
func (l *Listener) ExitVarList(ctx *genAntlr.VarListContext) {}

// EnterProgramDescription is called when production programDescription is entered.
func (l *Listener) EnterProgramDescription(ctx *genAntlr.ProgramDescriptionContext) {
	//if loop, ok := ctx.GetChild(0).(*genAntlr.LoopContext); ok {
	//	l.EnterLoop(loop)
	//}
}

// ExitProgramDescription is called when production programDescription is exited.
func (l *Listener) ExitProgramDescription(ctx *genAntlr.ProgramDescriptionContext) {}

// EnterAssign is called when production assign is entered.
func (l *Listener) EnterAssign(ctx *genAntlr.AssignContext) {}

// ExitAssign is called when production assign is exited.
func (l *Listener) ExitAssign(ctx *genAntlr.AssignContext) {
	id := ctx.Id().GetText()
	l.ids[id] = l.pop()
}

// EnterExpression is called when production expression is entered.
func (l *Listener) EnterExpression(ctx *genAntlr.ExpressionContext) {
	//fmt.Println(ctx.GetText())
}

// ExitExpression is called when production expression is exited.
func (l *Listener) ExitExpression(ctx *genAntlr.ExpressionContext) {
	if len(l.stack) == 0 {
		return
	}
	start := ctx.GetStart().GetTokenType()
	if start == courseWorkGrammarParserSUB {
		value := -l.pop()
		l.push(value)
		return
	}
}

// EnterSubexpression is called when production subexpression is entered.
func (l *Listener) EnterSubexpression(ctx *genAntlr.SubexpressionContext) {
}

// ExitSubexpression is called when production subexpression is exited.
func (l *Listener) ExitSubexpression(ctx *genAntlr.SubexpressionContext) {
	children := ctx.GetChildren()
	if len(children) == 3 {
		switch tt := children[1].GetChild(0).GetPayload().(type) {
		case *antlr.BaseParserRuleContext:
			return
		case *antlr.CommonToken:
			childToken := tt.GetTokenType()
			switch childToken {
			case courseWorkGrammarParserADD:
				value := l.calc(children[1].GetChild(0).GetPayload().(*antlr.CommonToken).GetText())
				l.push(value)
			case courseWorkGrammarParserMUL:
				value := l.calc(children[1].GetChild(0).GetPayload().(*antlr.CommonToken).GetText())
				l.push(value)
			case courseWorkGrammarParserSUB:
				value := l.calc(children[1].GetChild(0).GetPayload().(*antlr.CommonToken).GetText())
				l.push(value)
			}
		}
	}
}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (l *Listener) EnterUnaryOperator(ctx *genAntlr.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (l *Listener) ExitUnaryOperator(ctx *genAntlr.UnaryOperatorContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (l *Listener) EnterBinaryOperator(ctx *genAntlr.BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (l *Listener) ExitBinaryOperator(ctx *genAntlr.BinaryOperatorContext) {}

// EnterOperand is called when production operand is entered.
func (l *Listener) EnterOperand(ctx *genAntlr.OperandContext) {
	for i, v := range ctx.GetChildren() {
		child := v.GetChild(i).(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		switch child {
		case courseWorkGrammarParserINTEGER:
			value, err := strconv.Atoi(v.GetChild(i).GetPayload().(*antlr.CommonToken).GetText())
			if err != nil {
				fmt.Println(err)
				return
			}
			l.push(value)
		case courseWorkGrammarParserLITERAL:
			value := l.ids[v.GetChild(i).GetPayload().(*antlr.CommonToken).GetText()]
			l.push(value)
		}
	}
}

// ExitOperand is called when production operand is exited.
func (l *Listener) ExitOperand(ctx *genAntlr.OperandContext) {}

// EnterId is called when production id is entered.
func (l *Listener) EnterId(ctx *genAntlr.IdContext) {
	_, ok := l.ids[ctx.GetText()]
	if !ok {
		l.errorHandler.Error(errorhandler.CustomError{
			Line:    ctx.GetStart().GetLine(),
			Message: fmt.Sprintf("переменная '%s' не определена", ctx.GetText()),
		})
	}
}

// ExitId is called when production id is exited.
func (l *Listener) ExitId(ctx *genAntlr.IdContext) {}

// EnterConst is called when production const is entered.
func (l *Listener) EnterConst(ctx *genAntlr.ConstContext) {}

// ExitConst is called when production const is exited.
func (l *Listener) ExitConst(ctx *genAntlr.ConstContext) {}

// EnterLoop is called when production loop is entered.
func (l *Listener) EnterLoop(ctx *genAntlr.LoopContext) {}

// ExitLoop is called when production loop is exited.
func (l *Listener) ExitLoop(ctx *genAntlr.LoopContext) {
	id := ctx.GetChild(1).(*genAntlr.AssignContext).Id().GetText()
	idValue := l.ids[id]
	toLoop := l.pop()

	for temp := idValue; temp < toLoop; temp++ {
		l.EnterProgramDescription(ctx.GetChild(5).(*genAntlr.ProgramDescriptionContext))
	}
}

// EnterRead is called when production read is entered.
func (l *Listener) EnterRead(ctx *genAntlr.ReadContext) {}

// ExitRead is called when production read is exited.
func (l *Listener) ExitRead(ctx *genAntlr.ReadContext) {}

// EnterWrite is called when production write is entered.
func (l *Listener) EnterWrite(ctx *genAntlr.WriteContext) {}

// ExitWrite is called when production write is exited.
func (l *Listener) ExitWrite(ctx *genAntlr.WriteContext) {}
