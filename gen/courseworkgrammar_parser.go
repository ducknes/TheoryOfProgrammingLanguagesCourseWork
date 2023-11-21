// Code generated from /Users/ilyaantonov/Downloads/3 курс/ТЯП/cwp/courseWorkGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // courseWorkGrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type courseWorkGrammarParser struct {
	*antlr.BaseParser
}

var CourseWorkGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func courseworkgrammarParserInit() {
	staticData := &CourseWorkGrammarParserStaticData
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
		"program", "progDeclare", "varDeclare", "varList", "programDescription",
		"assign", "expression", "subexpression", "unaryOperator", "binaryOperator",
		"operand", "id", "const", "loop", "read", "write",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 132, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 49, 8, 3, 10, 3, 12, 3, 52, 9, 3, 1, 4, 4,
		4, 55, 8, 4, 11, 4, 12, 4, 56, 1, 4, 1, 4, 1, 4, 4, 4, 62, 8, 4, 11, 4,
		12, 4, 63, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6,
		75, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 83, 8, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 3, 10, 100, 8, 10, 1, 11, 5, 11, 103, 8, 11, 10, 11,
		12, 11, 106, 9, 11, 1, 12, 5, 12, 109, 8, 12, 10, 12, 12, 12, 112, 9, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 0, 1, 14, 16,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 1, 1, 0,
		4, 6, 127, 0, 32, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6,
		45, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 65, 1, 0, 0, 0, 12, 74, 1, 0, 0,
		0, 14, 82, 1, 0, 0, 0, 16, 93, 1, 0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 99,
		1, 0, 0, 0, 22, 104, 1, 0, 0, 0, 24, 110, 1, 0, 0, 0, 26, 113, 1, 0, 0,
		0, 28, 121, 1, 0, 0, 0, 30, 126, 1, 0, 0, 0, 32, 33, 3, 4, 2, 0, 33, 34,
		3, 2, 1, 0, 34, 1, 1, 0, 0, 0, 35, 36, 5, 15, 0, 0, 36, 37, 3, 8, 4, 0,
		37, 38, 5, 16, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40, 5, 18, 0, 0, 40, 41, 3,
		6, 3, 0, 41, 42, 5, 19, 0, 0, 42, 43, 5, 17, 0, 0, 43, 44, 5, 14, 0, 0,
		44, 5, 1, 0, 0, 0, 45, 50, 3, 22, 11, 0, 46, 47, 5, 20, 0, 0, 47, 49, 3,
		22, 11, 0, 48, 46, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 7, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 55, 3, 10,
		5, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57,
		1, 0, 0, 0, 57, 62, 1, 0, 0, 0, 58, 62, 3, 26, 13, 0, 59, 62, 3, 28, 14,
		0, 60, 62, 3, 30, 15, 0, 61, 54, 1, 0, 0, 0, 61, 58, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 9, 1, 0, 0, 0, 65, 66, 3, 22, 11, 0, 66, 67, 5,
		7, 0, 0, 67, 68, 3, 12, 6, 0, 68, 69, 5, 14, 0, 0, 69, 11, 1, 0, 0, 0,
		70, 71, 3, 16, 8, 0, 71, 72, 3, 14, 7, 0, 72, 75, 1, 0, 0, 0, 73, 75, 3,
		14, 7, 0, 74, 70, 1, 0, 0, 0, 74, 73, 1, 0, 0, 0, 75, 13, 1, 0, 0, 0, 76,
		77, 6, 7, -1, 0, 77, 78, 5, 8, 0, 0, 78, 79, 3, 12, 6, 0, 79, 80, 5, 9,
		0, 0, 80, 83, 1, 0, 0, 0, 81, 83, 3, 20, 10, 0, 82, 76, 1, 0, 0, 0, 82,
		81, 1, 0, 0, 0, 83, 90, 1, 0, 0, 0, 84, 85, 10, 1, 0, 0, 85, 86, 3, 18,
		9, 0, 86, 87, 3, 14, 7, 2, 87, 89, 1, 0, 0, 0, 88, 84, 1, 0, 0, 0, 89,
		92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 15, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 5, 0, 0, 94, 17, 1, 0, 0, 0, 95, 96,
		7, 0, 0, 0, 96, 19, 1, 0, 0, 0, 97, 100, 3, 22, 11, 0, 98, 100, 3, 24,
		12, 0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 21, 1, 0, 0, 0, 101,
		103, 5, 1, 0, 0, 102, 101, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 23, 1, 0, 0, 0, 106, 104, 1, 0,
		0, 0, 107, 109, 5, 2, 0, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0,
		110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 25, 1, 0, 0, 0, 112, 110,
		1, 0, 0, 0, 113, 114, 5, 12, 0, 0, 114, 115, 3, 10, 5, 0, 115, 116, 5,
		13, 0, 0, 116, 117, 3, 12, 6, 0, 117, 118, 5, 21, 0, 0, 118, 119, 3, 8,
		4, 0, 119, 120, 5, 22, 0, 0, 120, 27, 1, 0, 0, 0, 121, 122, 5, 10, 0, 0,
		122, 123, 5, 8, 0, 0, 123, 124, 3, 6, 3, 0, 124, 125, 5, 9, 0, 0, 125,
		29, 1, 0, 0, 0, 126, 127, 5, 11, 0, 0, 127, 128, 5, 8, 0, 0, 128, 129,
		3, 6, 3, 0, 129, 130, 5, 9, 0, 0, 130, 31, 1, 0, 0, 0, 10, 50, 56, 61,
		63, 74, 82, 90, 99, 104, 110,
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

// courseWorkGrammarParserInit initializes any static state used to implement courseWorkGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewcourseWorkGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CourseWorkGrammarParserInit() {
	staticData := &CourseWorkGrammarParserStaticData
	staticData.once.Do(courseworkgrammarParserInit)
}

// NewcourseWorkGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewcourseWorkGrammarParser(input antlr.TokenStream) *courseWorkGrammarParser {
	CourseWorkGrammarParserInit()
	this := new(courseWorkGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CourseWorkGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "courseWorkGrammar.g4"

	return this
}

// courseWorkGrammarParser tokens.
const (
	courseWorkGrammarParserEOF           = antlr.TokenEOF
	courseWorkGrammarParserLITERAL       = 1
	courseWorkGrammarParserINTEGER       = 2
	courseWorkGrammarParserWS            = 3
	courseWorkGrammarParserADD           = 4
	courseWorkGrammarParserSUB           = 5
	courseWorkGrammarParserMUL           = 6
	courseWorkGrammarParserASSIGN        = 7
	courseWorkGrammarParserLEFT_BRACKET  = 8
	courseWorkGrammarParserRIGHT_BRACKET = 9
	courseWorkGrammarParserREAD          = 10
	courseWorkGrammarParserWRITE         = 11
	courseWorkGrammarParserFOR           = 12
	courseWorkGrammarParserTO            = 13
	courseWorkGrammarParserSEMICOLON     = 14
	courseWorkGrammarParserBEGIN         = 15
	courseWorkGrammarParserEND           = 16
	courseWorkGrammarParserINT           = 17
	courseWorkGrammarParserVAR           = 18
	courseWorkGrammarParserCOLON         = 19
	courseWorkGrammarParserCOMMA         = 20
	courseWorkGrammarParserDO            = 21
	courseWorkGrammarParserEND_FOR       = 22
)

// courseWorkGrammarParser rules.
const (
	courseWorkGrammarParserRULE_program            = 0
	courseWorkGrammarParserRULE_progDeclare        = 1
	courseWorkGrammarParserRULE_varDeclare         = 2
	courseWorkGrammarParserRULE_varList            = 3
	courseWorkGrammarParserRULE_programDescription = 4
	courseWorkGrammarParserRULE_assign             = 5
	courseWorkGrammarParserRULE_expression         = 6
	courseWorkGrammarParserRULE_subexpression      = 7
	courseWorkGrammarParserRULE_unaryOperator      = 8
	courseWorkGrammarParserRULE_binaryOperator     = 9
	courseWorkGrammarParserRULE_operand            = 10
	courseWorkGrammarParserRULE_id                 = 11
	courseWorkGrammarParserRULE_const              = 12
	courseWorkGrammarParserRULE_loop               = 13
	courseWorkGrammarParserRULE_read               = 14
	courseWorkGrammarParserRULE_write              = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDeclare() IVarDeclareContext
	ProgDeclare() IProgDeclareContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) VarDeclare() IVarDeclareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclareContext)
}

