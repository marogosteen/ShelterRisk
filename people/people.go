package people

type People struct {
	PeopleList []Person
}

func NewPeople(peopleCount int) *People{
	people := People{
		PeopleList: make([]Person, peopleCount),
	}
	for i:=0; i<100; i++{
		people.PeopleList[i] = *NewPerson(i)
	}
	return &people
}
