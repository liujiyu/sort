package main

import "fmt"

func main() {
	arr := []int{1, 3, 3, 3, 3, 3, 3, 3, 3, 32, 9, 4, 5}
	HeapSort(arr)
	fmt.Println(arr)
}

//堆排序 时间复杂度 Nlog(N) 不稳定
func HeapSort(arr []int) {
	count := len(arr)
	if count <= 1 {
		return
	}
	buildHeap(arr)
	for i := 1; i < count; i++ {
		arr[0], arr[count-i] = arr[count-i], arr[0]
		adjustHeap(arr[:count-i], 0)
	}

}

func buildHeap(arr []int) {
	//左子节点 2*n+1 右子节点 2*n+2  父节点 (n-1)/2
	count := len(arr)
	for i := count / 2; i >= 0; i-- {
		adjustHeap(arr, i)
	}
}

//调整为大顶堆 137286
func adjustHeap(arr []int, index int) {
	count := len(arr)
	pos := index
	for pos*2+2 < count {
		if arr[index] >= arr[index*2+1] && arr[index] >= arr[index*2+2] {
			break
		}
		if arr[index] < arr[index*2+1] {
			arr[index], arr[index*2+1] = arr[index*2+1], arr[index]
			pos = index*2 + 1
		}
		if arr[index] < arr[index*2+2] {
			arr[index], arr[index*2+2] = arr[index*2+2], arr[index]
			pos = index*2 + 2
		}
		index = pos
	}
}
