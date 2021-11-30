package simulation

import (
	"fmt"

	"example/OSURisk/person"
)

type Simulation struct {
	MapSize           person.Position
	GridCapacity      int
	EndSec            int
	People            []person.PersonModel
	MoversPositionMap map[person.Position][]person.PersonModel
}

func (s *Simulation) Run(diffSec int) {
	// TODO new書くべきでは？？
	s.MoversPositionMapInitialize()
	personOrder := getPersonOder(len(s.People))

	eatCount := 0
	hogecount := 0
	for currentSec := 0; currentSec <= s.EndSec; currentSec += diffSec {
		day := currentSec / (3600 * 17)
		
			if currentSec >= hogecount*100000 {
				movercount2 := 0
				infectedcount := 0
				for _, p := range s.People {
					fmt.Printf("%+v,\n", p)
					if p.InfectionStatus != person.Health {
						infectedcount++
					}
					if p.LifeAction != person.Stay {
						movercount2++
					}
				}
				movercount1 := 0
				for key, pl := range s.MoversPositionMap {
					movercount1 += len(pl)
					fmt.Printf("%v %+v\n", key, pl)
				}
				fmt.Println()
				fmt.Println("sec", currentSec)
				fmt.Println("infectedcount", infectedcount)
				fmt.Println("movercount", movercount1)
				fmt.Println("movercount2", movercount2)
				hogecount++
			}
		

		if currentSec >= day*3600*17+(1*3600)+0 && eatCount == 0 {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(1*3600)+900 && eatCount == 1 {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(1*3600)+1800 && eatCount == 2 {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(1*3600)+2700 && eatCount == 3 {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(6*3600)+0 && eatCount == 4 {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(6*3600)+900 && eatCount == 5 {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(6*3600)+1800 && eatCount == 6 {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(6*3600)+2700 && eatCount == 7 {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(12*3600)+0 && eatCount == 8 {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(12*3600)+900 && eatCount == 9 {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(12*3600)+1800 && eatCount == 10 {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec >= day*3600*17+(12*3600)+2700 && eatCount == 11 {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Eat
				p.LifeActionElapsedTime = 0
				p.PassedCount = 0
				s.People[p.Id] = p
			}
			eatCount++
		} else if currentSec < day*3600*17+(1*3600)+0 && eatCount == 12 {
			eatCount = 0
		}

		var (
			nextPersonOder  []int
			congestedPeople []int
		)
		congestedPeopleCount := 0

		for {
			for _, id := range personOrder {
				p := s.People[id]
				// TODO 後で消す
				// if p.NowPosition == p.HomePosition && p.LifeAction == person.GoHome && p.HomePosition.X == 0 && p.NowPosition.X == 0 {
				// 	fmt.Printf("%v %+v\n", p.IsDone(), p)
				// 	if !p.IsDone() {
				// 		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n")
				// 	}
				// }

				// 目的地に到達、次の目的地。
				if p.IsDone() {
					p.SetNextDistination()

					// TODO 後で消す
					if p.NowPosition == p.HomePosition && p.LifeAction == person.GoHome && p.HomePosition.X == 0 && p.NowPosition.X == 0 {
						fmt.Printf("%v %+v\n", p.IsDone(), p)
					}

				}

				// LifeActionに合わせた動作。
				var nextPosition person.Position
				switch p.LifeAction {
				case person.Stay:
					p.Stay(diffSec)
					nextPosition = p.NowPosition
				case person.Stroll:
					nextPosition = p.Stroll(diffSec, s.MapSize)
				default:
					nextPosition = p.Move(s.MapSize)
				}

				// 渋滞による移動制限。移動できなかったPersonは残しておき、再度移動させる
				if len(s.MoversPositionMap[nextPosition]) >= s.GridCapacity && p.LifeAction != person.Stay {
					// fmt.Printf("len:%v\nposition:%v\n", len(s.MoversPositionMap[nextPosition]), nextPosition)
					nextPosition = p.NowPosition
					congestedPeople = append(congestedPeople, p.Id)
					continue
				}

				// 制限されなかったPerson.Idを処理。
				nextPersonOder = append(nextPersonOder, p.Id)

				for index, bar := range s.MoversPositionMap[p.NowPosition] {
					if bar.Id == p.Id {
						s.MoversPositionMap[p.NowPosition] = append(s.MoversPositionMap[p.NowPosition][:index], s.MoversPositionMap[p.NowPosition][index+1:]...)
						if len(s.MoversPositionMap[p.NowPosition]) == 0 {
							delete(s.MoversPositionMap, p.NowPosition)
						}
						break
					}
				}

				p.NowPosition = nextPosition
				if p.LifeAction != person.Stay {
					s.MoversPositionMap[p.NowPosition] = append(s.MoversPositionMap[p.NowPosition], p)
				}

				s.People[id] = p
			}

			// 移動制限されたPersonの再移動。
			if len(congestedPeople) != congestedPeopleCount {
				personOrder = congestedPeople
				congestedPeopleCount = len(congestedPeople)
				congestedPeople = nil
				continue
			}

			// TODO 動作確認
			nextPersonOder = append(nextPersonOder, congestedPeople...)
			personOrder = nextPersonOder
			break
		}

		s.infectionJudge()
	}

	movercount2 := 0
	infectedcount := 0
	for _, p := range s.People {
		fmt.Printf("%+v,\n", p)
		if p.InfectionStatus != person.Health {
			infectedcount++
		}
		if p.LifeAction != person.Stay {
			movercount2++
		}
	}
	movercount1 := 0
	for key, pl := range s.MoversPositionMap {
		movercount1 += len(pl)
		fmt.Printf("%v %+v\n", key, pl)
	}
	fmt.Println()
	fmt.Println("infectedcount", infectedcount)
	fmt.Println("movercount", movercount1)
	fmt.Println("movercount2", movercount2)
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
