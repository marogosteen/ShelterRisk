package simulations

import (
	"example/OSURisk/coodinate"
	"example/OSURisk/people"
)

type Simulation struct {
	MapSize coodinate.Coodinate
	EndSec  int
	// People     *[]people.Person
	People []people.Person
}

// Map size (21, 13)
// func NewGymSimulation(p people.People) *Simulation {
// 	gymSimulation := Simulation{
// 		CurrentSec: 0,
// 		EndSec:     428400, //17時間×７日 (17hour × 60min × 60sec × 7days)
// 		People:     p,
// 	}
// 	return &gymSimulation
// }

func (s *Simulation) Run(diffSec int) {
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		for index, person := range s.People {
			person.Move(s.MapSize)
			s.People[index] = person
		}
		s.infectionTest()
	}
}

func (s *Simulation) infectionTest() {
	positionsMap := make(map[coodinate.Coodinate][]people.Person)
	infectedCountMap := make(map[coodinate.Coodinate]int)

	for _, person := range s.People {
		key := person.NowPosition
		positionsMap[key] = append(positionsMap[key], person)
		if person.InfectionStatus != people.EnumInfectionStatus.Health {
			infectedCountMap[key]++
		}
	}

	for key, position := range positionsMap {
		if infectedCountMap[key] < 1 {
			continue
		}
		for _, person := range position {
			for i := 0; i < infectedCountMap[key]; i++ {
				if person.InfectionStatus != people.EnumInfectionStatus.Health {
					break
				}
				s.People[person.Id].InfectionStatus = person.InfectionTest()
			}
		}
	}
}

//tauch＆goでシミュレーションする。
//乱数で決めるため、全体からの％いる？
//bathroom:30/930(3.325%),CB:20/930(2.150%),CC:20/930(2.150%),
//乱数でぶらぶらして、目的地に行くのは現実てきか？
