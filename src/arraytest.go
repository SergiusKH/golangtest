package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	arr := make([]int, 3)
	copy(arr, x)
	fmt.Println(arr)
	arr2 := x[3:5]
	fmt.Println(arr2)

	slice := append(x, 5, 10)
	fmt.Println(slice)

	var min int = x[0]

	for _, value := range slice {
		if min > value {
			min = value
		}
	}

	fmt.Printf("Min element array X: %d", min)
}
