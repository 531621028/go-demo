package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(convert("A", 1))
}
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result bytes.Buffer
	var size = 2*numRows - 2
	for row := 0; row < numRows; row++ {
		for nextIndex := row; nextIndex < len(s); {
			result.WriteString(string(s[nextIndex]))
			if row != 0 && row != numRows-1 {
				middle := nextIndex + size - 2*row
				if middle < len(s) {
					result.WriteString(string(s[middle]))
				}
			}
			nextIndex = nextIndex + size
		}
	}
	return result.String()
}
