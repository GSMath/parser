package parser

import (
	"testing"

	"github.com/GSMath/lexer/tokenizer"
)

func TestMakeNode(t *testing.T) {
	token_string := "12ia*12"
	tokens := tokenizer.TokenizeString(token_string)
	for i, token := range tokens {
		node := MakeNode(token, []ParserRule{ParserRule{
			direction: Left,
			rule: []tokenizer.Token{tokenizer.Token{
				TokenType: tokenizer.Wildcard,
			}},
			action: nil,
		}})
		if node == nil {
			t.Error("Token ", token, " for token of ", token_string[i], " did not yield a Node.")
		}
		if node.token.TokenType != token.TokenType {
			t.Error("Node is alterning token type")
		}
	}

}

func TestParseString(t *testing.T) {
	node := ParseString("2•(3^√2 + 2x) + ∑ßi ≈ 3", StandardGrammar)
	if node == nil {
		t.Error("Node is nil")
	}
	if len(node.branches) != 2 {
		t.Error("Expected 3 branches, found ", len(node.branches))
	}
}
