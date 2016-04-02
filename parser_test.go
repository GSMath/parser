package parser

import "testing"

func TestParseString(t *testing.T) {
	node := ParseString("3•(3^√2 + 2x)x", StandardGrammar)
	if node == nil {
		t.Error("Node is nil")
	}
	if len(node.branches) != 3 {
		t.Error("Expected 3 branches, found ", len(node.branches))
	}

}
