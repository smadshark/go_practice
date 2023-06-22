package main

import "fmt"

func main() {
	var aSlice []float64
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)

	fmt.Println(aSlice, "with length", len(aSlice))

	bSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(bSlice[0:5])

	byteSlice := make([]byte, 12)

	fmt.Println(byteSlice)

	byteSlice = []byte("hello")

	fmt.Println(byteSlice)

	fmt.Printf("byteSlice: %s\n", byteSlice)

	fmt.Println("byteSlice:", string(byteSlice))
}
