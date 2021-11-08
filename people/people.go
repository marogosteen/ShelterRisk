package people

type People struct {
	PeopleList []Person
}

func NewPeople(peopleCount int) *People {
	people := People{
		PeopleList: make([]Person, peopleCount),
	}
	for i := 0; i < 100; i++ {
		people.PeopleList[i] = *NewPerson(i)
	}
	return &people
}

func (p *People) Move() {
	for _, person := range p.PeopleList {
		person.Move()
	}
}

// TODO IncubationからInfectionになるプログラム
func (p *People) InfectionTest() {
	for _, person := range p.PeopleList {
		person.InfectionTest()
	}
}
