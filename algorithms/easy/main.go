package main

import (
	"fmt"
	"sort"
)

func main() {

	//values := []int32{99, 40, 53, 31, 92, 68, 17, 70, 100, 16, 26, 82, 72, 89, 19, 14, 56, 7, 26, 69, 8, 44, 51, 88, 24, 34, 40, 70, 90, 68, 95, 95, 28, 39, 71, 75, 31, 17, 96, 60, 98, 98, 33, 35, 68, 84, 17, 11, 76, 17, 45, 61, 72, 76, 18, 67, 55, 81, 57, 43, 45, 96, 58, 49, 4, 61, 38, 66, 82, 16, 44, 100, 50, 19, 82, 15, 72, 5, 81, 97, 94, 70, 7, 92, 75, 55, 1, 87, 4, 9, 92, 35, 83, 20, 53, 8, 90, 2, 92, 82}
	//plusMinus(values)
	//staircase(66)
	//miniMaxSum(values)
	//candles := birthdayCakeCandles(values)

	//students := gradingStudents(values)

	//values2 := []int32{1, 3, 2, 6, 1, 2}
	//pairs := divisibleSumPairs(100, 32, values)
	//pairs2 := divisibleSumPairs(6, 3, values2)
	//fmt.Println(pairs)
	//values := []int32{10, 5, 20, 20, 4 ,5 ,2 ,25, 1}
	//values := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	//records := breakingRecords(values)
	//fmt.Print(records)
	fmt.Print(catAndMouse(1, 2, 3))
}

/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
Print the decimal value of each fraction on a new line with  places after the decimal.
Note: This challenge introduces precision problems. The test cases are scaled to six decimal places,
though answers with absolute error of up to  are acceptable.
*/
// Complete the plusMinus function below.
func plusMinus(arr []int32) {

	var positve float32
	var negative float32
	var zero float32

	for _, value := range arr {
		if value > 0 {
			positve++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%06f\n", positve/float32(len(arr)))
	fmt.Printf("%06f\n", negative/float32(len(arr)))
	fmt.Printf("%06f\n", zero/float32(len(arr)))
}

/*
	Staircase detail

	This is a staircase of size :
	   #
	  ##
	 ###
	####
	Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.
	Write a program that prints a staircase of size .
	Function Description
	Complete the staircase function in the editor below.
	staircase has the following parameter(s):
	int n: an integer
	Print
	Print a staircase as described above.
	Input Format
	A single integer,n , denoting the size of the staircase.
*/
func staircase(n int32) {
	var printf = "%" + fmt.Sprintf("%ds\n", n)
	var print string
	for i := 1; i <= int(n); i++ {
		for j := 0; j < i; j++ {
			print = print + "#"
		}
		fmt.Printf(printf, print)
		print = ""
	}
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {

	var min int
	var max int
	var temp int
	for i, _ := range arr {
		temp = 0
		for j, second := range arr {
			if i != j {
				temp += int(second)
			}
		}
		if temp < min || min == 0 {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}
	fmt.Printf("%d %d", min, max)
}

/*
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {

	sort.Slice(candles, func(i, j int) bool { return candles[i] < candles[j] })
	result := Filter(candles, func(val int) bool {
		return val == int(candles[len(candles)-1])
	})
	return int32(len(result))
}
func Filter(arr []int32, cond func(int) bool) []int {
	var result []int
	for i := range arr {
		if cond(int(arr[i])) {
			result = append(result, int(arr[i]))
		}
	}
	return result
}

/*
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) (result []int32) {
	for _, grade := range grades {
		if grade < 38 {
			result = append(result, grade)
			continue
		}
		var val int
		if val = (int(grade) / 5) + 1; grade%5 == 0 {
			val = int(grade) / 5
		}

		diff := 5 * val
		if diff-int(grade) < 3 {
			result = append(result, int32(diff))
		} else {
			result = append(result, grade)
		}
	}
	return result
}

func simpleArraySum(ar []int64) int64 {
	var result int64
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	return result

}

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	var alice int32
	var bob int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if b[i] > a[i] {
			bob++
		}
	}
	return []int32{alice, bob}
}

func diagonalDifference(arr [][]int32) int32 {
	var right int32
	var left int32

	var indiceLeft = len(arr) - 1

	for i := 0; i < len(arr); i++ {
		right = right + arr[i][i]
		left = left + arr[i][indiceLeft]
		indiceLeft--
	}
	return (left) - (right)
}

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) (r int32) {

	for i := 0; i < int(n); i++ {
		for j := i; j < int(n); j++ {
			if i == j {
				continue
			}
			if (ar[j]+ar[i])%k == 0 {
				r++
			}
		}
	}
	return r
}

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	temp := []int32{0, scores[0], 0, 0}
	for i, score := range scores {
		if score > temp[0] {
			temp[0] = score
			if i > 0 {
				temp[2] = temp[2] + 1
			}
		}
		if score < temp[1] {
			temp[1] = score
			if i > 0 {
				temp[3] = temp[3] + 1
			}
		}
	}
	return []int32{temp[2], temp[3]}
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {

	diff1 := x - z
	diff2 := y - z
	if diff1 < 0 {
		diff1 = diff1 * -1
	}
	if diff2 < 0 {
		diff2 = diff2 * -1
	}

	if diff1 == diff2 {
		return "Mouse C"
	} else if diff1 > diff2 {
		return "Cat B"
	}
	return "Cat A"
}
