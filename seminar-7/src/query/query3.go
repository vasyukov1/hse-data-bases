package query

import "gorm.io/gorm"

type Response3 struct {
	PlayerName string `gorm:"column:player_name"`
	OlympicID  string `gorm:"column:olympic_id"`
}

func Query3(db *gorm.DB) []Response3 {
	var res []Response3

	db.Table("players AS p").
		Select("p.name AS player_name, e.olympic_id").
		Joins("JOIN results AS r ON p.player_id = r.player_id").
		Joins("JOIN events AS e ON r.event_id = e.event_id").
		Where("r.medal in ('GOLD', 'SILVER', 'BRONZE')").
		Group("p.name, e.olympic_id").
		Scan(&res)

	return res
}
