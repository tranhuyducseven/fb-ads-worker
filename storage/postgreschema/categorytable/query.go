package categorytable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *Category) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id string) (*Category, error) {
	var record Category

	_, err := db.QueryOne(&record, `SELECT * FROM categories WHERE id = ?`, id)
	return &record, err
}
