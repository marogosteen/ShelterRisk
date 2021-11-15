package person

import (
	"math/rand"
)

type Position struct {
	Y int
	X int
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
