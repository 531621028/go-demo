package main

import (
	// demo module的名称，/headfirst/magazine对应module下的路径
	"demo/headfirst/magazine"
	"fmt"
)

func main() {
	s := new(magazine.Subscriber)
	fmt.Print(s)
}
