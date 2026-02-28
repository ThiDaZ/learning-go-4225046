package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("Input text:", str)
}
