package main

import (
	"fmt"
	"learn-golang-bwa/calculation"
)

func main() {
	fmt.Println("Test Print")
	// sentence := TestFucn()
	// fmt.Println(sentence)

	result := calculation.Add(3, 6)
	fmt.Println(result)

	//Variable Definition
	var name string = "hai"
	var name1 string
	name1 = "Cek ssatu"
	age := 30
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(age)

	//Conditional IF
	umur := 10
	score := 80
	var grade string
	if umur > 10 {
		fmt.Println("Boleh main game")
	}

	if umur > 10 {
		fmt.Println("Boleh main game")
	} else {
		fmt.Println("Tidak Boleh main game")
	}

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else {
		grade = "A"
	}
	fmt.Println(grade)

	//Conditional Switch
	number := 5
	switch number {
	case 1:
		fmt.Println("Sate")
	case 2:
		fmt.Println("Dua")
	default:
		fmt.Println("Nukk")
	}

	//Looping
	index := 1
	title := "OI AH UW LU"
	//For
	for i := 1; i <= 5; i++ {
		fmt.Println("KWKWKWKW", i)
	}
	//While
	for index < 5 {
		fmt.Println("KWKWKWKW", index)
		index++
	}
	//FOREACH with index
	for index, letter := range title {
		fmt.Println(index, string(letter))
	}

	//FOREACH without index
	for _, letter := range title {
		fmt.Println(string(letter))
	}
}