func (s *ProgramContext) ProgDeclare() IProgDeclareContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgDeclareContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgDeclareContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, courseWorkGrammarParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.VarDeclare()
	}
	{
		p.SetState(33)
		p.ProgDeclare()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgDeclareContext is an interface to support dynamic dispatch.
type IProgDeclareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN() antlr.TerminalNode
	ProgramDescription() IProgramDescriptionContext
	END() antlr.TerminalNode

	// IsProgDeclareContext differentiates from other interfaces.
	IsProgDeclareContext()
}

type ProgDeclareContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgDeclareContext() *ProgDeclareContext {
	var p = new(ProgDeclareContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_progDeclare
	return p
}

func InitEmptyProgDeclareContext(p *ProgDeclareContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_progDeclare
}

func (*ProgDeclareContext) IsProgDeclareContext() {}

func NewProgDeclareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgDeclareContext {
	var p = new(ProgDeclareContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_progDeclare

	return p
}

func (s *ProgDeclareContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgDeclareContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserBEGIN, 0)
}

func (s *ProgDeclareContext) ProgramDescription() IProgramDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramDescriptionContext)
}

func (s *ProgDeclareContext) END() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserEND, 0)
}

func (s *ProgDeclareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgDeclareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgDeclareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterProgDeclare(s)
	}
}

func (s *ProgDeclareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitProgDeclare(s)
	}
}

func (s *ProgDeclareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitProgDeclare(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) ProgDeclare() (localctx IProgDeclareContext) {
	localctx = NewProgDeclareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, courseWorkGrammarParserRULE_progDeclare)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(courseWorkGrammarParserBEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.ProgramDescription()
	}
	{
		p.SetState(37)
		p.Match(courseWorkGrammarParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclareContext is an interface to support dynamic dispatch.
type IVarDeclareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	VarList() IVarListContext
	COLON() antlr.TerminalNode
	INT() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsVarDeclareContext differentiates from other interfaces.
	IsVarDeclareContext()
}

type VarDeclareContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclareContext() *VarDeclareContext {
	var p = new(VarDeclareContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_varDeclare
	return p
}

func InitEmptyVarDeclareContext(p *VarDeclareContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_varDeclare
}

func (*VarDeclareContext) IsVarDeclareContext() {}

func NewVarDeclareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclareContext {
	var p = new(VarDeclareContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_varDeclare

	return p
}

func (s *VarDeclareContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclareContext) VAR() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserVAR, 0)
}

func (s *VarDeclareContext) VarList() IVarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *VarDeclareContext) COLON() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserCOLON, 0)
}

func (s *VarDeclareContext) INT() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserINT, 0)
}

func (s *VarDeclareContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserSEMICOLON, 0)
}

func (s *VarDeclareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterVarDeclare(s)
	}
}

func (s *VarDeclareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitVarDeclare(s)
	}
}

func (s *VarDeclareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitVarDeclare(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) VarDeclare() (localctx IVarDeclareContext) {
	localctx = NewVarDeclareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, courseWorkGrammarParserRULE_varDeclare)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(courseWorkGrammarParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.VarList()
	}
	{
		p.SetState(41)
		p.Match(courseWorkGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(courseWorkGrammarParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(courseWorkGrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarListContext is an interface to support dynamic dispatch.
type IVarListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllId() []IIdContext
	Id(i int) IIdContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVarListContext differentiates from other interfaces.
	IsVarListContext()
}

type VarListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarListContext() *VarListContext {
	var p = new(VarListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_varList
	return p
}

func InitEmptyVarListContext(p *VarListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_varList
}

func (*VarListContext) IsVarListContext() {}

func NewVarListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarListContext {
	var p = new(VarListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_varList

	return p
}

func (s *VarListContext) GetParser() antlr.Parser { return s.parser }

func (s *VarListContext) AllId() []IIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdContext); ok {
			len++
		}
	}

	tst := make([]IIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdContext); ok {
			tst[i] = t.(IIdContext)
			i++
		}
	}

	return tst
}

func (s *VarListContext) Id(i int) IIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *VarListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(courseWorkGrammarParserCOMMA)
}

func (s *VarListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserCOMMA, i)
}

