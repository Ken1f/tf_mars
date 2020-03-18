package main

import (
	"fmt"
)

type player struct {
	res   resource
	cards *card
	act   *action
	tags  tag
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

type tag struct {
	building int
	space    int
	power    int
	science  int
	jovian   int
	earth    int
	planet   int
	microbe  int
	animal   int
	city     int
	event    int
}

type card struct {
	name string
}

type action struct {
	name string
}

func (p *player) Init() {
	(*p).res.Init()
	p.cards = nil
	p.act = nil
	(*p).tags.Init()
}

func (r *resource) Init() {
	r.terraforming = 20
	r.mc = 20
	r.mcProd = 0
	r.plant = 0
	r.plantProd = 0
	r.steel = 0
	r.steelProd = 0
	r.titanium = 0
	r.titaniumProd = 0
	r.energy = 0
	r.energyProd = 0
	r.heat = 0
	r.heatProd = 0
	r.victoryPoint = 0
}

func (t *tag) Init() {
	t.building = 0
	t.space = 0
	t.power = 0
	t.science = 0
	t.jovian = 0
	t.earth = 0
	t.planet = 0
	t.microbe = 0
	t.animal = 0
	t.city = 0
	t.event = 0
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

type mars struct {
	gen   int
	o2    int
	temp  int
	ocean int
}

func (m *mars) Init() {
	m.gen = 1
	m.o2 = 0
	m.temp = 0
	m.ocean = -30
}

func (r *resource) Production() {
	r.mc = r.mc + r.mcProd
	r.plant = r.plant + r.plantProd
	r.steel = r.steel + r.steelProd
	r.titanium = r.titanium + r.titaniumProd
	r.energy = r.energy + r.energyProd
	r.heat = r.heat + r.heatProd
	// draw 4 new cards
}

func (m *mars) isTerraform() bool {
	if m.o2 == 14 && m.temp == 8 && m.ocean == 9 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Printf("Hello Mars\n")

	var p1, p2 player
	var m mars
	m.Init()
	p1.Init()
	p2.Init()
	p1.print()
}
