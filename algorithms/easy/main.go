package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int32{73, 67, 38, 33}
	//plusMinus(values)
	//staircase(66)
	//miniMaxSum(values)
	//candles := birthdayCakeCandles(values)

	students := gradingStudents(values)
	println(students)
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

	var indiceLeft = len(arr)-1

	for i := 0; i < len(arr); i++ {
		right =right + arr[i][i]
		left = left + arr[i][indiceLeft]
		indiceLeft--
	}


	return (left)- (right)
}