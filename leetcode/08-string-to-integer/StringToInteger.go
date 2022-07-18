package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	//fmt.Println(myAtoi("-4193"))
	//fmt.Println(myAtoi("    -4193"))
	//fmt.Println(myAtoi("4193 with words"))
	//fmt.Println(myAtoi("00000-42a1234"))
	//fmt.Println(myAtoi("  -0012a42"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("   -+5-"))
	fmt.Println(myAtoi("9223372036854775808"))
}

// "00000-42a1234" -> 0
// "  -0012a42" -> -12
// "   +0 123" -> 0
func myAtoi(s string) int {
	var validateInput bytes.Buffer
	isNegative := false
	hasValidateInput := false
	for index, ch := range []byte(s) {
		if ch == ' ' {
			if hasValidateInput {
				break
			} else {
				continue
			}
		}
		if ch == '-' || ch == '+' {
			if hasValidateInput {
				break
			}
			if ch == '-' {
				isNegative = true
			}
			if index+1 < len([]byte(s)) && !isDigital(s[index+1]) {
				return 0
			}
		} else if isDigital(ch) {
			hasValidateInput = true
			validateInput.WriteByte(ch)
		} else {
			if validateInput.Len() == 0 {
				return 0
			}
			break
		}
	}
	if validateInput.Len() == 0 {
		return 0
	}
	hasNoZeroDigital := false
	var result int64 = 0
	for _, ch := range []byte(validateInput.String()) {
		if ch == '0' && !hasNoZeroDigital {
			continue
		} else {
			hasNoZeroDigital = true
		}
		if isNegative {
			result = result*10 + int64(ch-'0')*-1
		} else {
			result = result*10 + int64(ch-'0')
		}
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		if result < math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(result)
}

func isDigital(i byte) bool {
	return i >= '0' && i <= '9'
}
