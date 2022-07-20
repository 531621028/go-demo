package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var result = 0
	var tmp = x
	for tmp != 0 {
		result = result*10 + tmp%10
		tmp = tmp / 10
	}
	return x == result

}
