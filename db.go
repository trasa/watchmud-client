package main

type RaceData struct {
	RaceId        int32
	RaceGroupName string
	RaceName      string
}

type ClassData struct {
	ClassId           int32
	ClassName         string
	AbilityPreference *AbilityPreferenceList
}

type AbilityPreferenceList struct {
	Preferences []string
}

type Database struct {
	Races   map[int32]RaceData
	Classes map[int32]ClassData
}

var database = &Database{}

func (db *Database) SetRaces(races []RaceData) {
	db.Races = make(map[int32]RaceData)
	for _, r := range races {
		db.Races[r.RaceId] = r
	}
}

func (db *Database) SetClasses(data []ClassData) {
	db.Classes = make(map[int32]ClassData)
	for _, c := range data {
		db.Classes[c.ClassId] = c
	}
}
