package main

import (
	"fmt"
	"math/rand"
	"time"

	"example/OSURisk/people"
	"example/OSURisk/simulations"
)

func main() {
	diffSec := 3
	rand.Seed(time.Now().Unix())
	people := people.NewPeople(100, 2)
	diningSimulation := simulations.NewDiningSimulation(*people)
	// for _, person := range diningSimulation.People.PersonList {
	// 	fmt.Println(person)
	// }
	diningSimulation.Run(diffSec)
	// simulations.NewGymSimulation()
	for _, person := range diningSimulation.People.PersonList {
		fmt.Println(person)
	}
	fmt.Printf("\nDone\n")
}
