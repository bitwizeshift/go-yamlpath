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
    "", "'$'", "'@'", "'.'", "'*'", "'..'", "'['", "']'", "':'", "'?'", 
    "'('", "')'", "','", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "", 
    "", "", "", "'null'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "NAME", "NUMBER", "STRING", "BOOLEAN", "NULL", "WS", "COMMENT",
  }
  staticData.RuleNames = []string{
    "yamlPath", "root", "current", "selector", "dotSelector", "bracketSelector", 
    "bracketExpression", "slice", "filter", "union", "expression", "subexpression", 
    "currentSubPath", "value", "quotedName", "wildcard",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 25, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 1, 1, 1, 1, 1, 2, 1, 
	2, 1, 3, 1, 3, 3, 3, 46, 8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 
	54, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 63, 8, 6, 1, 
	7, 3, 7, 66, 8, 7, 1, 7, 1, 7, 3, 7, 70, 8, 7, 1, 7, 1, 7, 3, 7, 74, 8, 
	7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 83, 8, 9, 1, 9, 1, 9, 
	1, 9, 3, 9, 88, 8, 9, 5, 9, 90, 8, 9, 10, 9, 12, 9, 93, 9, 9, 1, 10, 1, 
	10, 1, 10, 3, 10, 98, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 104, 8, 
	11, 1, 12, 1, 12, 5, 12, 108, 8, 12, 10, 12, 12, 12, 111, 9, 12, 1, 13, 
	1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 3, 2, 0, 4, 5, 19, 19, 1, 0, 
	13, 18, 1, 0, 20, 23, 120, 0, 32, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 4, 41, 
	1, 0, 0, 0, 6, 45, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 50, 1, 0, 0, 0, 12, 
	62, 1, 0, 0, 0, 14, 65, 1, 0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 82, 1, 0, 0, 
	0, 20, 94, 1, 0, 0, 0, 22, 103, 1, 0, 0, 0, 24, 105, 1, 0, 0, 0, 26, 112, 
	1, 0, 0, 0, 28, 114, 1, 0, 0, 0, 30, 116, 1, 0, 0, 0, 32, 36, 3, 2, 1, 
	0, 33, 35, 3, 6, 3, 0, 34, 33, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 
	1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 1, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 
	39, 40, 5, 1, 0, 0, 40, 3, 1, 0, 0, 0, 41, 42, 5, 2, 0, 0, 42, 5, 1, 0, 
	0, 0, 43, 46, 3, 8, 4, 0, 44, 46, 3, 10, 5, 0, 45, 43, 1, 0, 0, 0, 45, 
	44, 1, 0, 0, 0, 46, 7, 1, 0, 0, 0, 47, 48, 5, 3, 0, 0, 48, 49, 7, 0, 0, 
	0, 49, 9, 1, 0, 0, 0, 50, 53, 5, 6, 0, 0, 51, 54, 3, 12, 6, 0, 52, 54, 
	3, 30, 15, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 
	0, 55, 56, 5, 7, 0, 0, 56, 11, 1, 0, 0, 0, 57, 63, 3, 28, 14, 0, 58, 63, 
	5, 20, 0, 0, 59, 63, 3, 14, 7, 0, 60, 63, 3, 16, 8, 0, 61, 63, 3, 18, 9, 
	0, 62, 57, 1, 0, 0, 0, 62, 58, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 62, 60, 
	1, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 13, 1, 0, 0, 0, 64, 66, 5, 20, 0, 0, 
	65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 5, 
	8, 0, 0, 68, 70, 5, 20, 0, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 
	73, 1, 0, 0, 0, 71, 72, 5, 8, 0, 0, 72, 74, 5, 20, 0, 0, 73, 71, 1, 0, 
	0, 0, 73, 74, 1, 0, 0, 0, 74, 15, 1, 0, 0, 0, 75, 76, 5, 9, 0, 0, 76, 77, 
	5, 10, 0, 0, 77, 78, 3, 20, 10, 0, 78, 79, 5, 11, 0, 0, 79, 17, 1, 0, 0, 
	0, 80, 83, 3, 28, 14, 0, 81, 83, 5, 20, 0, 0, 82, 80, 1, 0, 0, 0, 82, 81, 
	1, 0, 0, 0, 83, 91, 1, 0, 0, 0, 84, 87, 5, 12, 0, 0, 85, 88, 3, 28, 14, 
	0, 86, 88, 5, 20, 0, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 90, 
	1, 0, 0, 0, 89, 84, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 
	91, 92, 1, 0, 0, 0, 92, 19, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 97, 3, 
	22, 11, 0, 95, 96, 7, 1, 0, 0, 96, 98, 3, 22, 11, 0, 97, 95, 1, 0, 0, 0, 
	97, 98, 1, 0, 0, 0, 98, 21, 1, 0, 0, 0, 99, 104, 3, 26, 13, 0, 100, 104, 
	5, 19, 0, 0, 101, 104, 3, 0, 0, 0, 102, 104, 3, 24, 12, 0, 103, 99, 1, 
	0, 0, 0, 103, 100, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0, 
	0, 104, 23, 1, 0, 0, 0, 105, 109, 3, 4, 2, 0, 106, 108, 3, 6, 3, 0, 107, 
	106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 
	1, 0, 0, 0, 110, 25, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 7, 2, 
	0, 0, 113, 27, 1, 0, 0, 0, 114, 115, 5, 21, 0, 0, 115, 29, 1, 0, 0, 0, 
	116, 117, 5, 4, 0, 0, 117, 31, 1, 0, 0, 0, 13, 36, 45, 53, 62, 65, 69, 
	73, 82, 87, 91, 97, 103, 109,
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
	yamlpathParserEOF = antlr.TokenEOF
	yamlpathParserT__0 = 1
	yamlpathParserT__1 = 2
	yamlpathParserT__2 = 3
	yamlpathParserT__3 = 4
	yamlpathParserT__4 = 5
	yamlpathParserT__5 = 6
	yamlpathParserT__6 = 7
	yamlpathParserT__7 = 8
	yamlpathParserT__8 = 9
	yamlpathParserT__9 = 10
	yamlpathParserT__10 = 11
	yamlpathParserT__11 = 12
	yamlpathParserT__12 = 13
	yamlpathParserT__13 = 14
	yamlpathParserT__14 = 15
	yamlpathParserT__15 = 16
	yamlpathParserT__16 = 17
	yamlpathParserT__17 = 18
	yamlpathParserNAME = 19
	yamlpathParserNUMBER = 20
	yamlpathParserSTRING = 21
	yamlpathParserBOOLEAN = 22
	yamlpathParserNULL = 23
	yamlpathParserWS = 24
	yamlpathParserCOMMENT = 25
)

