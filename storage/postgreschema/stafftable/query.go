package stafftable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *Staff) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id string) (*Staff, error) {
	var record Staff

	_, err := db.QueryOne(&record, `SELECT * FROM staffs WHERE id = ?`, id)
	return &record, err
}
