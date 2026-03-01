package main

import "fmt"

func main() {

	i1, i2, i3 := 12 , 14, 16
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 12.5, 15.2, 4.9
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	total := floatSum + float64(intSum)
	fmt.Println("Total:", total)

	sub := float64(intSum) * floatSum
	fmt.Println("Value:", sub)

}
