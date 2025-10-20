// Code generated from ConditionExpr.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type ConditionExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ConditionExprLexerLexerStaticData struct {
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

func conditionexprlexerLexerInit() {
	staticData := &ConditionExprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'",
		"'&&'", "'||'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='", "'('",
		"')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "BOOLEAN", "IDENTIFIER", "NUMBER", "STRING", "PLUS", "MINUS",
		"MULTIPLY", "DIVIDE", "MODULO", "NOT", "AND", "OR", "GT", "GE", "LT",
		"LE", "EQ", "NE", "LPAREN", "RPAREN", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "BOOLEAN", "IDENTIFIER", "NUMBER", "STRING", "ESC", "UNICODE",
		"HEX", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "NOT", "AND",
		"OR", "GT", "GE", "LT", "LE", "EQ", "NE", "LPAREN", "RPAREN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 152, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 63, 8, 1, 1,
		2, 1, 2, 5, 2, 67, 8, 2, 10, 2, 12, 2, 70, 9, 2, 1, 3, 4, 3, 73, 8, 3,
		11, 3, 12, 3, 74, 1, 3, 1, 3, 4, 3, 79, 8, 3, 11, 3, 12, 3, 80, 3, 3, 83,
		8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 88, 8, 4, 10, 4, 12, 4, 91, 9, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 3, 5, 98, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 4, 24, 147,
		8, 24, 11, 24, 12, 24, 148, 1, 24, 1, 24, 0, 0, 25, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 0, 13, 0, 15, 0, 17, 6, 19, 7, 21, 8, 23, 9, 25, 10, 27, 11,
		29, 12, 31, 13, 33, 14, 35, 15, 37, 16, 39, 17, 41, 18, 43, 19, 45, 20,
		47, 21, 49, 22, 1, 0, 7, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 8, 0, 34, 34,
		47, 47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48,
		57, 65, 70, 97, 102, 3, 0, 9, 10, 13, 13, 32, 32, 157, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 62, 1, 0, 0, 0, 5,
		64, 1, 0, 0, 0, 7, 72, 1, 0, 0, 0, 9, 84, 1, 0, 0, 0, 11, 94, 1, 0, 0,
		0, 13, 99, 1, 0, 0, 0, 15, 105, 1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 109,
		1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 113, 1, 0, 0, 0, 25, 115, 1, 0, 0,
		0, 27, 117, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 122, 1, 0, 0, 0, 33, 125,
		1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37, 130, 1, 0, 0, 0, 39, 132, 1, 0, 0,
		0, 41, 135, 1, 0, 0, 0, 43, 138, 1, 0, 0, 0, 45, 141, 1, 0, 0, 0, 47, 143,
		1, 0, 0, 0, 49, 146, 1, 0, 0, 0, 51, 52, 5, 44, 0, 0, 52, 2, 1, 0, 0, 0,
		53, 54, 5, 116, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 117, 0, 0, 56, 63,
		5, 101, 0, 0, 57, 58, 5, 102, 0, 0, 58, 59, 5, 97, 0, 0, 59, 60, 5, 108,
		0, 0, 60, 61, 5, 115, 0, 0, 61, 63, 5, 101, 0, 0, 62, 53, 1, 0, 0, 0, 62,
		57, 1, 0, 0, 0, 63, 4, 1, 0, 0, 0, 64, 68, 7, 0, 0, 0, 65, 67, 7, 1, 0,
		0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 69, 6, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 73, 7, 2, 0, 0,
		72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 82, 1, 0, 0, 0, 76, 78, 5, 46, 0, 0, 77, 79, 7, 2, 0, 0, 78,
		77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0,
		0, 81, 83, 1, 0, 0, 0, 82, 76, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 8, 1,
		0, 0, 0, 84, 89, 5, 34, 0, 0, 85, 88, 3, 11, 5, 0, 86, 88, 8, 3, 0, 0,
		87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1,
		0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92,
		93, 5, 34, 0, 0, 93, 10, 1, 0, 0, 0, 94, 97, 5, 92, 0, 0, 95, 98, 7, 4,
		0, 0, 96, 98, 3, 13, 6, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98,
		12, 1, 0, 0, 0, 99, 100, 5, 117, 0, 0, 100, 101, 3, 15, 7, 0, 101, 102,
		3, 15, 7, 0, 102, 103, 3, 15, 7, 0, 103, 104, 3, 15, 7, 0, 104, 14, 1,
		0, 0, 0, 105, 106, 7, 5, 0, 0, 106, 16, 1, 0, 0, 0, 107, 108, 5, 43, 0,
		0, 108, 18, 1, 0, 0, 0, 109, 110, 5, 45, 0, 0, 110, 20, 1, 0, 0, 0, 111,
		112, 5, 42, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114, 5, 47, 0, 0, 114, 24,
		1, 0, 0, 0, 115, 116, 5, 37, 0, 0, 116, 26, 1, 0, 0, 0, 117, 118, 5, 33,
		0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 38, 0, 0, 120, 121, 5, 38, 0, 0,
		121, 30, 1, 0, 0, 0, 122, 123, 5, 124, 0, 0, 123, 124, 5, 124, 0, 0, 124,
		32, 1, 0, 0, 0, 125, 126, 5, 62, 0, 0, 126, 34, 1, 0, 0, 0, 127, 128, 5,
		62, 0, 0, 128, 129, 5, 61, 0, 0, 129, 36, 1, 0, 0, 0, 130, 131, 5, 60,
		0, 0, 131, 38, 1, 0, 0, 0, 132, 133, 5, 60, 0, 0, 133, 134, 5, 61, 0, 0,
		134, 40, 1, 0, 0, 0, 135, 136, 5, 61, 0, 0, 136, 137, 5, 61, 0, 0, 137,
		42, 1, 0, 0, 0, 138, 139, 5, 33, 0, 0, 139, 140, 5, 61, 0, 0, 140, 44,
		1, 0, 0, 0, 141, 142, 5, 40, 0, 0, 142, 46, 1, 0, 0, 0, 143, 144, 5, 41,
		0, 0, 144, 48, 1, 0, 0, 0, 145, 147, 7, 6, 0, 0, 146, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		150, 1, 0, 0, 0, 150, 151, 6, 24, 0, 0, 151, 50, 1, 0, 0, 0, 10, 0, 62,
		68, 74, 80, 82, 87, 89, 97, 148, 1, 6, 0, 0,
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

// ConditionExprLexerInit initializes any static state used to implement ConditionExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewConditionExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConditionExprLexerInit() {
	staticData := &ConditionExprLexerLexerStaticData
	staticData.once.Do(conditionexprlexerLexerInit)
}

// NewConditionExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewConditionExprLexer(input antlr.CharStream) *ConditionExprLexer {
	ConditionExprLexerInit()
	l := new(ConditionExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ConditionExprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ConditionExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ConditionExprLexer tokens.
const (
	ConditionExprLexerT__0       = 1
	ConditionExprLexerBOOLEAN    = 2
	ConditionExprLexerIDENTIFIER = 3
	ConditionExprLexerNUMBER     = 4
	ConditionExprLexerSTRING     = 5
	ConditionExprLexerPLUS       = 6
	ConditionExprLexerMINUS      = 7
	ConditionExprLexerMULTIPLY   = 8
	ConditionExprLexerDIVIDE     = 9
	ConditionExprLexerMODULO     = 10
	ConditionExprLexerNOT        = 11
	ConditionExprLexerAND        = 12
	ConditionExprLexerOR         = 13
	ConditionExprLexerGT         = 14
	ConditionExprLexerGE         = 15
	ConditionExprLexerLT         = 16
	ConditionExprLexerLE         = 17
	ConditionExprLexerEQ         = 18
	ConditionExprLexerNE         = 19
	ConditionExprLexerLPAREN     = 20
	ConditionExprLexerRPAREN     = 21
	ConditionExprLexerWS         = 22
)
