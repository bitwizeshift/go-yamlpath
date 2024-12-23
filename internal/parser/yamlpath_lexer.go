// Code generated from yamlpath.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type yamlpathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
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
    "", "'$'", "'@'", "'.'", "'*'", "'..'", "'['", "']'", "':'", "'?'", 
    "'('", "')'", "','", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "", 
    "", "", "", "'null'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WS", "COMMENT",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "T__17", "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WS", "COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 25, 171, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 
	1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 
	12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 
	1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 95, 8, 18, 10, 
	18, 12, 18, 98, 9, 18, 1, 19, 3, 19, 101, 8, 19, 1, 19, 4, 19, 104, 8, 
	19, 11, 19, 12, 19, 105, 1, 19, 1, 19, 4, 19, 110, 8, 19, 11, 19, 12, 19, 
	111, 3, 19, 114, 8, 19, 1, 19, 1, 19, 3, 19, 118, 8, 19, 1, 19, 4, 19, 
	121, 8, 19, 11, 19, 12, 19, 122, 3, 19, 125, 8, 19, 1, 20, 1, 20, 1, 20, 
	1, 20, 5, 20, 131, 8, 20, 10, 20, 12, 20, 134, 9, 20, 1, 20, 1, 20, 1, 
	21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 147, 
	8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 4, 23, 155, 8, 23, 11, 
	23, 12, 23, 156, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 165, 
	8, 24, 10, 24, 12, 24, 168, 9, 24, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 
	5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 
	13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 
	22, 45, 23, 47, 24, 49, 25, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 4, 
	0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101, 
	2, 0, 43, 43, 45, 45, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 
	2, 0, 10, 10, 13, 13, 183, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 
	0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 
	1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 
	21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 
	0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 
	0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 
	0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 
	0, 0, 0, 3, 53, 1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 59, 
	1, 0, 0, 0, 11, 62, 1, 0, 0, 0, 13, 64, 1, 0, 0, 0, 15, 66, 1, 0, 0, 0, 
	17, 68, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21, 72, 1, 0, 0, 0, 23, 74, 1, 
	0, 0, 0, 25, 76, 1, 0, 0, 0, 27, 79, 1, 0, 0, 0, 29, 82, 1, 0, 0, 0, 31, 
	84, 1, 0, 0, 0, 33, 86, 1, 0, 0, 0, 35, 89, 1, 0, 0, 0, 37, 92, 1, 0, 0, 
	0, 39, 100, 1, 0, 0, 0, 41, 126, 1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 148, 
	1, 0, 0, 0, 47, 154, 1, 0, 0, 0, 49, 160, 1, 0, 0, 0, 51, 52, 5, 36, 0, 
	0, 52, 2, 1, 0, 0, 0, 53, 54, 5, 64, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 
	46, 0, 0, 56, 6, 1, 0, 0, 0, 57, 58, 5, 42, 0, 0, 58, 8, 1, 0, 0, 0, 59, 
	60, 5, 46, 0, 0, 60, 61, 5, 46, 0, 0, 61, 10, 1, 0, 0, 0, 62, 63, 5, 91, 
	0, 0, 63, 12, 1, 0, 0, 0, 64, 65, 5, 93, 0, 0, 65, 14, 1, 0, 0, 0, 66, 
	67, 5, 58, 0, 0, 67, 16, 1, 0, 0, 0, 68, 69, 5, 63, 0, 0, 69, 18, 1, 0, 
	0, 0, 70, 71, 5, 40, 0, 0, 71, 20, 1, 0, 0, 0, 72, 73, 5, 41, 0, 0, 73, 
	22, 1, 0, 0, 0, 74, 75, 5, 44, 0, 0, 75, 24, 1, 0, 0, 0, 76, 77, 5, 61, 
	0, 0, 77, 78, 5, 61, 0, 0, 78, 26, 1, 0, 0, 0, 79, 80, 5, 33, 0, 0, 80, 
	81, 5, 61, 0, 0, 81, 28, 1, 0, 0, 0, 82, 83, 5, 60, 0, 0, 83, 30, 1, 0, 
	0, 0, 84, 85, 5, 62, 0, 0, 85, 32, 1, 0, 0, 0, 86, 87, 5, 60, 0, 0, 87, 
	88, 5, 61, 0, 0, 88, 34, 1, 0, 0, 0, 89, 90, 5, 62, 0, 0, 90, 91, 5, 61, 
	0, 0, 91, 36, 1, 0, 0, 0, 92, 96, 7, 0, 0, 0, 93, 95, 7, 1, 0, 0, 94, 93, 
	1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 
	97, 38, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 101, 5, 45, 0, 0, 100, 99, 
	1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 104, 7, 2, 
	0, 0, 103, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 
	105, 106, 1, 0, 0, 0, 106, 113, 1, 0, 0, 0, 107, 109, 5, 46, 0, 0, 108, 
	110, 7, 2, 0, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 
	1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 107, 1, 0, 
	0, 0, 113, 114, 1, 0, 0, 0, 114, 124, 1, 0, 0, 0, 115, 117, 7, 3, 0, 0, 
	116, 118, 7, 4, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 
	120, 1, 0, 0, 0, 119, 121, 7, 2, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 
	1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 
	0, 0, 124, 115, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 40, 1, 0, 0, 0, 
	126, 132, 5, 34, 0, 0, 127, 131, 8, 5, 0, 0, 128, 129, 5, 92, 0, 0, 129, 
	131, 9, 0, 0, 0, 130, 127, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 134, 
	1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 
	0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 34, 0, 0, 136, 42, 1, 0, 0, 0, 
	137, 138, 5, 116, 0, 0, 138, 139, 5, 114, 0, 0, 139, 140, 5, 117, 0, 0, 
	140, 147, 5, 101, 0, 0, 141, 142, 5, 102, 0, 0, 142, 143, 5, 97, 0, 0, 
	143, 144, 5, 108, 0, 0, 144, 145, 5, 115, 0, 0, 145, 147, 5, 101, 0, 0, 
	146, 137, 1, 0, 0, 0, 146, 141, 1, 0, 0, 0, 147, 44, 1, 0, 0, 0, 148, 149, 
	5, 110, 0, 0, 149, 150, 5, 117, 0, 0, 150, 151, 5, 108, 0, 0, 151, 152, 
	5, 108, 0, 0, 152, 46, 1, 0, 0, 0, 153, 155, 7, 6, 0, 0, 154, 153, 1, 0, 
	0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 
	157, 158, 1, 0, 0, 0, 158, 159, 6, 23, 0, 0, 159, 48, 1, 0, 0, 0, 160, 
	161, 5, 47, 0, 0, 161, 162, 5, 47, 0, 0, 162, 166, 1, 0, 0, 0, 163, 165, 
	8, 7, 0, 0, 164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 
	0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 
	169, 170, 6, 24, 0, 0, 170, 50, 1, 0, 0, 0, 14, 0, 96, 100, 105, 111, 113, 
	117, 122, 124, 130, 132, 146, 156, 166, 1, 6, 0, 0,
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
	yamlpathLexerT__0 = 1
	yamlpathLexerT__1 = 2
	yamlpathLexerT__2 = 3
	yamlpathLexerT__3 = 4
	yamlpathLexerT__4 = 5
	yamlpathLexerT__5 = 6
	yamlpathLexerT__6 = 7
	yamlpathLexerT__7 = 8
	yamlpathLexerT__8 = 9
	yamlpathLexerT__9 = 10
	yamlpathLexerT__10 = 11
	yamlpathLexerT__11 = 12
	yamlpathLexerT__12 = 13
	yamlpathLexerT__13 = 14
	yamlpathLexerT__14 = 15
	yamlpathLexerT__15 = 16
	yamlpathLexerT__16 = 17
	yamlpathLexerT__17 = 18
	yamlpathLexerNAME = 19
	yamlpathLexerNUMBER = 20
	yamlpathLexerSTRING = 21
	yamlpathLexerBOOLEAN = 22
	yamlpathLexerNULL = 23
	yamlpathLexerWS = 24
	yamlpathLexerCOMMENT = 25
)

