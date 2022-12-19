package main

import "fmt"

func main() {
	//fmt.Println(isMatch("1", "1"))       // true
	//fmt.Println(isMatch("ab", "ab"))     // true
	//fmt.Println(isMatch("aab", "a*b"))   // true
	//fmt.Println(isMatch("aba", "a*.."))  // true
	//fmt.Println(isMatch("aab", "c*a*b")) //true
	//fmt.Println(isMatch("ab", ".*c"))    // false
	//fmt.Println(isMatch("aaa", "aaaa"))  // false
	//fmt.Println(isMatch("aaa", ".*")) // true
	//fmt.Println(isMatch("a", "ab*"))  // true
	//fmt.Println(isMatch("ab", ".*c")) // true
	fmt.Println(isMatch("bbbba", ".*a*a")) // true
}

func isMatch(s string, p string) bool {

	var sIndex = 0
	var pIndex = 0
	for ; pIndex < len(p); {
		if sIndex >= len(s) {
			break
		}
		if p[pIndex] == s[sIndex] || p[pIndex] == '.' {
			pIndex++
			sIndex++
			continue
		}
		if p[pIndex] == '*' {
			if pIndex > 0 && (s[sIndex] == p[pIndex-1] || p[pIndex-1] == '.') {
				return isMatch(s[sIndex:], p[pIndex+1:]) || // 匹配0个
					isMatch(s[sIndex+1:], p[pIndex+1:]) || // 匹配一个
					isMatch(s[sIndex+1:], p[pIndex-1:]) // 匹配多个
			} else {
				return isMatch(s[sIndex:], p[pIndex+1:])
			}
		}
		pIndex++
	}
	if sIndex >= len(s) && pIndex >= len(p) {
		return true
	} else if pIndex < len(p) {
		return isOptional(p[pIndex:])
	}
	return false
}

func isOptional(p string) bool {
	if len(p) == 1 {
		return p[0] == '*'
	} else if len(p) == 2 {
		return p[1] == '*'
	} else {
		return isOptional(p[2:])
	}
}
