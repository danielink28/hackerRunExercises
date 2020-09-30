package main

import "fmt"

func main() {

	values1 := []int32{4, 5, 3, 7, 2}
	fmt.Print(quickSort(values1))
}

// Complete the quickSort function below.
func quickSort(arr []int32) (result []int32) {
	var pivote = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivote {
			result = append(result, arr[i])
		}
	}
	result = append(result, pivote)
	for i := 1; i < len(arr); i++ {
		if arr[i] > pivote {
			result = append(result, arr[i])
		}
	}
	return result
}
