package parser

import "testing"

func TestTokenizer(t *testing.T) {
	arg := `
	{
		"user" : "string"
	}`
	tkn := Tokenizer{json: []rune(arg)}

	got := tkn.Exec()

	t.Log(got)
}

// arg := `
// {
// 	"user" : {
// 	  "name": "Taro",
// 	  "age": 30,
// 	  "languages": ["Japanese", "English"],
// 	  "active": true
// 	}
// }`
