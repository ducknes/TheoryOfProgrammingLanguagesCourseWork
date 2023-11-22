package errorhandler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

const (
	ProgramError = iota
	StackError
)

type CustomError struct {
	Line     int
	Position int
	Type     int
	Message  string
}

type CustomErrorHandler struct {
	*antlr.DefaultErrorListener
	errors []CustomError
}

func (ceh *CustomErrorHandler) SyntaxError(_ antlr.Recognizer, _ interface{}, line, pos int, msg string, _ antlr.RecognitionException) {
	ceh.Error(CustomError{
		Line:     line,
		Position: pos,
		Type:     ProgramError,
		Message:  msg,
	})
}

func (ceh *CustomErrorHandler) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (ceh *CustomErrorHandler) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (ceh *CustomErrorHandler) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ *antlr.ATNConfigSet) {
}

func New() *CustomErrorHandler {
	return &CustomErrorHandler{
		errors: make([]CustomError, 0),
	}
}

func (ceh *CustomErrorHandler) AddError(line, pos int, msg string) {
	ceh.errors = append(ceh.errors, CustomError{
		Line:     line,
		Position: pos,
		Message:  msg,
	})
}

func (ceh *CustomErrorHandler) Error(error CustomError) {
	switch error.Type {
	case ProgramError:
		fmt.Printf("PROGRAM ERROR!!! line: %d, pos: %d, msg: %s", error.Line, error.Position, error.Message)
		os.Exit(1)
	case StackError:
		fmt.Printf("STACK ERROR!!! %s", error.Message)
		os.Exit(1)
	default:
		os.Exit(1)
	}

}
