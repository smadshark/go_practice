package main

import "fmt"

func main() {
	const (
		Jan = iota
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)

	const (
		A = iota * 100
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
