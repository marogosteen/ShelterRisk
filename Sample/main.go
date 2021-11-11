package main

import (
	"fmt"
)

func main() {
	type foo struct {
		y int
		x int
	}

	dimSize := 100
	step := 0

	for inputID := 99; inputID < 101; inputID++ {
		if dimSize%inputID == 0 {
			step = dimSize / inputID
		} else {
			step = dimSize/inputID + 1
		}

		peoplecount := 0
		count := 0
		for i := 0; i < 10; i += 2 {
			for j := 0; j < 20; j++ {
				if count%step == 0 {
					fmt.Println(foo{y: i, x: j / 2})
					peoplecount++
				}
				count++
			}
		}
		//var step float64
		//step = float64(inputID) / float64(dimSize)
		//割り切れる分を先に作って割り切れない分をスライドさせる。↑使う

		// peoplecount := 0
		// count := 0.0
		// for i := 0; i < 10; i += 2 {
		// 	for j := 0; j < 20; j++ {
		// 		count += step
		// 		if count >= 1.0 {
		// 			//fmt.Println(foo{y: i, x: j / 2})
		// 			peoplecount++
		// 			count -= 1.0
		// 		}

		// 	}
		// }

		fmt.Println("count", count)
		fmt.Println("peoplecount", peoplecount)
		fmt.Println("inputID", inputID)
		fmt.Println("step", step)
		fmt.Println()
	}

}

// インプットIDを５０通り
// forループでかこう。
