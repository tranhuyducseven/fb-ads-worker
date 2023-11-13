package customertable

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Create(db *pg.DB, data *Customer) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetById(db *pg.DB, id string) (Customer, error) {
	var record Customer

	_, err := db.QueryOne(&record, `SELECT * FROM customers WHERE id = ?`, id)
	return record, err
}

func UpdateById(db *pg.DB, id string, data *Customer) error {
	var record Customer

	query := orm.NewUpdateQuery(orm.NewQuery(nil, data).Where("id = ?", id), false)

	_, err := db.Query(&record, query)
	return err
}
