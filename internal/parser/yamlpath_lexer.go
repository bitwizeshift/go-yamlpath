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
    "", "'$'", "'@'", "'.'", "'..'", "'['", "']'", "':'", "'?'", "'('", 
    "')'", "','", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "", "", 
    "", "", "'null'", "'*'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WILDCARD", "WS", 
    "COMMENT",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WILDCARD", "WS", "COMMENT",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 25, 171, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 
	1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 
	11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 
	1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 93, 8, 17, 10, 17, 12, 17, 96, 
	9, 17, 1, 18, 3, 18, 99, 8, 18, 1, 18, 4, 18, 102, 8, 18, 11, 18, 12, 18, 
	103, 1, 18, 1, 18, 4, 18, 108, 8, 18, 11, 18, 12, 18, 109, 3, 18, 112, 
	8, 18, 1, 18, 1, 18, 3, 18, 116, 8, 18, 1, 18, 4, 18, 119, 8, 18, 11, 18, 
	12, 18, 120, 3, 18, 123, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 129, 
	8, 19, 10, 19, 12, 19, 132, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 
	20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 145, 8, 20, 1, 21, 1, 21, 
	1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 4, 23, 155, 8, 23, 11, 23, 12, 
	23, 156, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 165, 8, 24, 10, 
	24, 12, 24, 168, 9, 24, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7, 4, 
	9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 
	29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 
	47, 24, 49, 25, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 
	90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 
	45, 45, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 
	13, 13, 183, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 
	1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 
	15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 
	0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 
	0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 
	0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 
	0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 
	1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 
	62, 1, 0, 0, 0, 13, 64, 1, 0, 0, 0, 15, 66, 1, 0, 0, 0, 17, 68, 1, 0, 0, 
	0, 19, 70, 1, 0, 0, 0, 21, 72, 1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 77, 
	1, 0, 0, 0, 27, 80, 1, 0, 0, 0, 29, 82, 1, 0, 0, 0, 31, 84, 1, 0, 0, 0, 
	33, 87, 1, 0, 0, 0, 35, 90, 1, 0, 0, 0, 37, 98, 1, 0, 0, 0, 39, 124, 1, 
	0, 0, 0, 41, 144, 1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 151, 1, 0, 0, 0, 
	47, 154, 1, 0, 0, 0, 49, 160, 1, 0, 0, 0, 51, 52, 5, 36, 0, 0, 52, 2, 1, 
	0, 0, 0, 53, 54, 5, 64, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 
	6, 1, 0, 0, 0, 57, 58, 5, 46, 0, 0, 58, 59, 5, 46, 0, 0, 59, 8, 1, 0, 0, 
	0, 60, 61, 5, 91, 0, 0, 61, 10, 1, 0, 0, 0, 62, 63, 5, 93, 0, 0, 63, 12, 
	1, 0, 0, 0, 64, 65, 5, 58, 0, 0, 65, 14, 1, 0, 0, 0, 66, 67, 5, 63, 0, 
	0, 67, 16, 1, 0, 0, 0, 68, 69, 5, 40, 0, 0, 69, 18, 1, 0, 0, 0, 70, 71, 
	5, 41, 0, 0, 71, 20, 1, 0, 0, 0, 72, 73, 5, 44, 0, 0, 73, 22, 1, 0, 0, 
	0, 74, 75, 5, 61, 0, 0, 75, 76, 5, 61, 0, 0, 76, 24, 1, 0, 0, 0, 77, 78, 
	5, 33, 0, 0, 78, 79, 5, 61, 0, 0, 79, 26, 1, 0, 0, 0, 80, 81, 5, 60, 0, 
	0, 81, 28, 1, 0, 0, 0, 82, 83, 5, 62, 0, 0, 83, 30, 1, 0, 0, 0, 84, 85, 
	5, 60, 0, 0, 85, 86, 5, 61, 0, 0, 86, 32, 1, 0, 0, 0, 87, 88, 5, 62, 0, 
	0, 88, 89, 5, 61, 0, 0, 89, 34, 1, 0, 0, 0, 90, 94, 7, 0, 0, 0, 91, 93, 
	7, 1, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 
	94, 95, 1, 0, 0, 0, 95, 36, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 5, 
	45, 0, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 
	100, 102, 7, 2, 0, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 
	101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 111, 1, 0, 0, 0, 105, 107, 
	5, 46, 0, 0, 106, 108, 7, 2, 0, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 
	0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 
	111, 105, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 122, 1, 0, 0, 0, 113, 
	115, 7, 3, 0, 0, 114, 116, 7, 4, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 
	1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 119, 7, 2, 0, 0, 118, 117, 1, 0, 
	0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 
	121, 123, 1, 0, 0, 0, 122, 113, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 
	38, 1, 0, 0, 0, 124, 130, 5, 34, 0, 0, 125, 129, 8, 5, 0, 0, 126, 127, 
	5, 92, 0, 0, 127, 129, 9, 0, 0, 0, 128, 125, 1, 0, 0, 0, 128, 126, 1, 0, 
	0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 
	131, 133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 134, 5, 34, 0, 0, 134, 
	40, 1, 0, 0, 0, 135, 136, 5, 116, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 
	5, 117, 0, 0, 138, 145, 5, 101, 0, 0, 139, 140, 5, 102, 0, 0, 140, 141, 
	5, 97, 0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 115, 0, 0, 143, 145, 
	5, 101, 0, 0, 144, 135, 1, 0, 0, 0, 144, 139, 1, 0, 0, 0, 145, 42, 1, 0, 
	0, 0, 146, 147, 5, 110, 0, 0, 147, 148, 5, 117, 0, 0, 148, 149, 5, 108, 
	0, 0, 149, 150, 5, 108, 0, 0, 150, 44, 1, 0, 0, 0, 151, 152, 5, 42, 0, 
	0, 152, 46, 1, 0, 0, 0, 153, 155, 7, 6, 0, 0, 154, 153, 1, 0, 0, 0, 155, 
	156, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 
	1, 0, 0, 0, 158, 159, 6, 23, 0, 0, 159, 48, 1, 0, 0, 0, 160, 161, 5, 47, 
	0, 0, 161, 162, 5, 47, 0, 0, 162, 166, 1, 0, 0, 0, 163, 165, 8, 7, 0, 0, 
	164, 163, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 
	167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 170, 
	6, 24, 0, 0, 170, 50, 1, 0, 0, 0, 14, 0, 94, 98, 103, 109, 111, 115, 120, 
	122, 128, 130, 144, 156, 166, 1, 6, 0, 0,
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
	yamlpathLexerNAME = 18
	yamlpathLexerNUMBER = 19
	yamlpathLexerSTRING = 20
	yamlpathLexerBOOLEAN = 21
	yamlpathLexerNULL = 22
	yamlpathLexerWILDCARD = 23
	yamlpathLexerWS = 24
	yamlpathLexerCOMMENT = 25
)

