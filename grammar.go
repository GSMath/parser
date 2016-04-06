package parser

import "github.com/GoSym/lexer/tokenizer"

const (
	Left  = -1
	Right = 1
)

type ParserRule struct {
	direction int
	rule      []tokenizer.Token
	action    func([]*Node) *Node
}

func rule_tokens(rule_string []rune) []tokenizer.Token {
	var tokens []tokenizer.Token
	tokens = tokenizer.TokenizeString(string(rule_string))
	for i := 0; i < len(tokens); i++ {
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
		i++
	}
	return tokens
}

func MakeRule(direction int, rule_string string, fx func([]*Node) *Node) ParserRule {
	var rule ParserRule = ParserRule{direction: direction, action: fx}
	rule.rule = rule_tokens([]rune(rule_string))
	return rule
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
