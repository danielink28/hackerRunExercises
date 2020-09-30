package main

import "fmt"

func main() {
	//values1 := []int32{2, 3, 4, 5, 1}
	values1 := []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	insertionSort1(10, values1)
}

func insertionSort1(n int32, arr []int32) {
	var temp = arr[n-1]
	for j := n - 2; j >= 0; j-- {
		if j > 0 {
			if temp < arr[j] {
				arr[j+1] = arr[j]
				printArray(arr)
			} else {
				arr[j+1] = temp
				printArray(arr)
				break
			}
		} else if temp < arr[j] {
			arr[j+1] = arr[j]
			printArray(arr)
			arr[j] = temp
			printArray(arr)
		}else{
			arr[j+1] = temp
			printArray(arr)
		}
	}
}

func printArray(arr []int32) {
	for i, value := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(value)
	}
	fmt.Println()
}
