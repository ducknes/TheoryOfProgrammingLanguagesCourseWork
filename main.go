package main

import (
	"TheoryOfProgrammingLanguagesCourseWork/customVisitor"
	"TheoryOfProgrammingLanguagesCourseWork/errorhandler"
	genAntlr "TheoryOfProgrammingLanguagesCourseWork/gen"
	_ "embed"
	"github.com/antlr4-go/antlr/v4"
)

//go:embed testProg.txt
var program string

func main() {

	customErrorHandler := errorhandler.New()

	visitor := customVisitor.New(customErrorHandler)

	sim := program[len(program)-1]

	if sim != ';' {
		customErrorHandler.Error(errorhandler.CustomError{
			Message: "лишние символы после конца программы или нет ;",
		})
	}

	is := antlr.NewInputStream(program)
	lexer := genAntlr.NewcourseWorkGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := genAntlr.NewcourseWorkGrammarParser(stream)
	parser.AddErrorListener(customErrorHandler)
	tree := parser.Program()

	visitor.Visit(tree)
}