// yamlpathParser rules.
const (
	yamlpathParserRULE_yamlPath = 0
	yamlpathParserRULE_root = 1
	yamlpathParserRULE_current = 2
	yamlpathParserRULE_selector = 3
	yamlpathParserRULE_dotSelector = 4
	yamlpathParserRULE_bracketSelector = 5
	yamlpathParserRULE_bracketExpression = 6
	yamlpathParserRULE_slice = 7
	yamlpathParserRULE_filter = 8
	yamlpathParserRULE_union = 9
	yamlpathParserRULE_expression = 10
	yamlpathParserRULE_subexpression = 11
	yamlpathParserRULE_currentSubPath = 12
	yamlpathParserRULE_value = 13
	yamlpathParserRULE_quotedName = 14
	yamlpathParserRULE_wildcard = 15
)

// IYamlPathContext is an interface to support dynamic dispatch.
type IYamlPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Root() IRootContext
	AllSelector() []ISelectorContext
	Selector(i int) ISelectorContext

	// IsYamlPathContext differentiates from other interfaces.
	IsYamlPathContext()
}

type YamlPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYamlPathContext() *YamlPathContext {
	var p = new(YamlPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_yamlPath
	return p
}

func InitEmptyYamlPathContext(p *YamlPathContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_yamlPath
}

func (*YamlPathContext) IsYamlPathContext() {}

func NewYamlPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YamlPathContext {
	var p = new(YamlPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_yamlPath

	return p
}

func (s *YamlPathContext) GetParser() antlr.Parser { return s.parser }

func (s *YamlPathContext) Root() IRootContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRootContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRootContext)
}

