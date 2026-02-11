package query

import "gorm.io/gorm"

type Response2 struct {
	EventID string `gorm:"column:event_id"`
	Name    string `gorm:"column:name"`
}

func Query2(db *gorm.DB) []Response2 {
	var res []Response2

	db.Table("events AS e").
		Select("e.event_id, e.name").
		Joins("JOIN results AS r ON e.event_id = r.event_id").
		Where("e.is_team_event = FALSE").
		Where("r.medal = 'GOLD'").
		Group("e.event_id, e.name").
		Having("COUNT(*) >= 2").
		Scan(&res)

	return res
}
