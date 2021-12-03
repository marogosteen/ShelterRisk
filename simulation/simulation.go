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

type SimulationModel struct {
	MapSize        person.Position
	GridCapacity   int
	currentDate    time.Time
	finishDate     time.Time
	People         People
	MoversPosition MoversPosition
}

func NewSimulationModel(
	mapSize person.Position, gridCapacity int, simulationDays int, people People,
) (s SimulationModel, err error) {
	if mapSize.Y < 11 {
		return SimulationModel{}, errors.New("mapSizeのYが11未満")
	} else if mapSize.X < 11 {
		return SimulationModel{}, errors.New("mapSizeのXが11未満")
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

	s = SimulationModel{
		MapSize:        mapSize,
		GridCapacity:   gridCapacity,
		currentDate:    currentDate,
		finishDate:     finishDate,
		People:         people,
		MoversPosition: moversPosition,
	}
	return s, nil
}

func (s *SimulationModel) ShowInfo() {
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

func (s *SimulationModel) Run(interval time.Duration) {
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

		s.People = s.setMealTimes(interval)

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

const (
	breakFastHour   int = 7
	breakFastSecond int = breakFastHour * 60 * 60
	lunchHour       int = 12
	lunchSecond     int = lunchHour * 60 * 60
	dinnerHour      int = 18
	dinnerSecond    int = dinnerHour * 60 * 60
)

func (s *SimulationModel) setMealTimes(interval time.Duration) People {
	// TODO 正しい時間に反応してる？？

	// TODO これ正しい？？秒に直せてる？？

	newPeople := s.People
	mealGroupCount := (len(newPeople)-1) / 25
	secondsFromMidnight := (s.currentDate.Hour()*60 + s.currentDate.Minute()) * 60

	var nextMealSecond int
	if s.currentDate.Hour() <= breakFastHour {
		nextMealSecond = breakFastSecond
	} else if s.currentDate.Hour() <= lunchHour {
		nextMealSecond = breakFastSecond
	} else if s.currentDate.Hour() <= dinnerHour {
		nextMealSecond = breakFastSecond
	}

	groupDurationSecond := 15 * 60

	for xxx := 0; xxx < mealGroupCount; xxx++ {
		nextMealSecond__ := nextMealSecond + groupDurationSecond*xxx

		if secondsFromMidnight >= nextMealSecond__ && nextMealSecond__+int(interval.Seconds()) > secondsFromMidnight {
			for _, p := range s.People[xxx*25 : (xxx+1)*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}
		}
	}

	return newPeople

	/*
		// 食事時間を超え、超過時間がInterval以内
		if secondsFromMidnight >= breakFastSecond && breakFastSecond+int(interval.Seconds()) > secondsFromMidnight {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}
		} else if secondsFromMidnight >= breakFastSecond+15*60 && breakFastSecond+int(interval.Seconds())+15*60 > secondsFromMidnight {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}
		} else if secondsFromMidnight >= breakFastSecond+30*60 && breakFastSecond+int(interval.Seconds())+30*60 > secondsFromMidnight {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}
		} else if secondsFromMidnight >= breakFastSecond+45*60 && breakFastSecond+int(interval.Seconds())+45*60 > secondsFromMidnight {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= lunchSecond && lunchSecond+int(interval.Seconds()) > secondsFromMidnight {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= lunchSecond+15*60 && lunchSecond+int(interval.Seconds())+15*60 > secondsFromMidnight {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= lunchSecond+30*60 && lunchSecond+int(interval.Seconds())+30*60 > secondsFromMidnight {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= lunchSecond+45*60 && lunchSecond+int(interval.Seconds())+45*60 > secondsFromMidnight {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= dinnerSecond && dinnerSecond+int(interval.Seconds()) > secondsFromMidnight {
			for _, p := range s.People[0*25 : 1*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= dinnerSecond+15*60 && dinnerSecond+int(interval.Seconds())+15*60 > secondsFromMidnight {
			for _, p := range s.People[1*25 : 2*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= dinnerSecond+30*60 && dinnerSecond+int(interval.Seconds())+30*60 > secondsFromMidnight {
			for _, p := range s.People[2*25 : 3*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}

		} else if secondsFromMidnight >= dinnerSecond+45*60 && dinnerSecond+int(interval.Seconds())+45*60 > secondsFromMidnight {
			for _, p := range s.People[3*25 : 4*25] {
				p.LifeAction = person.Meal
				p.LifeActionElapsedSec = 0
				p.PassedCount = 0
				newPeople[p.Id] = p
			}
		}

		return newPeople
	*/
}

// simulationのcurrentDateを1日進め、Hourを午前6時にしたDateを返す。
func (s *SimulationModel) nextDate() (nextDate time.Time) {
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
func (s *SimulationModel) infectionJudge() {
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
