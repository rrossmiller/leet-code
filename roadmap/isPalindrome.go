package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))

}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := regexp.MustCompile(`\s`)
	s = r.ReplaceAllString(s, "")

	r = regexp.MustCompile(`[^a-z0-9]`)
	s = r.ReplaceAllString(s, "")

	cnt := utf8.RuneCountInString(s)
	for i := 0; i < cnt; i++ {
		if s[i] != s[cnt-i-1] {
			return false
		}
	}
	return true
}
