package simulation

import (
	"example/OSURisk/person"
	"fmt"
)

type Simulation struct {
	MapSize           person.Position
	GridCapacity      int
	EndSec            int
	People            []person.PersonModel
	MoversPositionMap map[person.Position][]person.PersonModel
}

func (s *Simulation) Run(diffSec int) {
	s.MoversPositionMapInitialize()
	personOrder := getPersonOder(len(s.People))
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		fmt.Println("sec", currentSec)
		for _, person := range s.People {
			fmt.Printf("%+v,\n", person)
		}
		var (
			nextPersonOder  []int
			congestedPeople []int
		)
		congestedPeopleCount := 0

		for {
			for _, id := range personOrder {
				p := s.People[id]
				// 目的地に到達、次の目的地。
				if p.IsDone() {
					p.SetNextDistination()
				}

				// LifeActionに合わせた動作。
				var nextPosition person.Position
				switch p.LifeAction {
				case person.Stay:
					// TODO nextPosition返す?
					p.Stay(diffSec)
				case person.Stroll:
					// nextPostion返す
					nextPosition = p.Stroll(diffSec, s.MapSize)
				default:
					// nextpostion返す
					nextPosition = p.Move(s.MapSize)
				}

				// こっから下Stayと移動者がごっちゃになってない？？

				// Stay以外！
				// 渋滞による移動制限。移動できなかったPersonは残しておき、再度移動させる
				if len(s.MoversPositionMap[p.NowPosition]) > s.GridCapacity {
					nextPosition = p.NowPosition
					congestedPeople = append(congestedPeople, p.Id)
					continue
				}

				// Stay含む!
				// 制限されなかったPerson.Idを処理。
				nextPersonOder = append(nextPersonOder, p.Id)

				p.NowPosition = nextPosition
				// TODO s.MoversPositionMapのPopとAppend　Moverのみ！
				s.MoversPositionMap[p.NowPosition] = append(s.MoversPositionMap[p.NowPosition], p)
				for index, bar := range s.MoversPositionMap[p.NowPosition] {
					if bar.Id == p.Id {
						s.MoversPositionMap[p.NowPosition] = append(s.MoversPositionMap[p.NowPosition][:index], s.MoversPositionMap[p.NowPosition][index+1:]...)
						break
					}
				}
				s.People[id] = p
			}

			// 移動制限されたPersonの再移動。
			if len(congestedPeople) != congestedPeopleCount {
				personOrder = congestedPeople
				continue
			}

			// TODO 動作確認
			nextPersonOder = append(nextPersonOder, congestedPeople...)
			personOrder = nextPersonOder
			break
		}

		s.infectionJudge()
	}
}

// LifeActionがStay以外のPersonのみ保持する
func (s *Simulation) MoversPositionMapInitialize() {
	positionsMap := make(map[person.Position][]person.PersonModel)
	for _, p := range s.People {
		if p.LifeAction != person.Stay {
			positionsMap[p.NowPosition] = append(positionsMap[p.NowPosition], p)
		}
	}
	s.MoversPositionMap = positionsMap
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

	for position, people := range s.MoversPositionMap {
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

func getPersonOder(personCount int) []int {
	movementOder := make([]int, personCount)
	for index := 0; index < len(movementOder); index++ {
		movementOder[index] = index
	}
	return movementOder
}
