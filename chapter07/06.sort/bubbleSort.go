package main

import "fmt"

func bubbleSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
			if !flag {
				break
			}
		}
		if !flag {
			break
		}
	}
}
func main() {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	bubbleSort(array)
	fmt.Println(array)
}
