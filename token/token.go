package token

type Token struct {
	Type TokenType
	Lexeme string
	Linenum int
}

type TokenType int

const (
	Number TokenType = iota
	Character
	Identifier
	Punctuation
	Keyword
	Symbol
	Unknown
	WhiteSpace
	EOT
)
