package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}
	//array := []int32{3, -7, 0}
	difference := minimumAbsoluteDifference(array)
	fmt.Print("Minima diferencia ", difference)
}

func minimumAbsoluteDifference(arr []int32) (result int32) {
	 sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	if arr[0]-arr[1] < 0 {
		result = (arr[0] - arr[1]) * -1
	} else {
		result = arr[0] - arr[1]
	}
	for i := 0; i < len(arr)-1; i++ {
		newArray := arr[i+1:10]
		binary(newArray, arr[i], &result)
	}
	return result
}

func binary(arr []int32, valorPosition int32, min *int32) {
	for j := 0; j < len(arr); j++ {
		rest := valorPosition - arr[j]
		if rest < 0 {
			rest = rest * -1
		}
		if *min > rest {
			*min = rest
		}
	}
}