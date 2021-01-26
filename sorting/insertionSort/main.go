package main

func myinsertionsort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		valueToSort := arr[i]
		// j := i-1
		for i > 0 && arr[i-1] > valueToSort {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
	return arr
}
