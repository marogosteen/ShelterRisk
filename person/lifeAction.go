package person

import (
	"math/rand"
)

// トイレや食事などの生活する上での活動を表現したEnum。
type LifeAction int

const (
	Stay LifeAction = iota
	CheckBoard
	ChangeClthes
	BathRoom
	Eat
	GoHome
)

// ランダムで決まるActionのMap。強制に設定させるEatやGobackは含まれない。
var probabilityMap map[LifeAction]float32

func init() {
	// ランダムで決まるActionのそれぞれの確率
	probabilityMap = map[LifeAction]float32{
		Stay:         0.90,
		CheckBoard:   0.04,
		ChangeClthes: 0.03,
		BathRoom:     0.03,
	}
}

// probabilityMapからランダムにActionを決める。
func GetRandomAction() LifeAction {
	randNum := rand.Float32()
	var actionProbability float32 = 0.0
	var resultAction LifeAction
	for action, value := range probabilityMap {
		actionProbability += value
		if actionProbability > randNum {
			resultAction = action
		}
	}
	return resultAction
}
