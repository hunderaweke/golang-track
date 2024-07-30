package main

import (
	"strings"
)

func countWords(s string) map[string]int {
	cnt := make(map[string]int)
	words := strings.Split(s, " ")
	for _, w := range words {
		cnt[strings.ToLower(w)]++
	}
	return cnt
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	rs := Reverse(s)
	return strings.Compare(s, rs) == 0
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
