package simulations

type Simulation struct {
	Map        [][]int
	CurrentSec int
	EndSec     int
}

func generateMap(xSize int, ySize int) [][]int{
	mapYSize := 21
	mapXSize := 13
	simulationMap := make([][]int, mapYSize)
	for i := 0; i < mapYSize; i++ {
		simulationMap[i] = make([]int, mapXSize)
	}
	return simulationMap
}

// Map size (21, 13)
func NewGymSimulation() *Simulation {
	xMapSize := 13
	yMapSize := 21
	gymSimulation := Simulation{
		Map: generateMap(xMapSize, yMapSize),
		CurrentSec: 0,
		EndSec:     428400, //17時間×７日 (17hour × 60min × 60sec × 7days)
	}
	return &gymSimulation
}

// 3回/1日 実施
// Map size (10, 10)
func NewDiningSimulation() *Simulation {
	xMapSize := 10
	yMapSize := 10
	diningSimulation := Simulation{
		Map:        generateMap(xMapSize, yMapSize),
		CurrentSec: 0,
		EndSec:     1800,
	}
	return &diningSimulation
}

func (s *Simulation) Run(diffSec int) {
	currentSec := 0
}
//tauch＆goでシミュレーションする。
//乱数で決めるため、全体からの％いる？
//bathroom:30/930(3.325%),CB:20/930(2.150%),CC:20/930(2.150%),
//乱数でぶらぶらして、目的地に行くのは現実てきか？
