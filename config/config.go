package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	MapSizeY             int
	MapSizeX             int
	GridCapacity         int
	PeopleCount          int
	InfectedCount        int
	TimeInterval         int
	InfectionProbability float64
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	simulationSection := cfg.Section("Simulation")

	Config = ConfigList{
		MapSizeY:             simulationSection.Key("MapSizeY").MustInt(),
		MapSizeX:             simulationSection.Key("MapSizeX").MustInt(),
		GridCapacity:         simulationSection.Key("GridCapacity").MustInt(),
		PeopleCount:          simulationSection.Key("People").MustInt(),
		InfectedCount:        simulationSection.Key("Infected").MustInt(),
		TimeInterval:         simulationSection.Key("TimeInterval").MustInt(),
		InfectionProbability: simulationSection.Key("InfectionProbability").MustFloat64(),
	}
}
