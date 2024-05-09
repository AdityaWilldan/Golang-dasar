package main

import "fmt"

func BubbleSort(arr []int) {
	var temp int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func SelectionSort(arr []int) {
	var temp int
	var min_idx int

	for i := 0; i < len(arr)-1; i++ {
		min_idx = i
		for j := 1 + i; j < len(arr); j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min_idx]
		arr[min_idx] = temp
	}
}

func main() {

	arr := []int{12, 2, 3, 1, 4, 5, 7, 8, 9, 10}

	BubbleSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Print("\n")

	SelectionSort(arr)
	for k := 0; k < len(arr); k++ {
		fmt.Println(arr[k])
	}

}
