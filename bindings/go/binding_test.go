package tree_sitter_linkerscript_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_linkerscript "github.com/tree-sitter/tree-sitter-linkerscript/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_linkerscript.Language())
	if language == nil {
		t.Errorf("Error loading Linkerscript grammar")
	}
}
