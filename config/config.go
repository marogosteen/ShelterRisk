package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	MapSizeY      int
	MapSizeX      int
	PeopleCount   int
	InfectedCount int
	TimeInterval  int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	simulationSection := cfg.Section("Simulation")

	Config = ConfigList{
		MapSizeY:      simulationSection.Key("Y").MustInt(),
		MapSizeX:      simulationSection.Key("X").MustInt(),
		PeopleCount:   simulationSection.Key("People").MustInt(),
		InfectedCount: simulationSection.Key("Infected").MustInt(),
		TimeInterval:  simulationSection.Key("TimeInterval").MustInt(),
	}
}
