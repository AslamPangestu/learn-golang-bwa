package main

import (
	"fmt"
	"learn-golang-bwa/calculation"
)

func main() {
	fmt.Println("Test Print")
	sentence := TestFucn()

	result := calculation.Add(3, 6)

	fmt.Println(result)

	fmt.Println(sentence)
}
