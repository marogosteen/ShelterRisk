package person

import (
	"log"
	"math"
	"math/rand"
	"time"

	"example/OSURisk/config"
)

// 一人の人間を表現したStruct。
type PersonModel struct {
	// Person ID
	Id int
	// 現在地
	NowPosition Position
	// 住居スペースの座標
	HomePosition Position
	// 目的地の通過数
	PassedCount int
	// 感染状況
	InfectionStatus InfectionStatus
	// 生活活動
	LifeAction LifeAction
	// 時間ベースのLife Actionの経過時間
	LifeActionElapsedSec float64
}

func NewPerson(id int, homePosition Position) (p *PersonModel) {
	p = &PersonModel{
		Id:              id,
		HomePosition:    homePosition,
		NowPosition:     homePosition,
		InfectionStatus: Health,
		LifeAction:      Stay,
	}
	p.setLifeAction()
	return p
}

// personのLifeActionElapsedTimeをinterval秒加算する。
func (p *PersonModel) Stay(interval time.Duration) {
	p.LifeActionElapsedSec += interval.Seconds()
}

// ランダムでPerson.NowPositionを周囲８方に変える。1%の確率でp.NowPositionが変化しない。
func (p *PersonModel) Stroll(interval time.Duration, mapSize Position) (nextPosition Position) {
	p.LifeActionElapsedSec += interval.Seconds()

	if 0.01 > rand.Float32() {
		return p.NowPosition
	}

	for {
		var (
			yCourse int
			xCourse int
		)

		// 移動方向を示す-1から1までの乱数。
		yCourse = rand.Intn(2+1) - 1
		xCourse = rand.Intn(2+1) - 1
		if xCourse == 0 && yCourse == 0 {
			continue
		}

		nextPosition = Position{
			Y: p.NowPosition.Y + yCourse,
			X: p.NowPosition.X + xCourse,
		}

		isCollision := collisionDetection(nextPosition, mapSize)
		if isCollision {
			continue
		}

		break
	}

	return nextPosition
}

// PersonのNowPositionをdistination方向に変化させる。
func (p *PersonModel) Move(mapSize Position) (nextPosition Position) {
	// 目的地座標を選定する。
	var distination Position
	if p.LifeAction == GoHome {
		distination = p.HomePosition
	} else {
		distination = getDistination(p)
	}

	if distination == p.NowPosition {
		log.Fatalf("distinationとnowPositionが同じ座標です。\nperson: %+v\n", p)
	}

	// 目的地との差分。
	diffY := distination.Y - p.NowPosition.Y
	diffX := distination.X - p.NowPosition.X
	absDiffY := int(math.Abs(float64(diffY)))
	absDiffX := int(math.Abs(float64(diffX)))
	yCourse := 0
	xCourse := 0

	// 縦方向と横方向の-1から1の移動量に直す。
	if !(diffY == 0) {
		yCourse = diffY / absDiffY
	}
	if !(diffX == 0) {
		xCourse = diffX / absDiffX
	}

	if absDiffY == absDiffX {
		nextPosition = Position{
			p.NowPosition.Y + yCourse,
			p.NowPosition.X + xCourse,
		}
	} else if absDiffY > absDiffX {
		nextPosition = Position{
			p.NowPosition.Y + yCourse,
			p.NowPosition.X,
		}
	} else {
		nextPosition = Position{
			p.NowPosition.Y,
			p.NowPosition.X + xCourse,
		}
	}

	isCollision := collisionDetection(nextPosition, mapSize)
	if isCollision {
		log.Fatalln("Moveでcollisionの値がTrueになりました。")
	}

	return nextPosition
}

// LifeActionが完了したかをboolで返す
func (p *PersonModel) IsDone() (isDone bool) {
	isDone = false

	switch p.LifeAction {
	// 時間ベースのLifeActionの完了の判定
	case Stay, Stroll:
		if p.LifeActionElapsedSec > TimeRequired[p.LifeAction] {
			isDone = true
		}
	// HomePositionに戻ったか判定
	case GoHome:
		if p.NowPosition == getDistination(p) {
			isDone = true
		}
	// 目的地到達の判定
	default:
		distination := getDistination(p)
		if p.NowPosition == distination {
			isDone = true
		}
	}

	return isDone
}

// 次のDistinationをSetする。最終目標地に到達した場合は、Actionを変更する。
func (p *PersonModel) SetDistination() {
	p.PassedCount++

	var isGoaled bool
	switch p.LifeAction {
	case Stay, Stroll, GoHome:
		isGoaled = true
	default:
		isGoaled = p.PassedCount == getPassedPoint(p)
	}

	// Stay以外は連続でActionしない
	if isGoaled {
		p.PassedCount = 0
		nextLifeAction := p.setLifeAction()
		if p.LifeAction == Stay {
			p.LifeAction = nextLifeAction
			return
		}

		for nextLifeAction == p.LifeAction {
			nextLifeAction = p.setLifeAction()
		}

		p.LifeAction = nextLifeAction
	}
}

// 次のActionとDistinationをSetする。ActionがGoHomeでない場合（現在地がHomePositionでない場合）は、
// StayイベントがGoHomeとなる。
func (p *PersonModel) setLifeAction() LifeAction {
	// LifeAction変更時、PassedCountとLifeActionElapsedSecは必ず０に初期化される。
	p.PassedCount = 0
	p.LifeActionElapsedSec = 0
	var nextLifeAction LifeAction

	switch p.LifeAction {
	// 現在のLifeActionがMealである場合は強制的にGoHomeとなる。
	case Meal:
		nextLifeAction = GoHome
	default:
		nextLifeAction = getRandomAction()
		// nextLifeActionがStayとなったとき、現在地がHomePositionでなければGoHomeとなる。
		// じゃあGoHomeがついたときはどうする？ -> 強制Stay？ -> じゃあ食事の時後のGoHome後は？
		if nextLifeAction == Stay && p.NowPosition != p.HomePosition {
			nextLifeAction = GoHome
		}
	}

	return nextLifeAction
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
