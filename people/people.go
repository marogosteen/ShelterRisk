package people

import "math/rand"

func GeneratePeople(peopleCount int, infectedPersonCount int) []Person {
	people := make([]Person, peopleCount)
	for i := 0; i < 100; i++ {
		people[i] = NewPerson(i)
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
	for i:=0; i<len(people); i++ {
		idList = append(idList, i)
	}
	for i:=0; i<infectedPersonCount; i++ {
		idIndex := rand.Intn(len(idList))
		id := idList[idIndex]
		people[id].InfectionStatus = EnumInfectionStatus.Infection
		idList = append(idList[:idIndex], idList[idIndex+1:]...)
	}
}