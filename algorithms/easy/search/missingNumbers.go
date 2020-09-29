package main

import (
	"fmt"
	"sort"
)

//clave valor tabla hash

func main() {
	//values1 := []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	//values2 := []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}
	values1 := []int32{11, 4, 11, 7, 13, 4, 12, 11, 10, 14}
	values2 := []int32{11, 4, 11, 7, 3, 7, 10, 13, 4, 8, 12, 11, 10, 14, 12}

	numbers := missingNumbers(values1, values2)
	fmt.Println(numbers)
}

func missingNumbers(arr []int32, brr []int32) (result []int32) {
	var table = make(map[int32]int)
	var table2 = make(map[int32]int)

	for _, value := range arr {
		table[value] = table[value] + 1
	}
	for _, value := range brr {
		table2[value] = table2[value] + 1
	}
	if len(arr) > len(brr) {
		for valor := range table {
			if table[valor] != table2[valor] {
				result = append(result, valor)
			}
		}
	} else {
		for valor := range table2 {
			if table[valor] != table2[valor] {
				result = append(result, valor)
			}
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
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

/*
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
*/
