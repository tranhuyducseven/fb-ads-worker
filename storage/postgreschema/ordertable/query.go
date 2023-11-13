package ordertable

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Create(db *pg.DB, data *Order) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id int64) (*Order, error) {
	var record Order

	_, err := db.QueryOne(&record, `SELECT * FROM orders WHERE id = ?`, id)
	return &record, err
}

func GetAllOrderInDate(db *pg.DB, date string) ([]Order, error) {
	var users []Order
	_, err := db.Query(&users, `SELECT * FROM orders WHERE created_date_time like ?`, date)
	return users, err
}

func GetOrders(db *pg.DB) ([]Order, error) {
	var users []Order
	_, err := db.Query(&users, `SELECT * FROM orders ORDER BY id desc`)
	return users, err
}

func UpdateById(db *pg.DB, id int64, data *Order) error {
	var record Order

	query := orm.NewUpdateQuery(orm.NewQuery(nil, data).Where("id = ?", id), false)

	_, err := db.Query(&record, query)
	return err
}

func DeleteById(db *pg.DB, id int64) error {
	var record Order

	query := orm.NewDeleteQuery(orm.NewQuery(nil, &record).Where("id = ?", id))
	_, err := db.Query(&record, query)
	return err
}
