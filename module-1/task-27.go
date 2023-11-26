package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	return merge(mergeSort(array[:middle]), mergeSort(array[middle:]))
}

func merge(left, right []int) []int {
	result := []int{}
	leftIdx, rightIdx := 0, 0

	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] <= right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx++
		} else {
			result = append(result, right[rightIdx])
			rightIdx++
		}
	}

	for leftIdx < len(left) {
		result = append(result, left[leftIdx])
		leftIdx++
	}

	for rightIdx < len(right) {
		result = append(result, right[rightIdx])
		rightIdx++
	}

	return result
}

func main() {
	var array = make([]int, 5)
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&a)
		array[i] = a
	}

	fmt.Print(mergeSort(array)[len(array)-1])
}
