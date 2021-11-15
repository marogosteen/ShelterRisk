package person

import (
	"example/OSURisk/config"
	"math/rand"
)

// 一人の人間を表現したStruct。
type Person struct {
	Id                    int             // ID
	NowPosition           Position        // 現在地
	HomePosition          Position        // スタート地点
	Distination           Position        // 目的地
	PassedCount           int             // 目的地の通過数
	InfectionStatus       InfectionStatus // 感染状況
	LifeAction            LifeAction      // 生活活動
	LifeActionElapsedTime int
}

func (p *Person) Stay(diffSec int) {
	p.LifeActionElapsedTime += diffSec
}

func (p *Person) Stroll(diffSec int, mapSize Position) {
	p.LifeActionElapsedTime += diffSec

	var nextPosition Position
	for {
		// TODO Moveはこっちに移動したい
		nextPosition = p.NowPosition.Move(p.Distination)
		isCollision := collisionDetection(nextPosition, mapSize)
		if !isCollision {
			break
		}
	}

	p.NowPosition = nextPosition
}

// TODO 指向性持たせたい
// PersonのNowPositionをdistination方向に変化させる。
func (p *Person) Move(mapSize Position) {
	var nextPosition Position
	for {
		// TODO Moveはこっちに移動したい
		nextPosition = p.NowPosition.Move(p.Distination)
		isCollision := collisionDetection(nextPosition, mapSize)
		if !isCollision {
			break
		}
	}

	p.NowPosition = nextPosition
}

// LifeActionが完了したかをboolで返す
func (p *Person) IsDone() (isDone bool) {
	isDone = false
	switch p.LifeAction {
	case Stay, Stroll:
		if p.LifeActionElapsedTime > NecessaryTimeMap[p.LifeAction] {
			isDone = true
		}
	default:
		distination := DistinationListMap[p.LifeAction][p.PassedCount]
		if p.NowPosition == distination {
			isDone = true
		}
	}
	return isDone
}

// 次のDistinationをSetする。最終目標地に到達した場合は、Actionを変更する。
func (p *Person) SetNextDistination() {
	p.PassedCount++

	var isGoaled bool
	switch p.LifeAction {
	case Stay, Stroll:
		isGoaled = true
	default:
		isGoaled = p.PassedCount == len(DistinationListMap[p.LifeAction])
	}

	if isGoaled {
		p.setNextLifeAction()
	}
}

// 次のActionとDistinationをSetする。ActionがGoHomeでない場合（現在地がHomePositionでない場合）は、
// StayイベントがGoHomeとなる。
func (p *Person) setNextLifeAction() {
	p.PassedCount = 0
	p.LifeActionElapsedTime = 0
	nextLifeAction := GetRandomAction()
	if p.LifeAction != GoHome && nextLifeAction == Stay {
		nextLifeAction = GoHome
	}
	p.LifeAction = nextLifeAction
	p.Distination = DistinationListMap[p.LifeAction][p.PassedCount]
}

// MapSize以上に移動しているかを判定する
func collisionDetection(nextPosition Position, mapSize Position) bool {
	collision := mapSize.X < nextPosition.X ||
		mapSize.Y < nextPosition.Y ||
		0 > nextPosition.X ||
		0 > nextPosition.Y
	return collision
}

// Configで設定した確率で感染者と判定する。
func (p *Person) InfectionJudge() InfectionStatus {
	infectionProbability := config.Config.InfectionProbability
	if infectionProbability > rand.Float64() {
		return Infection
	}
	return Health
}
