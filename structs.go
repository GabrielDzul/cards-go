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

//Form 1 of initializing a struct
//alex := person{"Alex", "Anderson"}

//Form 2 of initializing a struct
//alan := person{firstName: "Alan", lastName: "Anders"}

//Form 3 of initializing a struct
//var adan person
//adan.firstName = "Adan"
//fmt.Println(adan)
//fmt.Printf("%+v", adan)

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