func (s *VarListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterVarList(s)
	}
}

func (s *VarListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitVarList(s)
	}
}

func (s *VarListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitVarList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) VarList() (localctx IVarListContext) {
	localctx = NewVarListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, courseWorkGrammarParserRULE_varList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Id()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == courseWorkGrammarParserCOMMA {
		{
			p.SetState(46)
			p.Match(courseWorkGrammarParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Id()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgramDescriptionContext is an interface to support dynamic dispatch.
type IProgramDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLoop() []ILoopContext
	Loop(i int) ILoopContext
	AllRead() []IReadContext
	Read(i int) IReadContext
	AllWrite() []IWriteContext
	Write(i int) IWriteContext
	AllAssign() []IAssignContext
	Assign(i int) IAssignContext

	// IsProgramDescriptionContext differentiates from other interfaces.
	IsProgramDescriptionContext()
}

type ProgramDescriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramDescriptionContext() *ProgramDescriptionContext {
	var p = new(ProgramDescriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_programDescription
	return p
}

func InitEmptyProgramDescriptionContext(p *ProgramDescriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_programDescription
}

func (*ProgramDescriptionContext) IsProgramDescriptionContext() {}

func NewProgramDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramDescriptionContext {
	var p = new(ProgramDescriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_programDescription

	return p
}

func (s *ProgramDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramDescriptionContext) AllLoop() []ILoopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILoopContext); ok {
			len++
		}
	}

	tst := make([]ILoopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILoopContext); ok {
			tst[i] = t.(ILoopContext)
			i++
		}
	}

	return tst
}

func (s *ProgramDescriptionContext) Loop(i int) ILoopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *ProgramDescriptionContext) AllRead() []IReadContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReadContext); ok {
			len++
		}
	}

	tst := make([]IReadContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReadContext); ok {
			tst[i] = t.(IReadContext)
			i++
		}
	}

	return tst
}

func (s *ProgramDescriptionContext) Read(i int) IReadContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *ProgramDescriptionContext) AllWrite() []IWriteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWriteContext); ok {
			len++
		}
	}

	tst := make([]IWriteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWriteContext); ok {
			tst[i] = t.(IWriteContext)
			i++
		}
	}

	return tst
}

func (s *ProgramDescriptionContext) Write(i int) IWriteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWriteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWriteContext)
}

func (s *ProgramDescriptionContext) AllAssign() []IAssignContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignContext); ok {
			len++
		}
	}

	tst := make([]IAssignContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignContext); ok {
			tst[i] = t.(IAssignContext)
			i++
		}
	}

	return tst
}

func (s *ProgramDescriptionContext) Assign(i int) IAssignContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ProgramDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterProgramDescription(s)
	}
}

func (s *ProgramDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitProgramDescription(s)
	}
}

func (s *ProgramDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitProgramDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) ProgramDescription() (localctx IProgramDescriptionContext) {
	localctx = NewProgramDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, courseWorkGrammarParserRULE_programDescription)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7298) != 0) {
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case courseWorkGrammarParserLITERAL, courseWorkGrammarParserASSIGN:
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(53)
						p.Assign()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(56)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case courseWorkGrammarParserFOR:
			{
				p.SetState(58)
				p.Loop()
			}

		case courseWorkGrammarParserREAD:
			{
				p.SetState(59)
				p.Read()
			}

		case courseWorkGrammarParserWRITE:
			{
				p.SetState(60)
				p.Write()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserASSIGN, 0)
}

