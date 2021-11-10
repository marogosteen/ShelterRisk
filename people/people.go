package people

import "math/rand"

type People struct {
	PersonList []Person
}

func NewPeople(peopleCount int, infectedPersonCount int) *People {
	people := People{
		PersonList: make([]Person, peopleCount),
	}
	for i := 0; i < 100; i++ {
		people.PersonList[i] = *NewPerson(i)
	}
	people.setInfected(infectedPersonCount)
	return &people
}

func (p *People) setInfected(infectedPersonCount int) {
	// TODO 感染者数がシミュレーション人数より多い場合はエラー
	// if len(p.PersonList) < infectedPersonCount{
	// 	panic()
	// }
	var idList []int
	for i:=0; i<len(p.PersonList); i++ {
		idList = append(idList, i)
	}
	for i:=0; i<infectedPersonCount; i++ {
		index := rand.Intn(len(idList))
		p.PersonList[idList[index]].InfectionStatus = EnumInfectionStatus.Infection
		idList = append(idList[:index], idList[index+1:]...)
	}
}
