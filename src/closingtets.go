package main

import "fmt"

// defer -  который позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая
// recover - останавливает панику и возвращает значение, которое было передано функции panic
// panic - чтобы сгенерировать ошибку выполнения

func umn() func() int  {
	i := int(1)
	return func() (ret int) {
		ret = i
		i *= 2
		return
	}
}

func main() {
	nextEven := umn()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
