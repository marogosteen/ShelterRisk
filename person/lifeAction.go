package person

import (
	"math/rand"
)

// トイレや食事などの生活する上での活動を表現したEnum。
type LifeAction string

const (
	Stay         LifeAction = "Stay"
	Stroll       LifeAction = "Stroll"
	CheckBoard   LifeAction = "CheckBoard"
	ChangeClthes LifeAction = "ChangeClthes"
	BathRoom     LifeAction = "BathRoom"
	Meal         LifeAction = "Meal"
	GoHome       LifeAction = "GoHome"
)

var (
	// KeyをLifeAction、Valueを確率値としたMap。強制に設定させるEatやGobackは含まれない。
	probabilityMap   map[LifeAction]float32
	NecessaryTimeMap map[LifeAction]float64
)

func init() {
	// TODO Test実装して、１になるか確認できるようにするべき。
	// ランダムで決まるActionのそれぞれの確率
	probabilityMap = map[LifeAction]float32{
		Stay:         0.89,
		Stroll:       0.03,
		CheckBoard:   0.02,
		ChangeClthes: 0.03,
		BathRoom:     0.03,
	}

	NecessaryTimeMap = map[LifeAction]float64{
		Stay:   3600.0,
		Stroll: 600.0,
	}
}

// probabilityMapからランダムにLifeActionを決める。
func getRandomAction() LifeAction {
	randNum := rand.Float32()
	var actionProbability float32 = 0.0
	var resultAction LifeAction
	probabilityMapKey := []LifeAction{Stay, BathRoom, ChangeClthes, CheckBoard, Stroll}
	for _, lifeAction := range probabilityMapKey {
		value := probabilityMap[lifeAction]
		actionProbability += value
		if actionProbability > randNum {
			resultAction = lifeAction
			break
		}
	}
	return resultAction
}
