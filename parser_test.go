package parser

import "testing"

func TestParseString(t *testing.T) {
	node := ParseString("(3 ^ 2 • 2 • 2) + (4 • 2)", StandardGrammar)
	if node == nil {
		t.Error("Node is nil")
	}
	if len(node.branches) != 2 {
		t.Error("Expected 2 branches, found ", len(node.branches))
	}
}
