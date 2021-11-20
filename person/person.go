package person

import (
	"math"
	"math/rand"

	"example/OSURisk/config"
)

// 一人の人間を表現したStruct。
type PersonModel struct {
	Id           int      // ID
	NowPosition  Position // 現在地
	HomePosition Position // スタート地点
	// Distination           Position        // 目的地
	PassedCount           int             // 目的地の通過数
	InfectionStatus       InfectionStatus // 感染状況
	LifeAction            LifeAction      // 生活活動
	LifeActionElapsedTime int
}

func NewPerson(id int, homePosition Position) (p *PersonModel) {
	p = &PersonModel{
		Id:              id,
		HomePosition:    homePosition,
		NowPosition:     homePosition,
		InfectionStatus: Health,
		LifeAction:      Stay,
	}
	p.setNextLifeAction()
	return p
}

// diffSec分、personのLifeActionElapsedTimeを加算する。
func (p *PersonModel) Stay(diffSec int) {
	p.LifeActionElapsedTime += diffSec
}

// ランダムでPerson.NowPositionを周囲８方に変える。1%の確率でp.NowPositionが変化しない。
func (p *PersonModel) Stroll(diffSec int, mapSize Position) (nextPosition Position) {
	p.LifeActionElapsedTime += diffSec

	if 0.01 > rand.Float32() {
		return p.NowPosition
	}

	for {
		var (
			x_course int
			y_course int
		)

		x_course = rand.Intn(2+1) - 1
		y_course = rand.Intn(2+1) - 1
		if x_course == 0 && y_course == 0 {
			continue
		}

		nextPosition = Position{
			X: p.NowPosition.X + x_course,
			Y: p.NowPosition.Y + y_course,
		}

		isCollision := collisionDetection(nextPosition, mapSize)
		if isCollision {
			continue
		}

		break
	}

	return nextPosition
}

// TODO 指向性持たせたい
// PersonのNowPositionをdistination方向に変化させる。
func (p *PersonModel) Move(mapSize Position) (nextPosition Position) {
	var distination Position
	if p.LifeAction == GoHome {
		distination = p.HomePosition
	}else {
		distination = getDistination(p)
	}
	for {
		hogeY := distination.Y - p.NowPosition.Y
		hogeX := distination.X - p.NowPosition.X

		if int(math.Abs(float64(hogeY))) > int(math.Abs(float64(hogeX))) {
			var fuga int
			if hogeY > 0 {
				fuga = 1
			} else {
				fuga = -1
			}
			nextPosition = Position{p.NowPosition.Y + fuga, p.NowPosition.X}
		} else {
			var fuga int
			if hogeX > 0 {
				fuga = 1
			} else {
				fuga = -1
			}
			nextPosition = Position{p.NowPosition.Y, p.NowPosition.X + fuga}
		}

		isCollision := collisionDetection(nextPosition, mapSize)
		if !isCollision {
			break
		}
	}

	return nextPosition
}

// LifeActionが完了したかをboolで返す
func (p *PersonModel) IsDone() (isDone bool) {
	isDone = false
	switch p.LifeAction {
	case Stay, Stroll:
		if p.LifeActionElapsedTime > NecessaryTimeMap[p.LifeAction] {
			isDone = true
		}
	default:
		distination := getDistination(p)
		if p.NowPosition == distination {
			isDone = true
		}
	}

	return isDone
}

// 次のDistinationをSetする。最終目標地に到達した場合は、Actionを変更する。
func (p *PersonModel) SetNextDistination() {
	p.PassedCount++

	var isGoaled bool
	switch p.LifeAction {
	case Stay, Stroll:
		isGoaled = true
	default:
		isGoaled = p.PassedCount == getPassedPoint(p)
	}

	if isGoaled {
		p.setNextLifeAction()
	}
}

// 次のActionとDistinationをSetする。ActionがGoHomeでない場合（現在地がHomePositionでない場合）は、
// StayイベントがGoHomeとなる。
func (p *PersonModel) setNextLifeAction() {
	p.PassedCount = 0
	p.LifeActionElapsedTime = 0
	nextLifeAction := getRandomAction()
	if nextLifeAction == Stay && p.LifeAction != GoHome && p.LifeAction != Stay {
		nextLifeAction = GoHome
	}
	p.LifeAction = nextLifeAction
	// p.Distination = DistinationListMap[p.LifeAction][p.PassedCount]
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
func (p *PersonModel) InfectionJudge() InfectionStatus {
	infectionProbability := config.Config.InfectionProbability
	if infectionProbability > rand.Float64() {
		return Infection
	}
	return Health
}
