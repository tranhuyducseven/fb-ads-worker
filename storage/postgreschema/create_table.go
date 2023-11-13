package postgreschema

import (
	"data-pipeline/storage/postgreschema/categorytable"
	"data-pipeline/storage/postgreschema/customerproductordertable"
	"data-pipeline/storage/postgreschema/customertable"
	"data-pipeline/storage/postgreschema/depotproducttable"
	"data-pipeline/storage/postgreschema/depottable"
	"data-pipeline/storage/postgreschema/orderproducttable"
	"data-pipeline/storage/postgreschema/ordertable"
	"data-pipeline/storage/postgreschema/productexptable"
	"data-pipeline/storage/postgreschema/producttable"
	"data-pipeline/storage/postgreschema/stafftable"
	"data-pipeline/storage/postgreschema/tokentable"

	"github.com/ponlv/go-kit/postgresql"
)

func CreateAllTable() {

	db := postgresql.GetDB()
	if db == nil {
		return
	}

	models := []interface{}{
		(*producttable.Product)(nil),
		(*depotproducttable.DepotProduct)(nil),
		(*depottable.Depot)(nil),
		(*ordertable.Order)(nil),
		(*orderproducttable.OrderProduct)(nil),
		(*customertable.Customer)(nil),
		(*categorytable.Category)(nil),
		(*stafftable.Staff)(nil),
		(*productexptable.ProductExp)(nil),
		(*customerproductordertable.CustomerProductOrder)(nil),
		(*tokentable.Token)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(nil)
		if err != nil {
			continue
		}
	}

}
