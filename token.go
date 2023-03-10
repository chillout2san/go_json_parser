package parser

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LEFT_BRACE TokenType = iota
	RIGHT_BRACE
	DOUBLE_QUOTE
	COLON
	String
)
