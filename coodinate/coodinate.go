package coodinate

import "math/rand"

type Coodinate struct {
	Y int
	X int
}

func (c *Coodinate) Move() Coodinate {
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

	return Coodinate{
		X: c.X + x_course,
		Y: c.Y + y_course,
	}
}
