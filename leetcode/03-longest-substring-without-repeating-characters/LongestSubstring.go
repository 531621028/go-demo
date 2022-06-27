package main

import "fmt"

func main() {
	fmt.Print(lengthOfLongestSubstring("tmmzuxt"))
}

// 滑动窗口 O(N*K)
func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	var left = 0  // 符合条件字符串的最左边的元素（包含）
	var right = 0 // 符合条件字符串的最左边的元素（包含）
	var charSet = make(map[uint8]int)
	for right < len(s) {
		duplicateLocation, ok := charSet[s[right]]
		// 在字符串最左边拥有重复的字符串才认为重复
		// "tmmzuxt",当right=7的时候，left=2，duplicateLocation = 1,则不在最大重复的字符串中
		if !ok || duplicateLocation < left {
			charSet[s[right]] = right
			currentLength := right - left + 1
			if maxLength < currentLength {
				maxLength = currentLength
			}
		} else {
			// 更新最新重复字符的位置
			charSet[s[right]] = right
			// 左侧边界右移
			left = duplicateLocation + 1
		}
		right++
	}
	return maxLength
}

// 滑动窗口 O(N*K)
//func lengthOfLongestSubstring(s string) int {
//	var maxLength = 0
//	var left = 0
//	var right = 0
//	var charSet = make(map[uint8]int)
//	for right < len(s) {
//		charSet[s[right]]++
//		for charSet[s[right]] > 1 {
//			charSet[s[left]]--
//			left++
//		}
//		right++
//		if maxLength < right-left {
//			maxLength = right - left
//		}
//
//	}
//	return maxLength
//}

// 暴力法-使用hashSet减少查找重复字符串的时间复杂度 O(N*N)
//func lengthOfLongestSubstring(s string) int {
//	var maxLength = 0
//	for index := 0; index < len(s); index++ {
//		var length = 0
//		var charSet = make(map[uint8]bool)
//		for position := index; position < len(s); position++ {
//			if !charSet[s[position]] {
//				charSet[s[position]] = true
//				length++
//			} else {
//				break
//			}
//		}
//		if maxLength < length {
//			maxLength = length
//		}
//	}
//	return maxLength
//}
