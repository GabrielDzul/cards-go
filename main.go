package main

func main() {
	//cards := newDeck()
	//cards.shuffle()
	//cards.print()

	//evenOrOdd()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "test@test",
			zipCode: 12345,
		},
	}

	jim.updateName("Jimmy")

	jim.print()
}
