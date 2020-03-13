package main

import (
	"fmt"
)

type player struct {
	res	resource
	cards	*card
	act	*action
}

type resource struct {
	terraforming int
	mc           int
	mcProd       int
	plant        int
	plantProd    int
	steel        int
	steelProd    int
	titanium     int
	titniumProd  int
	energy       int
	energyProd   int
	heat         int
	heatProd     int
	victoryPoint int
}

type card struct {
	name string
}

type action struct {
	name string
}

func (p *player) init() {
	(*p).res.init()
	p.cards = nil
	p.act	= nil
}

func (r *resource) init() {
	r.terraforming = 20
	r.mc = 20
	r.mcProd = 0
	r.plant = 0
	r.plantProd = 0
	r.steel = 0
	r.steelProd = 0
	r.titanium = 0
	r.titniumProd = 0
	r.energy = 0
	r.energyProd = 0
	r.heat = 0
	r.heatProd = 0
	r.victoryPoint = 0
}

func main() {
	var p1, p2 player
	p1.init()
	p2.init()

	fmt.Printf("Hello Mars\n")
}
