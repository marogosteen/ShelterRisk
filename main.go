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
	simu run MoversPositionMapの移動の表現
	[]int...の動作確認
	10人だと住居スペースが縦に並ぶ。
	Moveまだ
	Eatがまだ　Eatは絶対ホームに戻る
	Simulation.PositionsMapの動作確認
	渋滞の表現。
		Actionの目的地のindex0に渋滞者の現在地を追加する。
		Moveは常に実行前に目的地をPassedでチェックする。
	LifeActionElapsedTimeのInitilizeがいるのでは？？
	stay と　Strollは乱数で前後させたいな。
*/

func main() {
	var config = config.Config
	rand.Seed(time.Now().Unix())

	//simulationの設定
	simulation := simulation.Simulation{
		// TODO 11マス*11マス以下のmapSize指定はError吐くべきでは？
		MapSize:      person.Position{Y: config.MapSizeY, X: config.MapSizeX},
		GridCapacity: config.GridCapacity,
		EndSec:       428400,
		People:       simulation.GeneratePeople(config.PeopleCount, config.InfectedCount),
	}
	simulation.Run(config.TimeInterval)

	fmt.Printf("\nDone!\n")
}
