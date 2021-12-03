package simulation

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"example/OSURisk/person"
)

type MoversPosition map[person.Position][]person.PersonModel

type Simulation struct {
	MapSize        person.Position
	GridCapacity   int
	currentDate    time.Time
	finishDate     time.Time
	People         People
	MoversPosition MoversPosition
}

func NewSimulation(
	mapSize person.Position, gridCapacity int, simulationDays int, people People,
) (s Simulation, err error) {
	if mapSize.Y < 11 {
		return Simulation{}, errors.New("mapSizeのYが11未満")
	} else if mapSize.X < 11 {
		return Simulation{}, errors.New("mapSizeのXが11未満")
	}

	hour := simulationDays * 24
	strHour := strconv.Itoa(int(hour)) + "h"
	addHour, err := time.ParseDuration(strHour)
	if err != nil {
		log.Fatal(err)
	}

	NowDate := time.Now()
	currentDate := time.Date(
		NowDate.Year(), NowDate.Month(), 1, 6, 0, 0, 0, NowDate.Location(),
	)
	finishDate := currentDate.Add(addHour)
	moversPosition := GenerateMoversPosition(people)

	s = Simulation{
		MapSize:        mapSize,
		GridCapacity:   gridCapacity,
		currentDate:    currentDate,
		finishDate:     finishDate,
		People:         people,
		MoversPosition: moversPosition,
	}
	return s, nil
}

func (s *Simulation) ShowInfo() {
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
	fmt.Println("\nmoversposition")
	movercount1 := 0
	for key, pl := range s.MoversPosition {
		movercount1 += len(pl)
		fmt.Printf("%v %+v\n", key, pl)
	}
	fmt.Println()
	fmt.Println("Date", s.currentDate)
	fmt.Println("infectedcount", infectedcount)
	fmt.Println("movercount", movercount1)
	if !(movercount1 == movercount2) {
		log.Fatal(errors.New("movercountが一致しない"))
	}

	fmt.Printf("\n\n")
}

func (s *Simulation) Run(interval time.Duration) {
	s.ShowInfo()
	personOrder := getPersonOder(len(s.People))

	for ; ; s.currentDate = s.currentDate.Add(interval) {
		// 1日を23時までとし、過ぎるとnextDate()をCallする。
		if s.currentDate.Hour() >= 23 {
			s.currentDate = s.nextDate()
			s.ShowInfo()
		}

		if !s.currentDate.Before(s.finishDate) {
			break
		}

		s.People = s.checkMealTimes()

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
					p.Stay(interval)
					nextPosition = p.NowPosition
				case person.Stroll:
					nextPosition = p.Stroll(interval, s.MapSize)
				default:
					nextPosition = p.Move(s.MapSize)
				}

				// 渋滞による移動制限。移動できなかったPersonは残しておき、再度移動させる
				if len(s.MoversPosition[nextPosition]) >= s.GridCapacity && p.LifeAction != person.Stay {
					// fmt.Printf("len:%v\nposition:%v\n", len(s.MoversPositionMap[nextPosition]), nextPosition)
					nextPosition = p.NowPosition
					congestedPeople = append(congestedPeople, p.Id)
					continue
				}

				// 制限されなかったPerson.Idを処理。
				nextPersonOder = append(nextPersonOder, p.Id)

				for index, bar := range s.MoversPosition[p.NowPosition] {
					if bar.Id == p.Id {
						s.MoversPosition[p.NowPosition] = append(s.MoversPosition[p.NowPosition][:index], s.MoversPosition[p.NowPosition][index+1:]...)
						if len(s.MoversPosition[p.NowPosition]) == 0 {
							delete(s.MoversPosition, p.NowPosition)
						}
						break
					}
				}

				p.NowPosition = nextPosition
				if p.LifeAction != person.Stay {
					s.MoversPosition[p.NowPosition] = append(s.MoversPosition[p.NowPosition], p)
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

			nextPersonOder = append(nextPersonOder, congestedPeople...)
			personOrder = nextPersonOder
			break
		}

		s.infectionJudge()
	}
}

func (s *Simulation) checkMealTimes() People {
	return s.People
}

/*
func (s *Simulation) __checkMealTimes() People {
	// TODO 正しい時間に反応してる？？
	elapsedScounds := (s.currentDate.Hour()*60 + s.currentDate.Minute()) * 60
	if elapsedScounds >= 7*60 && eatCount == 0 {
		for _, p := range s.People[0*25 : 1*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 7*60+15 && eatCount == 1 {
		for _, p := range s.People[1*25 : 2*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 7*60+30 && eatCount == 2 {
		for _, p := range s.People[2*25 : 3*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 7*60+45 && eatCount == 3 {
		for _, p := range s.People[3*25 : 4*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 12*60 && eatCount == 4 {
		for _, p := range s.People[0*25 : 1*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 12*60+15 && eatCount == 5 {
		for _, p := range s.People[1*25 : 2*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 12*60+30 && eatCount == 6 {
		for _, p := range s.People[2*25 : 3*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 12*60+45 && eatCount == 7 {
		for _, p := range s.People[3*25 : 4*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 18*60+0 && eatCount == 8 {
		for _, p := range s.People[0*25 : 1*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 18*60+15 && eatCount == 9 {
		for _, p := range s.People[1*25 : 2*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 18*60+30 && eatCount == 10 {
		for _, p := range s.People[2*25 : 3*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds >= 18*60+45 && eatCount == 11 {
		for _, p := range s.People[3*25 : 4*25] {
			p.LifeAction = person.Eat
			p.LifeActionElapsedSec = 0
			p.PassedCount = 0
			s.People[p.Id] = p
		}
		eatCount++
	} else if elapsedScounds < 7*60+0 && eatCount == 12 {
		eatCount = 0
	}

	return newPeople
}
*/

// simulationのcurrentDateを1日進め、Hourを午前6時にしたDateを返す。
func (s *Simulation) nextDate() (nextDate time.Time) {
	nextDate = time.Date(
		s.currentDate.Year(), s.currentDate.Month(), s.currentDate.Day()+1,
		6, 0, 0, 0, s.currentDate.Location(),
	)
	return nextDate
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

	for position, people := range s.MoversPosition {
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

func GenerateMoversPosition(people People) MoversPosition {
	moversPosition := make(MoversPosition)
	for _, p := range people {
		if p.LifeAction != person.Stay {
			moversPosition[p.NowPosition] = append(moversPosition[p.NowPosition], p)
		}
	}
	return moversPosition
}

func getPersonOder(personCount int) []int {
	movementOder := make([]int, personCount)
	for index := 0; index < len(movementOder); index++ {
		movementOder[index] = index
	}
	return movementOder
}
