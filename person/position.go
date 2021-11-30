package person

import (
	"example/OSURisk/config"
)

type Position struct {
	Y int
	X int
}

// 移動するLifeAction毎の目的地
var distinationListMap map[LifeAction][]Position

func init() {
	mapSizeX := config.Config.MapSizeX
	mapSizeY := config.Config.MapSizeY

	distinationListMap = map[LifeAction][]Position{
		CheckBoard:   {Position{Y: mapSizeY, X: mapSizeX}},
		ChangeClthes: {Position{Y: mapSizeY / 2, X: mapSizeX}},
		BathRoom:     {Position{Y: 0, X: mapSizeX}},
		Eat: {
			Position{Y: mapSizeY - 5, X: 0},
			Position{Y: mapSizeY, X: 0},
			Position{Y: mapSizeY, X: 5},
		},
	}
}

func getPassedPoint(p *PersonModel) int {
	return len(distinationListMap[p.LifeAction])
}

func getDistination(p *PersonModel) (distination Position) {
	if p.LifeAction == GoHome {
		distination = p.HomePosition
	} else {
		distination = distinationListMap[p.LifeAction][p.PassedCount]
	}
	return distination
}
