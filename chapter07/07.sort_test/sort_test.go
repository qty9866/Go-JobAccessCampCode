package sort_test

import (
	"fmt"
	"testing"
)

func BubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	array := []int{8, 9, 6, 5, 1, 2, 3, 4}
	BubbleSort(array)
	fmt.Println(array)
}

func InsertSort(array []int) {
	length := len(array)
	for i := 1; i < length; i++ {
		tmp := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if tmp < array[j] {
				array[j+1] = array[j]
			} else {
				break
			}
			array[j] = tmp
		}
	}
}

func TestInsertSort(t *testing.T) {
	array := []int{8, 9, 6, 5, 1, 2, 3, 4, 7}
	InsertSort(array)
	fmt.Println(array)
}

func MergeSort(array []int) {
	length := len(array)
	mergesort(array, 0, length-1)
}

func mergesort(array []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergesort(array, start, mid)
	mergesort(array, mid+1, end)
	merge(array, start, mid, end)
}

func merge(array []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if array[i] > array[j] {
			tmpArr[k] = array[j]
			j++
		} else {
			tmpArr[k] = array[i]
			i++
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

func TestMergeSort(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 16}
	MergeSort(array)
	fmt.Println(array)
}

func QuickSort(array []int, left, right int) {
	if left < right {
		loc := partition(array, left, right)
		QuickSort(array, left, loc-1)
		QuickSort(array, loc+1, right)
	}
}

func partition(array []int, left, right int) int {
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
	array[left], array[i] = array[i], array[left]
	return i
}

func TestQuickSort(t *testing.T) {
	array := []int{5, 1, 4, 3, 6, 2, 9, 8, 7, 0, 16}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
