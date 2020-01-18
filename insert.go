package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 9, 4, 5}
	InsertSort(arr)
	fmt.Println(arr)
}

//插入排序
func InsertSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}

	}

}
