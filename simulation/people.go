package simulation

import (
	"math/rand"

	"example/OSURisk/person"
)

// Configで指定された人数のpersonを生成し、スライスにする。
func GeneratePeople(peopleCount int, infectedPersonCount int) []person.PersonModel {
	people := make([]person.PersonModel, peopleCount)
	for id := 0; id < len(people); id++ {
		homePosition := getLivingPosition(peopleCount, id)
		people[id] = *person.NewPerson(id, homePosition)
		// lifeAction := person.GetRandomAction()
		// distinationList := person.DistinationListMap[lifeAction]
		// people[id] = person.PersonModel{
		// 	Id:                    id,
		// 	HomePosition:          homePosition,
		// 	NowPosition:           homePosition,
		// 	Distination:           distinationList[0],
		// 	PassedCount:           0,
		// 	InfectionStatus:       person.Health,
		// 	LifeAction:            lifeAction,
		// 	LifeActionElapsedTime: 0,
		// }
	}
	setInfected(people, infectedPersonCount)
	return people
}

// 指定された人数のPersonランダムに感染者に変更する
func setInfected(people []person.PersonModel, infectedPersonCount int) {
	// TODO 感染者数がシミュレーション人数より多い場合はエラー
	// if len(p.PersonList) < infectedPersonCount{
	// 	panic()
	// }
	var HelthIdList []int
	for i := 0; i < len(people); i++ {
		HelthIdList = append(HelthIdList, i)
	}
	for i := 0; i < infectedPersonCount; i++ {
		idIndex := rand.Intn(len(HelthIdList))
		id := HelthIdList[idIndex]
		people[id].InfectionStatus = person.Infection
		HelthIdList = append(HelthIdList[:idIndex], HelthIdList[idIndex+1:]...)
	}
}

// TODO diviend < divisor のエラー処理
func getLivingPosition(peopleCount int, personId int) person.Position {
	livingSpaceCapacity := 100
	yLivingSpaceCapacity := livingSpaceCapacity / 10
	xLivingSpaceCapacity := livingSpaceCapacity / yLivingSpaceCapacity
	byDivisible := getByDivisible(livingSpaceCapacity, peopleCount)
	whetherRemainder := personId / byDivisible
	step := livingSpaceCapacity / byDivisible

	livingPositionId := (personId-whetherRemainder*byDivisible)*step +
		step/2*whetherRemainder
	// 住居スペースが１マス２人であることを表現するために/2をした
	livingPositionId /= 1

	return person.Position{
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
