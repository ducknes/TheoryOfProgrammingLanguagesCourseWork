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

func (d *CustomErrorHandler) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, _ string, _ antlr.RecognitionException) {
	os.Exit(1)
}

func (d *CustomErrorHandler) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (d *CustomErrorHandler) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (d *CustomErrorHandler) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ *antlr.ATNConfigSet) {
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
