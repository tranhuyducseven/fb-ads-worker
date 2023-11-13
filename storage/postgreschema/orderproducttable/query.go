package orderproducttable

import (
	"github.com/go-pg/pg/v10"
)

func Create(db *pg.DB, data *OrderProduct) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetByOrderIdAndProductId(db *pg.DB, product_id string) (*OrderProduct, error) {
	var record OrderProduct

	_, err := db.QueryOne(&record, `SELECT * FROM order_products WHERE order_product_id = ?`, product_id)
	return &record, err
}

func GetProducts(db *pg.DB, order_id int64) ([]*OrderProduct, error) {
	var users []*OrderProduct
	_, err := db.Query(&users, `SELECT * FROM order_products WHERE order_id = ?`, order_id)
	return users, err
}
