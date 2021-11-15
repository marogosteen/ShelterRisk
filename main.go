package main

import (
	"fmt"
	"math/rand"
	"time"

	"example/OSURisk/config"
	"example/OSURisk/person"
	"example/OSURisk/simulation"
)

/*
TODO
	1マス4人まだ
	10人だと住居スペースが縦に並ぶ。
	Stayの実装がまだ。
		Stayの感染は、しない。させない。
	Eatがまた
	peopleのソースはSimulationに移すべき。
*/

func main() {
	var config = config.Config
	rand.Seed(time.Now().Unix())

	simulation := simulation.Simulation{
		// TODO 11マス*11マス以下のmapSize指定はError吐くべきでは？
		MapSize: person.Position{Y: config.MapSizeY, X: config.MapSizeX},
		EndSec:  428400,
		People:  simulation.GeneratePeople(config.PeopleCount, config.InfectedCount),
	}
	simulation.GymRun(config.TimeInterval)

	fmt.Printf("\nDone!\n")
}
