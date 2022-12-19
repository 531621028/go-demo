package main

type MethodFunc func(int, int) int

func (f MethodFunc) Add(x, y int) int {
	return f(x, y)
}
