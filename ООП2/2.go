package main

import (
	"fmt"
)

type Vehicle interface {
	Start()
	Stop()
}

type Car struct {
	brand string
}

func (c Car) Start() {
	fmt.Printf("Starting the %s car\n", c.brand)
}

func (c Car) Stop() {
	fmt.Printf("Stopping the %s car\n", c.brand)
}

type Motorcycle struct {
	brand string
}

func (m Motorcycle) Start() {
	fmt.Printf("Starting the %s motorcycle\n", m.brand)
}

func (m Motorcycle) Stop() {
	fmt.Printf("Stopping the %s motorcycle\n", m.brand)
}

func main() {
	car := Car{brand: "Toyota"}
	motorcycle := Motorcycle{brand: "Honda"}

	car.Start()
	car.Stop()

	motorcycle.Start()
	motorcycle.Stop()
}
