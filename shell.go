package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 3, 2, 9, 4, 5, 10, 99, 15, 18, 14, 13, 36, 25, 33, 66, 55, 44, 22, 78, 65, 22, 88, 55, 42, 19, 87, 81, 97, 98, 64}
	//for i := 0; i < 10000; i++ {
	//	arr = append(arr, i)
	//}
	now := time.Now()
	for i := 0; i < 100000; i++ {
		ShellSort(arr, 10)
		//ShellSort2(arr)
	}
	fmt.Println(time.Since(now))
	fmt.Println(arr)
}

//希尔排序
func ShellSort(arr []int, gap int) {
	count := len(arr)
	if count <= 1 {
		return
	}
	if gap > count/2 {
		gap = count / 2
	}
	for gap >= 1 {
		for j := gap; j < count; j++ {
			for x := j; x >= gap; x -= gap {
				if arr[x] < arr[x-gap] {
					arr[x], arr[x-gap] = arr[x-gap], arr[x]
				}
			}

		}
		gap /= 3
	}

}

func ShellSort2(a []int) {
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
