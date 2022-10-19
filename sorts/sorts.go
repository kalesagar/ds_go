package sorts

import "fmt"

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	fmt.Println("selection sorted array: ", arr)
	return arr
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	for pass := 0; pass < n-1; pass++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	fmt.Println("Buble sorted array: ", arr)
	return arr
}

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println("Insertion sorted: ", arr)
	return arr
}

func QuickSort(arr []int, low, high int) {
	if high < low {
		return
	}
	pivot := high
	if high > low {
		i := low
		for j := low; j < high; j++ {
			if arr[j] < arr[pivot] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
				i = i + 1
			}
		}
		fmt.Println("partitions created: ", arr[low:pivot], arr[pivot:high])
		pivot = i
		QuickSort(arr, low, pivot+1)
		QuickSort(arr, pivot, high)
	}
	fmt.Println("Quick sorted arr: ", arr)
}
