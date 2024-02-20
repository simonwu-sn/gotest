package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	rand.Seed(time.Now().Unix())

	arr := rand.Perm(10000000)

	start := time.Now()
	sorted := mergeSort(arr)
	elapsed := time.Since(start)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			panic("Array is not sorted")
		}
	}

	fmt.Printf("Array is sorted, time taken %v\n", elapsed)
}
