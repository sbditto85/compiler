package lexer

import (
	"testing"
	"github.com/sbditto85/compiler/token"
)

func shouldNotError(t *testing.T, e error) {
	if e != nil {
		t.Fatal("Should Not have errored, but did!")
	}
}

func verifyToken(t *testing.T,tok *token.Token,tokTypeString string, tokType token.TokenType,tokLexeme string,tokLine int) {
	if tok == nil {
		t.Errorf("Token is nil, expected %s\n",tokTypeString)
		return
	}
	if tok.Type != tokType {
		t.Errorf("Expected Type %s %d, received %d",tokTypeString,tokType,tok.Type)
	}
	if tok.Lexeme != tokLexeme {
		t.Errorf("Expected Lexeme %s %s, received %s",tokTypeString,tokLexeme,tok.Lexeme)
	}
	if tok.Linenum != tokLine {
		t.Errorf("Expected Line Num %s %d, received %d",tokTypeString,tokLine,tok.Linenum)
	}
}

func testNextToken(l *Lexer,t *testing.T,tokTypeString string, tokType token.TokenType,tokLexeme string,tokLine int) {
	//t.Logf("%s ::: %s\n",tokTypeString,tokLexeme)
	tok,e := l.GetNextToken()
	shouldNotError(t,e)
	verifyToken(t,tok,tokTypeString,tokType,tokLexeme,tokLine)
	tok,e = l.GetCurrentToken()
	shouldNotError(t,e)
	verifyToken(t,tok,tokTypeString,tokType,tokLexeme,tokLine)
}

func TestParseValidMultipleLines(t *testing.T) {
	
	var s []string
	s = append(s,"class fun {")
	s = append(s,"\tint apple[] = {19348,1234};")
	s = append(s,"\tpublic function ap(lets,dothis){")
	s = append(s,"\t\tprint 'this' + lets + 'do' + dothis + 'yo' + apple;")
	s = append(s,"\t}")
	s = append(s,`}'c''\n'_`)

	l := NewLexer()
	
	l.LoadStrings(s)

	testNextToken(l,t,"Keyword",     token.Keyword,     "class",      0)
	testNextToken(l,t,"Identifier",  token.Identifier,  "fun",        0)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          0)
	//go to next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "int",        1)
	testNextToken(l,t,"Identifier",  token.Identifier,  "apple",      1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "[",          1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "]",          1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "=",          1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          1)
	testNextToken(l,t,"Number",      token.Number,      "19348",      1)
	testNextToken(l,t,"Punctuation", token.Punctuation, ",",          1)
	testNextToken(l,t,"Number",      token.Number,      "1234",       1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          1)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          1)
	//go to next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "public",     2)
	testNextToken(l,t,"Identifier",  token.Identifier,  "function",   2)
	testNextToken(l,t,"Identifier",  token.Identifier,  "ap",         2)
	testNextToken(l,t,"Symbol",      token.Symbol,      "(",          2)
	testNextToken(l,t,"Identifier",  token.Identifier,  "lets",       2)
	testNextToken(l,t,"Punctuation", token.Punctuation, ",",          2)
	testNextToken(l,t,"Identifier",  token.Identifier,  "dothis",     2)
	testNextToken(l,t,"Symbol",      token.Symbol,      ")",          2)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          2)
	//go to next line
	testNextToken(l,t,"Identifier",  token.Identifier,  "print",      3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Keyword",     token.Keyword,     "this",       3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "lets",       3)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "do",         3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "dothis",     3)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "yo",         3)
	testNextToken(l,t,"Punctuation", token.Punctuation, "'",          3)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "apple",      3)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          3)
	//go to next line
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          4)
	//go to next line
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          5)
	testNextToken(l,t,"Character",   token.Character,   "'c'",        5)
	testNextToken(l,t,"Character",   token.Character,   `'\n'`,       5)
	testNextToken(l,t,"Unknown",     token.Unknown,     "_",          5)
	//end of file (next line)
	testNextToken(l,t,"EOT",         token.EOT,         "",           6)
}

func TestReadFileTestfile(t *testing.T) {
	l := NewLexer()
	
	l.ReadFile("testfile")

	testNextToken(l,t,"Keyword",     token.Keyword,     "class",      1)
	testNextToken(l,t,"Identifier",  token.Identifier,  "funone",     1)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          1)
	//next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "private",    3)
	testNextToken(l,t,"Keyword",     token.Keyword,     "string",     3)
	testNextToken(l,t,"Identifier",  token.Identifier,  "s",          3)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          3)
	//next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "public",     4)
	testNextToken(l,t,"Keyword",     token.Keyword,     "int",        4)
	testNextToken(l,t,"Identifier",  token.Identifier,  "i",          4)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          4)
	//next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "public",     6)
	testNextToken(l,t,"Identifier",  token.Identifier,  "func",       6)
	testNextToken(l,t,"Symbol",      token.Symbol,      "(",          6)
	testNextToken(l,t,"Identifier",  token.Identifier,  "param",      6)
	testNextToken(l,t,"Punctuation", token.Punctuation, ",",          6)
	testNextToken(l,t,"Identifier",  token.Identifier,  "param",      6)
	testNextToken(l,t,"Symbol",      token.Symbol,      ")",          6)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          6)
	//next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "int",        7)
	testNextToken(l,t,"Identifier",  token.Identifier,  "y",          7)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          7)
	//next line
	testNextToken(l,t,"Keyword",     token.Keyword,     "while",      8)
	testNextToken(l,t,"Symbol",      token.Symbol,      "(",          8)
	testNextToken(l,t,"Identifier",  token.Identifier,  "param",      8)
	testNextToken(l,t,"Symbol",      token.Symbol,      ">",          8)
	testNextToken(l,t,"Identifier",  token.Identifier,  "y",          8)
	testNextToken(l,t,"Symbol",      token.Symbol,      ")",          8)
	testNextToken(l,t,"Symbol",      token.Symbol,      "{",          8)
	//next line
	testNextToken(l,t,"Identifier",  token.Identifier,  "y",          9)
	testNextToken(l,t,"Symbol",      token.Symbol,      "=",          9)
	testNextToken(l,t,"Identifier",  token.Identifier,  "param",      9)
	testNextToken(l,t,"Symbol",      token.Symbol,      "*",          9)
	testNextToken(l,t,"Identifier",  token.Identifier,  "y",          9)
	testNextToken(l,t,"Symbol",      token.Symbol,      "+",          9)
	testNextToken(l,t,"Identifier",  token.Identifier,  "param",      9)
	testNextToken(l,t,"Punctuation", token.Punctuation, ";",          9)
	//next line
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          10)
	//next line
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          11)
	//next line
	testNextToken(l,t,"Symbol",      token.Symbol,      "}",          12)
	//end of file (next line)
	testNextToken(l,t,"EOT",         token.EOT,         "",           13)	
}
