package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sbditto85/compiler/token"
	"os"
	"regexp"
)

var WhiteSpace = regexp.MustCompile(`^(\s)+(.*)`)
var Number = regexp.MustCompile(`^([-]?\d+)(\s)*(.*)`) // only supports integers
var Character = regexp.MustCompile(`^('\\?.')(\s*)(.*)`)
var Identifier = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9_]{0,79})(\s*)(.*)`) // make sure the next 79 chars isn't something else
var Punctuation = regexp.MustCompile(`^([;:,'"])(\s)*(.*)`)
var Keyword = regexp.MustCompile(`^(atoi|bool|class|char|cin|cout|else|false|if|int|itoa|main|new|null|object|public|private|return|string|this|true|void|while)(.*)`)
var CantFollowKeyword = regexp.MustCompile(`^[A-Za-z0-9]`)
var ClearWhiteSpace = regexp.MustCompile(`^(\s*)(.*)`)
var Symbol = regexp.MustCompile(`^([-+*/]|[<>=!]{1,2}|[&|]{2}|[\[\]]|[{}()])(\s)*(.*)`)
var Comment = regexp.MustCompile(`^//.*$`)

//for comparing if the line is done or not
var emptyString = []byte("")

type Lexer struct {
	file        []string
	cur         *token.Token
	peek        *token.Token
	curline     []byte
	curFullLine string
	curlinenum  int
}

func NewLexer() *Lexer {
	l := Lexer{}
	return &l
}

func (l *Lexer) GetCurFullLine() string {
	return l.curFullLine
}

//Puts the contents of file in to the file string slice
//stolen and modified from http://stackoverflow.com/a/18479916/706882
func (l *Lexer) ReadFile(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.file = append(l.file, scanner.Text())
	}

	if len(l.file) > 0 {
		l.curline = []byte(l.file[0])
	}

	return scanner.Err()
}

func (l *Lexer) LoadStrings(s []string) {
	l.file = s
	if len(s) > 0 {
		l.curline = []byte(s[0])
	}
}

func (l *Lexer) GetCurrentToken() (*token.Token, error) {
	if l.cur != nil {
		return l.cur, nil
	}
	t, err := l.GetNextToken(false) //if we haven't then get one
	return t, err
}

func testType(l *Lexer, r *regexp.Regexp, tokType token.TokenType, isPeek bool) (*token.Token, bool) {
	if res := r.FindSubmatch(l.curline); len(res) == 4 {
		tok := &token.Token{}
		tok.Type = tokType
		tok.Lexeme = string(res[1])
		tok.Linenum = l.curlinenum
		if isPeek {
			l.peek = tok
		} else {
			l.cur = tok
			l.curline = res[3]
		}
		return tok, true
	}
	return nil, false
}

func testKeyword(l *Lexer, r *regexp.Regexp, tokType token.TokenType, isPeek bool) (*token.Token, bool) {
	if res := r.FindSubmatch(l.curline); len(res) == 3 && !CantFollowKeyword.Match(res[2]) {
		if res2 := ClearWhiteSpace.FindSubmatch(res[2]); len(res2) == 3 {
			tok := &token.Token{}
			tok.Type = tokType
			tok.Lexeme = string(res[1])
			tok.Linenum = l.curlinenum
			if isPeek {
				l.peek = tok
			} else {
				l.cur = tok
				l.curline = res2[2]
			}
			return tok, true
		}
	}
	return nil, false
}

func testEOT(l *Lexer) (*token.Token, error) {
	if len(l.curline) == 0 {
		l.curlinenum++
		if len(l.file) > l.curlinenum {
			l.curline = []byte(l.file[l.curlinenum])
		}
	}
	return nil, fmt.Errorf("")
}

func (l *Lexer) GetNextToken(expectSymbol bool) (*token.Token, error) {
	if res := WhiteSpace.FindSubmatch(l.curline); len(res) == 3 {
		l.curline = res[2]
		t, e := l.GetNextToken(expectSymbol)
		return t, e
	}

	//if the line is blank get the next one if nothing then return EOT
	if bytes.Equal(l.curline, emptyString) || Comment.Match(l.curline) {
		//can we get a new line?
		if l.loadNextLine() {
			tok, e := l.GetNextToken(expectSymbol)
			return tok, e
		}
		tok := &token.Token{Type: token.EOT, Lexeme: "", Linenum: l.curlinenum}
		l.cur = tok
		return tok, nil
	}

	if tok, found := testKeyword(l, Keyword, token.Keyword, false); found {
		return tok, nil
	}
	if tok, found := testType(l, Identifier, token.Identifier, false); found {
		return tok, nil
	}
	if tok, found := testType(l, Character, token.Character, false); found {
		return tok, nil
	}
	if expectSymbol {
		if tok, found := testType(l, Symbol, token.Symbol, false); found {
			return tok, nil
		}
		if tok, found := testType(l, Number, token.Number, false); found {
			return tok, nil
		}
	} else {
		if tok, found := testType(l, Number, token.Number, false); found {
			return tok, nil
		}
		if tok, found := testType(l, Symbol, token.Symbol, false); found {
			return tok, nil
		}
	}
	if tok, found := testType(l, Punctuation, token.Punctuation, false); found {
		return tok, nil
	}

	//Unkown is one character at a time
	if len(l.curline) > 0 {
		tok := &token.Token{Type: token.Unknown, Lexeme: string(l.curline[0]), Linenum: l.curlinenum}
		l.cur = tok
		l.curline = l.curline[1:]
		return tok, nil
	}

	//REALLY should never get here, but hey just in case right?
	fmt.Printf("%s\n", l.curline)
	fmt.Printf("%v\n", Identifier.Match(l.curline))
	fmt.Printf("%v\n", Identifier.FindSubmatch(l.curline))

	return nil, fmt.Errorf("Could not process Token")
}

func (l *Lexer) PeekNextToken() (*token.Token, error) {
	if res := WhiteSpace.FindSubmatch(l.curline); len(res) == 3 {
		l.curline = res[2]
		t, e := l.PeekNextToken()
		return t, e
	}

	//if the line is blank get the next one if nothing then return EOT
	if bytes.Equal(l.curline, emptyString) || Comment.Match(l.curline) {
		//can we get a new line?
		if l.loadNextLine() {
			tok, e := l.PeekNextToken()
			return tok, e
		}
		tok := &token.Token{Type: token.EOT, Lexeme: "", Linenum: l.curlinenum}
		l.peek = tok
		return tok, nil
	}

	if tok, found := testKeyword(l, Keyword, token.Keyword, true); found {
		return tok, nil
	}
	if tok, found := testType(l, Identifier, token.Identifier, true); found {
		return tok, nil
	}
	if tok, found := testType(l, Character, token.Character, true); found {
		return tok, nil
	}
	if tok, found := testType(l, Number, token.Number, true); found {
		return tok, nil
	}
	if tok, found := testType(l, Symbol, token.Symbol, true); found {
		return tok, nil
	}
	if tok, found := testType(l, Punctuation, token.Punctuation, true); found {
		return tok, nil
	}

	//Unkown is one character at a time
	if len(l.curline) > 0 {
		tok := &token.Token{Type: token.Unknown, Lexeme: string(l.curline[0]), Linenum: l.curlinenum}
		l.peek = tok
		return tok, nil
	}

	//REALLY should never get here, but hey just in case right?
	fmt.Printf("%s\n", l.curline)
	fmt.Printf("%v\n", Identifier.Match(l.curline))
	fmt.Printf("%v\n", Identifier.FindSubmatch(l.curline))

	return nil, fmt.Errorf("Could not process Token")
}

func (l *Lexer) loadNextLine() bool {
	l.curlinenum++
	if len(l.file) > l.curlinenum {
		l.curline = []byte(l.file[l.curlinenum])
		l.curFullLine = l.file[l.curlinenum]
		return true
	}
	return false
}
