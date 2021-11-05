package person

import (
	"fmt"
	"math/rand"
)

type coodinate struct {
	x int
	y int
}

type infectionStatus struct {
	Health     string // 健全
	Incubation string // 潜伏期間
	Infection  string // 感染
}

var foo = infectionStatus{
	Health:     "Health",
	Incubation: "Incubation",
	Infection:  "Infection",
}

type Person struct {
	Id               int       // ID
	CurrentCoodinate coodinate // 現在地
	StartCoodinate   coodinate // スタート地点
	EventelapsedTime int       // イベント経過時間
	InfectionStatus  string    // 感染状況
}

func NewPerson(inputId int) *Person {
	coodinate := coodinate{x: 0, y: 1}
	p := Person{
		Id:               inputId,
		CurrentCoodinate: coodinate,
		StartCoodinate:   coodinate,
		EventelapsedTime: 0,
		InfectionStatus:  foo.Health,
	}
	return &p
}

func (p *Person) ShowPerson() {
	fmt.Printf(
		"ID: %v, CCy:%v, CCx:%v, SCy:%v, SCx:%v, EventTime:%v\n",
		p.Id, p.CurrentCoodinate.y, p.CurrentCoodinate.x,
		p.StartCoodinate.y, p.StartCoodinate.x, p.EventelapsedTime,
	)
}

// TODO 壁判定
// TODO 指向性持たせたい
// TODO Simulation Mapを関与させたい
// func (p *Person) Move(xMax int, yMax int) {
func (p *Person) Move() {
	x_direction := rand.Intn(2+1) - 1
	y_direction := rand.Intn(2+1) - 1
	if x_direction == 0 && y_direction == 0{
		x_direction = rand.Intn(2+1) - 1
		y_direction = rand.Intn(2+1) - 1
	}

	p.CurrentCoodinate.x += rand.Intn(2+1) - 1
	p.CurrentCoodinate.y += rand.Intn(2+1) - 1
}

// TODO IncubationからInfectionになるプログラム
func (p *Person) InfectionTest() {
	infectionThreshold := 0.1
	if infectionThreshold > rand.Float64() {
		p.InfectionStatus = foo.Incubation
	}
}

// TODO ご飯は時間ベースでDecideする
func (p *Person) EventDecide() {
	//for i:=0; i<3; i++{
		//
	}
	//
	

	// stay:=0.75
	//
	// rand.Seed(Intn(4))
	// secretNumber= rand.Str(4)
	// for i :=0; i<