func (s *AssignContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserSEMICOLON, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, courseWorkGrammarParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Id()
	}
	{
		p.SetState(66)
		p.Match(courseWorkGrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Expression()
	}
	{
		p.SetState(68)
		p.Match(courseWorkGrammarParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOperator() IUnaryOperatorContext
	Subexpression() ISubexpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) UnaryOperator() IUnaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *ExpressionContext) Subexpression() ISubexpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubexpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubexpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, courseWorkGrammarParserRULE_expression)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.UnaryOperator()
		}
		{
			p.SetState(71)
			p.subexpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.subexpression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubexpressionContext is an interface to support dynamic dispatch.
type ISubexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	RIGHT_BRACKET() antlr.TerminalNode
	Operand() IOperandContext
	AllSubexpression() []ISubexpressionContext
	Subexpression(i int) ISubexpressionContext
	BinaryOperator() IBinaryOperatorContext

	// IsSubexpressionContext differentiates from other interfaces.
	IsSubexpressionContext()
}

type SubexpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubexpressionContext() *SubexpressionContext {
	var p = new(SubexpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_subexpression
	return p
}

func InitEmptySubexpressionContext(p *SubexpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_subexpression
}

func (*SubexpressionContext) IsSubexpressionContext() {}

func NewSubexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubexpressionContext {
	var p = new(SubexpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_subexpression

	return p
}

func (s *SubexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubexpressionContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserLEFT_BRACKET, 0)
}

func (s *SubexpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubexpressionContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserRIGHT_BRACKET, 0)
}

func (s *SubexpressionContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *SubexpressionContext) AllSubexpression() []ISubexpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubexpressionContext); ok {
			len++
		}
	}

	tst := make([]ISubexpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubexpressionContext); ok {
			tst[i] = t.(ISubexpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubexpressionContext) Subexpression(i int) ISubexpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubexpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubexpressionContext)
}

func (s *SubexpressionContext) BinaryOperator() IBinaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *SubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterSubexpression(s)
	}
}

func (s *SubexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitSubexpression(s)
	}
}

func (s *SubexpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitSubexpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Subexpression() (localctx ISubexpressionContext) {
	return p.subexpression(0)
}

func (p *courseWorkGrammarParser) subexpression(_p int) (localctx ISubexpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSubexpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISubexpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, courseWorkGrammarParserRULE_subexpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(77)
			p.Match(courseWorkGrammarParserLEFT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Expression()
		}
		{
			p.SetState(79)
			p.Match(courseWorkGrammarParserRIGHT_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(81)
			p.Operand()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSubexpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, courseWorkGrammarParserRULE_subexpression)
			p.SetState(84)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(85)
				p.BinaryOperator()
			}
			{
				p.SetState(86)
				p.subexpression(2)
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUB() antlr.TerminalNode

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_unaryOperator
	return p
}

func InitEmptyUnaryOperatorContext(p *UnaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_unaryOperator
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserSUB, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitUnaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, courseWorkGrammarParserRULE_unaryOperator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(courseWorkGrammarParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_binaryOperator
	return p
}

func InitEmptyBinaryOperatorContext(p *BinaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_binaryOperator
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserADD, 0)
}

func (s *BinaryOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserSUB, 0)
}

func (s *BinaryOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserMUL, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, courseWorkGrammarParserRULE_binaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	Const_() IConstContext

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *OperandContext) Const_() IConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, courseWorkGrammarParserRULE_operand)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Id()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Const_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLITERAL() []antlr.TerminalNode
	LITERAL(i int) antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) AllLITERAL() []antlr.TerminalNode {
	return s.GetTokens(courseWorkGrammarParserLITERAL)
}

func (s *IdContext) LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserLITERAL, i)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, courseWorkGrammarParserRULE_id)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(101)
				p.Match(courseWorkGrammarParserLITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstContext is an interface to support dynamic dispatch.
type IConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINTEGER() []antlr.TerminalNode
	INTEGER(i int) antlr.TerminalNode

	// IsConstContext differentiates from other interfaces.
	IsConstContext()
}

type ConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstContext() *ConstContext {
	var p = new(ConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_const
	return p
}

func InitEmptyConstContext(p *ConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_const
}

func (*ConstContext) IsConstContext() {}

func NewConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstContext {
	var p = new(ConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_const

	return p
}

func (s *ConstContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(courseWorkGrammarParserINTEGER)
}

func (s *ConstContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserINTEGER, i)
}

func (s *ConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterConst(s)
	}
}

func (s *ConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitConst(s)
	}
}

func (s *ConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Const_() (localctx IConstContext) {
	localctx = NewConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, courseWorkGrammarParserRULE_const)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(107)
				p.Match(courseWorkGrammarParserINTEGER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Assign() IAssignContext
	TO() antlr.TerminalNode
	Expression() IExpressionContext
	DO() antlr.TerminalNode
	ProgramDescription() IProgramDescriptionContext
	END_FOR() antlr.TerminalNode

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserFOR, 0)
}

func (s *LoopContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *LoopContext) TO() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserTO, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) DO() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserDO, 0)
}

func (s *LoopContext) ProgramDescription() IProgramDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramDescriptionContext)
}

func (s *LoopContext) END_FOR() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserEND_FOR, 0)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, courseWorkGrammarParserRULE_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(courseWorkGrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Assign()
	}
	{
		p.SetState(115)
		p.Match(courseWorkGrammarParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Expression()
	}
	{
		p.SetState(117)
		p.Match(courseWorkGrammarParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.ProgramDescription()
	}
	{
		p.SetState(119)
		p.Match(courseWorkGrammarParserEND_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ() antlr.TerminalNode
	LEFT_BRACKET() antlr.TerminalNode
	VarList() IVarListContext
	RIGHT_BRACKET() antlr.TerminalNode

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_read
	return p
}

func InitEmptyReadContext(p *ReadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_read
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) READ() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserREAD, 0)
}

func (s *ReadContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserLEFT_BRACKET, 0)
}

func (s *ReadContext) VarList() IVarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *ReadContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserRIGHT_BRACKET, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitRead(s)
	}
}

func (s *ReadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitRead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Read() (localctx IReadContext) {
	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, courseWorkGrammarParserRULE_read)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(courseWorkGrammarParserREAD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(courseWorkGrammarParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.VarList()
	}
	{
		p.SetState(124)
		p.Match(courseWorkGrammarParserRIGHT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWriteContext is an interface to support dynamic dispatch.
type IWriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WRITE() antlr.TerminalNode
	LEFT_BRACKET() antlr.TerminalNode
	VarList() IVarListContext
	RIGHT_BRACKET() antlr.TerminalNode

	// IsWriteContext differentiates from other interfaces.
	IsWriteContext()
}

type WriteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteContext() *WriteContext {
	var p = new(WriteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_write
	return p
}

func InitEmptyWriteContext(p *WriteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = courseWorkGrammarParserRULE_write
}

func (*WriteContext) IsWriteContext() {}

func NewWriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteContext {
	var p = new(WriteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = courseWorkGrammarParserRULE_write

	return p
}

func (s *WriteContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteContext) WRITE() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserWRITE, 0)
}

func (s *WriteContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserLEFT_BRACKET, 0)
}

func (s *WriteContext) VarList() IVarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarListContext)
}

func (s *WriteContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(courseWorkGrammarParserRIGHT_BRACKET, 0)
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(courseWorkGrammarListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (s *WriteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case courseWorkGrammarVisitor:
		return t.VisitWrite(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *courseWorkGrammarParser) Write() (localctx IWriteContext) {
	localctx = NewWriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, courseWorkGrammarParserRULE_write)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(courseWorkGrammarParserWRITE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(courseWorkGrammarParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.VarList()
	}
	{
		p.SetState(129)
		p.Match(courseWorkGrammarParserRIGHT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *courseWorkGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *SubexpressionContext = nil
		if localctx != nil {
			t = localctx.(*SubexpressionContext)
		}
		return p.Subexpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *courseWorkGrammarParser) Subexpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
