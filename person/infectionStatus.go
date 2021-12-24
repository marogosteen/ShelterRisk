package person

type InfectionStatus string

const (
	Health     InfectionStatus = "Health"
	Incubation InfectionStatus = "Incubation"
	Infection  InfectionStatus = "Infection"
)
