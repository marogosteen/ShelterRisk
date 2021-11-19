package person

type InfectionStatus int

const (
	Health InfectionStatus = iota
	Incubation
	Infection
)
