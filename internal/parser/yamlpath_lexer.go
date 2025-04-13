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
		"", "'$'", "'@'", "'..'", "'.'", "'['", "']'", "','", "'*'", "':'",
		"'('", "')'", "'!'", "'not'", "'+'", "'-'", "'/'", "'%'", "'<='", "'<'",
		"'>'", "'>='", "'=='", "'!='", "'=~'", "'in'", "'nin'", "'subsetof'",
		"'&&'", "'and'", "'||'", "'or'", "'true'", "'false'", "'null'", "'{'",
		"'}'", "'i'", "'m'", "'s'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "IDENTIFIER", "NUMBER", "STRING", "REGEX", "WS",
		"COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "IDENTIFIER",
		"NUMBER", "STRING", "REGEX", "WS", "COMMENT", "ESC", "UNICODE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 293, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 211, 8, 39,
		10, 39, 12, 39, 214, 9, 39, 1, 40, 3, 40, 217, 8, 40, 1, 40, 4, 40, 220,
		8, 40, 11, 40, 12, 40, 221, 1, 40, 1, 40, 4, 40, 226, 8, 40, 11, 40, 12,
		40, 227, 3, 40, 230, 8, 40, 1, 40, 1, 40, 3, 40, 234, 8, 40, 1, 40, 4,
		40, 237, 8, 40, 11, 40, 12, 40, 238, 3, 40, 241, 8, 40, 1, 41, 1, 41, 1,
		41, 5, 41, 246, 8, 41, 10, 41, 12, 41, 249, 9, 41, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 5, 42, 256, 8, 42, 10, 42, 12, 42, 259, 9, 42, 1, 42, 1,
		42, 1, 43, 4, 43, 264, 8, 43, 11, 43, 12, 43, 265, 1, 43, 1, 43, 1, 44,
		1, 44, 1, 44, 1, 44, 5, 44, 274, 8, 44, 10, 44, 12, 44, 277, 9, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 45, 3, 45, 284, 8, 45, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 2, 247, 257, 0, 48, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40,
		81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 0, 93, 0, 95, 0, 1, 0, 9, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0,
		48, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 8, 0, 39, 39, 47, 47, 92, 92, 96, 96,
		102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102,
		304, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 97, 1, 0, 0,
		0, 3, 99, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 104, 1, 0, 0, 0, 9, 106, 1,
		0, 0, 0, 11, 108, 1, 0, 0, 0, 13, 110, 1, 0, 0, 0, 15, 112, 1, 0, 0, 0,
		17, 114, 1, 0, 0, 0, 19, 116, 1, 0, 0, 0, 21, 118, 1, 0, 0, 0, 23, 120,
		1, 0, 0, 0, 25, 122, 1, 0, 0, 0, 27, 126, 1, 0, 0, 0, 29, 128, 1, 0, 0,
		0, 31, 130, 1, 0, 0, 0, 33, 132, 1, 0, 0, 0, 35, 134, 1, 0, 0, 0, 37, 137,
		1, 0, 0, 0, 39, 139, 1, 0, 0, 0, 41, 141, 1, 0, 0, 0, 43, 144, 1, 0, 0,
		0, 45, 147, 1, 0, 0, 0, 47, 150, 1, 0, 0, 0, 49, 153, 1, 0, 0, 0, 51, 156,
		1, 0, 0, 0, 53, 160, 1, 0, 0, 0, 55, 169, 1, 0, 0, 0, 57, 172, 1, 0, 0,
		0, 59, 176, 1, 0, 0, 0, 61, 179, 1, 0, 0, 0, 63, 182, 1, 0, 0, 0, 65, 187,
		1, 0, 0, 0, 67, 193, 1, 0, 0, 0, 69, 198, 1, 0, 0, 0, 71, 200, 1, 0, 0,
		0, 73, 202, 1, 0, 0, 0, 75, 204, 1, 0, 0, 0, 77, 206, 1, 0, 0, 0, 79, 208,
		1, 0, 0, 0, 81, 216, 1, 0, 0, 0, 83, 242, 1, 0, 0, 0, 85, 252, 1, 0, 0,
		0, 87, 263, 1, 0, 0, 0, 89, 269, 1, 0, 0, 0, 91, 280, 1, 0, 0, 0, 93, 285,
		1, 0, 0, 0, 95, 291, 1, 0, 0, 0, 97, 98, 5, 36, 0, 0, 98, 2, 1, 0, 0, 0,
		99, 100, 5, 64, 0, 0, 100, 4, 1, 0, 0, 0, 101, 102, 5, 46, 0, 0, 102, 103,
		5, 46, 0, 0, 103, 6, 1, 0, 0, 0, 104, 105, 5, 46, 0, 0, 105, 8, 1, 0, 0,
		0, 106, 107, 5, 91, 0, 0, 107, 10, 1, 0, 0, 0, 108, 109, 5, 93, 0, 0, 109,
		12, 1, 0, 0, 0, 110, 111, 5, 44, 0, 0, 111, 14, 1, 0, 0, 0, 112, 113, 5,
		42, 0, 0, 113, 16, 1, 0, 0, 0, 114, 115, 5, 58, 0, 0, 115, 18, 1, 0, 0,
		0, 116, 117, 5, 40, 0, 0, 117, 20, 1, 0, 0, 0, 118, 119, 5, 41, 0, 0, 119,
		22, 1, 0, 0, 0, 120, 121, 5, 33, 0, 0, 121, 24, 1, 0, 0, 0, 122, 123, 5,
		110, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125, 5, 116, 0, 0, 125, 26, 1,
		0, 0, 0, 126, 127, 5, 43, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 5, 45, 0,
		0, 129, 30, 1, 0, 0, 0, 130, 131, 5, 47, 0, 0, 131, 32, 1, 0, 0, 0, 132,
		133, 5, 37, 0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 5, 60, 0, 0, 135, 136,
		5, 61, 0, 0, 136, 36, 1, 0, 0, 0, 137, 138, 5, 60, 0, 0, 138, 38, 1, 0,
		0, 0, 139, 140, 5, 62, 0, 0, 140, 40, 1, 0, 0, 0, 141, 142, 5, 62, 0, 0,
		142, 143, 5, 61, 0, 0, 143, 42, 1, 0, 0, 0, 144, 145, 5, 61, 0, 0, 145,
		146, 5, 61, 0, 0, 146, 44, 1, 0, 0, 0, 147, 148, 5, 33, 0, 0, 148, 149,
		5, 61, 0, 0, 149, 46, 1, 0, 0, 0, 150, 151, 5, 61, 0, 0, 151, 152, 5, 126,
		0, 0, 152, 48, 1, 0, 0, 0, 153, 154, 5, 105, 0, 0, 154, 155, 5, 110, 0,
		0, 155, 50, 1, 0, 0, 0, 156, 157, 5, 110, 0, 0, 157, 158, 5, 105, 0, 0,
		158, 159, 5, 110, 0, 0, 159, 52, 1, 0, 0, 0, 160, 161, 5, 115, 0, 0, 161,
		162, 5, 117, 0, 0, 162, 163, 5, 98, 0, 0, 163, 164, 5, 115, 0, 0, 164,
		165, 5, 101, 0, 0, 165, 166, 5, 116, 0, 0, 166, 167, 5, 111, 0, 0, 167,
		168, 5, 102, 0, 0, 168, 54, 1, 0, 0, 0, 169, 170, 5, 38, 0, 0, 170, 171,
		5, 38, 0, 0, 171, 56, 1, 0, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 110,
		0, 0, 174, 175, 5, 100, 0, 0, 175, 58, 1, 0, 0, 0, 176, 177, 5, 124, 0,
		0, 177, 178, 5, 124, 0, 0, 178, 60, 1, 0, 0, 0, 179, 180, 5, 111, 0, 0,
		180, 181, 5, 114, 0, 0, 181, 62, 1, 0, 0, 0, 182, 183, 5, 116, 0, 0, 183,
		184, 5, 114, 0, 0, 184, 185, 5, 117, 0, 0, 185, 186, 5, 101, 0, 0, 186,
		64, 1, 0, 0, 0, 187, 188, 5, 102, 0, 0, 188, 189, 5, 97, 0, 0, 189, 190,
		5, 108, 0, 0, 190, 191, 5, 115, 0, 0, 191, 192, 5, 101, 0, 0, 192, 66,
		1, 0, 0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 117, 0, 0, 195, 196, 5,
		108, 0, 0, 196, 197, 5, 108, 0, 0, 197, 68, 1, 0, 0, 0, 198, 199, 5, 123,
		0, 0, 199, 70, 1, 0, 0, 0, 200, 201, 5, 125, 0, 0, 201, 72, 1, 0, 0, 0,
		202, 203, 5, 105, 0, 0, 203, 74, 1, 0, 0, 0, 204, 205, 5, 109, 0, 0, 205,
		76, 1, 0, 0, 0, 206, 207, 5, 115, 0, 0, 207, 78, 1, 0, 0, 0, 208, 212,
		7, 0, 0, 0, 209, 211, 7, 1, 0, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0,
		0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 80, 1, 0, 0, 0,
		214, 212, 1, 0, 0, 0, 215, 217, 5, 45, 0, 0, 216, 215, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 220, 7, 2, 0, 0, 219, 218,
		1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0,
		0, 0, 222, 229, 1, 0, 0, 0, 223, 225, 5, 46, 0, 0, 224, 226, 7, 2, 0, 0,
		225, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227,
		228, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229, 223, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 240, 1, 0, 0, 0, 231, 233, 7, 3, 0, 0, 232, 234, 7, 4,
		0, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0,
		235, 237, 7, 2, 0, 0, 236, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0, 240, 231,
		1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 82, 1, 0, 0, 0, 242, 247, 5, 34,
		0, 0, 243, 246, 3, 91, 45, 0, 244, 246, 9, 0, 0, 0, 245, 243, 1, 0, 0,
		0, 245, 244, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 247,
		245, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 251,
		5, 34, 0, 0, 251, 84, 1, 0, 0, 0, 252, 257, 5, 47, 0, 0, 253, 256, 3, 91,
		45, 0, 254, 256, 9, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0,
		256, 259, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258,
		260, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 261, 5, 47, 0, 0, 261, 86,
		1, 0, 0, 0, 262, 264, 7, 5, 0, 0, 263, 262, 1, 0, 0, 0, 264, 265, 1, 0,
		0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 268, 6, 43, 0, 0, 268, 88, 1, 0, 0, 0, 269, 270, 5, 47, 0, 0, 270,
		271, 5, 47, 0, 0, 271, 275, 1, 0, 0, 0, 272, 274, 8, 6, 0, 0, 273, 272,
		1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 279, 6, 44, 0, 0,
		279, 90, 1, 0, 0, 0, 280, 283, 5, 92, 0, 0, 281, 284, 7, 7, 0, 0, 282,
		284, 3, 93, 46, 0, 283, 281, 1, 0, 0, 0, 283, 282, 1, 0, 0, 0, 284, 92,
		1, 0, 0, 0, 285, 286, 5, 117, 0, 0, 286, 287, 3, 95, 47, 0, 287, 288, 3,
		95, 47, 0, 288, 289, 3, 95, 47, 0, 289, 290, 3, 95, 47, 0, 290, 94, 1,
		0, 0, 0, 291, 292, 7, 8, 0, 0, 292, 96, 1, 0, 0, 0, 16, 0, 212, 216, 221,
		227, 229, 233, 238, 240, 245, 247, 255, 257, 265, 275, 283, 1, 0, 1, 0,
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
	yamlpathLexerT__0       = 1
	yamlpathLexerT__1       = 2
	yamlpathLexerT__2       = 3
	yamlpathLexerT__3       = 4
	yamlpathLexerT__4       = 5
	yamlpathLexerT__5       = 6
	yamlpathLexerT__6       = 7
	yamlpathLexerT__7       = 8
	yamlpathLexerT__8       = 9
	yamlpathLexerT__9       = 10
	yamlpathLexerT__10      = 11
	yamlpathLexerT__11      = 12
	yamlpathLexerT__12      = 13
	yamlpathLexerT__13      = 14
	yamlpathLexerT__14      = 15
	yamlpathLexerT__15      = 16
	yamlpathLexerT__16      = 17
	yamlpathLexerT__17      = 18
	yamlpathLexerT__18      = 19
	yamlpathLexerT__19      = 20
	yamlpathLexerT__20      = 21
	yamlpathLexerT__21      = 22
	yamlpathLexerT__22      = 23
	yamlpathLexerT__23      = 24
	yamlpathLexerT__24      = 25
	yamlpathLexerT__25      = 26
	yamlpathLexerT__26      = 27
	yamlpathLexerT__27      = 28
	yamlpathLexerT__28      = 29
	yamlpathLexerT__29      = 30
	yamlpathLexerT__30      = 31
	yamlpathLexerT__31      = 32
	yamlpathLexerT__32      = 33
	yamlpathLexerT__33      = 34
	yamlpathLexerT__34      = 35
	yamlpathLexerT__35      = 36
	yamlpathLexerT__36      = 37
	yamlpathLexerT__37      = 38
	yamlpathLexerT__38      = 39
	yamlpathLexerIDENTIFIER = 40
	yamlpathLexerNUMBER     = 41
	yamlpathLexerSTRING     = 42
	yamlpathLexerREGEX      = 43
	yamlpathLexerWS         = 44
	yamlpathLexerCOMMENT    = 45
)
