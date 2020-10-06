package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	array := []int32{5, 10, 7}
	//array := []int32{3, -7, 0}
	difference := marcsCakewalk(array)
	fmt.Print("Minima diferencia ", difference)
}

func marcsCakewalk(calorie []int32) (result int64) {

	sort.Slice(calorie, func(i, j int) bool { return calorie[i] > calorie[j] })
	for i := 0; i < len(calorie); i++ {
		pow := int64(math.Pow(2, float64(i)))
		result = result + (pow * int64(calorie[i]))
	}
	return result
}
