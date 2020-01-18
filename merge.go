package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 9, 4, 5}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

//归并排序 时间复杂度 Nlog(N) 稳定
func mergeSort(arr []int) []int {
	count := len(arr)
	if count <= 1 {
		return arr
	}
	mid := count / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left []int, right []int) []int {

	leftCount := len(left)
	rightCount := len(right)
	leftIndex := 0
	rightIndex := 0
	var arr []int
	for {
		if leftIndex >= leftCount && rightIndex >= rightCount {
			return arr
		}
		if leftIndex >= leftCount {
			arr = append(arr, right[rightIndex:]...)
			return arr
		}

		if rightIndex >= rightCount {
			arr = append(arr, left[leftIndex:]...)
			return arr
		}
		if left[leftIndex] > right[rightIndex] {
			arr = append(arr, right[rightIndex])
			rightIndex++
		} else {
			arr = append(arr, left[leftIndex])
			leftIndex++
		}

	}

}
