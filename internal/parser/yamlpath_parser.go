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
		"", "'$'", "'@'", "'..'", "'.'", "'['", "']'", "','", "'*'", "':'",
		"'?'", "'('", "')'", "'!'", "'not'", "'+'", "'-'", "'/'", "'%'", "'<='",
		"'<'", "'>'", "'>='", "'=='", "'!='", "'=~'", "'in'", "'nin'", "'subsetof'",
		"'&&'", "'and'", "'||'", "'or'", "'true'", "'false'", "'null'", "'{'",
		"'}'", "'i'", "'m'", "'s'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "IDENTIFIER", "NUMBER", "STRING", "REGEX",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"path", "expression", "bracketParam", "subexpression", "literal", "aggregation",
		"listEntries", "mapEntries", "invocation", "paramList", "identifier",
		"regex",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 198, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 34, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 44,
		8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 2, 5, 2, 52, 8, 2, 10, 2,
		12, 2, 55, 9, 2, 1, 2, 1, 2, 1, 2, 5, 2, 60, 8, 2, 10, 2, 12, 2, 63, 9,
		2, 1, 2, 1, 2, 3, 2, 67, 8, 2, 1, 2, 1, 2, 3, 2, 71, 8, 2, 1, 2, 1, 2,
		3, 2, 75, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 86, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 100, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 126, 8, 3, 10, 3, 12, 3, 129, 9, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 135, 8, 4, 1, 5, 1, 5, 3, 5, 139, 8, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 144, 8, 5, 1, 5, 1, 5, 3, 5, 148, 8, 5, 1, 6, 1, 6,
		1, 6, 5, 6, 153, 8, 6, 10, 6, 12, 6, 156, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 5, 7, 165, 8, 7, 10, 7, 12, 7, 168, 9, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 175, 8, 8, 1, 8, 1, 8, 3, 8, 179, 8, 8, 1, 9, 1,
		9, 1, 9, 5, 9, 184, 8, 9, 10, 9, 12, 9, 187, 9, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 5, 11, 193, 8, 11, 10, 11, 12, 11, 196, 9, 11, 1, 11, 1, 194, 2,
		2, 6, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 11, 1, 0, 1, 2,
		1, 0, 13, 14, 1, 0, 15, 16, 2, 0, 8, 8, 17, 18, 1, 0, 19, 22, 1, 0, 23,
		24, 1, 0, 26, 28, 1, 0, 29, 30, 1, 0, 31, 32, 1, 0, 33, 34, 1, 0, 38, 40,
		226, 0, 24, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 99, 1,
		0, 0, 0, 8, 134, 1, 0, 0, 0, 10, 147, 1, 0, 0, 0, 12, 149, 1, 0, 0, 0,
		14, 157, 1, 0, 0, 0, 16, 178, 1, 0, 0, 0, 18, 180, 1, 0, 0, 0, 20, 188,
		1, 0, 0, 0, 22, 190, 1, 0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 0, 0, 1,
		26, 1, 1, 0, 0, 0, 27, 28, 6, 1, -1, 0, 28, 29, 7, 0, 0, 0, 29, 45, 1,
		0, 0, 0, 30, 31, 10, 3, 0, 0, 31, 33, 5, 3, 0, 0, 32, 34, 3, 16, 8, 0,
		33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 44, 1, 0, 0, 0, 35, 36, 10,
		2, 0, 0, 36, 37, 5, 4, 0, 0, 37, 44, 3, 16, 8, 0, 38, 39, 10, 1, 0, 0,
		39, 40, 5, 5, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 6, 0, 0, 42, 44, 1,
		0, 0, 0, 43, 30, 1, 0, 0, 0, 43, 35, 1, 0, 0, 0, 43, 38, 1, 0, 0, 0, 44,
		47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0,
		0, 47, 45, 1, 0, 0, 0, 48, 53, 5, 43, 0, 0, 49, 50, 5, 7, 0, 0, 50, 52,
		5, 43, 0, 0, 51, 49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0,
		53, 54, 1, 0, 0, 0, 54, 86, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 61, 5,
		42, 0, 0, 57, 58, 5, 7, 0, 0, 58, 60, 5, 42, 0, 0, 59, 57, 1, 0, 0, 0,
		60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 86, 1,
		0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 86, 5, 8, 0, 0, 65, 67, 5, 42, 0, 0, 66,
		65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 5, 9, 0,
		0, 69, 71, 5, 42, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 74,
		1, 0, 0, 0, 72, 73, 5, 9, 0, 0, 73, 75, 5, 42, 0, 0, 74, 72, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 86, 1, 0, 0, 0, 76, 77, 5, 10, 0, 0, 77, 78, 5,
		11, 0, 0, 78, 79, 3, 6, 3, 0, 79, 80, 5, 12, 0, 0, 80, 86, 1, 0, 0, 0,
		81, 82, 5, 11, 0, 0, 82, 83, 3, 6, 3, 0, 83, 84, 5, 12, 0, 0, 84, 86, 1,
		0, 0, 0, 85, 48, 1, 0, 0, 0, 85, 56, 1, 0, 0, 0, 85, 64, 1, 0, 0, 0, 85,
		66, 1, 0, 0, 0, 85, 76, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 86, 5, 1, 0, 0,
		0, 87, 88, 6, 3, -1, 0, 88, 100, 3, 2, 1, 0, 89, 100, 3, 8, 4, 0, 90, 100,
		3, 10, 5, 0, 91, 92, 5, 11, 0, 0, 92, 93, 3, 6, 3, 0, 93, 94, 5, 12, 0,
		0, 94, 100, 1, 0, 0, 0, 95, 96, 7, 1, 0, 0, 96, 100, 3, 6, 3, 10, 97, 98,
		7, 2, 0, 0, 98, 100, 3, 2, 1, 0, 99, 87, 1, 0, 0, 0, 99, 89, 1, 0, 0, 0,
		99, 90, 1, 0, 0, 0, 99, 91, 1, 0, 0, 0, 99, 95, 1, 0, 0, 0, 99, 97, 1,
		0, 0, 0, 100, 127, 1, 0, 0, 0, 101, 102, 10, 8, 0, 0, 102, 103, 7, 3, 0,
		0, 103, 126, 3, 6, 3, 9, 104, 105, 10, 7, 0, 0, 105, 106, 7, 2, 0, 0, 106,
		126, 3, 6, 3, 8, 107, 108, 10, 6, 0, 0, 108, 109, 7, 4, 0, 0, 109, 126,
		3, 6, 3, 7, 110, 111, 10, 5, 0, 0, 111, 112, 7, 5, 0, 0, 112, 126, 3, 6,
		3, 6, 113, 114, 10, 3, 0, 0, 114, 115, 7, 6, 0, 0, 115, 126, 3, 6, 3, 4,
		116, 117, 10, 2, 0, 0, 117, 118, 7, 7, 0, 0, 118, 126, 3, 6, 3, 3, 119,
		120, 10, 1, 0, 0, 120, 121, 7, 8, 0, 0, 121, 126, 3, 6, 3, 2, 122, 123,
		10, 4, 0, 0, 123, 124, 5, 25, 0, 0, 124, 126, 3, 22, 11, 0, 125, 101, 1,
		0, 0, 0, 125, 104, 1, 0, 0, 0, 125, 107, 1, 0, 0, 0, 125, 110, 1, 0, 0,
		0, 125, 113, 1, 0, 0, 0, 125, 116, 1, 0, 0, 0, 125, 119, 1, 0, 0, 0, 125,
		122, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 7, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 135, 5, 43,
		0, 0, 131, 135, 5, 42, 0, 0, 132, 135, 7, 9, 0, 0, 133, 135, 5, 35, 0,
		0, 134, 130, 1, 0, 0, 0, 134, 131, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134,
		133, 1, 0, 0, 0, 135, 9, 1, 0, 0, 0, 136, 138, 5, 5, 0, 0, 137, 139, 3,
		12, 6, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0,
		0, 140, 148, 5, 6, 0, 0, 141, 143, 5, 36, 0, 0, 142, 144, 3, 14, 7, 0,
		143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		148, 5, 37, 0, 0, 146, 148, 3, 8, 4, 0, 147, 136, 1, 0, 0, 0, 147, 141,
		1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 11, 1, 0, 0, 0, 149, 154, 3, 8,
		4, 0, 150, 151, 5, 7, 0, 0, 151, 153, 3, 8, 4, 0, 152, 150, 1, 0, 0, 0,
		153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		13, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 5, 43, 0, 0, 158, 159,
		5, 9, 0, 0, 159, 166, 3, 10, 5, 0, 160, 161, 5, 7, 0, 0, 161, 162, 5, 43,
		0, 0, 162, 163, 5, 9, 0, 0, 163, 165, 3, 10, 5, 0, 164, 160, 1, 0, 0, 0,
		165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		15, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 179, 3, 20, 10, 0, 170, 179,
		5, 8, 0, 0, 171, 172, 3, 20, 10, 0, 172, 174, 5, 11, 0, 0, 173, 175, 3,
		18, 9, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 177, 5, 12, 0, 0, 177, 179, 1, 0, 0, 0, 178, 169, 1, 0, 0, 0, 178,
		170, 1, 0, 0, 0, 178, 171, 1, 0, 0, 0, 179, 17, 1, 0, 0, 0, 180, 185, 3,
		6, 3, 0, 181, 182, 5, 7, 0, 0, 182, 184, 3, 6, 3, 0, 183, 181, 1, 0, 0,
		0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		19, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 41, 0, 0, 189, 21, 1,
		0, 0, 0, 190, 194, 5, 44, 0, 0, 191, 193, 7, 10, 0, 0, 192, 191, 1, 0,
		0, 0, 193, 196, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0,
		195, 23, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 22, 33, 43, 45, 53, 61, 66,
		70, 74, 85, 99, 125, 127, 134, 138, 143, 147, 154, 166, 174, 178, 185,
		194,
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
	yamlpathParserRULE_bracketParam  = 2
	yamlpathParserRULE_subexpression = 3
	yamlpathParserRULE_literal       = 4
	yamlpathParserRULE_aggregation   = 5
	yamlpathParserRULE_listEntries   = 6
	yamlpathParserRULE_mapEntries    = 7
	yamlpathParserRULE_invocation    = 8
	yamlpathParserRULE_paramList     = 9
	yamlpathParserRULE_identifier    = 10
	yamlpathParserRULE_regex         = 11
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

