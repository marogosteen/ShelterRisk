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

	func (LivingSpeace) {
		var voo[][] int
		

	for i := 0; i < 5; i++ {
		for j:=0; j<11; {
			voo [i][j]=(i+2)+(j+2)
	}
	}
		fmt.Prentln(voo)
	}

	//
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
