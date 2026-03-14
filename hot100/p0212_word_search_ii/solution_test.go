package p0212_word_search_ii

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	got := findWords(board, words)
	sort.Strings(got)
	want := []string{"eat", "oath"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
