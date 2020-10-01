package main

import "fmt"

func main() {
	slice := []int32{63, 54, 17, 78, 43, 70, 32, 97, 16, 94, 74, 18, 60, 61, 35, 83, 13, 56, 75, 52, 70, 12, 24, 37, 17, 0, 16, 64, 34, 81, 82, 24, 69,
		2, 30, 61, 83, 37, 97, 16, 70, 53, 0, 61, 12, 17, 97, 67, 33, 30, 49, 70, 11, 40, 67, 94, 84, 60, 35, 58, 19, 81, 16, 14, 68, 46, 42, 81, 75, 87, 13,
		84, 33, 34, 14, 96, 7, 59, 17, 98, 79, 47, 71, 75, 8, 27, 73, 66, 64, 12, 29, 35, 80, 78, 80, 6, 5, 24, 49, 82}

	/*slice := []int32{63,25,73,1,98,73,56,84,86,57,16,83,8,25,81,56,9,53,98,67,99,12,83,89,80,91,39,86,76,85,74,39,25,
	90,59,10,94,32,44,3,89,30,27,79,46,96,27,32,18,21,92,69,81,40,40,34,68,78,24,87,42,69,23,41,78,22,6,90,99,89,
	50,30,20,1,43,3,70,95,33,46,44,9,69,48,33,60,65,16,82,67,61,32,21,79,75,75,13,87,70,33}*/

	//fmt.Println(countingSort(slice)) // [0 1 2 2 5 7]*/
	//slice := []int32{1, 1, 3, 2, 1}
	fmt.Println(countingSort(slice)) // [0 1 2 2 5 7]

	//slice1 := []int{1, 2, 3, 6, 4, 5, 4, 6, 7, 8}
	//fmt.Println(countingSort(slice1)) // [0 1 2 3 4 4 5 6 6 7 8]
	//fmt.Println(slice1)               // [1 2 3 6 4 5 4 6 7 8]

	//slice2 := []int{20, 370, 45, 75, 410, 1802, 24, 2, 66}
	//fmt.Println(countingSort(slice2))
	// [0 2 20 24 45 66 75 370 410 1802]

	//fmt.Println(slice)
	// [20 370 45 75 410 1802 24 2 66]
}
func countingSort(arr []int32) []int32 {
	k := GetMaxIntArray(arr)
	count := make([]int32, k+1)

	array := CountIntArray(arr)
	for i := 0; i < len(count); i++ {
		count[i] = array[int32(i)]
	}

	return count
}

// CountIntArray counts the element frequencies.
func CountIntArray(arr []int32) map[int32]int32 {
	model := make(map[int32]int32)
	for _, elem := range arr {
		// first element is already initialized with 0
		model[elem] += 1
	}
	return model
}

// Return the maximum value in an integer array.
func GetMaxIntArray(arr []int32) int32 {
	maxNum := arr[0]
	for _, elem := range arr {
		if maxNum < elem {
			maxNum = elem
		}
	}
	return maxNum
}