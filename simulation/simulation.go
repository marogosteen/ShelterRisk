package simulation

import "example/OSURisk/person"

type Simulation struct {
	MapSize      person.Position
	EndSec       int
	People       []person.Person
	positionsMap map[person.Position][]person.Person
}

func (s *Simulation) Run(diffSec int) {
	s.positionsMapInitialize()
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		// TODO 渋滞と整列 優先車の表現
		for i, p := range s.People {
			if p.IsDone() {
				p.SetNextDistination()
			}

			switch p.LifeAction {
			case person.Stay:
				p.Stay(diffSec)
			case person.Stroll:
				p.Stroll(diffSec, s.MapSize)
			default:
				p.Move(s.MapSize)
			}

			s.People[i] = p
		}
		s.infectionJudge()
	}
}

// LifeActionがStay以外のPersonのみ保持する
func (s *Simulation) positionsMapInitialize() {
	positionsMap := make(map[person.Position][]person.Person)
	for _, p := range s.People {
		if p.LifeAction != person.Stay {
			positionsMap[p.NowPosition] = append(positionsMap[p.NowPosition], p)
		}
	}
	s.positionsMap = positionsMap
}

/*
特定の条件を全て満たしたPersonに対して一定の確率で感染させる。
	[terms]
	同じ座標に2人以上がPersonが位置している。
	同じ座標上に位置するPersonの中に1人以上の感染者がいる。

*/
func (s *Simulation) infectionJudge() {
	infectedCountMap := make(map[person.Position]int)
	// 同じ座標上に位置したLifeActionがStayかつ感染しているPersonをカウントする。
	for _, p := range s.People {
		if p.InfectionStatus != person.Health && p.LifeAction != person.Stay {
			infectedCountMap[p.NowPosition]++
		}
	}

	for position, people := range s.positionsMap {
		for _, p := range people {
			// 同じ座標上にいる感染者数分の感染判定を行う。
			for i := 0; i < infectedCountMap[position]; i++ {
				// 既に感染している場合は飛ばす。
				if p.InfectionStatus != person.Health {
					break
				}
				s.People[p.Id].InfectionStatus = p.InfectionJudge()
			}
		}
	}
}
