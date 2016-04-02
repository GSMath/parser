package parser

import "github.com/GSMath/lexer/tokenizer"

func parse_bin_op(nodes []*Node) *Node {
	lhs, operator, rhs := nodes[0], nodes[1], nodes[2]
	if operator.token.TokenType != tokenizer.Operator {
		return nil
	}
	if lhs.token.TokenType != tokenizer.Operator &&
		rhs.token.TokenType != tokenizer.Operator {
		operator.branches = []*Node{lhs, rhs}
		return operator
	}
	operator.branches = []*Node{}
	for i, side := range [...]*Node{lhs, rhs} {
		i++
		if operator.token.Operator == '+' ||
			operator.token.Operator == []rune("•")[0] {
			if side.token.Operator == operator.token.Operator {
				for j, branch := range side.branches {
					j++
					operator.branches = append(operator.branches, branch)
				}
			} else {
				operator.branches = append(operator.branches, side)
			}
		} else {

			operator.branches = append(operator.branches, side)
		}
	}
	return operator
}

func parse_implicit_bin_op(nodes []*Node) *Node {
	lhs, rhs := nodes[0], nodes[1]
	operator := Node{
		token: tokenizer.Token{
			TokenType: tokenizer.Operator,
			operator:  []rune("•"),
		},
		branches: []*Node{},
	}
	if operator.token.TokenType != tokenizer.Operator {
		return nil
	}
	if lhs.token.TokenType != tokenizer.Operator &&
		rhs.token.TokenType != tokenizer.Operator {
		operator.branches = []*Node{lhs, rhs}
		return operator
	}
	operator.branches = []*Node{}
	for i, side := range [...]*Node{lhs, rhs} {
		i++
		if operator.token.Operator == '+' ||
			operator.token.Operator == []rune("•")[0] {
			if side.token.Operator == operator.token.Operator {
				for j, branch := range side.branches {
					j++
					operator.branches = append(operator.branches, branch)
				}
			} else {
				operator.branches = append(operator.branches, side)
			}
		} else {

			operator.branches = append(operator.branches, side)
		}
	}
	return operator
}
