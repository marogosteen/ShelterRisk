package person

import (
	"math/rand"

	"example/OSURisk/config"
)

type Position struct {
	Y int
	X int
}

// 移動するLifeAction毎の目的地
var DistinationListMap map[LifeAction][]Position

func init() {
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
}

// TODO personに持たせたい
func (c *Position) Move(distination Position) Position {
	var (
		x_course int
		y_course int
	)

	for i := 0; i < 2; i++ {
		x_course = rand.Intn(2+1) - 1
		y_course = rand.Intn(2+1) - 1
		if !(x_course == 0 && y_course == 0) {
			break
		}
	}

	return Position{
		X: c.X + x_course,
		Y: c.Y + y_course,
	}
}
