package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Print("Hi：" + m)
}

func main() {
	m := MyType("hukangkang")
	m.sayHi()
}
