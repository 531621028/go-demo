// 同一个目录下的文件只能是一个包名
package main

// chapter4/sources/function_as_first_class_citizen_3.go

type BinaryAdder interface {
	Add(int, int) int
}

type MyAdderFunc func(int, int) int

//func (f MyAdderFunc) Add(x, y int) int {
//	return f(x, y)
//}

// 方法定义要与类型定义放在同一个包内
func (f MyAdderFunc) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func test(fn func() int) int {
	return fn()
}

// 定义函数类型。
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

// func main() {
// 	var i BinaryAdder = MyAdderFunc(MyAdd)
// 	fmt.Println(i.Add(5, 6))
// 	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

// 	s2 := format(func(s string, x, y int) string {
// 		return fmt.Sprintf(s, x, y)
// 	}, "%d, %d", 10, 20)

// 	println(s1, s2)
// }
