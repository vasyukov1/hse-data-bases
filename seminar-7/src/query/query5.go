package query

import "gorm.io/gorm"

type Response5 struct {
	CountryName string  `gorm:"column:country_name"`
	Ratio       float64 `gorm:"column:ratio"`
}

func Query5(db *gorm.DB) []Response5 {
	var res []Response5

	db.Raw(`
		 SELECT c.name AS country_name,
               COUNT(r.medal)::float / c.population AS ratio
        FROM results r
        JOIN events e ON r.event_id = e.event_id
        JOIN olympics o ON o.olympic_id = e.olympic_id
        JOIN players p ON p.player_id = r.player_id
        JOIN countries c ON p.country_id = c.country_id
        WHERE o.year = 2000
          AND e.is_team_event = TRUE
          AND r.medal IS NOT NULL
        GROUP BY c.name, c.population
        ORDER BY ratio
        LIMIT 5
	`).Scan(&res)

	return res
}
