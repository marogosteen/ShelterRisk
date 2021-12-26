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

// 感染状況を表示する
func (s *SimulationModel) ShowInfo() {
	infectedcount := 0
	for _, p := range s.People {
		fmt.Printf("%+v,\n", p)
		if p.InfectionStatus != person.Health {
			infectedcount++
		}
	}
	movercount := 0
	for _, pl := range s.MoversPosition {
		movercount += len(pl)
	}
	fmt.Println()
	fmt.Println("Date", s.currentDate)
	fmt.Println("infectedcount", infectedcount)
	fmt.Println("movercount", movercount)

	fmt.Printf("\n\n")
}

func (s *SimulationModel) Run(interval time.Duration) {
	s.ShowInfo()
	m := newMealTime(s.People, interval)
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

		s.People = m.setMealTime(s.People, s.currentDate)

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
					p.SetDistination()
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
