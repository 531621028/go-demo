package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergeNums := mergeNums(nums1, nums2)
	totalSize := len(mergeNums)
	if (totalSize % 2) != 0 {
		return float64(mergeNums[totalSize/2])
	} else {
		return (float64(mergeNums[totalSize/2]) + float64(mergeNums[totalSize/2-1])) / 2
	}
}

func mergeNums(nums1 []int, nums2 []int) []int {
	// 定义一个切片
	result := make([]int, 0)
	index1 := 0
	index2 := 0
	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			result = append(result, nums1[index1])
			index1++
		} else {
			result = append(result, nums2[index2])
			index2++
		}
	}
	for index1 < len(nums1) {
		result = append(result, nums1[index1])
		index1++
	}
	for index2 < len(nums2) {
		result = append(result, nums2[index2])
		index2++
	}
	return result
}
