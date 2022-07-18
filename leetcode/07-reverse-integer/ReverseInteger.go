package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(-1234 % 10)
	fmt.Print(reverse(1534236469))
}

func reverse(x int) int {
	positive := math.MaxInt32 / 10
	negative := math.MinInt32 / 10
	var result = 0
	for left := x; left != 0; {
		if result > positive || result < negative {
			return 0
		}
		result = result*10 + left%10
		left = left / 10
	}
	return result
}
