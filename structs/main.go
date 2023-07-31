package main

import "fmt"

type car struct {
	Make       string
	Model      string
	Height     float32
	Width      float32
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

type truck struct {
	car
	backSeats int
}

func main() {
	println("Hello World")

	var myCar car = car{}

	myCar.Make = "Toyota"
	myCar.Model = "Vitz"
	myCar.Height = 4.0
	myCar.Width = 8.0
	myCar.FrontWheel.Radius = 4
	myCar.FrontWheel.Material = "Carbon Fibre"
	myCar.BackWheel.Material = "Carbon Fibre"
	myCar.BackWheel.Radius = 5

	fmt.Println(myCar)

	var hisCar = car{
		Make:   "Toyota",
		Model:  "Sorento",
		Height: 3,
		Width:  4,
		FrontWheel: Wheel{
			Radius:   4,
			Material: "Burgundy",
		},
		BackWheel: Wheel{
			Radius:   5,
			Material: "Carbon Fibre",
		},
	}

	fmt.Println(hisCar)

	var truck1 truck = truck{
		car: car{
			Make:   "Toyota",
			Model:  "Sorento",
			Height: 3,
			Width:  4,
			FrontWheel: Wheel{
				Radius:   4,
				Material: "Burgundy",
			},
			BackWheel: Wheel{
				Radius:   5,
				Material: "Carbon Fibre",
			},
		},
		backSeats: 4,
	}

	fmt.Print(truck1)

	//Calling functions within functions
	myFunc := func() {
		fmt.Println("HEllo World")
	}

	myFunc()

	rectangle1 := rectangle{
		Width:  4,
		Height: 10,
	}

	//Both methods work
	fmt.Println("Area2: ", rectangle1.area2(2, 4))
	fmt.Printf("Area %v", area(rectangle1))

}

//STRUCT METHODS
type rectangle struct {
	Width  int
	Height int
}

func area(r rectangle) int {
	return r.Height * r.Width
}

func (r rectangle) area2(x, y int) int {
	return r.Width*r.Height + (x * y)
}
