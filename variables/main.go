package main

import "fmt"

//We have to provide an entry point which will be the main function

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50 //constants which are immutable just as javascript
	var remainingTickets int = 0
	remainingTickets = conferenceTickets - remainingTickets

	fmt.Println("Welcome to our conference booking application")
	fmt.Print("Get your tickets here to attend")

	fmt.Println(" ", conferenceName)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)

	fmt.Printf("ConferenceTicket is %T\n", conferenceTickets)
	//asks user for their name

	//We can easily do this method without assigning the type declaration since it can read the type of the value assigned and derive the type
	//constants cannot be used like below
	values := 10
	fmt.Printf("Values are: %v", values)

	//if we assign an already typed value to a new variable it latches onto the type of it

	var i = values
	fmt.Printf("\ni is of type %T ", i)

	//can do it in the colon method as well

	j := i
	fmt.Printf("j is of type %T", j)

	// The standard sizes that should be used unless you have a specific need are:

	// 	int
	// 	uint
	// 	float64
	// 	complex128

	// Type conversions

	temperatureInt := 88
	temperatureFloat := float64(temperatureInt)

	fmt.Printf("\n%v", temperatureFloat)

	//Unless you have a good reason to, stick to the following types:

	// bool
	// string
	// int
	// uint
	// byte
	// rune
	// float64
	// complex128

	//RETURNING A CONCATENATED STRING
	var s = "Intelligent"

	var t = fmt.Sprintf("\nI am very %v", s)

	fmt.Println(t)

}
