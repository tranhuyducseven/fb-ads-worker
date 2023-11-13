package productexptable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *ProductExp) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id string) (*ProductExp, error) {
	var product ProductExp

	_, err := db.QueryOne(&product, `SELECT * FROM product_exps WHERE id = ?`, id)
	return &product, err
}

func GetByCode(db *pg.DB, code string) (ProductExp, error) {
	var product ProductExp

	_, err := db.QueryOne(&product, `SELECT * FROM product_exps WHERE product_code = ?`, code)
	return product, err
}
