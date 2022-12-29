package main

import "fmt"

func main() {
	name := "胡康康"

	for i, s := range name {
		fmt.Printf("%d,%s", i, string(s))
	}
}
