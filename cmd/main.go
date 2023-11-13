package main

import (
	"data-pipeline/service/vnpost"
	"fmt"
)

func main() {

	vnPostData, err := vnpost.GetJourney("EH759147566VN")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(vnPostData.OrderStatusHistoryDtoList))
}
