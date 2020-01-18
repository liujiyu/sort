package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 9, 4, 5}
	BubbleSort(arr)
	fmt.Println(arr)
}

//冒泡排序 时间复杂度 n^2 稳定
func BubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

}
