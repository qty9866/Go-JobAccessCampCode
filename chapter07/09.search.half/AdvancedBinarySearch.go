package main

import "fmt"

// BinarySearchFirstEqual  二分查找：查找第一个第一个值等于给定值的元素
func BinarySearchFirstEqual(arr []int, value int) int {
	length := len(arr)
	if length == 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// BinarySearchLastEqual 二分查找：查找最后一个值等于给定值的元素
func BinarySearchLastEqual(arr []int, value int) int {
	length := len(arr)
	low := 0
	high := length - 1
	if length == 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == length-1 || arr[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// BinarySearchFirstBiggerOrEqual 二分查找：查找第一个值大于等于给定值的元素
func BinarySearchFirstBiggerOrEqual(arr []int, value int) int {
	length := len(arr)
	low := 0
	high := length - 1
	if length == 0 {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// BinarySearchLastLessOrEqual 二分查找：查找最后一个小于等于给定值的元素
func BinarySearchLastLessOrEqual(arr []int, value int) int {
	length := len(arr) - 1
	low := 0
	high := length - 1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > value {
			high = mid - 1
		} else {
			if mid == length-1 || arr[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 4, 5, 5, 5, 5, 9, 10, 11}
	index1 := BinarySearchFirstEqual(array, 5)
	index2 := BinarySearchLastEqual(array, 5)
	index3 := BinarySearchFirstBiggerOrEqual(array, 6)
	index4 := BinarySearchLastLessOrEqual(array, 6)
	fmt.Println("数组中的第一个5下标：", index1)
	fmt.Println("数组中的最后一个5下标：", index2)
	fmt.Println("数组中的第一个大于等于6的下标：", index3)
	fmt.Println("数组中最后一个小于等于6的下标：", index4)
}
