package main

import "fmt"

func main() {

	slice2 := []int32{19, 10, 12, 10, 24, 25, 22}
	fmt.Println(countingSort2(slice2))
	// [0 2 20 24 45 66 75 370 410 1802]

	//fmt.Println(slice)
	// [20 370 45 75 410 1802 24 2 66]
}
func countingSort2(arr []int32) []int32 {
	k := GetMaxIntArray2(arr)
	count := make([]int32, k+1)

	array := CountIntArray2(arr)
	for i := 0; i < len(count); i++ {
		count[i] = array[int32(i)]
	}
	var result []int32
	for i := 0; i < len(count); i++ {
		for array[int32(i)] > 0{
			result = append(result,int32(i))
			array[int32(i)] = array[int32(i)] -1
		}
	}

	return result
}

// CountIntArray counts the element frequencies.
func CountIntArray2(arr []int32) map[int32]int32 {
	model := make(map[int32]int32)
	for _, elem := range arr {
		// first element is already initialized with 0
		model[elem] += 1
	}
	return model
}

// Return the maximum value in an integer array.
func GetMaxIntArray2(arr []int32) int32 {
	maxNum := arr[0]
	for _, elem := range arr {
		if maxNum < elem {
			maxNum = elem
		}
	}
	return maxNum
}
