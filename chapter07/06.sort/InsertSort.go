package main

import "fmt"

func insertSort(array []int) {
	length := len(array)
	for i := 0; i < length; i++ {
		tmp := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > tmp {
				array[j+1] = array[j]
			} else {
				break
			}
			array[j] = tmp
		}
	}
}

func main() {
	array := []int{8, 3, 5, 1, 2, 6, 4, 9, 7, 11}
	insertSort(array)
	fmt.Println(array)
}
