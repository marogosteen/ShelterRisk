package people

import (
	"fmt"
	"math/rand"
)

type Coodinate struct {
	x int
	y int
}

type infectionStatus struct {
	Health     string // 健全
	Incubation string // 潜伏期間
	Infection  string // 感染
}

var EnumInfectionStatus = infectionStatus{
	Health:     "Health",
	Incubation: "Incubation",
	Infection:  "Infection",
}

type Person struct {
	Id               int       // ID
	CurrentCoodinate Coodinate // 現在地
	StartCoodinate   Coodinate // スタート地点
	EventelapsedTime int       // イベント経過時間
	InfectionStatus  string    // 感染状況
}

func NewPerson(inputId int) *Person {
	coodinate := Coodinate{x: 0, y: 1}
	p := Person{
		Id:               inputId,
		CurrentCoodinate: coodinate,
		StartCoodinate:   coodinate,
		EventelapsedTime: 0,
		InfectionStatus:  EnumInfectionStatus.Health,
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
// func (p *Person) Move(xMax int, yMax int) {
func (p *Person) Move() *Person{
	var (
		x_direction int
		y_direction int
	)

	for i := 0; i < 2; i++ {
		x_direction = rand.Intn(2+1) - 1
		y_direction = rand.Intn(2+1) - 1
		if !(x_direction == 0 && y_direction == 0) {
			break
		}
	}

	p.CurrentCoodinate.x += x_direction
	p.CurrentCoodinate.y += y_direction

	return p
}

// TODO IncubationからInfectionになるプログラム
func (p *Person) InfectionTest() {
	infectionThreshold := 0.1
	if infectionThreshold > rand.Float64() {
		p.InfectionStatus = EnumInfectionStatus.Incubation
	}
}

// TODO ご飯は時間ベースでDecideする
func (p *Person) EventDecide() {

	//for i:=0; i<3; i++{
	//
}

//
