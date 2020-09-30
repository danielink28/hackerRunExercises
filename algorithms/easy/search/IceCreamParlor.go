package main

/*
Sunny and Johnny like to pool their money and go to the ice cream parlor.
Johnny never buys the same flavor that Sunny does. The only other rule they have is that they spend all of their money.
Given a list of prices for the flavors of ice cream, select the two that will cost all of the money they have
*/
func main() {
	values1 := []int32{1, 3, 4, 5, 6}
	parlor := icecreamParlor(6, values1)
	println(parlor)
}

func icecreamParlor(m int32, arr []int32) (result []int32) {

	for i, value := range arr {
		for j, val := range arr {
			if  i != j && (value + val) == m {
				return []int32{int32(i + 1), int32(j + 1)}
			}
		}
	}
	return result
}

func binarySearch(matches []int32, sum int32, value int32) int {
	var indexA = 0

	var indexB = len(matches) - 1
	for indexA <= indexB {
		var mid = (indexA + indexB) / 2
		if matches[mid] <= 0 {
			indexA = mid + 1
		} else {
			indexB = mid - 1
		}
	}
	return indexB
}
