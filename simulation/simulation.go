package simulation

import "example/OSURisk/person"

type Simulation struct {
	GymMap      [21][13]person.Person
	DiningMap   [10][10]person.Person // TODO 食堂サイズ考えてくる（Sugi）
	CurrentTime int
	EndTime     int
}

type DiningSimulation struct {
	DiningMap   [10][10]person.Person
	CurrentTime int
	EndTime     int
}

type GymSimulation struct {
	GymMap      [21][13]person.Person
	CurrentTime int
	EndTime     int
}

func NewDinigSimulation() *DiningSimulation {
	diningSimulation := DiningSimulation{
		DiningMap:   [10][10]person.Person{},
		CurrentTime: 0,
		EndTime:     1800,
	}
	return &diningSimulation
	//3回実施
}
func NewGymSimulation()*GymSimulation {
	gymSimulation:=GymSimulation{
		GymMap: [21][13]person.Person{},
		CurrentTime: 0,
		EndTime:     428400,//17時間×７日 
		
		//tauch＆goでシミュレーションする。
		//EndTimeいる？
		//乱数で決めるため、全体からの％いる？
		//bathroom:30/930(3.325%),CB:20/930(2.150%),CC:20/930(2.150%),
	}
	return &gymSimulation 
}


func NewSimulation() *Simulation {
	s := new(Simulation)
	s.CurrentTime = 0
	// TODO 終了秒設定する
	// s.EndTime =
	return s
}

func (s *Simulation) Run() {

}
