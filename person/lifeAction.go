package person

import (
	"example/OSURisk/config"
	"math/rand"
)

// トイレや食事などの生活する上での活動を表現したEnum。
type LifeAction int

const (
	Stay LifeAction = iota
	Stroll
	CheckBoard
	ChangeClthes
	BathRoom
	Eat
	GoHome
)

// ランダムで決まるActionのMap。強制に設定させるEatやGobackは含まれない。
var (
	probabilityMap map[LifeAction]float32
	// 移動するLifeAction毎の目的地
	DistinationListMap map[LifeAction][]Position
	NecessaryTimeMap   map[LifeAction]int
)

func init() {
	// ランダムで決まるActionのそれぞれの確率
	probabilityMap = map[LifeAction]float32{
		Stay:         0.88,
		Stroll:       0.03,
		CheckBoard:   0.03,
		ChangeClthes: 0.03,
		BathRoom:     0.03,
	}

	mapSizeX := config.Config.MapSizeX
	mapSizeY := config.Config.MapSizeY

	DistinationListMap = map[LifeAction][]Position{
		CheckBoard:   {Position{Y: mapSizeY, X: mapSizeX}},
		ChangeClthes: {Position{Y: mapSizeY / 2, X: mapSizeX}},
		BathRoom:     {Position{Y: 0, X: mapSizeX}},
		Eat: {
			Position{Y: mapSizeY - 5, X: 0},
			Position{Y: mapSizeY, X: 0},
			Position{Y: mapSizeY, X: 5},
		},
	}

	NecessaryTimeMap = map[LifeAction]int{
		Stay:   3600,
		Stroll: 600,
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
