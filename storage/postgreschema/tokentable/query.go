package tokentable

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func GetById(db *pg.DB, id string) (*Token, error) {
	var record Token

	_, err := db.QueryOne(&record, `SELECT * FROM tokens WHERE app_id = ?`, id)
	return &record, err
}

func UpdateById(db *pg.DB, id string, data *Token) error {
	var record Token

	query := orm.NewUpdateQuery(orm.NewQuery(nil, data).Where("app_id = ?", id), false)

	_, err := db.Query(&record, query)
	return err
}
