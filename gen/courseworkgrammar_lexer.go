// Code generated from /Users/ilyaantonov/Downloads/TheoryOfProgrammingLanguagesCourseWork-test/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type courseWorkGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CourseWorkGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func courseworkgrammarlexerLexerInit() {
	staticData := &CourseWorkGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'+'", "'-'", "'*'", "'='", "'('", "')'", "'READ'",
		"'WRITE'", "'FOR'", "'TO'", "';'", "'BEGIN'", "'END'", "'INTEGER'",
		"'VAR'", "':'", "','", "'DO'", "'END_FOR'",
	}
	staticData.SymbolicNames = []string{
		"", "LITERAL", "INTEGER", "WS", "ADD", "SUB", "MUL", "ASSIGN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "READ", "WRITE", "FOR", "TO", "SEMICOLON", "BEGIN",
		"END", "INT", "VAR", "COLON", "COMMA", "DO", "END_FOR",
	}
	staticData.RuleNames = []string{
		"LITERAL", "INTEGER", "WS", "ADD", "SUB", "MUL", "ASSIGN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "READ", "WRITE", "FOR", "TO", "SEMICOLON", "BEGIN",
		"END", "INT", "VAR", "COLON", "COMMA", "DO", "END_FOR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 125, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 51, 8, 2, 11, 2,
		12, 2, 52, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 0, 0, 22, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13,
		13, 32, 32, 125, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1,
		45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 50, 1, 0, 0, 0, 7, 56, 1, 0, 0, 0,
		9, 58, 1, 0, 0, 0, 11, 60, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 64, 1, 0,
		0, 0, 17, 66, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 79,
		1, 0, 0, 0, 25, 83, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0, 29, 88, 1, 0, 0, 0,
		31, 94, 1, 0, 0, 0, 33, 98, 1, 0, 0, 0, 35, 106, 1, 0, 0, 0, 37, 110, 1,
		0, 0, 0, 39, 112, 1, 0, 0, 0, 41, 114, 1, 0, 0, 0, 43, 117, 1, 0, 0, 0,
		45, 46, 7, 0, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 7, 1, 0, 0, 48, 4, 1, 0,
		0, 0, 49, 51, 7, 2, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50,
		1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 6, 2, 0, 0,
		55, 6, 1, 0, 0, 0, 56, 57, 5, 43, 0, 0, 57, 8, 1, 0, 0, 0, 58, 59, 5, 45,
		0, 0, 59, 10, 1, 0, 0, 0, 60, 61, 5, 42, 0, 0, 61, 12, 1, 0, 0, 0, 62,
		63, 5, 61, 0, 0, 63, 14, 1, 0, 0, 0, 64, 65, 5, 40, 0, 0, 65, 16, 1, 0,
		0, 0, 66, 67, 5, 41, 0, 0, 67, 18, 1, 0, 0, 0, 68, 69, 5, 82, 0, 0, 69,
		70, 5, 69, 0, 0, 70, 71, 5, 65, 0, 0, 71, 72, 5, 68, 0, 0, 72, 20, 1, 0,
		0, 0, 73, 74, 5, 87, 0, 0, 74, 75, 5, 82, 0, 0, 75, 76, 5, 73, 0, 0, 76,
		77, 5, 84, 0, 0, 77, 78, 5, 69, 0, 0, 78, 22, 1, 0, 0, 0, 79, 80, 5, 70,
		0, 0, 80, 81, 5, 79, 0, 0, 81, 82, 5, 82, 0, 0, 82, 24, 1, 0, 0, 0, 83,
		84, 5, 84, 0, 0, 84, 85, 5, 79, 0, 0, 85, 26, 1, 0, 0, 0, 86, 87, 5, 59,
		0, 0, 87, 28, 1, 0, 0, 0, 88, 89, 5, 66, 0, 0, 89, 90, 5, 69, 0, 0, 90,
		91, 5, 71, 0, 0, 91, 92, 5, 73, 0, 0, 92, 93, 5, 78, 0, 0, 93, 30, 1, 0,
		0, 0, 94, 95, 5, 69, 0, 0, 95, 96, 5, 78, 0, 0, 96, 97, 5, 68, 0, 0, 97,
		32, 1, 0, 0, 0, 98, 99, 5, 73, 0, 0, 99, 100, 5, 78, 0, 0, 100, 101, 5,
		84, 0, 0, 101, 102, 5, 69, 0, 0, 102, 103, 5, 71, 0, 0, 103, 104, 5, 69,
		0, 0, 104, 105, 5, 82, 0, 0, 105, 34, 1, 0, 0, 0, 106, 107, 5, 86, 0, 0,
		107, 108, 5, 65, 0, 0, 108, 109, 5, 82, 0, 0, 109, 36, 1, 0, 0, 0, 110,
		111, 5, 58, 0, 0, 111, 38, 1, 0, 0, 0, 112, 113, 5, 44, 0, 0, 113, 40,
		1, 0, 0, 0, 114, 115, 5, 68, 0, 0, 115, 116, 5, 79, 0, 0, 116, 42, 1, 0,
		0, 0, 117, 118, 5, 69, 0, 0, 118, 119, 5, 78, 0, 0, 119, 120, 5, 68, 0,
		0, 120, 121, 5, 95, 0, 0, 121, 122, 5, 70, 0, 0, 122, 123, 5, 79, 0, 0,
		123, 124, 5, 82, 0, 0, 124, 44, 1, 0, 0, 0, 2, 0, 52, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// courseWorkGrammarLexerInit initializes any static state used to implement courseWorkGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewcourseWorkGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CourseWorkGrammarLexerInit() {
	staticData := &CourseWorkGrammarLexerLexerStaticData
	staticData.once.Do(courseworkgrammarlexerLexerInit)
}

// NewcourseWorkGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewcourseWorkGrammarLexer(input antlr.CharStream) *courseWorkGrammarLexer {
	CourseWorkGrammarLexerInit()
	l := new(courseWorkGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CourseWorkGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "courseWorkGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// courseWorkGrammarLexer tokens.
const (
	courseWorkGrammarLexerLITERAL       = 1
	courseWorkGrammarLexerINTEGER       = 2
	courseWorkGrammarLexerWS            = 3
	courseWorkGrammarLexerADD           = 4
	courseWorkGrammarLexerSUB           = 5
	courseWorkGrammarLexerMUL           = 6
	courseWorkGrammarLexerASSIGN        = 7
	courseWorkGrammarLexerLEFT_BRACKET  = 8
	courseWorkGrammarLexerRIGHT_BRACKET = 9
	courseWorkGrammarLexerREAD          = 10
	courseWorkGrammarLexerWRITE         = 11
	courseWorkGrammarLexerFOR           = 12
	courseWorkGrammarLexerTO            = 13
	courseWorkGrammarLexerSEMICOLON     = 14
	courseWorkGrammarLexerBEGIN         = 15
	courseWorkGrammarLexerEND           = 16
	courseWorkGrammarLexerINT           = 17
	courseWorkGrammarLexerVAR           = 18
	courseWorkGrammarLexerCOLON         = 19
	courseWorkGrammarLexerCOMMA         = 20
	courseWorkGrammarLexerDO            = 21
	courseWorkGrammarLexerEND_FOR       = 22
)
