package main

import "fmt"

func main() {
	var intVal int
	var strVal string

	var (
		multiVal1 string  = "hello"
		multiVal2 int     = 2
		multiVal3 float32 = 2.2
	)

	fmt.Println(intVal, strVal, multiVal1, multiVal2, multiVal3)

	// short declare (전역X), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성 Good -> recommend
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

	shortVar3 = true

	fmt.Println("shortVar3: ", shortVar3)

	// Example
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}
}
