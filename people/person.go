package people

import (
	"example/OSURisk/coodinate"
	"math/rand"
)

// 一人の人間を表現したStruct。
type Person struct {
	Id              int                             // ID
	NowPosition     coodinate.Coodinate             // 現在地
	HomePosition    coodinate.Coodinate             // スタート地点
	Distination     coodinate.Coodinate             // 目的地
	InfectionStatus InfectionStatus // 感染状況
	LifeAction      LifeAction           // 生活活動
}

// 目的地をActionとMapSizeから求める。初期化時と、目的地に到達時に目的地を変更する。
// Eatの移動は複雑。
// ConfigでSetするのもあり。

// TODO 指向性持たせたい
// PersonのNowPositionを変化させる。
func (p *Person) Move(mapSize coodinate.Coodinate) {
	/*
		毎度DistinationをSetするのは無駄
		Personのfieldに持たすべき？？
	*/

	distination := p.setDistination()

	// 目的地に到達した場合、次のGoBackに変更する。HomePositionに戻った場合、次のActionをSet
	if p.NowPosition == distination {
		if p.LifeAction == GoBack {
			p.LifeAction = GetRandomAction()
		} else {
			p.LifeAction = GoBack
		}
		p.LifeAction = GoBack
	}

	var nextPosition coodinate.Coodinate
	for {
		nextPosition = p.NowPosition.Move(distination)
		isCollision := collisionDetection(nextPosition, mapSize)
		if !isCollision {
			break
		}
	}

	p.NowPosition = nextPosition
}

// 目的地に到達した場合、次のGoBackに変更する。HomePositionに戻った場合、次のActionをSet
func (p *Person) setDistination() coodinate.Coodinate {
	if p.LifeAction == GoBack {
		p.LifeAction = GetRandomAction()
	} else {
		p.LifeAction = GoBack
	}
	return coodinate.Coodinate{}
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
func (p *Person) InfectionJudge() InfectionStatus {
	infectionThreshold := 0.1
	if infectionThreshold > rand.Float64() {
		return Infection
	}
	return Health
}
