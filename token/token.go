package token

type Token struct {
	Type    TokenType
	Lexeme  string
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

func GetTokToStringMap() map[TokenType]string {
	TokTypeTrans := make(map[TokenType]string)
	TokTypeTrans[Number] = "Number"
	TokTypeTrans[Character] = "Character"
	TokTypeTrans[Identifier] = "Identifier"
	TokTypeTrans[Punctuation] = "Punctuation"
	TokTypeTrans[Keyword] = "Keyword"
	TokTypeTrans[Symbol] = "Symbol"
	TokTypeTrans[Unknown] = "Unknown"
	TokTypeTrans[WhiteSpace] = "WhiteSpace"
	TokTypeTrans[EOT] = "EOT"
	return TokTypeTrans
}
