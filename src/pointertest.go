package main

import "fmt"
// * (звёздочка), за которым следует тип хранимого значения
// &, который используется для получения адреса переменной

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func swap(x, y *int)  {
	*x, *y = *y, *x
}

func main() {
	x1 := 1
	y1 := 2
	fmt.Println(x1)
	fmt.Println(y1)
	swap(&x1, &y1)
	fmt.Println(x1)
	fmt.Println(y1)

	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
