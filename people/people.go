package people

import (
	"math/rand"

	"example/OSURisk/coodinate"
	"example/OSURisk/infectionStatus"
	"example/OSURisk/lifeAction"
)

func GeneratePeople(peopleCount int, infectedPersonCount int) []Person {

	people := make([]Person, peopleCount)
	for id := 0; id < len(people); id++ {
		livingPosition := getLivingPosition(peopleCount, id)
		lifeaction := lifeAction.GetRandomAction()
		// distination := 
		people[id] = Person{
			Id:              id,
			HomePosition:    livingPosition,
			NowPosition:     livingPosition,
			InfectionStatus: infectionStatus.Health,
			LifeAction:      lifeaction,
			// Distination: ,
		}
		// いやいや。。。Actionを元に行き先はわかるんじゃね？？
		// いや、しかも、MoveにはMapSizeがあるから、都合いいのでは？？
	}
	setInfected(people, infectedPersonCount)
	return people
}

func setInfected(people []Person, infectedPersonCount int) {
	// TODO 感染者数がシミュレーション人数より多い場合はエラー
	// if len(p.PersonList) < infectedPersonCount{
	// 	panic()
	// }
	var idList []int
	for i := 0; i < len(people); i++ {
		idList = append(idList, i)
	}
	for i := 0; i < infectedPersonCount; i++ {
		idIndex := rand.Intn(len(idList))
		id := idList[idIndex]
		people[id].InfectionStatus = infectionStatus.Infection
		idList = append(idList[:idIndex], idList[idIndex+1:]...)
	}
}

// TODO diviend < divisor のエラー処理
func getLivingPosition(peopleCount int, personId int) coodinate.Coodinate {
	livingSpaceCapacity := 100
	yLivingSpaceCapacity := livingSpaceCapacity / 10
	xLivingSpaceCapacity := livingSpaceCapacity / yLivingSpaceCapacity
	byDivisible := getByDivisible(livingSpaceCapacity, peopleCount)
	whetherRemainder := personId / byDivisible
	step := livingSpaceCapacity / byDivisible

	livingPositionId := (personId-whetherRemainder*byDivisible)*step +
		step/2*whetherRemainder
	// 住居スペースが１マス２人であることを表現するために/2をした
	livingPositionId /= 2

	return coodinate.Coodinate{
		// 廊下スペースの表現するために*2をした
		Y: livingPositionId / yLivingSpaceCapacity * 2,
		X: livingPositionId % xLivingSpaceCapacity,
	}
}

// TODO diviend < divisor のエラー処理
func getByDivisible(dividend int, divisor int) int {
	for ; ; divisor-- {
		if dividend%divisor == 0 {
			break
		}
	}
	return divisor
}
