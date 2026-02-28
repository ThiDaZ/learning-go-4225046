package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	aNumber := 12

	fmt.Println(str1, str2, str3)
	stringLength, err := fmt.Println("The value is", aNumber)
	if err == nil {
		fmt.Println("The string length is", stringLength)
	}
}
