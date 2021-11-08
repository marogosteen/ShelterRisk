package main

import (
	"example/OSURisk/simulations"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	simulations.NewDiningSimulation()
	simulations.NewGymSimulation()
}
