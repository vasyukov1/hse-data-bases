package query

import (
	"gorm.io/gorm"
)

type Response1 struct {
	Year           int `gorm:"column:birth_year"`
	PlayersCount   int `gorm:"column:players_count"`
	GoldMedalCount int `gorm:"column:gold_medal_count"`
}

func Query1(db *gorm.DB) []Response1 {
	var res []Response1

	db.Table("players AS p").
		Select(`
			EXTRACT(YEAR FROM p.birthdate) AS birth_year,
			COUNT(DISTINCT p.player_id) AS players_count,
			Count(CASE WHEN r.medal = 'GOLD' THEN 1 END) AS gold_medal_count
		`).
		Joins("JOIN results AS r ON p.player_id = r.player_id").
		Joins("JOIN events AS e ON r.event_id = e.event_id").
		Joins("JOIN olympics AS o ON e.olympic_id = o.olympic_id").
		Where("o.year = 2004").
		Group("birth_year").
		Order("birth_year").
		Scan(&res)

	return res
}
