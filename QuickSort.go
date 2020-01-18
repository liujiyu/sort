package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 2, 9, 4, 5}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

//快速排序 时间复杂度 nlog(n) 最坏的情况（基本有序） n^2  不稳定
func QuickSort(arr []int, left int, l int) {
	last := l
	if left >= last {
		return
	}
	flag := left
	for left <= last {
		if arr[left] < arr[flag] {
			arr[left], arr[flag] = arr[flag], arr[left]
			flag = left
			left++
		} else if arr[left] > arr[flag] {
			arr[left], arr[last] = arr[last], arr[left]
			last--
		} else {
			left++
		}

	}
	QuickSort(arr, 0, flag-1)
	QuickSort(arr, flag+1, l)

}
