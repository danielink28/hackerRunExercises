package main

func main() {
	values1 := []int32{40,50,60}
	values2 := []int32{5,8,12}
	getMoneySpent := getMoneySpent(values1, values2, 60)
	println(getMoneySpent)
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) (result int32) {
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			sum := keyboards[i] + drives[j]
			if sum <= b && sum > result {
				result = sum
			}
		}
	}
	if result == 0 {
		return -1
	}
	return result
}
