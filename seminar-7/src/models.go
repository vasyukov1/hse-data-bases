package main

import "time"

type Country struct {
	CountryID  string `gorm:"primaryKey;column:country_id"`
	Name       string
	AreaSqkm   int
	Population int
}

func (Country) TableName() string { return "countries" }

type Olympic struct {
	OlympicID string `gorm:"primaryKey;column:olympic_id"`
	CountryID string `gorm:"foreignKey;column:country_id"`
	City      string
	Year      int
	StartDate time.Time `gorm:"column:startdate"`
	EndDate   time.Time `gorm:"column:enddate"`
}

func (Olympic) TableName() string { return "olympics" }

type Players struct {
	PlayerID  string `gorm:"primaryKey;column:player_id"`
	Name      string
	CountryID string    `gorm:"foreignKey;column:country_id"`
	BirthDate time.Time `gorm:"column:birthdate"`
}

func (Players) TableName() string { return "players" }

type Events struct {
	EventID          string `gorm:"primaryKey;column:event_id"`
	Name             string
	EventType        string
	OlympicID        string `gorm:"foreignKey;column:olympic_id"`
	IsTeamEvent      bool
	NumPlayersInTeam int
	ResultNotedIn    string
}

func (Events) TableName() string { return "events" }

type Result struct {
	EventID  string `gorm:"foreignKey;column:event_id"`
	PlayerID string `gorm:"foreignKey;column:player_id"`
	Medal    string
	Result   float64
}

func (Result) TableName() string { return "results" }
