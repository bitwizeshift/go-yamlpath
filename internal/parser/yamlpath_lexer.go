// Code generated from yamlpath.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type yamlpathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var YamlpathLexerLexerStaticData struct {
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

func yamlpathlexerLexerInit() {
	staticData := &YamlpathLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'$'", "'@'", "'..'", "'.'", "'['", "']'", "':'", "'?'", "'('",
		"')'", "','", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'",
		"'||'", "'+'", "'-'", "'/'", "'in'", "'nin'", "'subsetof'", "'!'", "",
		"", "", "", "'null'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "NAME", "NUMBER", "STRING",
		"BOOLEAN", "NULL", "WILDCARD", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WILDCARD",
		"WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 219, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 5, 26, 141, 8, 26, 10, 26, 12, 26, 144, 9, 26, 1, 27, 3, 27, 147, 8,
		27, 1, 27, 4, 27, 150, 8, 27, 11, 27, 12, 27, 151, 1, 27, 1, 27, 4, 27,
		156, 8, 27, 11, 27, 12, 27, 157, 3, 27, 160, 8, 27, 1, 27, 1, 27, 3, 27,
		164, 8, 27, 1, 27, 4, 27, 167, 8, 27, 11, 27, 12, 27, 168, 3, 27, 171,
		8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 177, 8, 28, 10, 28, 12, 28, 180,
		9, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 193, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 32, 4, 32, 203, 8, 32, 11, 32, 12, 32, 204, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 5, 33, 213, 8, 33, 10, 33, 12, 33, 216, 9, 33,
		1, 33, 1, 33, 0, 0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 1, 0, 8,
		3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1,
		0, 48, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 34, 34,
		92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 231, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3,
		71, 1, 0, 0, 0, 5, 73, 1, 0, 0, 0, 7, 76, 1, 0, 0, 0, 9, 78, 1, 0, 0, 0,
		11, 80, 1, 0, 0, 0, 13, 82, 1, 0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 86, 1,
		0, 0, 0, 19, 88, 1, 0, 0, 0, 21, 90, 1, 0, 0, 0, 23, 92, 1, 0, 0, 0, 25,
		95, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29, 100, 1, 0, 0, 0, 31, 102, 1, 0,
		0, 0, 33, 105, 1, 0, 0, 0, 35, 108, 1, 0, 0, 0, 37, 111, 1, 0, 0, 0, 39,
		114, 1, 0, 0, 0, 41, 116, 1, 0, 0, 0, 43, 118, 1, 0, 0, 0, 45, 120, 1,
		0, 0, 0, 47, 123, 1, 0, 0, 0, 49, 127, 1, 0, 0, 0, 51, 136, 1, 0, 0, 0,
		53, 138, 1, 0, 0, 0, 55, 146, 1, 0, 0, 0, 57, 172, 1, 0, 0, 0, 59, 192,
		1, 0, 0, 0, 61, 194, 1, 0, 0, 0, 63, 199, 1, 0, 0, 0, 65, 202, 1, 0, 0,
		0, 67, 208, 1, 0, 0, 0, 69, 70, 5, 36, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72,
		5, 64, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5, 46, 0, 0, 74, 75, 5, 46, 0,
		0, 75, 6, 1, 0, 0, 0, 76, 77, 5, 46, 0, 0, 77, 8, 1, 0, 0, 0, 78, 79, 5,
		91, 0, 0, 79, 10, 1, 0, 0, 0, 80, 81, 5, 93, 0, 0, 81, 12, 1, 0, 0, 0,
		82, 83, 5, 58, 0, 0, 83, 14, 1, 0, 0, 0, 84, 85, 5, 63, 0, 0, 85, 16, 1,
		0, 0, 0, 86, 87, 5, 40, 0, 0, 87, 18, 1, 0, 0, 0, 88, 89, 5, 41, 0, 0,
		89, 20, 1, 0, 0, 0, 90, 91, 5, 44, 0, 0, 91, 22, 1, 0, 0, 0, 92, 93, 5,
		61, 0, 0, 93, 94, 5, 61, 0, 0, 94, 24, 1, 0, 0, 0, 95, 96, 5, 33, 0, 0,
		96, 97, 5, 61, 0, 0, 97, 26, 1, 0, 0, 0, 98, 99, 5, 60, 0, 0, 99, 28, 1,
		0, 0, 0, 100, 101, 5, 62, 0, 0, 101, 30, 1, 0, 0, 0, 102, 103, 5, 60, 0,
		0, 103, 104, 5, 61, 0, 0, 104, 32, 1, 0, 0, 0, 105, 106, 5, 62, 0, 0, 106,
		107, 5, 61, 0, 0, 107, 34, 1, 0, 0, 0, 108, 109, 5, 38, 0, 0, 109, 110,
		5, 38, 0, 0, 110, 36, 1, 0, 0, 0, 111, 112, 5, 124, 0, 0, 112, 113, 5,
		124, 0, 0, 113, 38, 1, 0, 0, 0, 114, 115, 5, 43, 0, 0, 115, 40, 1, 0, 0,
		0, 116, 117, 5, 45, 0, 0, 117, 42, 1, 0, 0, 0, 118, 119, 5, 47, 0, 0, 119,
		44, 1, 0, 0, 0, 120, 121, 5, 105, 0, 0, 121, 122, 5, 110, 0, 0, 122, 46,
		1, 0, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125, 5, 105, 0, 0, 125, 126, 5,
		110, 0, 0, 126, 48, 1, 0, 0, 0, 127, 128, 5, 115, 0, 0, 128, 129, 5, 117,
		0, 0, 129, 130, 5, 98, 0, 0, 130, 131, 5, 115, 0, 0, 131, 132, 5, 101,
		0, 0, 132, 133, 5, 116, 0, 0, 133, 134, 5, 111, 0, 0, 134, 135, 5, 102,
		0, 0, 135, 50, 1, 0, 0, 0, 136, 137, 5, 33, 0, 0, 137, 52, 1, 0, 0, 0,
		138, 142, 7, 0, 0, 0, 139, 141, 7, 1, 0, 0, 140, 139, 1, 0, 0, 0, 141,
		144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 54, 1,
		0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 147, 5, 45, 0, 0, 146, 145, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 150, 7, 2, 0, 0, 149,
		148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 159, 1, 0, 0, 0, 153, 155, 5, 46, 0, 0, 154, 156, 7, 2,
		0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0,
		157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 153, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 170, 1, 0, 0, 0, 161, 163, 7, 3, 0, 0, 162, 164,
		7, 4, 0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0,
		0, 0, 165, 167, 7, 2, 0, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0,
		168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170,
		161, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 56, 1, 0, 0, 0, 172, 178, 5,
		34, 0, 0, 173, 177, 8, 5, 0, 0, 174, 175, 5, 92, 0, 0, 175, 177, 9, 0,
		0, 0, 176, 173, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0,
		178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180,
		178, 1, 0, 0, 0, 181, 182, 5, 34, 0, 0, 182, 58, 1, 0, 0, 0, 183, 184,
		5, 116, 0, 0, 184, 185, 5, 114, 0, 0, 185, 186, 5, 117, 0, 0, 186, 193,
		5, 101, 0, 0, 187, 188, 5, 102, 0, 0, 188, 189, 5, 97, 0, 0, 189, 190,
		5, 108, 0, 0, 190, 191, 5, 115, 0, 0, 191, 193, 5, 101, 0, 0, 192, 183,
		1, 0, 0, 0, 192, 187, 1, 0, 0, 0, 193, 60, 1, 0, 0, 0, 194, 195, 5, 110,
		0, 0, 195, 196, 5, 117, 0, 0, 196, 197, 5, 108, 0, 0, 197, 198, 5, 108,
		0, 0, 198, 62, 1, 0, 0, 0, 199, 200, 5, 42, 0, 0, 200, 64, 1, 0, 0, 0,
		201, 203, 7, 6, 0, 0, 202, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207,
		6, 32, 0, 0, 207, 66, 1, 0, 0, 0, 208, 209, 5, 47, 0, 0, 209, 210, 5, 47,
		0, 0, 210, 214, 1, 0, 0, 0, 211, 213, 8, 7, 0, 0, 212, 211, 1, 0, 0, 0,
		213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215,
		217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 6, 33, 0, 0, 218, 68,
		1, 0, 0, 0, 14, 0, 142, 146, 151, 157, 159, 163, 168, 170, 176, 178, 192,
		204, 214, 1, 6, 0, 0,
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

// yamlpathLexerInit initializes any static state used to implement yamlpathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewyamlpathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func YamlpathLexerInit() {
	staticData := &YamlpathLexerLexerStaticData
	staticData.once.Do(yamlpathlexerLexerInit)
}

// NewyamlpathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewyamlpathLexer(input antlr.CharStream) *yamlpathLexer {
	YamlpathLexerInit()
	l := new(yamlpathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &YamlpathLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "yamlpath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// yamlpathLexer tokens.
const (
	yamlpathLexerT__0     = 1
	yamlpathLexerT__1     = 2
	yamlpathLexerT__2     = 3
	yamlpathLexerT__3     = 4
	yamlpathLexerT__4     = 5
	yamlpathLexerT__5     = 6
	yamlpathLexerT__6     = 7
	yamlpathLexerT__7     = 8
	yamlpathLexerT__8     = 9
	yamlpathLexerT__9     = 10
	yamlpathLexerT__10    = 11
	yamlpathLexerT__11    = 12
	yamlpathLexerT__12    = 13
	yamlpathLexerT__13    = 14
	yamlpathLexerT__14    = 15
	yamlpathLexerT__15    = 16
	yamlpathLexerT__16    = 17
	yamlpathLexerT__17    = 18
	yamlpathLexerT__18    = 19
	yamlpathLexerT__19    = 20
	yamlpathLexerT__20    = 21
	yamlpathLexerT__21    = 22
	yamlpathLexerT__22    = 23
	yamlpathLexerT__23    = 24
	yamlpathLexerT__24    = 25
	yamlpathLexerT__25    = 26
	yamlpathLexerNAME     = 27
	yamlpathLexerNUMBER   = 28
	yamlpathLexerSTRING   = 29
	yamlpathLexerBOOLEAN  = 30
	yamlpathLexerNULL     = 31
	yamlpathLexerWILDCARD = 32
	yamlpathLexerWS       = 33
	yamlpathLexerCOMMENT  = 34
)
