package simulations

import (
	"example/OSURisk/people"
)

// mtdbujy4tj6suf2xcvffg5tgyk7cyjpz22h4etfb4xma2tvphh6q
type Simulation struct {
	Map        [][]people.Person
	CurrentSec int
	EndSec     int
	People     people.People
}

// Map size (21, 13)
func NewGymSimulation(p people.People) *Simulation {
	gymSimulation := Simulation{
		CurrentSec: 0,
		EndSec:     428400, //17時間×７日 (17hour × 60min × 60sec × 7days)
		People:     p,
	}
	return &gymSimulation
}

// 3回/1日 実施
// Map size (10, 10)
func NewDiningSimulation(p people.People) *Simulation {
	diningSimulation := Simulation{
		CurrentSec: 0,
		EndSec:     1800,
		People:     p,
	}
	return &diningSimulation
}

func (s *Simulation) Run(diffSec int) {
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		for index, person := range s.People.PersonList {
			s.People.PersonList[index] = *person.Move()
		}
		s.infectionTest()
	}
}

func (s *Simulation) infectionTest() {
	positions := make(map[people.Coodinate][]people.Person)
	infectedCountMap := make(map[people.Coodinate]int)

	for _, person := range s.People.PersonList {
		key := person.CurrentCoodinate
		positions[key] = append(positions[key], person)
		if person.InfectionStatus != people.EnumInfectionStatus.Health || infectedCountMap[key] > 0 {
			infectedCountMap[key]++
		}
	}

	for key, onePosition := range positions {
		if infectedCountMap[key] < 1 {
			continue
		}
		for _, person := range onePosition {
			if person.InfectionStatus != people.EnumInfectionStatus.Health {
				continue
			}
			for i := 0; i < infectedCountMap[key]; i++ {
				person.InfectionTest()
			}
		}
	}
}

//tauch＆goでシミュレーションする。
//乱数で決めるため、全体からの％いる？
//bathroom:30/930(3.325%),CB:20/930(2.150%),CC:20/930(2.150%),
//乱数でぶらぶらして、目的地に行くのは現実てきか？
