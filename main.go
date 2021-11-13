package main

import (
	"fmt"
	"math/rand"
	"time"

	"example/OSURisk/config"
	"example/OSURisk/coodinate"
	"example/OSURisk/people"
	"example/OSURisk/simulation"
)

// TODO 10人だと住居スペースが縦に並ぶ。

func main() {
	var config = config.Config
	rand.Seed(time.Now().Unix())

	// 3回/1日 実施 Map size (11マス*11マス) (20m*20m) 1800sec
	// diningSimulation := simulations.Simulation{
	// 	MapSize: coodinate.Coodinate{X: 11, Y: 11},
	// 	EndSec:  1800,
	// 	People:  people.GeneratePeople(50, 2),
	// }
	// diningSimulation.Run(diffSec)

	//428400, //17時間×７日 (17hour × 60min × 60sec × 7days)
	// parameterはConfigで指定する。
	simulation := simulation.Simulation{
		// 11マス*11マス以下のmapSize指定はError吐くべきでは？ あるいは固定する。
		MapSize: coodinate.Coodinate{Y: config.MapSizeY, X: config.MapSizeX},
		EndSec:  428400,
		People:  people.GeneratePeople(config.PeopleCount, config.InfectedCount),
	}
	simulation.GymRun(config.TimeInterval)

	fmt.Printf("\nDone!\n")
}
