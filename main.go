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
	Eatがまた　Eatは絶対ホームに戻る
	peopleのソースはSimulationに移すべき。
	Simulation.PositionsMapの動作確認
	散歩まだ
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
	simulation.Run(config.TimeInterval)

	fmt.Printf("\nDone!\n")
}
