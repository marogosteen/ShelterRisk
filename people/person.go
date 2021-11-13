package people

import (
	"math/rand"

	"example/OSURisk/action"
	"example/OSURisk/coodinate"
	"example/OSURisk/infectionStatus"
)

// 一人の人間を表現したStruct。
type Person struct {
	Id              int                             // ID
	NowPosition     coodinate.Coodinate             // 現在地
	StartPosition   coodinate.Coodinate             // スタート地点
	InfectionStatus infectionStatus.InfectionStatus // 感染状況
	Action          action.Action                   // 生活活動
}

// TODO 指向性持たせたい
// PersonのNowPositionを変化させる。
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

// MapSize以上に移動しているかを判定する
func collisionDetection(nextPosition coodinate.Coodinate, mapSize coodinate.Coodinate) bool {
	collision := mapSize.X < nextPosition.X ||
		mapSize.Y < nextPosition.Y ||
		0 > nextPosition.X ||
		0 > nextPosition.Y
	return collision
}

// 一定の確率で感染者と判定する。
func (p *Person) InfectionJudge() infectionStatus.InfectionStatus {
	infectionThreshold := 0.1
	if infectionThreshold > rand.Float64() {
		return infectionStatus.Infection
	}
	return infectionStatus.Health
}
