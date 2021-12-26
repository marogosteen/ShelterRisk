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
	10人だと住居スペースが縦に並ぶ。
	Eatは絶対ホームに戻る
	LifeActionElapsedTimeのInitilizeがいるのでは？？
	stay と　Strollは乱数で前後させたいな。
*/

func main() {
	var config = config.Config
	rand.Seed(time.Now().Unix())

	people := simulation.NewPeople(config.PeopleCount)
	people.SetInfected(config.InfectedCount)
	mapSize := person.Position{Y: config.MapSizeY, X: config.MapSizeX}

	//simulationの設定
	s, err := simulation.NewSimulationModel(mapSize, config.GridCapacity, 7, people)
	if err != nil {
		log.Fatal(err)
	}

	interval, err := time.ParseDuration(strconv.Itoa(config.TimeInterval) + "s")
	if err != nil {
		log.Fatal(err)
	}
	s.Run(interval)

	fmt.Printf("\nDone!\n")
}
