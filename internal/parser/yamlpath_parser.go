// Code generated from yamlpath.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // yamlpath

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

type yamlpathParser struct {
	*antlr.BaseParser
}

var YamlpathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yamlpathParserInit() {
	staticData := &YamlpathParserStaticData
	staticData.LiteralNames = []string{
		"", "'$'", "'@'", "'..'", "'.'", "'['", "']'", "'*'", "':'", "'('",
		"')'", "'!'", "'not'", "'+'", "'-'", "'/'", "'%'", "'|'", "'<='", "'<'",
		"'>'", "'>='", "'=='", "'!='", "'=~'", "'in'", "'nin'", "'subsetof'",
		"'&&'", "'and'", "'||'", "'or'", "'true'", "'false'", "'null'", "'{'",
		"'}'", "','", "'i'", "'m'", "'s'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "IDENTIFIER", "NUMBER", "STRING", "REGEX",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"path", "expression", "indexParam", "subexpression", "literal", "listEntries",
		"mapEntries", "invocation", "paramList", "identifier", "regex",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 176, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 32, 8,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 42, 8, 1, 10,
		1, 12, 1, 45, 9, 1, 1, 2, 1, 2, 3, 2, 49, 8, 2, 1, 2, 1, 2, 3, 2, 53, 8,
		2, 1, 2, 1, 2, 3, 2, 57, 8, 2, 1, 2, 3, 2, 60, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 73, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 99,
		8, 3, 11, 3, 12, 3, 100, 1, 3, 1, 3, 1, 3, 5, 3, 106, 8, 3, 10, 3, 12,
		3, 109, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 117, 8, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 122, 8, 4, 1, 4, 3, 4, 125, 8, 4, 1, 5, 1, 5, 1, 5, 5,
		5, 130, 8, 5, 10, 5, 12, 5, 133, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 142, 8, 6, 10, 6, 12, 6, 145, 9, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 153, 8, 7, 1, 7, 1, 7, 3, 7, 157, 8, 7, 1, 8, 1, 8,
		1, 8, 5, 8, 162, 8, 8, 10, 8, 12, 8, 165, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		5, 10, 171, 8, 10, 10, 10, 12, 10, 174, 9, 10, 1, 10, 1, 172, 2, 2, 6,
		11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 11, 1, 0, 1, 2, 1, 0, 11,
		12, 1, 0, 13, 14, 2, 0, 7, 7, 15, 16, 1, 0, 18, 21, 1, 0, 22, 23, 1, 0,
		25, 27, 1, 0, 28, 29, 1, 0, 30, 31, 1, 0, 32, 33, 1, 0, 38, 40, 202, 0,
		22, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 72, 1, 0, 0, 0,
		8, 124, 1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 156,
		1, 0, 0, 0, 16, 158, 1, 0, 0, 0, 18, 166, 1, 0, 0, 0, 20, 168, 1, 0, 0,
		0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 6,
		1, -1, 0, 26, 27, 7, 0, 0, 0, 27, 43, 1, 0, 0, 0, 28, 29, 10, 3, 0, 0,
		29, 31, 5, 3, 0, 0, 30, 32, 3, 14, 7, 0, 31, 30, 1, 0, 0, 0, 31, 32, 1,
		0, 0, 0, 32, 42, 1, 0, 0, 0, 33, 34, 10, 2, 0, 0, 34, 35, 5, 4, 0, 0, 35,
		42, 3, 14, 7, 0, 36, 37, 10, 1, 0, 0, 37, 38, 5, 5, 0, 0, 38, 39, 3, 4,
		2, 0, 39, 40, 5, 6, 0, 0, 40, 42, 1, 0, 0, 0, 41, 28, 1, 0, 0, 0, 41, 33,
		1, 0, 0, 0, 41, 36, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0,
		43, 44, 1, 0, 0, 0, 44, 3, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 60, 5, 7,
		0, 0, 47, 49, 5, 42, 0, 0, 48, 47, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49,
		50, 1, 0, 0, 0, 50, 52, 5, 8, 0, 0, 51, 53, 5, 42, 0, 0, 52, 51, 1, 0,
		0, 0, 52, 53, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 55, 5, 8, 0, 0, 55, 57,
		5, 42, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0,
		58, 60, 3, 6, 3, 0, 59, 46, 1, 0, 0, 0, 59, 48, 1, 0, 0, 0, 59, 58, 1,
		0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 62, 6, 3, -1, 0, 62, 73, 3, 2, 1, 0, 63,
		73, 3, 8, 4, 0, 64, 65, 5, 9, 0, 0, 65, 66, 3, 6, 3, 0, 66, 67, 5, 10,
		0, 0, 67, 73, 1, 0, 0, 0, 68, 69, 7, 1, 0, 0, 69, 73, 3, 6, 3, 11, 70,
		71, 7, 2, 0, 0, 71, 73, 3, 2, 1, 0, 72, 61, 1, 0, 0, 0, 72, 63, 1, 0, 0,
		0, 72, 64, 1, 0, 0, 0, 72, 68, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 107,
		1, 0, 0, 0, 74, 75, 10, 9, 0, 0, 75, 76, 7, 3, 0, 0, 76, 106, 3, 6, 3,
		10, 77, 78, 10, 8, 0, 0, 78, 79, 7, 2, 0, 0, 79, 106, 3, 6, 3, 9, 80, 81,
		10, 6, 0, 0, 81, 82, 7, 4, 0, 0, 82, 106, 3, 6, 3, 7, 83, 84, 10, 5, 0,
		0, 84, 85, 7, 5, 0, 0, 85, 106, 3, 6, 3, 6, 86, 87, 10, 3, 0, 0, 87, 88,
		7, 6, 0, 0, 88, 106, 3, 6, 3, 4, 89, 90, 10, 2, 0, 0, 90, 91, 7, 7, 0,
		0, 91, 106, 3, 6, 3, 3, 92, 93, 10, 1, 0, 0, 93, 94, 7, 8, 0, 0, 94, 106,
		3, 6, 3, 2, 95, 98, 10, 7, 0, 0, 96, 97, 5, 17, 0, 0, 97, 99, 3, 6, 3,
		0, 98, 96, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 106, 1, 0, 0, 0, 102, 103, 10, 4, 0, 0, 103, 104, 5, 24,
		0, 0, 104, 106, 3, 20, 10, 0, 105, 74, 1, 0, 0, 0, 105, 77, 1, 0, 0, 0,
		105, 80, 1, 0, 0, 0, 105, 83, 1, 0, 0, 0, 105, 86, 1, 0, 0, 0, 105, 89,
		1, 0, 0, 0, 105, 92, 1, 0, 0, 0, 105, 95, 1, 0, 0, 0, 105, 102, 1, 0, 0,
		0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		7, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 125, 5, 43, 0, 0, 111, 125, 5,
		42, 0, 0, 112, 125, 7, 9, 0, 0, 113, 125, 5, 34, 0, 0, 114, 116, 5, 5,
		0, 0, 115, 117, 3, 10, 5, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0,
		117, 118, 1, 0, 0, 0, 118, 125, 5, 6, 0, 0, 119, 121, 5, 35, 0, 0, 120,
		122, 3, 12, 6, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123,
		1, 0, 0, 0, 123, 125, 5, 36, 0, 0, 124, 110, 1, 0, 0, 0, 124, 111, 1, 0,
		0, 0, 124, 112, 1, 0, 0, 0, 124, 113, 1, 0, 0, 0, 124, 114, 1, 0, 0, 0,
		124, 119, 1, 0, 0, 0, 125, 9, 1, 0, 0, 0, 126, 131, 3, 8, 4, 0, 127, 128,
		5, 37, 0, 0, 128, 130, 3, 8, 4, 0, 129, 127, 1, 0, 0, 0, 130, 133, 1, 0,
		0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 11, 1, 0, 0, 0,
		133, 131, 1, 0, 0, 0, 134, 135, 5, 43, 0, 0, 135, 136, 5, 8, 0, 0, 136,
		143, 3, 8, 4, 0, 137, 138, 5, 37, 0, 0, 138, 139, 5, 43, 0, 0, 139, 140,
		5, 8, 0, 0, 140, 142, 3, 8, 4, 0, 141, 137, 1, 0, 0, 0, 142, 145, 1, 0,
		0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 13, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 146, 157, 3, 18, 9, 0, 147, 157, 5, 43, 0, 0, 148,
		157, 5, 7, 0, 0, 149, 150, 3, 18, 9, 0, 150, 152, 5, 9, 0, 0, 151, 153,
		3, 16, 8, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 155, 5, 10, 0, 0, 155, 157, 1, 0, 0, 0, 156, 146, 1, 0, 0, 0,
		156, 147, 1, 0, 0, 0, 156, 148, 1, 0, 0, 0, 156, 149, 1, 0, 0, 0, 157,
		15, 1, 0, 0, 0, 158, 163, 3, 6, 3, 0, 159, 160, 5, 37, 0, 0, 160, 162,
		3, 6, 3, 0, 161, 159, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 163, 164, 1, 0, 0, 0, 164, 17, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0,
		166, 167, 5, 41, 0, 0, 167, 19, 1, 0, 0, 0, 168, 172, 5, 44, 0, 0, 169,
		171, 7, 10, 0, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 21, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 20, 31, 41, 43, 48, 52, 56, 59, 72, 100, 105, 107, 116, 121, 124,
		131, 143, 152, 156, 163, 172,
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

// yamlpathParserInit initializes any static state used to implement yamlpathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewyamlpathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YamlpathParserInit() {
	staticData := &YamlpathParserStaticData
	staticData.once.Do(yamlpathParserInit)
}

// NewyamlpathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewyamlpathParser(input antlr.TokenStream) *yamlpathParser {
	YamlpathParserInit()
	this := new(yamlpathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &YamlpathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "yamlpath.g4"

	return this
}

// yamlpathParser tokens.
const (
	yamlpathParserEOF        = antlr.TokenEOF
	yamlpathParserT__0       = 1
	yamlpathParserT__1       = 2
	yamlpathParserT__2       = 3
	yamlpathParserT__3       = 4
	yamlpathParserT__4       = 5
	yamlpathParserT__5       = 6
	yamlpathParserT__6       = 7
	yamlpathParserT__7       = 8
	yamlpathParserT__8       = 9
	yamlpathParserT__9       = 10
	yamlpathParserT__10      = 11
	yamlpathParserT__11      = 12
	yamlpathParserT__12      = 13
	yamlpathParserT__13      = 14
	yamlpathParserT__14      = 15
	yamlpathParserT__15      = 16
	yamlpathParserT__16      = 17
	yamlpathParserT__17      = 18
	yamlpathParserT__18      = 19
	yamlpathParserT__19      = 20
	yamlpathParserT__20      = 21
	yamlpathParserT__21      = 22
	yamlpathParserT__22      = 23
	yamlpathParserT__23      = 24
	yamlpathParserT__24      = 25
	yamlpathParserT__25      = 26
	yamlpathParserT__26      = 27
	yamlpathParserT__27      = 28
	yamlpathParserT__28      = 29
	yamlpathParserT__29      = 30
	yamlpathParserT__30      = 31
	yamlpathParserT__31      = 32
	yamlpathParserT__32      = 33
	yamlpathParserT__33      = 34
	yamlpathParserT__34      = 35
	yamlpathParserT__35      = 36
	yamlpathParserT__36      = 37
	yamlpathParserT__37      = 38
	yamlpathParserT__38      = 39
	yamlpathParserT__39      = 40
	yamlpathParserIDENTIFIER = 41
	yamlpathParserNUMBER     = 42
	yamlpathParserSTRING     = 43
	yamlpathParserREGEX      = 44
	yamlpathParserWS         = 45
	yamlpathParserCOMMENT    = 46
)

// yamlpathParser rules.
const (
	yamlpathParserRULE_path          = 0
	yamlpathParserRULE_expression    = 1
	yamlpathParserRULE_indexParam    = 2
	yamlpathParserRULE_subexpression = 3
	yamlpathParserRULE_literal       = 4
	yamlpathParserRULE_listEntries   = 5
	yamlpathParserRULE_mapEntries    = 6
	yamlpathParserRULE_invocation    = 7
	yamlpathParserRULE_paramList     = 8
	yamlpathParserRULE_identifier    = 9
	yamlpathParserRULE_regex         = 10
)

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) Expression() IExpressionContext {
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

func (s *PathContext) EOF() antlr.TerminalNode {
	return s.GetToken(yamlpathParserEOF, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, yamlpathParserRULE_path)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.expression(0)
	}
	{
		p.SetState(23)
		p.Match(yamlpathParserEOF)
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
	p.RuleIndex = yamlpathParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootExpressionContext struct {
	ExpressionContext
}

func NewRootExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootExpressionContext {
	var p = new(RootExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RootExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

type FieldExpressionContext struct {
	ExpressionContext
}

func NewFieldExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldExpressionContext {
	var p = new(FieldExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FieldExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExpressionContext) Expression() IExpressionContext {
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

func (s *FieldExpressionContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

type IndexExpressionContext struct {
	ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) Expression() IExpressionContext {
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

func (s *IndexExpressionContext) IndexParam() IIndexParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexParamContext)
}

type RecursiveExpressionContext struct {
	ExpressionContext
}

func NewRecursiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecursiveExpressionContext {
	var p = new(RecursiveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RecursiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursiveExpressionContext) Expression() IExpressionContext {
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

func (s *RecursiveExpressionContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (p *yamlpathParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *yamlpathParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, yamlpathParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewRootExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !(_la == yamlpathParserT__0 || _la == yamlpathParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRecursiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(29)
					p.Match(yamlpathParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(31)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(30)
						p.Invocation()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 2:
				localctx = NewFieldExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(34)
					p.Match(yamlpathParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(35)
					p.Invocation()
				}

			case 3:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(37)
					p.Match(yamlpathParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(38)
					p.IndexParam()
				}
				{
					p.SetState(39)
					p.Match(yamlpathParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
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

// IIndexParamContext is an interface to support dynamic dispatch.
type IIndexParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexParamContext differentiates from other interfaces.
	IsIndexParamContext()
}

type IndexParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexParamContext() *IndexParamContext {
	var p = new(IndexParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_indexParam
	return p
}

func InitEmptyIndexParamContext(p *IndexParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_indexParam
}

func (*IndexParamContext) IsIndexParamContext() {}

func NewIndexParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexParamContext {
	var p = new(IndexParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_indexParam

	return p
}

func (s *IndexParamContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexParamContext) CopyAll(ctx *IndexParamContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceIndexContext struct {
	IndexParamContext
}

func NewSliceIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceIndexContext {
	var p = new(SliceIndexContext)

	InitEmptyIndexParamContext(&p.IndexParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexParamContext))

	return p
}

func (s *SliceIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceIndexContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserNUMBER)
}

func (s *SliceIndexContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, i)
}

type ExpressionIndexContext struct {
	IndexParamContext
}

func NewExpressionIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIndexContext {
	var p = new(ExpressionIndexContext)

	InitEmptyIndexParamContext(&p.IndexParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexParamContext))

	return p
}

func (s *ExpressionIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIndexContext) Subexpression() ISubexpressionContext {
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

type WildcardIndexContext struct {
	IndexParamContext
}

func NewWildcardIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardIndexContext {
	var p = new(WildcardIndexContext)

	InitEmptyIndexParamContext(&p.IndexParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexParamContext))

	return p
}

func (s *WildcardIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *yamlpathParser) IndexParam() (localctx IIndexParamContext) {
	localctx = NewIndexParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, yamlpathParserRULE_indexParam)
	var _la int

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewWildcardIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(yamlpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSliceIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(47)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(50)
			p.Match(yamlpathParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(51)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserT__7 {
			{
				p.SetState(54)
				p.Match(yamlpathParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(55)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewExpressionIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
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
	p.RuleIndex = yamlpathParserRULE_subexpression
	return p
}

func InitEmptySubexpressionContext(p *SubexpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_subexpression
}

func (*SubexpressionContext) IsSubexpressionContext() {}

func NewSubexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubexpressionContext {
	var p = new(SubexpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_subexpression

	return p
}

func (s *SubexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubexpressionContext) CopyAll(ctx *SubexpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualitySubexpressionContext struct {
	SubexpressionContext
}

func NewEqualitySubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualitySubexpressionContext {
	var p = new(EqualitySubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *EqualitySubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualitySubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *EqualitySubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type ParenthesisSubexpressionContext struct {
	SubexpressionContext
}

func NewParenthesisSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisSubexpressionContext {
	var p = new(ParenthesisSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *ParenthesisSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisSubexpressionContext) Subexpression() ISubexpressionContext {
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

type AndSubexpressionContext struct {
	SubexpressionContext
}

func NewAndSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndSubexpressionContext {
	var p = new(AndSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *AndSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *AndSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type NegationSubexpressionContext struct {
	SubexpressionContext
}

func NewNegationSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationSubexpressionContext {
	var p = new(NegationSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *NegationSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationSubexpressionContext) Subexpression() ISubexpressionContext {
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

type PolaritySubexpressionContext struct {
	SubexpressionContext
}

func NewPolaritySubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PolaritySubexpressionContext {
	var p = new(PolaritySubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *PolaritySubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolaritySubexpressionContext) Expression() IExpressionContext {
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

type UnionSubexpressionContext struct {
	SubexpressionContext
}

func NewUnionSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionSubexpressionContext {
	var p = new(UnionSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *UnionSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *UnionSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type MatchSubexpressionContext struct {
	SubexpressionContext
}

func NewMatchSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchSubexpressionContext {
	var p = new(MatchSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *MatchSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchSubexpressionContext) Subexpression() ISubexpressionContext {
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

func (s *MatchSubexpressionContext) Regex() IRegexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexContext)
}

type RootSubexpressionContext struct {
	SubexpressionContext
}

func NewRootSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootSubexpressionContext {
	var p = new(RootSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *RootSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootSubexpressionContext) Expression() IExpressionContext {
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

type LiteralSubexpressionContext struct {
	SubexpressionContext
}

func NewLiteralSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralSubexpressionContext {
	var p = new(LiteralSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *LiteralSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSubexpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

type OrSubexpressionContext struct {
	SubexpressionContext
}

func NewOrSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrSubexpressionContext {
	var p = new(OrSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *OrSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *OrSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type MembershipSubexpressionContext struct {
	SubexpressionContext
}

func NewMembershipSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembershipSubexpressionContext {
	var p = new(MembershipSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *MembershipSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *MembershipSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type InequalitySubexpressionContext struct {
	SubexpressionContext
}

func NewInequalitySubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InequalitySubexpressionContext {
	var p = new(InequalitySubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *InequalitySubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InequalitySubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *InequalitySubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type AdditiveSubexpressionContext struct {
	SubexpressionContext
}

func NewAdditiveSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveSubexpressionContext {
	var p = new(AdditiveSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *AdditiveSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *AdditiveSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

type MultiplicativeSubexpressionContext struct {
	SubexpressionContext
}

func NewMultiplicativeSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeSubexpressionContext {
	var p = new(MultiplicativeSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *MultiplicativeSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeSubexpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *MultiplicativeSubexpressionContext) Subexpression(i int) ISubexpressionContext {
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

func (p *yamlpathParser) Subexpression() (localctx ISubexpressionContext) {
	return p.subexpression(0)
}

func (p *yamlpathParser) subexpression(_p int) (localctx ISubexpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSubexpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISubexpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, yamlpathParserRULE_subexpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserT__0, yamlpathParserT__1:
		localctx = NewRootSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(62)
			p.expression(0)
		}

	case yamlpathParserT__4, yamlpathParserT__31, yamlpathParserT__32, yamlpathParserT__33, yamlpathParserT__34, yamlpathParserNUMBER, yamlpathParserSTRING:
		localctx = NewLiteralSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Literal()
		}

	case yamlpathParserT__8:
		localctx = NewParenthesisSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(yamlpathParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.subexpression(0)
		}
		{
			p.SetState(66)
			p.Match(yamlpathParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__10, yamlpathParserT__11:
		localctx = NewNegationSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__10 || _la == yamlpathParserT__11) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(69)
			p.subexpression(11)
		}

	case yamlpathParserT__12, yamlpathParserT__13:
		localctx = NewPolaritySubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__12 || _la == yamlpathParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(71)
			p.expression(0)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(75)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&98432) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(76)
					p.subexpression(10)
				}

			case 2:
				localctx = NewAdditiveSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__12 || _la == yamlpathParserT__13) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)
					p.subexpression(9)
				}

			case 3:
				localctx = NewInequalitySubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3932160) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)
					p.subexpression(7)
				}

			case 4:
				localctx = NewEqualitySubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__21 || _la == yamlpathParserT__22) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(85)
					p.subexpression(6)
				}

			case 5:
				localctx = NewMembershipSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(87)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&234881024) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(88)
					p.subexpression(4)
				}

			case 6:
				localctx = NewAndSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(90)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__27 || _la == yamlpathParserT__28) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(91)
					p.subexpression(3)
				}

			case 7:
				localctx = NewOrSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__29 || _la == yamlpathParserT__30) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(94)
					p.subexpression(2)
				}

			case 8:
				localctx = NewUnionSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				p.SetState(98)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(96)
							p.Match(yamlpathParserT__16)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(97)
							p.subexpression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(100)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case 9:
				localctx = NewMatchSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(103)
					p.Match(yamlpathParserT__23)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(104)
					p.Regex()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullLiteralContext struct {
	LiteralContext
}

func NewNullLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralContext {
	var p = new(NullLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, 0)
}

type ListLiteralContext struct {
	LiteralContext
}

func NewListLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListLiteralContext {
	var p = new(ListLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ListEntries() IListEntriesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListEntriesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListEntriesContext)
}

type MapLiteralContext struct {
	LiteralContext
}

func NewMapLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapLiteralContext {
	var p = new(MapLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) MapEntries() IMapEntriesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapEntriesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapEntriesContext)
}

type BooleanLiteralContext struct {
	LiteralContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

type NumberLiteralContext struct {
	LiteralContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, 0)
}

func (p *yamlpathParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, yamlpathParserRULE_literal)
	var _la int

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserSTRING:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserNUMBER:
		localctx = NewNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__31, yamlpathParserT__32:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__31 || _la == yamlpathParserT__32) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case yamlpathParserT__33:
		localctx = NewNullLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(yamlpathParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__4:
		localctx = NewListLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Match(yamlpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13258564042784) != 0 {
			{
				p.SetState(115)
				p.ListEntries()
			}

		}
		{
			p.SetState(118)
			p.Match(yamlpathParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__34:
		localctx = NewMapLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.Match(yamlpathParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserSTRING {
			{
				p.SetState(120)
				p.MapEntries()
			}

		}
		{
			p.SetState(123)
			p.Match(yamlpathParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IListEntriesContext is an interface to support dynamic dispatch.
type IListEntriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

	// IsListEntriesContext differentiates from other interfaces.
	IsListEntriesContext()
}

type ListEntriesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListEntriesContext() *ListEntriesContext {
	var p = new(ListEntriesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_listEntries
	return p
}

func InitEmptyListEntriesContext(p *ListEntriesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_listEntries
}

func (*ListEntriesContext) IsListEntriesContext() {}

func NewListEntriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListEntriesContext {
	var p = new(ListEntriesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_listEntries

	return p
}

func (s *ListEntriesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListEntriesContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *ListEntriesContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
}

func (s *ListEntriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListEntriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) ListEntries() (localctx IListEntriesContext) {
	localctx = NewListEntriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, yamlpathParserRULE_listEntries)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Literal()
	}

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__36 {
		{
			p.SetState(127)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Literal()
		}

		p.SetState(133)
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

// IMapEntriesContext is an interface to support dynamic dispatch.
type IMapEntriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

	// IsMapEntriesContext differentiates from other interfaces.
	IsMapEntriesContext()
}

type MapEntriesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntriesContext() *MapEntriesContext {
	var p = new(MapEntriesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_mapEntries
	return p
}

func InitEmptyMapEntriesContext(p *MapEntriesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_mapEntries
}

func (*MapEntriesContext) IsMapEntriesContext() {}

func NewMapEntriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntriesContext {
	var p = new(MapEntriesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_mapEntries

	return p
}

func (s *MapEntriesContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntriesContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserSTRING)
}

func (s *MapEntriesContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, i)
}

func (s *MapEntriesContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *MapEntriesContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
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

	return t.(ILiteralContext)
}

func (s *MapEntriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) MapEntries() (localctx IMapEntriesContext) {
	localctx = NewMapEntriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, yamlpathParserRULE_mapEntries)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(yamlpathParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(yamlpathParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Literal()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__36 {
		{
			p.SetState(137)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(yamlpathParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Literal()
		}

		p.SetState(145)
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

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_invocation
	return p
}

func InitEmptyInvocationContext(p *InvocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_invocation
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) CopyAll(ctx *InvocationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedMemberInvocationContext struct {
	InvocationContext
}

func NewQuotedMemberInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedMemberInvocationContext {
	var p = new(QuotedMemberInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *QuotedMemberInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedMemberInvocationContext) STRING() antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, 0)
}

type WildcardInvocationContext struct {
	InvocationContext
}

func NewWildcardInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardInvocationContext {
	var p = new(WildcardInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *WildcardInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

type FunctionInvocationContext struct {
	InvocationContext
}

func NewFunctionInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionInvocationContext {
	var p = new(FunctionInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *FunctionInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvocationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionInvocationContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

type MemberInvocationContext struct {
	InvocationContext
}

func NewMemberInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberInvocationContext {
	var p = new(MemberInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *MemberInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberInvocationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (p *yamlpathParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, yamlpathParserRULE_invocation)
	var _la int

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Identifier()
		}

	case 2:
		localctx = NewQuotedMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewWildcardInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.Match(yamlpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFunctionInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Identifier()
		}
		{
			p.SetState(150)
			p.Match(yamlpathParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13258564074022) != 0 {
			{
				p.SetState(151)
				p.ParamList()
			}

		}
		{
			p.SetState(154)
			p.Match(yamlpathParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSubexpression() []ISubexpressionContext
	Subexpression(i int) ISubexpressionContext

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllSubexpression() []ISubexpressionContext {
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

func (s *ParamListContext) Subexpression(i int) ISubexpressionContext {
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

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, yamlpathParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.subexpression(0)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__36 {
		{
			p.SetState(159)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.subexpression(0)
		}

		p.SetState(165)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(yamlpathParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, yamlpathParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(yamlpathParserIDENTIFIER)
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

// IRegexContext is an interface to support dynamic dispatch.
type IRegexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGEX() antlr.TerminalNode

	// IsRegexContext differentiates from other interfaces.
	IsRegexContext()
}

type RegexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexContext() *RegexContext {
	var p = new(RegexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_regex
	return p
}

func InitEmptyRegexContext(p *RegexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_regex
}

func (*RegexContext) IsRegexContext() {}

func NewRegexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexContext {
	var p = new(RegexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_regex

	return p
}

func (s *RegexContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexContext) REGEX() antlr.TerminalNode {
	return s.GetToken(yamlpathParserREGEX, 0)
}

func (s *RegexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) Regex() (localctx IRegexContext) {
	localctx = NewRegexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, yamlpathParserRULE_regex)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(yamlpathParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(169)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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

func (p *yamlpathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 3:
		var t *SubexpressionContext = nil
		if localctx != nil {
			t = localctx.(*SubexpressionContext)
		}
		return p.Subexpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *yamlpathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *yamlpathParser) Subexpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
