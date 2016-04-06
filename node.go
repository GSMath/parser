package parser

import "github.com/GoSym/lexer/tokenizer"

type Node struct {
	token    tokenizer.Token
	branches []*Node
}

func (n *Node) IsOperator() bool {
	return n.token.TokenType == tokenizer.Operator
}

func (n *Node) IsRealNumeric() bool {
	return n.token.TokenType == tokenizer.RealNumeric
}

func (n *Node) IsImagNumeric() bool {
	return n.token.TokenType == tokenizer.ImagNumeric
}

func (n *Node) IsSymbol() bool {
	return n.token.TokenType == tokenizer.Symbol
}

func (n *Node) Numeric() float64 {
	return n.token.Numeric
}

func (n *Node) Symbol() string {
	return n.token.Symbolic
}

func (n *Node) Operator() rune {
	return n.token.Operator
}

func (n *Node) AddBranch(branch *Node) {
	n.branches = append(n.branches, branch)
}

func (n *Node) BranchAtIndex(index int) *Node {
	if index >= len(n.branches) {
		return nil
	}
	return n.branches[index]
}

func (n *Node) Equal(rhs *Node) bool {
	var areEqual bool = false
	if n == rhs {
		areEqual = true
		goto EXIT
	}
	if n.token != rhs.token {
		goto EXIT
	}
	if len(n.branches) != len(rhs.branches) {
		goto EXIT
	}
	areEqual = true
	for i := 0; i < len(n.branches); i++ {
		if n.branches[i].Equal(rhs.branches[i]) == false {
			areEqual = false
			break
		}
	}
EXIT:
	return areEqual
}

func parse_un_op(nodes []*Node) *Node {
	operator, rhs := nodes[0], nodes[1]
	if operator.token.TokenType != tokenizer.Operator {
		return nil
	}
	operator.branches = []*Node{rhs}
	return operator
}

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
		if operator.token.Operator == '/' {
			operator.token.Operator = []rune("•")[0]
			exponent := tokenizer.Token{TokenType: tokenizer.RealNumeric, Numeric: -1}
			token := tokenizer.Token{TokenType: tokenizer.Operator, Operator: '^'}
			power := &Node{token: exponent, branches: []*Node{}}
			temp := Node{token: token, branches: []*Node{rhs, power}}
			rhs = &temp
			operator = parse_bin_op([]*Node{lhs, operator, rhs})
		}
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
			Operator:  []rune("•")[0],
		},
		branches: []*Node{},
	}
	return parse_bin_op([]*Node{lhs, &operator, rhs})
}
