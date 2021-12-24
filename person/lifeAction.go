package person

import (
	"math/rand"
)

func init() {
	// TODO Test実装して、１になるか確認できるようにするべき。
	// ランダムで決まるActionのそれぞれの確率
	probability = map[LifeAction]float32{
		Stay:         0.89,
		Stroll:       0.03,
		CheckBoard:   0.02,
		ChangeClthes: 0.03,
		RestRoom:     0.03,
	}

	TimeRequired = map[LifeAction]float64{
		Stay:   3600.0,
		Stroll: 600.0,
	}
}

// トイレや食事などの生活する上での活動を表現したEnum。
type LifeAction string

const (
	Stay         LifeAction = "Stay"
	Stroll       LifeAction = "Stroll"
	CheckBoard   LifeAction = "CheckBoard"
	ChangeClthes LifeAction = "ChangeClthes"
	RestRoom     LifeAction = "RestRoom"
	Meal         LifeAction = "Meal"
	GoHome       LifeAction = "GoHome"
)

var (
	// KeyをLifeAction、Valueを確率値としたMap。強制に設定させるEatやGobackは含まれない。
	probability map[LifeAction]float32
	// KeyをLifeAction、ValueをStayとStrollの所要時間としたMap。
	TimeRequired map[LifeAction]float64
	// 確率で決まるLifeActionの配列
	key [5]LifeAction = [5]LifeAction{Stay, RestRoom, ChangeClthes, CheckBoard, Stroll}
)

// probabilityMapからランダムにLifeActionを決める。
func getRandomAction() LifeAction {
	randNum := rand.Float32()
	var sumP float32 = 0.0
	resultAction := key[len(key)-1]
	for _, lifeAction := range key {
		sumP += probability[lifeAction]
		if sumP > randNum {
			resultAction = lifeAction
			break
		}
	}

	return resultAction
}
