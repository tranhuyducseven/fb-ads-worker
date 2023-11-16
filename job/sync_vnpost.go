package job

import (
	"context"
	"data-pipeline/service/vnpost"
	"data-pipeline/storage/mongodb/ordercol"
	"fmt"
	"time"
)

func SyncVNPost() {

	orders, err := ordercol.FindAllEmptyVnPostID(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, order := range orders {

		// get vnpost data
		var vnPostData *vnpost.ItemVNPost
		for {
			v, err := vnpost.GetItem(order.OrderID)
			if err == nil {
				vnPostData = v
				break
			}

			fmt.Println("can't find order", order.OrderID)

			time.Sleep(10 * time.Second)
		}

		fmt.Println("find vnpost data success", vnPostData)
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
