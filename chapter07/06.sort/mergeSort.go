package main

import (
	"fmt"
)

func MergeSort(array []int) {
	length := len(array)
	mergeSort(array, 0, length-1)
}

func mergeSort(array []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(array, start, mid)
	mergeSort(array, mid+1, end)
	merge(array, start, mid, end)
}

func merge(array []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if array[i] <= array[j] {
			tmpArr[k] = array[i]
			i++
		} else {
			tmpArr[k] = array[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr[k] = array[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = array[j]
		k++
	}
	copy(array[start:end+1], tmpArr)
}

func main() {
	array := []int{4, 3, 1, 2, 8, 6, 0}
	MergeSort(array)
	fmt.Println(array)
}
