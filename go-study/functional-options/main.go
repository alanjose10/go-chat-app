package main

import "fmt"

type Car struct {
	Make         string
	Model        string
	Year         int
	Color        string
	EnginType    string
	Transmission string
	IsElectric   bool
}

type Option func(*Car)

func WithMake(make string) Option {
	return func(c *Car) {
		c.Make = make
	}
}

func WithModel(model string) Option {
	return func(c *Car) {
		c.Model = model
	}
}

func WithYear(year int) Option {
	return func(c *Car) {
		c.Year = year
	}
}

func getCar(opts ...Option) Car {
	defaultCar := Car{
		Make:         "defaultMake",
		Model:        "defaultModel",
		Year:         2023,
		Color:        "white",
		EnginType:    "default",
		Transmission: "default",
		IsElectric:   false,
	}

	for _, opt := range opts {
		opt(&defaultCar)
	}
	return defaultCar
}

func main() {

	car1 := getCar()
	fmt.Printf("%+v\n", car1)

	car2 := getCar(WithMake("VW"), WithModel("Golf"))
	fmt.Printf("%+v\n", car2)

	car3 := getCar(WithMake("VW"), WithYear(2024))
	fmt.Printf("%+v\n", car3)

}
