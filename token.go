package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LEFT_BRACE   TokenType = "left_brace"
	RIGHT_BRACE  TokenType = "right_brace"
	DOUBLE_QUOTE TokenType = "double_quote"
	COLON        TokenType = "colon"
	String       TokenType = "string"
)
