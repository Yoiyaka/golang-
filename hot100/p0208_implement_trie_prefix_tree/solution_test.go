package p0208_implement_trie_prefix_tree

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("expected to find 'apple'")
	}
	if trie.Search("app") {
		t.Error("expected not to find 'app'")
	}
	if !trie.StartsWith("app") {
		t.Error("expected prefix 'app' to exist")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("expected to find 'app' after insert")
	}
}
