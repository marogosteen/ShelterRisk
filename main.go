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
	Eatがまだ　Eatは絶対ホームに戻る
	Simulation.PositionsMapの動作確認
	渋滞の表現。
		Actionの目的地のindex0に渋滞者の現在地を追加する。
		Moveは常に実行前に目的地をPassedでチェックする。
	LifeActionElapsedTimeのInitilizeがいるのでは？？
	stay と　Strollは乱数で前後させたいな。
	move 斜め移動
*/

func main() {
	var config = config.Config
	rand.Seed(time.Now().Unix())

	people := simulation.NewPeople(config.PeopleCount)
	people.SetInfected(config.InfectedCount)
	moversPosition := simulation.GenerateMoversPosition(people)

	//simulationの設定
	simulation := simulation.Simulation{
		// TODO 11マス*11マス以下のmapSize指定はError吐くべきでは？
		MapSize:        person.Position{Y: config.MapSizeY, X: config.MapSizeX},
		GridCapacity:   config.GridCapacity,
		EndSec:         428400,
		People:         people,
		MoversPosition: moversPosition,
	}
	simulation.Run(config.TimeInterval)

	fmt.Printf("\nDone!\n")
}
