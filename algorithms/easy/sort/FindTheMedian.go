package main

import "fmt"

func main() {

	values1 := []int32{0, 1, 2, 4, 6, 5, 3}
	fmt.Print(findMedian(values1))
}

// Complete the findMedian function below.
func findMedian(arr []int32) int32 {
	root := Node{
		value: int(arr[0]),
	}
	for i := 1; i < len(arr); i++ {
		root.add(int(arr[i]))
	}
	var result []int
	root.printInOrder(&result)
	return int32(result[len(result)/2])
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) add(value int) {
	if value < n.value {
		if n.left != nil {
			n.left.add(value)
		} else {
			n.left = &Node{
				value: value,
			}

		}
	} else {
		if n.right != nil {
			n.right.add(value)
		} else {
			n.right = &Node{
				value: value,
			}
		}
	}
}

func (n *Node) printInOrder(result *[]int) {
	if n.left != nil {
		n.left.printInOrder(result)
	}
	*result = append(*result, n.value)
	if n.right != nil {
		n.right.printInOrder(result)
	}
}
