package main

func main() {

	//values1 := []int32{5, 2, 3, 4, 1}
	//values1 := []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854}
	//fmt.Print(closestNumbers(values1))
}
/*
// Complete the closestNumbers function below.
func closestNumbers(arr []int32) (result []int32) {
	root := Nodo{
		value: int(arr[0]),
	}

	for i := 1; i < len(arr); i++ {
		root.add(int(arr[i]))
	}
	var order []int
	root.printInOrder(&order)
	var minDiff int
	for i := 0; i < len(order)-1; i++ {
		diff :=  order[i+1] -order[i]
		if minDiff == 0 || diff < minDiff {
			minDiff = diff
			result = []int32{}
			result = append(result, int32(order[i]))
			result = append(result, int32(order[i+1]))
		} else if diff == minDiff {
			result = append(result, int32(order[i]))
			result = append(result, int32(order[i+1]))
		}
	}
	return result
}

type Nodo struct {
	value int
	left  *Nodo
	right *Nodo
}

func (n *Nodo) add(value int) {
	if value < n.value {
		if n.left != nil {
			n.left.add(value)
		} else {
			n.left = &Nodo{
				value: value,
			}
		}
	} else {
		if n.right != nil {
			n.right.add(value)
		} else {
			n.right = &Nodo{
				value: value,
			}
		}
	}
}

func (n *Nodo) printInOrder(result *[]int) {
	if n.left != nil {
		n.left.printInOrder(result)
	}
	*result = append(*result, n.value)
	if n.right != nil {
		n.right.printInOrder(result)
	}
}
*/