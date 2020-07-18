package main

import "fmt"

func main() {
	arr := []int{3, 66, 33, 9, 88, 5}
	BubbleSort(arr)
	fmt.Println(arr)
	index := BinaryFind(arr, 0, len(arr), 55)
	fmt.Println(index)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BinaryFind(arr []int, leftIndex, rightIndex, findVal int) int {
	// 二分查找,arr为从小到大排列, 找到返回索引位置,找不到返回-1
	if leftIndex > rightIndex {
		return -1
	}

	middle := (leftIndex + rightIndex) / 2
	if arr[middle] < findVal {
		return BinaryFind(arr, middle + 1, rightIndex, findVal)
	} else if arr[middle] > findVal {
		return BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else {
		return middle
	}
}
