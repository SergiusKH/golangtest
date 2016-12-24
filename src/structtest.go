package main

import (
	"fmt"
)

type Talk interface {
	hello()
}

type Person struct {
	Name string
}

func (p *Person) hello()  {
	fmt.Println("Hi,", p.Name)
}

func printAll(t ...Talk)  {
	for _, s := range t {
		s.hello()
	}
}

func main() {
	serg := Person{Name:"Sergey"}
	olga := Person{Name:"Olga"}

	p := Person{Name:"Dred"}
	p.hello()

	//var t Talk

	//t = serg
	//t.hello()

	//
	// x := make([]Talk, 2)
	//x[0], x[1] = serg, olga
	printAll(&serg, &olga)
}