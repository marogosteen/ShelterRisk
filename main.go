package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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
	mapSize := person.Position{Y: config.MapSizeY, X: config.MapSizeX}

	//simulationの設定
	s, err := simulation.NewSimulation(mapSize, config.GridCapacity, 7, people)
	if err != nil {
		log.Fatal(err)
	}

	interval, err := time.ParseDuration(strconv.Itoa(config.TimeInterval) + "s")
	if err != nil {
		log.Fatal(err)
	}
	s.Run(interval)

	movercount2 := 0
	infectedcount := 0
	for _, p := range s.People {
		fmt.Printf("%+v,\n", p)
		if p.InfectionStatus != person.Health {
			infectedcount++
		}
		if p.LifeAction != person.Stay {
			movercount2++
		}
	}
	movercount1 := 0
	for key, pl := range s.MoversPosition {
		movercount1 += len(pl)
		fmt.Printf("%v %+v\n", key, pl)
	}
	fmt.Println()
	fmt.Println("infectedcount", infectedcount)
	fmt.Println("movercount", movercount1)
	fmt.Println("movercount2", movercount2)

	fmt.Printf("\nDone!\n")
}
