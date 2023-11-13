package producttable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *Product) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id string) (*Product, error) {
	var product Product

	_, err := db.QueryOne(&product, `SELECT * FROM products WHERE id = ?`, id)
	return &product, err
}

func GetByCode(db *pg.DB, code string) (*Product, error) {
	var product Product

	_, err := db.QueryOne(&product, `SELECT * FROM products WHERE code = ?`, code)
	return &product, err
}
