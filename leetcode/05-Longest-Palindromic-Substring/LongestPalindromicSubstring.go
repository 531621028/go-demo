package main

import (
	"bytes"
	"fmt"
)

func main() {
	result := longestPalindrome("babad")
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	expandStr := expandString(s)
	var longest string
	for index := 1; index < len(expandStr); index = index + 1 {
		palindrome := searchPalindrome(expandStr, index)
		if len(palindrome) > len(longest) {
			longest = palindrome
		}
	}
	return clearExpandString(longest)
}

func expandString(s string) string {
	var b bytes.Buffer
	b.WriteString("#")
	for index := 0; index < len(s); index++ {
		b.WriteString(string(s[index]))
		b.WriteString("#")
	}
	return b.String()
}

func clearExpandString(s string) string {
	var b bytes.Buffer
	for index := 0; index < len(s); index++ {
		if string(s[index]) != "#" {
			b.WriteString(string(s[index]))
		}
	}
	return b.String()
}

func searchPalindrome(expandStr string, index int) string {
	var palindromeStr = ""
	for left, right := index-1, index+1; left >= 0 && right < len(expandStr); left, right = left-1, right+1 {
		if expandStr[left] != expandStr[right] {
			break
		}
		palindromeStr = expandStr[left+1 : right]
	}
	return palindromeStr
}
