package main

import (
	"fmt"
	"math/rand"
	"time"

	"example/OSURisk/coodinate"
	"example/OSURisk/people"
	"example/OSURisk/simulations"
)

func main() {
	diffSec := 3
	rand.Seed(time.Now().Unix())

	// 3回/1日 実施 Map size (11マス*11マス) (20m*20m) 1800sec
	diningSimulation := simulations.Simulation{
		MapSize: coodinate.Coodinate{X: 11, Y: 10},
		EndSec:  1800,
		People:  people.GeneratePeople(100, 2),
	}

	//428400, //17時間×７日 (17hour × 60min × 60sec × 7days)

	// diningSimulation := simulations.NewDiningSimulation(*people)
	// for _, person := range diningSimulation.People.PersonList {
	// 	fmt.Println(person)
	// }
	diningSimulation.Run(diffSec)
	// simulations.NewGymSimulation()
	infectionCount := 0
	for _, person := range diningSimulation.People {
		if person.InfectionStatus != people.EnumInfectionStatus.Health {
			infectionCount++
		}
		fmt.Printf("%+v\n", person)
	}
	fmt.Println(infectionCount)

	fmt.Printf("\nDone!\n")
}
