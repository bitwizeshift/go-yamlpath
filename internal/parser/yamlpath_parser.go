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
		"", "'..'", "'.'", "'['", "']'", "'('", "')'", "'!'", "'not'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'|'", "'<='", "'<'", "'>'", "'>='", "'=='",
		"'!='", "'=~'", "'in'", "'nin'", "'subsetof'", "'&&'", "'and'", "'||'",
		"'or'", "'$'", "'@'", "':'", "'true'", "'false'", "'null'", "'{'", "'}'",
		"','", "'i'", "'m'", "'s'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "IDENTIFIER", "NUMBER", "STRING", "REGEX",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"path", "expression", "term", "externalConstant", "indexParam", "literal",
		"listEntries", "mapEntries", "invocation", "paramList", "identifier",
		"regex",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 179, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 64, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 77, 8, 1, 11, 1,
		12, 1, 78, 1, 1, 1, 1, 1, 1, 5, 1, 84, 8, 1, 10, 1, 12, 1, 87, 9, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 93, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4,
		100, 8, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 4, 1, 4, 3, 4, 108, 8, 4, 1,
		4, 3, 4, 111, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 119, 8, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 124, 8, 5, 1, 5, 3, 5, 127, 8, 5, 1, 6, 1, 6, 1,
		6, 5, 6, 132, 8, 6, 10, 6, 12, 6, 135, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 144, 8, 7, 10, 7, 12, 7, 147, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 154, 8, 8, 1, 8, 1, 8, 3, 8, 158, 8, 8, 1, 9, 1, 9,
		1, 9, 5, 9, 163, 8, 9, 10, 9, 12, 9, 166, 9, 9, 1, 10, 1, 10, 3, 10, 170,
		8, 10, 1, 11, 1, 11, 5, 11, 174, 8, 11, 10, 11, 12, 11, 177, 9, 11, 1,
		11, 1, 175, 1, 2, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 11,
		1, 0, 7, 8, 1, 0, 9, 10, 1, 0, 11, 13, 1, 0, 15, 18, 1, 0, 19, 20, 1, 0,
		22, 24, 1, 0, 25, 26, 1, 0, 27, 28, 1, 0, 29, 30, 1, 0, 32, 33, 1, 0, 38,
		40, 206, 0, 24, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 92, 1, 0, 0, 0, 6, 94,
		1, 0, 0, 0, 8, 110, 1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 128, 1, 0, 0,
		0, 14, 136, 1, 0, 0, 0, 16, 157, 1, 0, 0, 0, 18, 159, 1, 0, 0, 0, 20, 169,
		1, 0, 0, 0, 22, 171, 1, 0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 0, 0, 1,
		26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0, 28, 38, 3, 4, 2, 0, 29, 30, 5,
		5, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 6, 0, 0, 32, 38, 1, 0, 0, 0, 33,
		34, 7, 0, 0, 0, 34, 38, 3, 2, 1, 11, 35, 36, 7, 1, 0, 0, 36, 38, 3, 2,
		1, 10, 37, 27, 1, 0, 0, 0, 37, 29, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37,
		35, 1, 0, 0, 0, 38, 85, 1, 0, 0, 0, 39, 40, 10, 9, 0, 0, 40, 41, 7, 2,
		0, 0, 41, 84, 3, 2, 1, 10, 42, 43, 10, 8, 0, 0, 43, 44, 7, 1, 0, 0, 44,
		84, 3, 2, 1, 9, 45, 46, 10, 6, 0, 0, 46, 47, 7, 3, 0, 0, 47, 84, 3, 2,
		1, 7, 48, 49, 10, 5, 0, 0, 49, 50, 7, 4, 0, 0, 50, 84, 3, 2, 1, 6, 51,
		52, 10, 3, 0, 0, 52, 53, 7, 5, 0, 0, 53, 84, 3, 2, 1, 4, 54, 55, 10, 2,
		0, 0, 55, 56, 7, 6, 0, 0, 56, 84, 3, 2, 1, 3, 57, 58, 10, 1, 0, 0, 58,
		59, 7, 7, 0, 0, 59, 84, 3, 2, 1, 2, 60, 61, 10, 15, 0, 0, 61, 63, 5, 1,
		0, 0, 62, 64, 3, 16, 8, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64,
		84, 1, 0, 0, 0, 65, 66, 10, 14, 0, 0, 66, 67, 5, 2, 0, 0, 67, 84, 3, 16,
		8, 0, 68, 69, 10, 13, 0, 0, 69, 70, 5, 3, 0, 0, 70, 71, 3, 8, 4, 0, 71,
		72, 5, 4, 0, 0, 72, 84, 1, 0, 0, 0, 73, 76, 10, 7, 0, 0, 74, 75, 5, 14,
		0, 0, 75, 77, 3, 2, 1, 0, 76, 74, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 84, 1, 0, 0, 0, 80, 81, 10, 4, 0, 0,
		81, 82, 5, 21, 0, 0, 82, 84, 3, 22, 11, 0, 83, 39, 1, 0, 0, 0, 83, 42,
		1, 0, 0, 0, 83, 45, 1, 0, 0, 0, 83, 48, 1, 0, 0, 0, 83, 51, 1, 0, 0, 0,
		83, 54, 1, 0, 0, 0, 83, 57, 1, 0, 0, 0, 83, 60, 1, 0, 0, 0, 83, 65, 1,
		0, 0, 0, 83, 68, 1, 0, 0, 0, 83, 73, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 84,
		87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 3, 1, 0, 0,
		0, 87, 85, 1, 0, 0, 0, 88, 93, 7, 8, 0, 0, 89, 93, 3, 6, 3, 0, 90, 93,
		3, 10, 5, 0, 91, 93, 3, 16, 8, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 5, 1, 0, 0, 0, 94, 95, 5,
		13, 0, 0, 95, 96, 3, 20, 10, 0, 96, 7, 1, 0, 0, 0, 97, 111, 5, 11, 0, 0,
		98, 100, 5, 42, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 103, 5, 31, 0, 0, 102, 104, 5, 42, 0, 0, 103, 102, 1,
		0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 106, 5, 31, 0,
		0, 106, 108, 5, 42, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		111, 1, 0, 0, 0, 109, 111, 3, 2, 1, 0, 110, 97, 1, 0, 0, 0, 110, 99, 1,
		0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 9, 1, 0, 0, 0, 112, 127, 5, 43, 0,
		0, 113, 127, 5, 42, 0, 0, 114, 127, 7, 9, 0, 0, 115, 127, 5, 34, 0, 0,
		116, 118, 5, 3, 0, 0, 117, 119, 3, 12, 6, 0, 118, 117, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 127, 5, 4, 0, 0, 121, 123,
		5, 35, 0, 0, 122, 124, 3, 14, 7, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1,
		0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 5, 36, 0, 0, 126, 112, 1, 0, 0,
		0, 126, 113, 1, 0, 0, 0, 126, 114, 1, 0, 0, 0, 126, 115, 1, 0, 0, 0, 126,
		116, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 127, 11, 1, 0, 0, 0, 128, 133, 3,
		10, 5, 0, 129, 130, 5, 37, 0, 0, 130, 132, 3, 10, 5, 0, 131, 129, 1, 0,
		0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 13, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 137, 5, 43, 0, 0, 137,
		138, 5, 31, 0, 0, 138, 145, 3, 10, 5, 0, 139, 140, 5, 37, 0, 0, 140, 141,
		5, 43, 0, 0, 141, 142, 5, 31, 0, 0, 142, 144, 3, 10, 5, 0, 143, 139, 1,
		0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0,
		0, 146, 15, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 158, 3, 20, 10, 0, 149,
		158, 5, 11, 0, 0, 150, 151, 3, 20, 10, 0, 151, 153, 5, 5, 0, 0, 152, 154,
		3, 18, 9, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 1, 0,
		0, 0, 155, 156, 5, 6, 0, 0, 156, 158, 1, 0, 0, 0, 157, 148, 1, 0, 0, 0,
		157, 149, 1, 0, 0, 0, 157, 150, 1, 0, 0, 0, 158, 17, 1, 0, 0, 0, 159, 164,
		3, 2, 1, 0, 160, 161, 5, 37, 0, 0, 161, 163, 3, 2, 1, 0, 162, 160, 1, 0,
		0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 19, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 170, 5, 41, 0, 0, 168,
		170, 5, 43, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 21,
		1, 0, 0, 0, 171, 175, 5, 44, 0, 0, 172, 174, 7, 10, 0, 0, 173, 172, 1,
		0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0,
		0, 176, 23, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 20, 37, 63, 78, 83, 85, 92,
		99, 103, 107, 110, 118, 123, 126, 133, 145, 153, 157, 164, 169, 175,
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
	yamlpathParserRULE_path             = 0
	yamlpathParserRULE_expression       = 1
	yamlpathParserRULE_term             = 2
	yamlpathParserRULE_externalConstant = 3
	yamlpathParserRULE_indexParam       = 4
	yamlpathParserRULE_literal          = 5
	yamlpathParserRULE_listEntries      = 6
	yamlpathParserRULE_mapEntries       = 7
	yamlpathParserRULE_invocation       = 8
	yamlpathParserRULE_paramList        = 9
	yamlpathParserRULE_identifier       = 10
	yamlpathParserRULE_regex            = 11
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
		p.SetState(24)
		p.expression(0)
	}
	{
		p.SetState(25)
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

type NegationExpressionContext struct {
	ExpressionContext
}

func NewNegationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationExpressionContext {
	var p = new(NegationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationExpressionContext) Expression() IExpressionContext {
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

type MatchExpressionContext struct {
	ExpressionContext
}

func NewMatchExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExpressionContext {
	var p = new(MatchExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MatchExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExpressionContext) Expression() IExpressionContext {
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

func (s *MatchExpressionContext) Regex() IRegexContext {
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

type PolarityExpressionContext struct {
	ExpressionContext
}

func NewPolarityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PolarityExpressionContext {
	var p = new(PolarityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PolarityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolarityExpressionContext) Expression() IExpressionContext {
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

type AdditiveExpressionContext struct {
	ExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type ParenthesisExpressionContext struct {
	ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
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

type MultiplicativeExpressionContext struct {
	ExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
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

type UnionExpressionContext struct {
	ExpressionContext
}

func NewUnionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionExpressionContext {
	var p = new(UnionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *UnionExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type OrExpressionContext struct {
	ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type AndExpressionContext struct {
	ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type MembershipExpressionContext struct {
	ExpressionContext
}

func NewMembershipExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembershipExpressionContext {
	var p = new(MembershipExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MembershipExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MembershipExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type InequalityExpressionContext struct {
	ExpressionContext
}

func NewInequalityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InequalityExpressionContext {
	var p = new(InequalityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InequalityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InequalityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InequalityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

type EqualityExpressionContext struct {
	ExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
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

type TermExpressionContext struct {
	ExpressionContext
}

func NewTermExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermExpressionContext {
	var p = new(TermExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TermExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExpressionContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserT__2, yamlpathParserT__10, yamlpathParserT__12, yamlpathParserT__28, yamlpathParserT__29, yamlpathParserT__31, yamlpathParserT__32, yamlpathParserT__33, yamlpathParserT__34, yamlpathParserIDENTIFIER, yamlpathParserNUMBER, yamlpathParserSTRING:
		localctx = NewTermExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(28)
			p.Term()
		}

	case yamlpathParserT__4:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(yamlpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.expression(0)
		}
		{
			p.SetState(31)
			p.Match(yamlpathParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__6, yamlpathParserT__7:
		localctx = NewNegationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(33)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__6 || _la == yamlpathParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(34)
			p.expression(11)
		}

	case yamlpathParserT__8, yamlpathParserT__9:
		localctx = NewPolarityExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__8 || _la == yamlpathParserT__9) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(36)
			p.expression(10)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(40)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)
					p.expression(10)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(43)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__8 || _la == yamlpathParserT__9) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)
					p.expression(9)
				}

			case 3:
				localctx = NewInequalityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(46)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&491520) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(47)
					p.expression(7)
				}

			case 4:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(49)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__18 || _la == yamlpathParserT__19) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(50)
					p.expression(6)
				}

			case 5:
				localctx = NewMembershipExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(52)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)
					p.expression(4)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(55)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__24 || _la == yamlpathParserT__25) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(56)
					p.expression(3)
				}

			case 7:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(58)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__26 || _la == yamlpathParserT__27) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(59)
					p.expression(2)
				}

			case 8:
				localctx = NewRecursiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(61)
					p.Match(yamlpathParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(63)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(62)
						p.Invocation()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 9:
				localctx = NewFieldExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(66)
					p.Match(yamlpathParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(67)
					p.Invocation()
				}

			case 10:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(69)
					p.Match(yamlpathParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.IndexParam()
				}
				{
					p.SetState(71)
					p.Match(yamlpathParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 11:
				localctx = NewUnionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(74)
							p.Match(yamlpathParserT__13)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(75)
							p.expression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(78)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case 12:
				localctx = NewMatchExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					p.Match(yamlpathParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)
					p.Regex()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyAll(ctx *TermContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExternalConstantTermContext struct {
	TermContext
}

func NewExternalConstantTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExternalConstantTermContext {
	var p = new(ExternalConstantTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *ExternalConstantTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalConstantTermContext) ExternalConstant() IExternalConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExternalConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExternalConstantContext)
}

type LiteralTermContext struct {
	TermContext
}

func NewLiteralTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralTermContext {
	var p = new(LiteralTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *LiteralTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTermContext) Literal() ILiteralContext {
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

type RootTermContext struct {
	TermContext
}

func NewRootTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootTermContext {
	var p = new(RootTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *RootTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

type InvocationTermContext struct {
	TermContext
}

func NewInvocationTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationTermContext {
	var p = new(InvocationTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *InvocationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationTermContext) Invocation() IInvocationContext {
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

func (p *yamlpathParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, yamlpathParserRULE_term)
	var _la int

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRootTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__28 || _la == yamlpathParserT__29) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewExternalConstantTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.ExternalConstant()
		}

	case 3:
		localctx = NewLiteralTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Literal()
		}

	case 4:
		localctx = NewInvocationTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Invocation()
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

// IExternalConstantContext is an interface to support dynamic dispatch.
type IExternalConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsExternalConstantContext differentiates from other interfaces.
	IsExternalConstantContext()
}

type ExternalConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternalConstantContext() *ExternalConstantContext {
	var p = new(ExternalConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_externalConstant
	return p
}

func InitEmptyExternalConstantContext(p *ExternalConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_externalConstant
}

func (*ExternalConstantContext) IsExternalConstantContext() {}

func NewExternalConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternalConstantContext {
	var p = new(ExternalConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_externalConstant

	return p
}

func (s *ExternalConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternalConstantContext) Identifier() IIdentifierContext {
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

func (s *ExternalConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) ExternalConstant() (localctx IExternalConstantContext) {
	localctx = NewExternalConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, yamlpathParserRULE_externalConstant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(yamlpathParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Identifier()
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

func (s *ExpressionIndexContext) Expression() IExpressionContext {
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
	p.EnterRule(localctx, 8, yamlpathParserRULE_indexParam)
	var _la int

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewWildcardIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSliceIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(98)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(101)
			p.Match(yamlpathParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(102)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserT__30 {
			{
				p.SetState(105)
				p.Match(yamlpathParserT__30)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(106)
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
			p.SetState(109)
			p.expression(0)
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
	p.EnterRule(localctx, 10, yamlpathParserRULE_literal)
	var _la int

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserSTRING:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
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
			p.SetState(113)
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
			p.SetState(114)
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
			p.SetState(115)
			p.Match(yamlpathParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__2:
		localctx = NewListLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)
			p.Match(yamlpathParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13258564042760) != 0 {
			{
				p.SetState(117)
				p.ListEntries()
			}

		}
		{
			p.SetState(120)
			p.Match(yamlpathParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__34:
		localctx = NewMapLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
			p.Match(yamlpathParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserSTRING {
			{
				p.SetState(122)
				p.MapEntries()
			}

		}
		{
			p.SetState(125)
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
	p.EnterRule(localctx, 12, yamlpathParserRULE_listEntries)
	var _la int

	p.EnterOuterAlt(localctx, 1)
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

	for _la == yamlpathParserT__36 {
		{
			p.SetState(129)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Literal()
		}

		p.SetState(135)
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
	p.EnterRule(localctx, 14, yamlpathParserRULE_mapEntries)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(yamlpathParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(yamlpathParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Literal()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__36 {
		{
			p.SetState(139)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(yamlpathParserT__30)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Literal()
		}

		p.SetState(147)
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
	p.EnterRule(localctx, 16, yamlpathParserRULE_invocation)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Identifier()
		}

	case 2:
		localctx = NewWildcardInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunctionInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Identifier()
		}
		{
			p.SetState(151)
			p.Match(yamlpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15459197923240) != 0 {
			{
				p.SetState(152)
				p.ParamList()
			}

		}
		{
			p.SetState(155)
			p.Match(yamlpathParserT__5)
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
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

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

func (s *ParamListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *yamlpathParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, yamlpathParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.expression(0)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__36 {
		{
			p.SetState(160)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.expression(0)
		}

		p.SetState(166)
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

func (s *IdentifierContext) CopyAll(ctx *IdentifierContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlainIdentifierContext struct {
	IdentifierContext
}

func NewPlainIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainIdentifierContext {
	var p = new(PlainIdentifierContext)

	InitEmptyIdentifierContext(&p.IdentifierContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifierContext))

	return p
}

func (s *PlainIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(yamlpathParserIDENTIFIER, 0)
}

type QuotedIdentifierContext struct {
	IdentifierContext
}

func NewQuotedIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedIdentifierContext {
	var p = new(QuotedIdentifierContext)

	InitEmptyIdentifierContext(&p.IdentifierContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifierContext))

	return p
}

func (s *QuotedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedIdentifierContext) STRING() antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, 0)
}

func (p *yamlpathParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, yamlpathParserRULE_identifier)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserIDENTIFIER:
		localctx = NewPlainIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(yamlpathParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserSTRING:
		localctx = NewQuotedIdentifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Match(yamlpathParserSTRING)
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
	p.EnterRule(localctx, 22, yamlpathParserRULE_regex)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(yamlpathParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(175)
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
				p.SetState(172)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(177)
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

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *yamlpathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
