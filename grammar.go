package parser

import "github.com/GSMath/lexer/tokenizer"

const (
	Left  = -1
	Right = 1
)

type Node struct {
	token    tokenizer.Token
	branches []*Node
}

type ParserRule struct {
	direction int
	rule      []rune
	action    func([]*Node) *Node
}

func (rule ParserRule) Tokens() []tokenizer.Token {
	var tokens []tokenizer.Token
	token_func := tokenizer.TokenizeString(string(rule.rule))
	tokens = make([]tokenizer.Token, len(token_func))
	for i, f := range token_func {
		tokens[i] = f(tokenizer.GetToken)[0]
		if tokens[i].TokenType == tokenizer.Symbol {
			switch tokens[i].Symbolic {
			case "W":
				tokens[i].TokenType = tokenizer.Wildcard
			case "E":
				tokens[i].TokenType = tokenizer.Expression
			case "N":
				tokens[i].TokenType = tokenizer.Numeric
			}
		}
	}
	return tokens
}

func MakeRule(direction int, tokens string, fx func([]*Node) *Node) ParserRule {
	return ParserRule{
		direction: direction,
		rule:      []rune(tokens),
		action:    fx,
	}
}

var StandardGrammar = []ParserRule{
	MakeRule(Left, "E^E", parse_bin_op),
	MakeRule(Left, "E•E", parse_bin_op),
	MakeRule(Left, "N E", parse_implicit_bin_op),
	MakeRule(Left, "E N", parse_implicit_bin_op),
	MakeRule(Left, "E/E", parse_bin_op),
	MakeRule(Left, "E+E", parse_bin_op),
	MakeRule(Left, "E±E", parse_bin_op),
	MakeRule(Left, "E∓E", parse_bin_op),
	MakeRule(Left, "∑E", parse_un_op),
	MakeRule(Left, "∏E", parse_un_op),
	MakeRule(Left, "∂E", parse_un_op),
	MakeRule(Left, "√E", parse_un_op),
	MakeRule(Left, "E≈E", parse_bin_op),
	MakeRule(Left, "E=E", parse_bin_op),
	MakeRule(Left, "E≠E", parse_bin_op),
	MakeRule(Left, "E≡E", parse_bin_op),
	MakeRule(Left, "E≢E", parse_bin_op),
}
