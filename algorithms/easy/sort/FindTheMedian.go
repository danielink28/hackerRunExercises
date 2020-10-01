package main

import "fmt"

func main() {

	values1 := []int32{0, 1, 2, 4, 6, 5, 3}
	fmt.Println(findMedian(values1))
}

// Complete the findMedian function below.
func findMedian(arr []int32) int32 {
	root := Nodo{
		value: int(arr[0]),
	}

	for i := 1; i < len(arr); i++ {
		root.add(int(arr[i]))
	}
	var result []int
	root.printInOrder(&result)
	return int32(result[len(result)/2])
}

