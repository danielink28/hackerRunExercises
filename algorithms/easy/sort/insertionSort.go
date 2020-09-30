package main

import "fmt"

func main() {
	values1 := []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	fmt.Print(runningTime(values1))
}

func runningTime(arr []int32) []int32 {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
	return arr
}
