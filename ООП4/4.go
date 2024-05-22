package main

import "fmt"

type Pump interface {
	PumpAction()
}

type BasePump struct {
	Type string
}

func (p BasePump) PumpAction() {
	fmt.Printf("%s насос\n", p.Type)
}

type PlungerPump struct {
	BasePump
}

func (p PlungerPump) PumpAction() {
	fmt.Printf("%s насос делает плунжерное движение\n", p.Type)
}

type CentrifugalPump struct {
	BasePump
}

func (c CentrifugalPump) PumpAction() {
	fmt.Printf("%s насос вращается по центробежной схеме\n", c.Type)
}

func main() {
	plungerPump := PlungerPump{BasePump: BasePump{Type: "Плунжерный"}}
	centrifugalPump := CentrifugalPump{BasePump: BasePump{Type: "Центробежный"}}

	pumps := []Pump{plungerPump, centrifugalPump}

	for _, pump := range pumps {
		pump.PumpAction()
	}
}
