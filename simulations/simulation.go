package simulations

import (
	"example/OSURisk/coodinate"
	"example/OSURisk/infectionStatus"
	"example/OSURisk/people"
)

type Simulation struct {
	MapSize coodinate.Coodinate
	EndSec  int
	People []people.Person
}

func (s *Simulation) GymRun(diffSec int) {
	// Eventから目的地をセット

	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		for index, person := range s.People {
			person.Move(s.MapSize)
			s.People[index] = person
		}
		s.infectionJudge()
	}
}

func (s *Simulation) DiningRun(diffSec int) {
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		for index, person := range s.People {
			person.Move(s.MapSize)
			s.People[index] = person
		}
		s.infectionJudge()
	}
}

/* 
特定の条件を全て満たしたPersonに対して一定の確率で感染させる。
	[terms]
	同じ座標に2人以上がPersonが位置している。
	同じ座標上に位置するPersonの中に1人以上の感染者がいる。

*/
func (s *Simulation) infectionJudge() {
	positionsMap := make(map[coodinate.Coodinate][]people.Person)
	infectedCountMap := make(map[coodinate.Coodinate]int)

	// 同じ座標上に位置するPersonと感染したPersonをそれぞれカウントする。
	for _, person := range s.People {
		key := person.NowPosition
		positionsMap[key] = append(positionsMap[key], person)
		if person.InfectionStatus != infectionStatus.Health {
			infectedCountMap[key]++
		}
	}

	for key, position := range positionsMap {
		if infectedCountMap[key] < 1 {
			continue
		}
		for _, person := range position {
			for i := 0; i < infectedCountMap[key]; i++ {
				if person.InfectionStatus != infectionStatus.Health {
					break
				}
				s.People[person.Id].InfectionStatus = person.InfectionJudge()
			}
		}
	}
}