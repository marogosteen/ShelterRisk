package core

import "example/OSURisk/people"

type Simulation struct {
	MapSize people.Position
	EndSec  int
	People  []people.Person
}

func (s *Simulation) GymRun(diffSec int) {
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		for index, person := range s.People {
			if person.IsReach() {
				person.SetNextDistination()
			}
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
	positionsMap := make(map[people.Position][]people.Person)
	infectedCountMap := make(map[people.Position]int)

	// 同じ座標上に位置するPersonと感染したPersonをそれぞれカウントする。
	for _, person := range s.People {
		key := person.NowPosition
		positionsMap[key] = append(positionsMap[key], person)
		if person.InfectionStatus != people.Health {
			infectedCountMap[key]++
		}
	}

	for key, position := range positionsMap {
		if infectedCountMap[key] < 1 {
			continue
		}
		for _, person := range position {
			for i := 0; i < infectedCountMap[key]; i++ {
				if person.InfectionStatus != people.Health {
					break
				}
				s.People[person.Id].InfectionStatus = person.InfectionJudge()
			}
		}
	}
}
