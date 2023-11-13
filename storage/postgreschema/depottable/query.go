package depottable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *Depot) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id int64) (*Depot, error) {
	var record Depot

	_, err := db.QueryOne(&record, `SELECT * FROM depots WHERE id = ?`, id)
	return &record, err
}
