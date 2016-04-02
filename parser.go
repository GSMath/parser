package parser

import "github.com/GSMath/lexer/tokenizer"

func check_left_rule(node_l int, nodes []*Node, rule ParserRule) (int, int, bool) {
	var start, end int
	var found bool
	var token tokenizer.Token
	rule_tokens := rule.Tokens()
	found = true
	token_l := len(rule_tokens)
	for i := 0; i < node_l-(token_l-1); i++ {
		found = true
		start = i
		for j := 0; j < token_l; j++ {
			token = nodes[i+j].token
			if token.TokenType == tokenizer.Operator &&
				len(nodes[i+j].branches) > 0 {
				token.TokenType = tokenizer.Expression
			}
			if token.Equivalent(rule_tokens[j]) == false {
				found = false
				break
			}
		}
		if found == true {
			end = i + token_l
			goto EXIT
		}
	}
	found = false
EXIT:
	return start, end, found
}

func check_right_rule(node_l int, nodes []*Node, rule ParserRule) (int, int, bool) {
	var start, end int
	var found bool
	var token tokenizer.Token
	rule_tokens := rule.Tokens()
	found = true
	token_l := len(rule_tokens)
	for i := node_l - 1; i >= token_l-1; i-- {
		end = i
		for j := 0; j < token_l; j++ {
			token = nodes[i+j].token
			if token.TokenType == tokenizer.Operator &&
				len(nodes[i+j].branches) > 0 {
				token.TokenType = tokenizer.Expression
			}
			if nodes[i-j].token.Equivalent(rule_tokens[token_l-j-1]) == false {
				found = false
				break
			}
		}
		if found == true {
			start = i - token_l
			goto EXIT
		}
		found = true
	}
	found = false
EXIT:
	return start, end, found
}

func check_rule(node_l int, nodes []*Node, rule ParserRule) (int, int, bool) {
	var start, end int
	var found bool
	if rule.direction == Left {
		start, end, found = check_left_rule(node_l, nodes, rule)
	}
	if rule.direction == Right {
		start, end, found = check_right_rule(node_l, nodes, rule)
	}
	return start, end, found
}

func MakeNode(token_func func(int) []tokenizer.Token, rules []ParserRule) *Node {
	var node Node
	token := token_func(tokenizer.GetToken)[0]
	switch token.TokenType {
	case tokenizer.Symbol:
		node.token = token
		node.branches = []*Node{}
	case tokenizer.RealNumeric:
		node.token = token
		node.branches = []*Node{}
	case tokenizer.ImagNumeric:
		node.token = token
		node.branches = []*Node{}
	case tokenizer.Operator:
		node.token = token
		node.branches = []*Node{}
	case tokenizer.Empty:
		node.token = token
		node.branches = []*Node{}
	case tokenizer.Subexpression:
		node = *(ParseString(token.Subexpression, rules))
	}
	return &node
}

func ParseString(expression string, rules []ParserRule) *Node {
	var node *Node
	tokenized := tokenizer.TokenizeString(expression)
	nodes := make([]*Node, len(tokenized))
	for i := 0; i < len(tokenized); i++ {
		nodes[i] = MakeNode(tokenized[i], rules)
	}
	length := len(nodes)
	for {
		var start int
		var end int
		var found bool
		var i int
		found = false
		for i = 0; i < len(rules); i++ {
			start, end, found = check_rule(length, nodes, rules[i])
			if found == true {
				break
			}
		}
		if found {
			node = rules[i].action(nodes[start:end])
			nodes[start] = node
			for j := 0; end+j < length; j++ {
				nodes[start+j+1] = nodes[end+j]
			}
			length -= end - (start + 1)
		} else {
			break
		}
	}
	if length > 1 {
		node = nil
	}
	return node
}
