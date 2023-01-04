package main

import (
	"fmt"
	"runtime"
)

func main() {
	//startDate := time.Now()

	//if len(os.Args) != 2 {
	//	fmt.Println("Usage: dates parse_string")
	//	return
	//}
	//dateString := os.Args[1]
	//
	//d, err := time.Parse("02 January 2006", dateString)
	//if err == nil {
	//	fmt.Println("Full:", d)
	//	fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	//}

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
}
