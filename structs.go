package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

//Form 1 of initializing a struct
//alex := person{"Alex", "Anderson"}

//Form 2 of initializing a struct
//alan := person{firstName: "Alan", lastName: "Anders"}

//Form 3 of initializing a struct
//var adan person
//adan.firstName = "Adan"
//fmt.Println(adan)
//fmt.Printf("%+v", adan)

// jim := person{
// 	firstName: "Jim",
// 	lastName:  "Party",
// 	contactInfo: contactInfo{
// 		email:   "test@test",
// 		zipCode: 12345,
// 	},
// }

// jim.updateName("Jimmy")

// jim.print()

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
