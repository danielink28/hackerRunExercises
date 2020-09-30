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
