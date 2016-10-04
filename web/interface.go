package main

import "fmt"

type Roller interface {
	Roll()
}

// Car is an object
type Car struct{}

func (c Car) Roll() {
	fmt.Println("Car: je roule")
}

// Bike is an object
type Bike struct{}

func (b Bike) Roll() {
	fmt.Println("Bike: I'm rolling...")
}

func RunVehicule(r Roller) {
	r.Roll()
}

func main() {
	RunVehicule(Bike{})
	RunVehicule(Car{})
}
