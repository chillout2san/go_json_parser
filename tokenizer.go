package parser

import (
	"regexp"
)

type Tokenizer struct {
	json  []rune
	index int
}

func (tkn *Tokenizer) getNext() rune {
	tkn.index = tkn.index + 1
	return tkn.json[tkn.index]
}

func (tkn *Tokenizer) Exec() []Token {
	tks := make([]Token, 0, len(tkn.json))

	for i, char := range tkn.json {
		if i < tkn.index {
			continue
		}

		switch char {
		case '{':
			tkn.index++
			tks = append(tks, Token{
				Type:    LEFT_BRACE,
				Literal: "{",
			})
		case '}':
			tkn.index++
			tks = append(tks, Token{
				Type:    LEFT_BRACE,
				Literal: "}",
			})
		case '"':
			tkn.index++
			tks = append(tks, Token{
				Type:    DOUBLE_QUOTE,
				Literal: "\"",
			})
		case ':':
			tkn.index++
			tks = append(tks, Token{
				Type:    COLON,
				Literal: ":",
			})
		case '\r' | '\n' | '\t':
			tkn.index++
			continue
		// 記号でない時
		default:
			strRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
			if strRegex.MatchString(string(char)) {
				t := Token{
					Type:    String,
					Literal: string(char),
				}
				for {
					if strRegex.MatchString(string(tkn.getNext())) {
						t.Literal = t.Literal + string(char)
						continue
					}
					break
				}
				tks = append(tks, t)
				continue
			}
		}
	}

	return tks
}
