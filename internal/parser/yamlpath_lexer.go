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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "IDENTIFIER",
		"NUMBER", "STRING", "REGEX", "WS", "COMMENT", "ESC", "UNICODE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 297, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40,
		1, 40, 5, 40, 215, 8, 40, 10, 40, 12, 40, 218, 9, 40, 1, 41, 3, 41, 221,
		8, 41, 1, 41, 4, 41, 224, 8, 41, 11, 41, 12, 41, 225, 1, 41, 1, 41, 4,
		41, 230, 8, 41, 11, 41, 12, 41, 231, 3, 41, 234, 8, 41, 1, 41, 1, 41, 3,
		41, 238, 8, 41, 1, 41, 4, 41, 241, 8, 41, 11, 41, 12, 41, 242, 3, 41, 245,
		8, 41, 1, 42, 1, 42, 1, 42, 5, 42, 250, 8, 42, 10, 42, 12, 42, 253, 9,
		42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 5, 43, 260, 8, 43, 10, 43, 12, 43,
		263, 9, 43, 1, 43, 1, 43, 1, 44, 4, 44, 268, 8, 44, 11, 44, 12, 44, 269,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 278, 8, 45, 10, 45, 12,
		45, 281, 9, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 3, 46, 288, 8, 46, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 2, 251, 261, 0, 49,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 0,
		95, 0, 97, 0, 1, 0, 9, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43,
		45, 45, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 8, 0, 39, 39,
		47, 47, 92, 92, 96, 96, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48,
		57, 65, 70, 97, 102, 308, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 1, 99, 1, 0, 0, 0, 3, 101, 1, 0, 0, 0, 5, 103,
		1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 108, 1, 0, 0, 0, 11, 110, 1, 0, 0, 0,
		13, 112, 1, 0, 0, 0, 15, 114, 1, 0, 0, 0, 17, 116, 1, 0, 0, 0, 19, 118,
		1, 0, 0, 0, 21, 120, 1, 0, 0, 0, 23, 122, 1, 0, 0, 0, 25, 126, 1, 0, 0,
		0, 27, 128, 1, 0, 0, 0, 29, 130, 1, 0, 0, 0, 31, 132, 1, 0, 0, 0, 33, 134,
		1, 0, 0, 0, 35, 136, 1, 0, 0, 0, 37, 139, 1, 0, 0, 0, 39, 141, 1, 0, 0,
		0, 41, 143, 1, 0, 0, 0, 43, 146, 1, 0, 0, 0, 45, 149, 1, 0, 0, 0, 47, 152,
		1, 0, 0, 0, 49, 155, 1, 0, 0, 0, 51, 158, 1, 0, 0, 0, 53, 162, 1, 0, 0,
		0, 55, 171, 1, 0, 0, 0, 57, 174, 1, 0, 0, 0, 59, 178, 1, 0, 0, 0, 61, 181,
		1, 0, 0, 0, 63, 184, 1, 0, 0, 0, 65, 189, 1, 0, 0, 0, 67, 195, 1, 0, 0,
		0, 69, 200, 1, 0, 0, 0, 71, 202, 1, 0, 0, 0, 73, 204, 1, 0, 0, 0, 75, 206,
		1, 0, 0, 0, 77, 208, 1, 0, 0, 0, 79, 210, 1, 0, 0, 0, 81, 212, 1, 0, 0,
		0, 83, 220, 1, 0, 0, 0, 85, 246, 1, 0, 0, 0, 87, 256, 1, 0, 0, 0, 89, 267,
		1, 0, 0, 0, 91, 273, 1, 0, 0, 0, 93, 284, 1, 0, 0, 0, 95, 289, 1, 0, 0,
		0, 97, 295, 1, 0, 0, 0, 99, 100, 5, 36, 0, 0, 100, 2, 1, 0, 0, 0, 101,
		102, 5, 64, 0, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 46, 0, 0, 104, 105,
		5, 46, 0, 0, 105, 6, 1, 0, 0, 0, 106, 107, 5, 46, 0, 0, 107, 8, 1, 0, 0,
		0, 108, 109, 5, 91, 0, 0, 109, 10, 1, 0, 0, 0, 110, 111, 5, 93, 0, 0, 111,
		12, 1, 0, 0, 0, 112, 113, 5, 42, 0, 0, 113, 14, 1, 0, 0, 0, 114, 115, 5,
		58, 0, 0, 115, 16, 1, 0, 0, 0, 116, 117, 5, 40, 0, 0, 117, 18, 1, 0, 0,
		0, 118, 119, 5, 41, 0, 0, 119, 20, 1, 0, 0, 0, 120, 121, 5, 33, 0, 0, 121,
		22, 1, 0, 0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125,
		5, 116, 0, 0, 125, 24, 1, 0, 0, 0, 126, 127, 5, 43, 0, 0, 127, 26, 1, 0,
		0, 0, 128, 129, 5, 45, 0, 0, 129, 28, 1, 0, 0, 0, 130, 131, 5, 47, 0, 0,
		131, 30, 1, 0, 0, 0, 132, 133, 5, 37, 0, 0, 133, 32, 1, 0, 0, 0, 134, 135,
		5, 124, 0, 0, 135, 34, 1, 0, 0, 0, 136, 137, 5, 60, 0, 0, 137, 138, 5,
		61, 0, 0, 138, 36, 1, 0, 0, 0, 139, 140, 5, 60, 0, 0, 140, 38, 1, 0, 0,
		0, 141, 142, 5, 62, 0, 0, 142, 40, 1, 0, 0, 0, 143, 144, 5, 62, 0, 0, 144,
		145, 5, 61, 0, 0, 145, 42, 1, 0, 0, 0, 146, 147, 5, 61, 0, 0, 147, 148,
		5, 61, 0, 0, 148, 44, 1, 0, 0, 0, 149, 150, 5, 33, 0, 0, 150, 151, 5, 61,
		0, 0, 151, 46, 1, 0, 0, 0, 152, 153, 5, 61, 0, 0, 153, 154, 5, 126, 0,
		0, 154, 48, 1, 0, 0, 0, 155, 156, 5, 105, 0, 0, 156, 157, 5, 110, 0, 0,
		157, 50, 1, 0, 0, 0, 158, 159, 5, 110, 0, 0, 159, 160, 5, 105, 0, 0, 160,
		161, 5, 110, 0, 0, 161, 52, 1, 0, 0, 0, 162, 163, 5, 115, 0, 0, 163, 164,
		5, 117, 0, 0, 164, 165, 5, 98, 0, 0, 165, 166, 5, 115, 0, 0, 166, 167,
		5, 101, 0, 0, 167, 168, 5, 116, 0, 0, 168, 169, 5, 111, 0, 0, 169, 170,
		5, 102, 0, 0, 170, 54, 1, 0, 0, 0, 171, 172, 5, 38, 0, 0, 172, 173, 5,
		38, 0, 0, 173, 56, 1, 0, 0, 0, 174, 175, 5, 97, 0, 0, 175, 176, 5, 110,
		0, 0, 176, 177, 5, 100, 0, 0, 177, 58, 1, 0, 0, 0, 178, 179, 5, 124, 0,
		0, 179, 180, 5, 124, 0, 0, 180, 60, 1, 0, 0, 0, 181, 182, 5, 111, 0, 0,
		182, 183, 5, 114, 0, 0, 183, 62, 1, 0, 0, 0, 184, 185, 5, 116, 0, 0, 185,
		186, 5, 114, 0, 0, 186, 187, 5, 117, 0, 0, 187, 188, 5, 101, 0, 0, 188,
		64, 1, 0, 0, 0, 189, 190, 5, 102, 0, 0, 190, 191, 5, 97, 0, 0, 191, 192,
		5, 108, 0, 0, 192, 193, 5, 115, 0, 0, 193, 194, 5, 101, 0, 0, 194, 66,
		1, 0, 0, 0, 195, 196, 5, 110, 0, 0, 196, 197, 5, 117, 0, 0, 197, 198, 5,
		108, 0, 0, 198, 199, 5, 108, 0, 0, 199, 68, 1, 0, 0, 0, 200, 201, 5, 123,
		0, 0, 201, 70, 1, 0, 0, 0, 202, 203, 5, 125, 0, 0, 203, 72, 1, 0, 0, 0,
		204, 205, 5, 44, 0, 0, 205, 74, 1, 0, 0, 0, 206, 207, 5, 105, 0, 0, 207,
		76, 1, 0, 0, 0, 208, 209, 5, 109, 0, 0, 209, 78, 1, 0, 0, 0, 210, 211,
		5, 115, 0, 0, 211, 80, 1, 0, 0, 0, 212, 216, 7, 0, 0, 0, 213, 215, 7, 1,
		0, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0,
		216, 217, 1, 0, 0, 0, 217, 82, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 221,
		5, 45, 0, 0, 220, 219, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0,
		0, 0, 222, 224, 7, 2, 0, 0, 223, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0,
		225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 233, 1, 0, 0, 0, 227,
		229, 5, 46, 0, 0, 228, 230, 7, 2, 0, 0, 229, 228, 1, 0, 0, 0, 230, 231,
		1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0,
		0, 0, 233, 227, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 244, 1, 0, 0, 0,
		235, 237, 7, 3, 0, 0, 236, 238, 7, 4, 0, 0, 237, 236, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 241, 7, 2, 0, 0, 240, 239,
		1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0,
		0, 0, 243, 245, 1, 0, 0, 0, 244, 235, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0,
		245, 84, 1, 0, 0, 0, 246, 251, 5, 34, 0, 0, 247, 250, 3, 93, 46, 0, 248,
		250, 9, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 248, 1, 0, 0, 0, 250, 253,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 254, 1, 0,
		0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 5, 34, 0, 0, 255, 86, 1, 0, 0, 0,
		256, 261, 5, 47, 0, 0, 257, 260, 3, 93, 46, 0, 258, 260, 9, 0, 0, 0, 259,
		257, 1, 0, 0, 0, 259, 258, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 262,
		1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 261, 1, 0,
		0, 0, 264, 265, 5, 47, 0, 0, 265, 88, 1, 0, 0, 0, 266, 268, 7, 5, 0, 0,
		267, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269,
		270, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 6, 44, 0, 0, 272, 90,
		1, 0, 0, 0, 273, 274, 5, 47, 0, 0, 274, 275, 5, 47, 0, 0, 275, 279, 1,
		0, 0, 0, 276, 278, 8, 6, 0, 0, 277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0,
		0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281,
		279, 1, 0, 0, 0, 282, 283, 6, 45, 0, 0, 283, 92, 1, 0, 0, 0, 284, 287,
		5, 92, 0, 0, 285, 288, 7, 7, 0, 0, 286, 288, 3, 95, 47, 0, 287, 285, 1,
		0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 94, 1, 0, 0, 0, 289, 290, 5, 117, 0,
		0, 290, 291, 3, 97, 48, 0, 291, 292, 3, 97, 48, 0, 292, 293, 3, 97, 48,
		0, 293, 294, 3, 97, 48, 0, 294, 96, 1, 0, 0, 0, 295, 296, 7, 8, 0, 0, 296,
		98, 1, 0, 0, 0, 16, 0, 216, 220, 225, 231, 233, 237, 242, 244, 249, 251,
		259, 261, 269, 279, 287, 1, 0, 1, 0,
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
	yamlpathLexerT__39      = 40
	yamlpathLexerIDENTIFIER = 41
	yamlpathLexerNUMBER     = 42
	yamlpathLexerSTRING     = 43
	yamlpathLexerREGEX      = 44
	yamlpathLexerWS         = 45
	yamlpathLexerCOMMENT    = 46
)
