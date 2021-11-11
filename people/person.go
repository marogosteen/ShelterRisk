package people

import (
	"math/rand"

	"example/OSURisk/coodinate"
)

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
	Id               int                 // ID
	NowPosition      coodinate.Coodinate // 現在地
	StartCoodinate   coodinate.Coodinate // スタート地点
	EventElapsedTime int                 // イベント経過時間
	InfectionStatus  string              // 感染状況
}

// TODO 壁判定
// TODO 指向性持たせたい
func (p *Person) Move(mapSize coodinate.Coodinate) {
	var nextPosition coodinate.Coodinate
	for {
		nextPosition = p.NowPosition.Move()
		isCollision := collisionDetection(nextPosition, mapSize)
		if !isCollision {
			break
		}
	}
	p.NowPosition = nextPosition
}

func collisionDetection(nextPosition coodinate.Coodinate, mapSize coodinate.Coodinate) bool {
	collision := mapSize.X < nextPosition.X ||
		mapSize.Y < nextPosition.Y ||
		0 > nextPosition.X ||
		0 > nextPosition.Y
	return collision
}

// TODO IncubationからInfectionになるプログラム
func (p *Person) InfectionTest() string{
	infectionThreshold := 0.1
	if infectionThreshold > rand.Float64() {
		return EnumInfectionStatus.Incubation
	}
	return EnumInfectionStatus.Health
}

// TODO ご飯は時間ベースでDecideする
func (p *Person) EventDecide() {

	//for i:=0; i<3; i++{
	//
}
