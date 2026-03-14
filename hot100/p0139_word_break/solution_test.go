package p0139_word_break

import "testing"

func TestWordBreak(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet", "code"}) {
		t.Error("expected true")
	}
	if !wordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Error("expected true")
	}
	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Error("expected false")
	}
}