func (s *YamlPathContext) AllSelector() []ISelectorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorContext); ok {
			len++
		}
	}

	tst := make([]ISelectorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorContext); ok {
			tst[i] = t.(ISelectorContext)
			i++
		}
	}

	return tst
}

func (s *YamlPathContext) Selector(i int) ISelectorContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *YamlPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YamlPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) YamlPath() (localctx IYamlPathContext) {
	localctx = NewYamlPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, yamlpathParserRULE_yamlPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Root()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == yamlpathParserT__2 || _la == yamlpathParserT__5 {
		{
			p.SetState(33)
			p.Selector()
		}


		p.SetState(38)
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


// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }
func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, yamlpathParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(yamlpathParserT__0)
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


// ICurrentContext is an interface to support dynamic dispatch.
type ICurrentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCurrentContext differentiates from other interfaces.
	IsCurrentContext()
}

type CurrentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentContext() *CurrentContext {
	var p = new(CurrentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_current
	return p
}

func InitEmptyCurrentContext(p *CurrentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_current
}

func (*CurrentContext) IsCurrentContext() {}

func NewCurrentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentContext {
	var p = new(CurrentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_current

	return p
}

func (s *CurrentContext) GetParser() antlr.Parser { return s.parser }
func (s *CurrentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Current() (localctx ICurrentContext) {
	localctx = NewCurrentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, yamlpathParserRULE_current)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(yamlpathParserT__1)
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


// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DotSelector() IDotSelectorContext
	BracketSelector() IBracketSelectorContext

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) DotSelector() IDotSelectorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDotSelectorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDotSelectorContext)
}

func (s *SelectorContext) BracketSelector() IBracketSelectorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketSelectorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketSelectorContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, yamlpathParserRULE_selector)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.DotSelector()
		}


	case yamlpathParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.BracketSelector()
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


// IDotSelectorContext is an interface to support dynamic dispatch.
type IDotSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsDotSelectorContext differentiates from other interfaces.
	IsDotSelectorContext()
}

type DotSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotSelectorContext() *DotSelectorContext {
	var p = new(DotSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_dotSelector
	return p
}

func InitEmptyDotSelectorContext(p *DotSelectorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_dotSelector
}

func (*DotSelectorContext) IsDotSelectorContext() {}

func NewDotSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotSelectorContext {
	var p = new(DotSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_dotSelector

	return p
}

func (s *DotSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *DotSelectorContext) NAME() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNAME, 0)
}

func (s *DotSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) DotSelector() (localctx IDotSelectorContext) {
	localctx = NewDotSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, yamlpathParserRULE_dotSelector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(yamlpathParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 524336) != 0)) {
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


// IBracketSelectorContext is an interface to support dynamic dispatch.
type IBracketSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BracketExpression() IBracketExpressionContext
	Wildcard() IWildcardContext

	// IsBracketSelectorContext differentiates from other interfaces.
	IsBracketSelectorContext()
}

type BracketSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketSelectorContext() *BracketSelectorContext {
	var p = new(BracketSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketSelector
	return p
}

func InitEmptyBracketSelectorContext(p *BracketSelectorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketSelector
}

func (*BracketSelectorContext) IsBracketSelectorContext() {}

func NewBracketSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketSelectorContext {
	var p = new(BracketSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_bracketSelector

	return p
}

func (s *BracketSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketSelectorContext) BracketExpression() IBracketExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketExpressionContext)
}

func (s *BracketSelectorContext) Wildcard() IWildcardContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWildcardContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWildcardContext)
}

