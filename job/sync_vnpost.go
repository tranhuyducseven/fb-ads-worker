package job

import (
	"context"
	"data-pipeline/service/vnpost"
	"data-pipeline/storage/mongodb/ordercol"
	"fmt"
)

func SyncVNPost() {

	orders, err := ordercol.FindAllEmptyVnPostID(context.Background())
	if err != nil {
		return
	}

	for _, order := range orders {
		vnPostData, err := vnpost.GetItem(order.OrderID)
		if err != nil {
			fmt.Println(err, order.OrderID)
			continue
		}

		if vnPostData.ItemCode != "" {
			order.VNPostID = vnPostData.ItemCode

			_, err := ordercol.Update(context.Background(), order)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("update order success ", order.OrderID)
		}
	}
}
