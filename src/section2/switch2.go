package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("50 이상 100 미만 ", i)
	case i >= 25 && i < 50:
		fmt.Println("25 이상 50 미만 ", i)
	default:
		fmt.Println("default ", i)
	}

	switch j := rand.Intn(1000); j % 2 {
	case 0:
		fmt.Println("짝수: ", j)
	case 1:
		fmt.Println("홀수: ", j)
	default:
		fmt.Println("Never: ", j)
	}
}
