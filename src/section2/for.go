package main

import "fmt"

func main() {
	loc := []string{"Seoul", "Busan", "Incheon"}

	for _, name := range loc {
		fmt.Println(name)
	}

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
	}

	fmt.Println(sum2)

	// LABEL
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println(i, j)
		}
	}

Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println(i, j)
		}
	}
}
