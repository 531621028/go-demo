package main

import "fmt"

func main() {
	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	// len(nums) 获取nums的长度
	// make用来构建指定的类型
	// 映射两种初始换的方式
	// 1) make(map[int]int, len(nums))
	// 2) map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	var result = make(map[int]int, len(nums))
	for index, num := range nums {
		// value是map中的值，exist表示对应的key是否存在值
		value, exist := result[target-num]
		//也能像下面这样再一个语句中执行
		//if j, ok := result[target-num]; ok {
		if exist {
			return []int{value, index}
		}
		result[num] = index
	}
	return nil
}
