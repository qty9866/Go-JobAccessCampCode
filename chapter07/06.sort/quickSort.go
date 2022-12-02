package main

import "fmt"

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
	array[i], array[left] = array[left], array[i]
	return i
}

func main() {
	array := []int{4, 3, 6, 9, 8, 1, 2, 0}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
