//conditionals

package main

import (
	"errors"
	"fmt"
)

func main() {

	length := []int{1, 2, 3, 4, 5}
	//or write the other way

	var length2 = []int{1, 2, 3, 4, 5}

	if len(length) > len(length2) {
		fmt.Print("First Array is bigger")
	} else {
		fmt.Print("Nvm")
	}

	//Other alternative
	if newValue := len(length2); newValue > 1 {
		fmt.Println("New value is gooood")
	}

	if stringValue := "Intermediate"; len(stringValue) > 1 {
		fmt.Print("length of stringValue is : ", len(stringValue))
	}

	fmt.Println("\n", add(3, 5))

	//if we are calling a function with multiple returns, we can ignore one return by simply stating _

	l, _ := getPoint()
	fmt.Println("Get Point first Coordination", l)

	a, _ := getCoords(3, 4)
	fmt.Println(a)

	var p, b, s = calculator(3, 4)
	fmt.Printf("%v %v %v \n", p, b, s)
}

func add(x int, y int) int {
	return x + y
}

// The second parameter shows the returns
func getPoint() (x int, y int) { // x and y are indexes for the return values -> When calling it can be simply any alias
	return 3, 4 //we are returning two values
}

// named return values
func getCoords(x int, y int) (int, int) {
	return x, y
}

func calculator(x, y float32) (float32, float32, error) {
	if x/y == 0 {
		return 0, 0.0, errors.New("Cant divide by zero")
	}

	mul := x * y
	div := float32(x / y)
	return mul, div, nil
}
