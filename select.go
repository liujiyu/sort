package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 9, 4, 5}
	SelectSort(arr)
	fmt.Println(arr)
}

//选择排序 时间复杂度:n^2 不稳定
func SelectSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

}
