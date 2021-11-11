package people

import (
	"math/rand"

	"example/OSURisk/coodinate"
)

func GeneratePeople(peopleCount int, infectedPersonCount int) []Person {
	people := make([]Person, peopleCount)
	for id := 0; id < len(people); id++ {
		livingPosition := getLivingPosition(peopleCount, id)
		people[id] = Person{
			Id:               id,
			StartCoodinate:   livingPosition,
			NowPosition:      livingPosition,
			EventElapsedTime: 0,
			InfectionStatus:  EnumInfectionStatus.Health,
		}
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
		people[id].InfectionStatus = EnumInfectionStatus.Infection
		idList = append(idList[:idIndex], idList[idIndex+1:]...)
	}
}

// TODO diviend < divisor のエラー処理
func getLivingPosition(peopleCount int, currentId int) coodinate.Coodinate {
	livingSpaceCapacity := 100
	yLivingSpaceCapacity := livingSpaceCapacity / 10
	xLivingSpaceCapacity := livingSpaceCapacity / yLivingSpaceCapacity
	byDivisible := getByDivisible(livingSpaceCapacity, peopleCount)
	whetherRemainder := currentId / byDivisible
	step := livingSpaceCapacity / byDivisible

	livingSpaceId := (currentId-whetherRemainder*byDivisible)*step + step/2*whetherRemainder

	return coodinate.Coodinate{
		Y: livingSpaceId / yLivingSpaceCapacity * 2,
		X: livingSpaceId % xLivingSpaceCapacity,
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