func (s *BracketSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) BracketSelector() (localctx IBracketSelectorContext) {
	localctx = NewBracketSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, yamlpathParserRULE_bracketSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(yamlpathParserT__5)
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

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserT__7, yamlpathParserT__8, yamlpathParserNUMBER, yamlpathParserSTRING:
		{
			p.SetState(51)
			p.BracketExpression()
		}


	case yamlpathParserT__3:
		{
			p.SetState(52)
			p.Wildcard()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(55)
		p.Match(yamlpathParserT__6)
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


// IBracketExpressionContext is an interface to support dynamic dispatch.
type IBracketExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QuotedName() IQuotedNameContext
	NUMBER() antlr.TerminalNode
	Slice() ISliceContext
	Filter() IFilterContext
	Union() IUnionContext

	// IsBracketExpressionContext differentiates from other interfaces.
	IsBracketExpressionContext()
}

type BracketExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketExpressionContext() *BracketExpressionContext {
	var p = new(BracketExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketExpression
	return p
}

func InitEmptyBracketExpressionContext(p *BracketExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_bracketExpression
}

func (*BracketExpressionContext) IsBracketExpressionContext() {}

func NewBracketExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_bracketExpression

	return p
}

func (s *BracketExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketExpressionContext) QuotedName() IQuotedNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedNameContext)
}

func (s *BracketExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, 0)
}

func (s *BracketExpressionContext) Slice() ISliceContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *BracketExpressionContext) Filter() IFilterContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *BracketExpressionContext) Union() IUnionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) BracketExpression() (localctx IBracketExpressionContext) {
	localctx = NewBracketExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, yamlpathParserRULE_bracketExpression)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.QuotedName()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Slice()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(60)
			p.Filter()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)
			p.Union()
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


// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_slice
	return p
}

func InitEmptySliceContext(p *SliceContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_slice
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserNUMBER)
}

func (s *SliceContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, i)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Slice() (localctx ISliceContext) {
	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, yamlpathParserRULE_slice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == yamlpathParserNUMBER {
		{
			p.SetState(64)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(67)
		p.Match(yamlpathParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == yamlpathParserNUMBER {
		{
			p.SetState(68)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == yamlpathParserT__7 {
		{
			p.SetState(71)
			p.Match(yamlpathParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
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


// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_filter
	return p
}

func InitEmptyFilterContext(p *FilterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_filter
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, yamlpathParserRULE_filter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(yamlpathParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
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
		p.Expression()
	}
	{
		p.SetState(78)
		p.Match(yamlpathParserT__10)
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


// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQuotedName() []IQuotedNameContext
	QuotedName(i int) IQuotedNameContext
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	var p = new(UnionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_union
	return p
}

func InitEmptyUnionContext(p *UnionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_union
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	var p = new(UnionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) AllQuotedName() []IQuotedNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQuotedNameContext); ok {
			len++
		}
	}

	tst := make([]IQuotedNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQuotedNameContext); ok {
			tst[i] = t.(IQuotedNameContext)
			i++
		}
	}

	return tst
}

func (s *UnionContext) QuotedName(i int) IQuotedNameContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedNameContext)
}

func (s *UnionContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(yamlpathParserNUMBER)
}

func (s *UnionContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, i)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Union() (localctx IUnionContext) {
	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, yamlpathParserRULE_union)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserSTRING:
		{
			p.SetState(80)
			p.QuotedName()
		}


	case yamlpathParserNUMBER:
		{
			p.SetState(81)
			p.Match(yamlpathParserNUMBER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == yamlpathParserT__11 {
		{
			p.SetState(84)
			p.Match(yamlpathParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case yamlpathParserSTRING:
			{
				p.SetState(85)
				p.QuotedName()
			}


		case yamlpathParserNUMBER:
			{
				p.SetState(86)
				p.Match(yamlpathParserNUMBER)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}


		p.SetState(93)
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


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSubexpression() []ISubexpressionContext
	Subexpression(i int) ISubexpressionContext

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

func InitEmptyExpressionContext(p *ExpressionContext)  {
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

func (s *ExpressionContext) AllSubexpression() []ISubexpressionContext {
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

func (s *ExpressionContext) Subexpression(i int) ISubexpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubexpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, yamlpathParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Subexpression()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 516096) != 0) {
		{
			p.SetState(95)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 516096) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(96)
			p.Subexpression()
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


// ISubexpressionContext is an interface to support dynamic dispatch.
type ISubexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	NAME() antlr.TerminalNode
	YamlPath() IYamlPathContext
	CurrentSubPath() ICurrentSubPathContext

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

func InitEmptySubexpressionContext(p *SubexpressionContext)  {
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

func (s *SubexpressionContext) Value() IValueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SubexpressionContext) NAME() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNAME, 0)
}

func (s *SubexpressionContext) YamlPath() IYamlPathContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYamlPathContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYamlPathContext)
}

func (s *SubexpressionContext) CurrentSubPath() ICurrentSubPathContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrentSubPathContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrentSubPathContext)
}

func (s *SubexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Subexpression() (localctx ISubexpressionContext) {
	localctx = NewSubexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, yamlpathParserRULE_subexpression)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case yamlpathParserNUMBER, yamlpathParserSTRING, yamlpathParserBOOLEAN, yamlpathParserNULL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Value()
		}


	case yamlpathParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(yamlpathParserNAME)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case yamlpathParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.YamlPath()
		}


	case yamlpathParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.CurrentSubPath()
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


// ICurrentSubPathContext is an interface to support dynamic dispatch.
type ICurrentSubPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Current() ICurrentContext
	AllSelector() []ISelectorContext
	Selector(i int) ISelectorContext

	// IsCurrentSubPathContext differentiates from other interfaces.
	IsCurrentSubPathContext()
}

type CurrentSubPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentSubPathContext() *CurrentSubPathContext {
	var p = new(CurrentSubPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_currentSubPath
	return p
}

func InitEmptyCurrentSubPathContext(p *CurrentSubPathContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_currentSubPath
}

func (*CurrentSubPathContext) IsCurrentSubPathContext() {}

func NewCurrentSubPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentSubPathContext {
	var p = new(CurrentSubPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_currentSubPath

	return p
}

func (s *CurrentSubPathContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrentSubPathContext) Current() ICurrentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrentContext)
}

func (s *CurrentSubPathContext) AllSelector() []ISelectorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorContext); ok {
			len++
		}
	}

	tst := make([]ISelectorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorContext); ok {
			tst[i] = t.(ISelectorContext)
			i++
		}
	}

	return tst
}

