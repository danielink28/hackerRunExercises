package main

import (
	"fmt"
	"math/big"
)

func main() {
	//31415926535897932384626433832795
	//18446744073709551615
	values1 := []string{"31415926535897932384626433832795", "1", "3", "10", "3", "5"}
	//values1 := []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854}
	fmt.Print(bigSorting(values1))
}

// Complete the bigSorting function below.
func bigSorting(unsorted []string) []string {

	n := new(big.Int)
	n, _ = n.SetString(unsorted[0], 10)
	nodo := Nodo{
		value: *n,
	}
	for i := 1; i < len(unsorted); i++ {
		n := new(big.Int)
		var setString, _ = n.SetString(unsorted[i], 10)
		nodo.add(*setString)
	}
	var result []string
	nodo.printInOrder(&result)
	return result
}

type Nodo struct {
	value big.Int
	left  *Nodo
	right *Nodo
}

func (n *Nodo) add(value big.Int) {
	if n.value.Cmp(&value) == 1 {
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

func (n *Nodo) printInOrder(result *[]string) {
	if n.left != nil {
		n.left.printInOrder(result)
	}
	*result = append(*result, n.value.String())
	if n.right != nil {
		n.right.printInOrder(result)
	}
}
