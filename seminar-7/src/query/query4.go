package query

import "gorm.io/gorm"

type Response4 struct {
	CountryName string  `gorm:"column:country_name"`
	Ratio       float64 `gorm:"column:ratio"`
}

func Query4(db *gorm.DB) []Response4 {
	var res []Response4

	db.Raw(`
		SELECT c.name AS country_name,
		       AVG(CASE WHEN LOWER(p.name) IN ('a','e','i','o','u','y') THEN 1 ELSE 0 END) AS ratio
		FROM countries c
		JOIN players p ON c.country_id = p.country_id
		GROUP BY c.name
		ORDER BY ratio DESC
		LIMIT 1
	`).Scan(&res)

	return res
}
