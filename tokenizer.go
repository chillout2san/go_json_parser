package parser

import (
	"regexp"
)

type Tokenizer struct {
	json    []rune
	counter int
}

var strRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

func (tkn *Tokenizer) getNext() rune {
	next := tkn.counter + 1
	return tkn.json[next]
}

func (tkn *Tokenizer) Exec() []Token {
	tks := make([]Token, 0, len(tkn.json))

	for i, char := range tkn.json {
		if i < tkn.counter {
			continue
		}

		switch char {
		case '{':
			tkn.counter++
			tks = append(tks, Token{
				Type:    LEFT_BRACE,
				Literal: "{",
			})
		case '}':
			tkn.counter++
			tks = append(tks, Token{
				Type:    LEFT_BRACE,
				Literal: "}",
			})
		case '"':
			tkn.counter++
			tks = append(tks, Token{
				Type:    DOUBLE_QUOTE,
				Literal: "\"",
			})
		case ':':
			tkn.counter++
			tks = append(tks, Token{
				Type:    COLON,
				Literal: ":",
			})
		case ' ', '\r', '\n', '\t':
			tkn.counter++
			continue
		default:
			if strRegex.MatchString(string(char)) {
				t := Token{
					Type:    String,
					Literal: string(char),
				}
				for {
					nc := tkn.getNext()
					if strRegex.MatchString(string(nc)) {
						tkn.counter++
						t.Literal = t.Literal + string(nc)
						continue
					}
					break
				}
				tkn.counter++
				tks = append(tks, t)
				continue
			}
			continue
		}
	}

	return tks
}
