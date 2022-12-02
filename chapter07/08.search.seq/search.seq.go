package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var totalCompare int = 0

//总比较次数 4570000000
//总共耗时 7.253375s

func main() {
	//size := 1000

	arr := sampleData
	QuickSort(arr, 0, len(arr)-1)
	start := time.Now()
	for i := 0; i < 2000000; i++ {
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
	}
	fmt.Println("总比较次数", totalCompare)
	finish := time.Now()
	fmt.Println("总共耗时", finish.Sub(start))
}

func search(arrp *[]int64, targetNum int64) bool {
	for _, v := range *arrp {
		totalCompare++
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(50))
		arr = append(arr, n.Int64())
	}
	return arr
}

func QuickSort(array []int64, left, right int) {
	if left < right {
		loc := partition(array, left, right)
		QuickSort(array, left, loc-1)
		QuickSort(array, loc+1, right)
	}
}

func partition(array []int64, left, right int) int {
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
