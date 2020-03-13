package main

import (
	"fmt"
)

type player struct {
	res   resource
	cards *card
	act   *action
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
	titaniumProd int
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
	p.act = nil
}

func (r *resource) init() {
	r.terraforming =	20
	r.mc =			20
	r.mcProd =		0
	r.plant =		0
	r.plantProd =		0
	r.steel =		0
	r.steelProd =		0
	r.titanium =		0
	r.titaniumProd =	0
	r.energy =		0
	r.energyProd =		0
	r.heat =		0
	r.heatProd =		0
	r.victoryPoint =	0
}

func (p *player) print() {
	(*p).res.print()
	fmt.Printf("%v\n", p.cards)
	fmt.Printf("%v\n", p.act)
}

func (r *resource) print() {
	fmt.Printf("tf:%2d	mc:%2d 	plant:%2d	steel:%3d 	ti:%3d 	en:%2d	heat:%3d\n",
		r.terraforming, r.mc, r.plant, r.steel, r.titanium, r.energy, r.heat)
	fmt.Printf("vp:%2d 	mcP:%2d plantP:%2d	steelP:%2d	tiP:%2d enP:%2d	heatP:%2d\n",
		r.victoryPoint, r.mcProd, r.plantProd, r.steelProd, r.titaniumProd, r.energyProd, r.heatProd)
}

func main() {
	fmt.Printf("Hello Mars\n")

	var p1, p2 player
	p1.init()
	p2.init()
	p1.print()
}
