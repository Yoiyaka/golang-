package p0211_design_add_and_search_words_data_structure

import "testing"

func TestWordDictionary(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	if wd.Search("pad") {
		t.Error("expected false for 'pad'")
	}
	if !wd.Search("bad") {
		t.Error("expected true for 'bad'")
	}
	if !wd.Search(".ad") {
		t.Error("expected true for '.ad'")
	}
	if !wd.Search("b..") {
		t.Error("expected true for 'b..'")
	}
}
