package main

import (
	"fmt"
	"time"
)

// 循环实现二分查找
func binarySearch1(array []int64, target int64) int64 {
	length := len(array)
	if length == 0 {
		return -1
	}
	low := 0
	high := length - 1
	for low <= high {
		totalCompare++
		mid := low + (high-low)/2
		if array[mid] == target {
			return int64(mid)
		} else if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 递归的实现
func binarySearch2(array []int64, target int64) int64 {
	length := int64(len(array))
	return bs(array, int64(0), length-1, target)
}

func bs(array []int64, low, high, target int64) int64 {
	if low >= high {
		return -1
	}
	totalCompare++
	mid := (low + high) / 2
	if array[mid] == target {
		return mid
	} else if array[mid] > target {
		return bs(array, low, mid-1, target)
	} else {
		return bs(array, mid+1, high, target)
	}
}

var totalCompare int64 = 0

/*
binarySearch1
总比较次数 74000000
总共耗时 137.6562ms

binarySearch2
总比较次数 70000000
总共耗时 105.8981ms
*/

func main() {
	arr := SampleData
	QuickSort(arr, 0, int64(len(arr)-1))
	start := time.Now()
	for i := 0; i < 200*10000; i++ {
		binarySearch2(arr, 501)
		binarySearch2(arr, 888)
		binarySearch2(arr, 900)
		binarySearch2(arr, 3)
	}
	fmt.Println("总比较次数", totalCompare)
	finish := time.Now()
	fmt.Println("总共耗时", finish.Sub(start))
}

func QuickSort(array []int64, left, right int64) {
	if left < right {
		loc := Partition(array, left, right)
		QuickSort(array, left, loc-1)
		QuickSort(array, loc+1, right)
	}
}

func Partition(array []int64, left, right int64) int64 {
	i := left + 1
	j := right
	for i < j {
		if array[i] > array[left] {
			array[i], array[j] = array[j], array[i]
			j--
		} else {
			i++
		}
	}
	if array[i] >= array[left] {
		i--
	}
	array[i], array[left] = array[left], array[i]
	return i
}
