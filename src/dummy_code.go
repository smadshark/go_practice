package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	hv := os.Getenv("hello")
	fmt.Println(hv)
	lowerStr := strings.ToLower(hv)
	fmt.Println(lowerStr)
}
