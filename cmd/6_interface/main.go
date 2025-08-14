package main

import "fmt"

type GasEngine struct {
	cylinders uint8
	fuelType  string
}

func (g GasEngine) start() {
	fmt.Println("Starting gas engine with", g.cylinders, "cylinders and fuel type", g.fuelType)
}

type PetrolEngine struct {
	capacity uint8
}

func (p *PetrolEngine) setCapacity(capacity uint8) {
	p.capacity = capacity
}

func (p PetrolEngine) getCapacity() uint8 {
	return p.capacity
}

func (p PetrolEngine) start() {
	fmt.Println("Starting petrol engine with", p.capacity, "liters of capacity")
}

type Engine interface {
	start()
}

func startAnyEngine(engine Engine) {
	engine.start()
}

func main() {
	gasEngine := GasEngine{
		cylinders: 4,
		fuelType:  "Petrol",
	}
	gasEngine.start()

	petrolEngine := PetrolEngine{60}
	petrolEngine.setCapacity(50)
	fmt.Println(petrolEngine.getCapacity())
	petrolEngine.start()
	startAnyEngine(gasEngine)
	startAnyEngine(petrolEngine)
}
