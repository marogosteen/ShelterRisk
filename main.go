package main

import (
	"example/OSURisk/person"
	"example/OSURisk/simulation"
	"fmt"
	"math/rand"
	"time"
)

// "fmt"
// "example/OSURisk/person"
// "example/OSURisk/lifeevent"

func main() {
	rand.Seed(time.Now().Unix())
	simulation := simulation.NewSimulation()

	fmt.Println(simulation)

	// Person群の初期化

	//最初の場所
	noma := person.NewPerson(999)
	sugi := person.NewPerson(990)
	fmt.Printf("noma 1:")
	noma.ShowPerson()
	fmt.Printf("sugi 1:")
	sugi.ShowPerson()
	fmt.Println()

	noma.Move()
	sugi.Move()
	fmt.Printf("noma 2:")
	noma.ShowPerson()
	fmt.Printf("sugi 2:")
	sugi.ShowPerson()

}

//make_kuku()
//fizz_buzz()
// zihanki(1000)
//suisoku_game()
//hairetsu()

//go build -o main.exe
