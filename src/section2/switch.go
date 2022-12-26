package main

import "fmt"

func main() {
	a := -7

	switch {
	case a < 0:
		fmt.Println("음수")
	case a == 0:
		fmt.Println("0")
	case a > 0:
		fmt.Println("양수")
	}

	switch b := 28; {
	case b < 0:
		fmt.Println("음수")
	case b == 0:
		fmt.Println("0")
	case b > 0:
		fmt.Println("양수")
	}

	switch c := "hello"; c {
	case "hello":
		fmt.Println("hello")
	case "bye":
		fmt.Println("bye")
	}

	yy := "y"
	xx := "x"

	vv := "y"

	switch vv {
	case yy:
		fmt.Println("same yy")
	case xx:
		fmt.Println("same xx")
	default:
		fmt.Println("not matching")
	}
}
