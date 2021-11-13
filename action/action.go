package action

import "math/rand"

type Action int

const (
	stay Action = iota
	checkBoard
	changeClthes
	bathRoom
	eat
)

var probabilityMap map[Action]float32 = map[Action]float32{
	stay:         0.92,
	checkBoard:   0.04,
	changeClthes: 0.04,
	bathRoom:     0.04,
	eat:          0.04,
}

func GetRandomAction() Action {
	randNum := rand.Float32()
	var actionProbability float32 = 0.0
	var resultAction Action
	for action, value := range probabilityMap {
		actionProbability += value
		if actionProbability > randNum {
			resultAction = action
		}
	}
	return resultAction
}
