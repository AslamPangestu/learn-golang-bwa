package main

import (
	"fmt"
	"learn-golang-bwa/calculation"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) showUser() string {
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	isAvailable bool
}

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + "S.T"
}

type Luas interface {
	CalculateLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) CalculateLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) CalculateLuas() int {
	return persegiPanjang.Lebar * persegiPanjang.Panjang
}
func main() {

	student := Student{1, "Dian", 3.75}
	fmt.Println(student)
	student.graduate()
	fmt.Println(student)
	numA := 1
	numB := &numA

	fmt.Println(numA)
	fmt.Println(numB)
	fmt.Println(*numB)

	*numB = 2
	fmt.Println(*numB)
	fmt.Println(numA)

	user := User{}
	user.ID = 1
	user.FirstName = "Budi"
	user.LastName = "Anduk"
	user.Email = "laall@va.ca"
	user1 := User{
		ID:        2,
		FirstName: "Dini",
	}
	user2 := User{2, "Dini", "Dian", "alaal@da.ca", true}
	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)

	users := []User{user1, user}
	group := Group{
		"Gamer", user, users, true,
	}
	username1 := user.showUser()
	// username1 := showUser(user)

	fmt.Println(username1)
	fmt.Println(group)

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

	sentence := printTest("OLLOLO")
	fmt.Println(sentence)
	addRes := add(1, 3)
	fmt.Println(addRes)

	luas, keliling := calculate(10, 5)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func printTest(sentence string) string {
	fmt.Println(sentence)
	newSentence := sentence + " belajar Golang"
	return newSentence
}

func add(a, b int) int {
	return a + b
}

func calculate(p, l int) (int, int) {
	luas := p * l
	kel := 2 * (p + l)

	return luas, kel
}
