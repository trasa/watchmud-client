package main

type RaceData struct {
	RaceId        int32
	RaceGroupName string
	RaceName      string
}

type Database struct {
	Races map[int32]RaceData
}

var database = &Database{}

func (db *Database) SetRaces(races []RaceData) {
	db.Races = make(map[int32]RaceData)
	for _, r := range races {
		db.Races[r.RaceId] = r
	}
}
