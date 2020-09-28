package main

import (
	"fmt"
	"sort"
)
//clave valor tabla hash

func main() {
	values1 := []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	values2 := []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}
	numbers := missingNumbers(values1, values2)
	fmt.Println(numbers)
}

func missingNumbers(arr []int32, brr []int32) []int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	sort.Slice(brr, func(i, j int) bool { return brr[i] < brr[j] })

	//arr = deleteRepeated(arr)
	//brr = deleteRepeated(arr)

	if len(arr) > len(brr) {
		return validate(arr, brr)
	}
	return validate(brr, arr)
	return nil
}

func deleteRepeated(arr []int32) []int32 {
	return Filter(arr, func(val int32, result []int32) bool {
		var add = true
		for _, value := range result {
			if val == value {
				add = false
			}
		}
		return add
	})
}

func Filter(arr []int32, cond func(int32, []int32) bool) (result []int32) {
	for i := range arr {
		if cond(arr[i], result) {
			result = append(result, arr[i])
		}
	}
	return result
}

func validate(arr1 []int32, brr2 []int32) (result []int32) {
	for _, val1 := range arr1 {
		var lost = true
		for _, val2 := range brr2 {
			if val1 == val2 {
				lost = false
				remove(brr2, val2)
				break
			}
		}
		if lost {
			result = append(result, val1)
		}
	}
	return result
}
