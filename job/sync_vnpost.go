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

		retries := 0

		var vnPostData *vnpost.ItemVNPost
		for {
			if retries > 10 {
				break
			}
			v, err := vnpost.GetItem(order.OrderID)
			if err == nil {
				vnPostData = v
				break
			}

			fmt.Println("can't find order", order.OrderID, err)

			time.Sleep(2 * time.Second)
			retries += 1
		}

		if vnPostData == nil {
			continue
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
