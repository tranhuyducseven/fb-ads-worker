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
			continue
		}

		if vnPostData.ItemCode != "" {
			order.VNPostID = vnPostData.ItemCode

			_, err := ordercol.Update(context.Background(), order)
			if err != nil {
				continue
			}

			fmt.Println("update order success ", order.OrderID)
		}
	}
}
