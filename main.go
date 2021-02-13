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

	//ARRAY
	var languanges [5]string //Definition
	//Set Value
	languanges[0] = "A"
	languanges[1] = "B"
	languanges[2] = "C"
	languanges[3] = "D"
	languanges[4] = "E"
	//Definition + Set Value
	tests := [5]string{"A", "B", "C", "D", "E"}
	//Definition + Set Value without define length
	testSatu := [...]string{"AAAA", "BBB"}

	length := len(languanges)
	fmt.Println(length)
	fmt.Println(tests)
	fmt.Println(languanges)
	fmt.Println(testSatu)

	//PRINT ARRAY
	for _, item := range languanges {
		fmt.Println(item)
	}

	//SLICE
	var gamConsoles []string                 //Define
	gameConsoles2 := []string{"PS5", "XBOX"} //Define + Set Vakye
	gamConsoles = append(gamConsoles, "PS4") //Set Value
	fmt.Println(gamConsoles)
	fmt.Println(gameConsoles2)

	//MAP
	var testMap map[string]int                                         //Definition
	testMap = map[string]int{}                                         //Inisiaton
	testMap2 := map[string]string{"KWKWKWK": "AHAHA", "UWUWU": "HEHE"} //Definition + Initiation
	//Set Value
	testMap["ruby"] = 1
	testMap["js"] = 2
	testMap["go"] = 3

	fmt.Println(testMap)
	fmt.Println(testMap2)
	fmt.Println(testMap["ruby"])

	for key, value := range testMap {
		fmt.Println(key, value)
	}
	delete(testMap, "go")           //Delete Value
	val, isExist := testMap["ters"] //check is map key exits
	fmt.Println(val, isExist)

	students := []map[string]string{
		{"name": "Budi", "score": "A"},
		{"name": "Deni", "score": "B"},
	}

	for _, item := range students {
		fmt.Println(item["name"], item["score"])
	}
}
