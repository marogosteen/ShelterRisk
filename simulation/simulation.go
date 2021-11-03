package simulation

import "example/OSURisk/person"

type Simulation struct {
	GymMap    [21][13]person.Person
	DiningMap [10][10]person.Person // TODO 食堂サイズ考えてくる（Sugi）
	CurrentTime int
	EndTime   int
}

func NewSimulation() *Simulation{
	s := new(Simulation)
	s.CurrentTime = 0
	// TODO 終了秒設定する 
	// s.EndTime = 
	return s
}

func (s *Simulation) Run() {
	

}
