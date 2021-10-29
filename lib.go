package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_kuku() {
	var kuku_grid [2][9]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 9; j++ {
			kuku_grid[i][j] = (i + 7) * (j + 1)
		}
	}

	fmt.Println(kuku_grid)

}

func suisoku_game() {
	var (
		userinput    int
		secretNumber int
	)
	var quit = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)

	// 無限ループさせたい
	for quit != true {
		// 入力を受け付けたい
		fmt.Printf("Please enter a number ")
		fmt.Scan(&userinput)

		// 入力値とシークレットが等価かどうか調べる
		if userinput == secretNumber {
			fmt.Println("you won")
			quit = true
		} else if userinput < secretNumber {
			fmt.Println("you low")
		} else if userinput > secretNumber {
			fmt.Println("you high")

		}
	}
}

func zihanki(money int) {
	var juice int

	for i := 0; i < 2; i++ {
		// 入力を受け付けたい
		fmt.Printf("Please enter a number ")
		fmt.Scan(&juice)

		money = money - juice
		fmt.Println("お釣り >> ", money)
	}

	count_coin(money)

	fmt.Printf("\nDone!!\n")
}

func count_coin(money int) {
	fmt.Println("500円 >> ", money/500, "枚")
	money = money - ((money / 500) * 500)

	fmt.Println("100円 >> ", money/100, "枚")
	money = money - ((money / 100) * 100)

	fmt.Println("50円 >> ", money/50, "枚")
	money = money - ((money / 50) * 50)

	fmt.Println("10円 >> ", money/10, "枚")
	money = money - ((money / 10) * 10)
}

func add(v1 int, v2 int) int {
	result := v1 + v2
	return result
}

func fizz_buzz() {

	for i := 1; i < 36; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz ")
		} else {
			fmt.Println(i)
		}
	}
}