func (s *CurrentSubPathContext) Selector(i int) ISelectorContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *CurrentSubPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentSubPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) CurrentSubPath() (localctx ICurrentSubPathContext) {
	localctx = NewCurrentSubPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, yamlpathParserRULE_currentSubPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Current()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == yamlpathParserT__2 || _la == yamlpathParserT__5 {
		{
			p.SetState(106)
			p.Selector()
		}


		p.SetState(111)
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


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNUMBER, 0)
}

func (s *ValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(yamlpathParserBOOLEAN, 0)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(yamlpathParserNULL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, yamlpathParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 15728640) != 0)) {
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


// IQuotedNameContext is an interface to support dynamic dispatch.
type IQuotedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsQuotedNameContext differentiates from other interfaces.
	IsQuotedNameContext()
}

type QuotedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedNameContext() *QuotedNameContext {
	var p = new(QuotedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_quotedName
	return p
}

func InitEmptyQuotedNameContext(p *QuotedNameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_quotedName
}

func (*QuotedNameContext) IsQuotedNameContext() {}

func NewQuotedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedNameContext {
	var p = new(QuotedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_quotedName

	return p
}

func (s *QuotedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedNameContext) STRING() antlr.TerminalNode {
	return s.GetToken(yamlpathParserSTRING, 0)
}

func (s *QuotedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) QuotedName() (localctx IQuotedNameContext) {
	localctx = NewQuotedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, yamlpathParserRULE_quotedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(yamlpathParserSTRING)
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


// IWildcardContext is an interface to support dynamic dispatch.
type IWildcardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWildcardContext differentiates from other interfaces.
	IsWildcardContext()
}

type WildcardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWildcardContext() *WildcardContext {
	var p = new(WildcardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_wildcard
	return p
}

func InitEmptyWildcardContext(p *WildcardContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = yamlpathParserRULE_wildcard
}

func (*WildcardContext) IsWildcardContext() {}

func NewWildcardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WildcardContext {
	var p = new(WildcardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = yamlpathParserRULE_wildcard

	return p
}

func (s *WildcardContext) GetParser() antlr.Parser { return s.parser }
func (s *WildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WildcardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *yamlpathParser) Wildcard() (localctx IWildcardContext) {
	localctx = NewWildcardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, yamlpathParserRULE_wildcard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(yamlpathParserT__3)
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


