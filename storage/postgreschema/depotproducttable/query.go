package depotproducttable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *DepotProduct) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetByProductId(db *pg.DB, id string) (*DepotProduct, error) {
	var record DepotProduct

	_, err := db.QueryOne(&record, `SELECT * FROM products WHERE product_id = ?`, id)
	return &record, err
}
