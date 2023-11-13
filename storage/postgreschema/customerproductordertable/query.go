package customerproductordertable

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Create(db *pg.DB, data *CustomerProductOrder) error {
	_, err := db.Model(data).Insert()
	return err
}

func GetOrders(db *pg.DB) ([]CustomerProductOrder, error) {
	var users []CustomerProductOrder
	_, err := db.Query(&users, `SELECT * FROM customer_product_orders ORDER BY order_id desc`)
	return users, err
}

func GetById(db *pg.DB, id string) (*CustomerProductOrder, error) {
	var record CustomerProductOrder

	_, err := db.QueryOne(&record, `SELECT * FROM customer_product_orders WHERE customer_product_order_id = ?`, id)
	return &record, err
}

func UpdateByCode(db *pg.DB, id string, data *CustomerProductOrder) error {
	var record CustomerProductOrder

	query := orm.NewUpdateQuery(orm.NewQuery(nil, data).Where("customer_product_order_id = ?", id), false)

	_, err := db.Query(&record, query)
	return err
}

func UpdateByOrderProductId(db *pg.DB, data *CustomerProductOrder) error {
	var record CustomerProductOrder

	query := orm.NewUpdateQuery(orm.NewQuery(nil, data).Where("order_id = ? AND product_id = ?", data.OrderId, data.ProductId), false)

	_, err := db.Query(&record, query)
	return err
}
func QueryRanking(db *pg.DB) ([]Ranking, error) {
	var record []Ranking

	_, err := db.Query(&record, `
	with distinct_od as (
	select
		distinct cpo.customer_id ,
		cpo.order_id ,
		cpo.status_name ,
		TO_TIMESTAMP(created_date_time ,
		'YYYY/MM/DD/HH24:MI:ss') as datee
	from
		customer_product_orders cpo
	order by
		cpo.customer_id ,
		cpo.order_id ),
		success_od as (
	select
		distinct cpo.customer_id ,
		cpo.order_id ,
		cpo.status_name ,
		TO_TIMESTAMP(created_date_time ,
		'YYYY/MM/DD/HH24:MI:ss') as datee
	from
		customer_product_orders cpo
	where
		cpo.status_name = 'Thành công'
	order by
		cpo.customer_id ,
		cpo.order_id ),
		product_record as (
	select
		cpo.*,
		TO_TIMESTAMP(cpo.created_date_time ,
		'YYYY/MM/DD/HH24:MI:ss') as created_timestamp
	from
		customer_product_orders cpo 
		),
	ranks_success as (
	select
		success_od.*,
		rank() over(partition by success_od.customer_id
	order by
		success_od.datee) as ranking_success_order
	from
		success_od),
	ranks_all as (
	select
		distinct_od.*,
		rank() over(partition by distinct_od.customer_id
	order by
		distinct_od.datee) as ranking_all_order
	from
		distinct_od)
	select 
		distinct_od.*,
		product_record.product_code,
		product_record.product_id,
		ranks_all.ranking_all_order,
		extract('day' from success_od.datee - lag(success_od.datee) over (partition by success_od.customer_id order by success_od.datee)) as days_since_last_success_order,
		extract('day' from distinct_od.datee - lag(distinct_od.datee) over (partition by distinct_od.customer_id order by distinct_od.datee)) as days_since_last_order,
		ranks_success.ranking_success_order,
		row_number() over(partition by product_record.order_id
	order by
		product_record.created_timestamp ) as ranking_product_order
	from
		distinct_od
	left join ranks_success on
		distinct_od.customer_id = ranks_success.customer_id
		and distinct_od.datee = ranks_success.datee
	left join product_record on
		distinct_od.customer_id = product_record.customer_id
		and distinct_od.datee = product_record.created_timestamp
	left join ranks_all on
		distinct_od.customer_id = ranks_all.customer_id
		and distinct_od.datee = ranks_all.datee
	left join success_od on
		distinct_od.customer_id = success_od.customer_id
		and distinct_od.datee = success_od.datee	
	group by
		distinct_od.customer_id,
		distinct_od.order_id,
		distinct_od.status_name,
		distinct_od.datee,
		product_record.created_timestamp,
		product_record.order_id,
		product_record.product_code,
		product_record.product_id,
		ranks_all.ranking_all_order,
		ranks_success.ranking_success_order,
		ranking_product_order,
		success_od.datee,
		success_od.customer_id
	order by
		distinct_od.customer_id ,
		distinct_od.order_id
	`)
	return record, err
}