func (s *IndexExpressionContext) BracketParam() IBracketParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketParamContext)
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
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(_la == yamlpathParserT__0 || _la == yamlpathParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(45)
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
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRecursiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(31)
					p.Match(yamlpathParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(33)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(32)
						p.Invocation()
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case 2:
				localctx = NewFieldExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(yamlpathParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.Invocation()
				}

			case 3:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					p.Match(yamlpathParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)
					p.BracketParam()
				}
				{
					p.SetState(41)
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
		p.SetState(47)
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

// IBracketParamContext is an interface to support dynamic dispatch.
type IBracketParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBracketParamContext differentiates from other interfaces.
	IsBracketParamContext()
}

type BracketParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketParamContext() *BracketParamContext {
	var p = new(BracketParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketParam
	return p
}

func InitEmptyBracketParamContext(p *BracketParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketParam
}

func (*BracketParamContext) IsBracketParamContext() {}

func NewBracketParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketParamContext {
	var p = new(BracketParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_bracketParam

	return p
}

func (s *BracketParamContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketParamContext) CopyAll(ctx *BracketParamContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BracketParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnionStringBracketContext struct {
	BracketParamContext
}

func NewUnionStringBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionStringBracketContext {
	var p = new(UnionStringBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *UnionStringBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionStringBracketContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserSTRING)
}

func (s *UnionStringBracketContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, i)
}

type WildcardBracketContext struct {
	BracketParamContext
}

func NewWildcardBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardBracketContext {
	var p = new(WildcardBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *WildcardBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

type FilterBracketContext struct {
	BracketParamContext
}

func NewFilterBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterBracketContext {
	var p = new(FilterBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *FilterBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterBracketContext) Subexpression() ISubexpressionContext {
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

type ScriptBracketContext struct {
	BracketParamContext
}

func NewScriptBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScriptBracketContext {
	var p = new(ScriptBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *ScriptBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptBracketContext) Subexpression() ISubexpressionContext {
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

type UnionNumberBracketContext struct {
	BracketParamContext
}

func NewUnionNumberBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionNumberBracketContext {
	var p = new(UnionNumberBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *UnionNumberBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionNumberBracketContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserNUMBER)
}

func (s *UnionNumberBracketContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, i)
}

type SliceBracketContext struct {
	BracketParamContext
}

func NewSliceBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceBracketContext {
	var p = new(SliceBracketContext)

	InitEmptyBracketParamContext(&p.BracketParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*BracketParamContext))

	return p
}

func (s *SliceBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceBracketContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserNUMBER)
}

func (s *SliceBracketContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, i)
}

func (p *yamlpathParser) BracketParam() (localctx IBracketParamContext) {
	localctx = NewBracketParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, yamlpathParserRULE_bracketParam)
	var _la int

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnionStringBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == yamlpathParserT__6 {
			{
				p.SetState(49)
				p.Match(yamlpathParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(50)
				p.Match(yamlpathParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewUnionNumberBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == yamlpathParserT__6 {
			{
				p.SetState(57)
				p.Match(yamlpathParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(58)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		localctx = NewWildcardBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(yamlpathParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSliceBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(65)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(68)
			p.Match(yamlpathParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserNUMBER {
			{
				p.SetState(69)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserT__8 {
			{
				p.SetState(72)
				p.Match(yamlpathParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		localctx = NewFilterBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Match(yamlpathParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.subexpression(0)
		}
		{
			p.SetState(79)
			p.Match(yamlpathParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewScriptBracketContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(81)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.subexpression(0)
		}
		{
			p.SetState(83)
			p.Match(yamlpathParserT__11)
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

type AggregationSubexpressionContext struct {
	SubexpressionContext
}

func NewAggregationSubexpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AggregationSubexpressionContext {
	var p = new(AggregationSubexpressionContext)

	InitEmptySubexpressionContext(&p.SubexpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubexpressionContext))

	return p
}

func (s *AggregationSubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationSubexpressionContext) Aggregation() IAggregationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregationContext)
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRootSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(88)
			p.expression(0)
		}

	case 2:
		localctx = NewLiteralSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Literal()
		}

	case 3:
		localctx = NewAggregationSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Aggregation()
		}

	case 4:
		localctx = NewParenthesisSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.subexpression(0)
		}
		{
			p.SetState(93)
			p.Match(yamlpathParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewNegationSubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__12 || _la == yamlpathParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(96)
			p.subexpression(10)
		}

	case 6:
		localctx = NewPolaritySubexpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__14 || _la == yamlpathParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(98)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(102)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&393472) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.subexpression(9)
				}

			case 2:
				localctx = NewAdditiveSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__14 || _la == yamlpathParserT__15) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(106)
					p.subexpression(8)
				}

			case 3:
				localctx = NewInequalitySubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(108)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(109)
					p.subexpression(7)
				}

			case 4:
				localctx = NewEqualitySubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(111)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__22 || _la == yamlpathParserT__23) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(112)
					p.subexpression(6)
				}

			case 5:
				localctx = NewMembershipSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&469762048) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)
					p.subexpression(4)
				}

			case 6:
				localctx = NewAndSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(117)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__28 || _la == yamlpathParserT__29) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)
					p.subexpression(3)
				}

			case 7:
				localctx = NewOrSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(120)
					_la = p.GetTokenStream().LA(1)

					if !(_la == yamlpathParserT__30 || _la == yamlpathParserT__31) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(121)
					p.subexpression(2)
				}

			case 8:
				localctx = NewMatchSubexpressionContext(p, NewSubexpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, yamlpathParserRULE_subexpression)
				p.SetState(122)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(123)
					p.Match(yamlpathParserT__24)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(124)
					p.Regex()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserSTRING:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
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
			p.SetState(131)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__32, yamlpathParserT__33:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			_la = p.GetTokenStream().LA(1)

			if !(_la == yamlpathParserT__32 || _la == yamlpathParserT__33) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case yamlpathParserT__34:
		localctx = NewNullLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.Match(yamlpathParserT__34)
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

// IAggregationContext is an interface to support dynamic dispatch.
type IAggregationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAggregationContext differentiates from other interfaces.
	IsAggregationContext()
}

type AggregationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregationContext() *AggregationContext {
	var p = new(AggregationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) CopyAll(ctx *AggregationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListAggregationContext struct {
	AggregationContext
}

func NewListAggregationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListAggregationContext {
	var p = new(ListAggregationContext)

	InitEmptyAggregationContext(&p.AggregationContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggregationContext))

	return p
}

func (s *ListAggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAggregationContext) ListEntries() IListEntriesContext {
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

type MapAggregationContext struct {
	AggregationContext
}

func NewMapAggregationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapAggregationContext {
	var p = new(MapAggregationContext)

	InitEmptyAggregationContext(&p.AggregationContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggregationContext))

	return p
}

func (s *MapAggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapAggregationContext) MapEntries() IMapEntriesContext {
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

type LiteralAggregationContext struct {
	AggregationContext
}

func NewLiteralAggregationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralAggregationContext {
	var p = new(LiteralAggregationContext)

	InitEmptyAggregationContext(&p.AggregationContext)
	p.parser = parser
	p.CopyAll(ctx.(*AggregationContext))

	return p
}

func (s *LiteralAggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralAggregationContext) Literal() ILiteralContext {
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

func (p *yamlpathParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, yamlpathParserRULE_aggregation)
	var _la int

	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserT__4:
		localctx = NewListAggregationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(yamlpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13254269075456) != 0 {
			{
				p.SetState(137)
				p.ListEntries()
			}

		}
		{
			p.SetState(140)
			p.Match(yamlpathParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__35:
		localctx = NewMapAggregationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(yamlpathParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == yamlpathParserSTRING {
			{
				p.SetState(142)
				p.MapEntries()
			}

		}
		{
			p.SetState(145)
			p.Match(yamlpathParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case yamlpathParserT__32, yamlpathParserT__33, yamlpathParserT__34, yamlpathParserNUMBER, yamlpathParserSTRING:
		localctx = NewLiteralAggregationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Literal()
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
		p.SetState(149)
		p.Literal()
	}

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__6 {
		{
			p.SetState(150)
			p.Match(yamlpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Literal()
		}

		p.SetState(156)
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
	AllAggregation() []IAggregationContext
	Aggregation(i int) IAggregationContext

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

func (s *MapEntriesContext) AllAggregation() []IAggregationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggregationContext); ok {
			len++
		}
	}

	tst := make([]IAggregationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggregationContext); ok {
			tst[i] = t.(IAggregationContext)
			i++
		}
	}

	return tst
}

func (s *MapEntriesContext) Aggregation(i int) IAggregationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
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

	return t.(IAggregationContext)
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
		p.SetState(157)
		p.Match(yamlpathParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(yamlpathParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Aggregation()
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__6 {
		{
			p.SetState(160)
			p.Match(yamlpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(yamlpathParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(yamlpathParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Aggregation()
		}

		p.SetState(168)
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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Identifier()
		}

	case 2:
		localctx = NewWildcardInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(yamlpathParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunctionInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Identifier()
		}
		{
			p.SetState(172)
			p.Match(yamlpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13322988677158) != 0 {
			{
				p.SetState(173)
				p.ParamList()
			}

		}
		{
			p.SetState(176)
			p.Match(yamlpathParserT__11)
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
	p.EnterRule(localctx, 18, yamlpathParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.subexpression(0)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == yamlpathParserT__6 {
		{
			p.SetState(181)
			p.Match(yamlpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.subexpression(0)
		}

		p.SetState(187)
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
	p.EnterRule(localctx, 20, yamlpathParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
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
	p.EnterRule(localctx, 22, yamlpathParserRULE_regex)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(yamlpathParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(191)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

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
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
