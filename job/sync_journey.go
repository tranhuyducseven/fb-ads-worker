package job

import (
	"context"
	larkaction "data-pipeline/service/larkservice/action"
	"data-pipeline/service/vnpost"
	"data-pipeline/storage/mongodb/ordercol"
	"data-pipeline/storage/postgreschema/tokentable"
	"fmt"
	"github.com/ponlv/go-kit/postgresql"
	"sync"
	"time"
)

const (
	botAppId    = "cli_a59d6ea23f78d010"
	baseAppId   = "AZx5bu2aPaXWHWsuRU0lvpF3gif"
	baseTableId = "tblU4JxyvTXwCoSi"
	baseViewId  = "vewYzSYKYY"
)

func SyncVnPostJourney() {
	db := postgresql.GetDB()

	// get access token from database
	var token *tokentable.Token
	for {
		t, err := tokentable.GetById(db, botAppId)
		if err == nil {
			token = t
			break
		}

		fmt.Println(err)
	}

	// check if access token is expired
	if token.AccessTokenExpireTime <= time.Now().Unix() {
		for {

			// renew token form database
			token, err := tokentable.GetById(db, botAppId)
			if err != nil {
				time.Sleep(10 * time.Second)
				continue
			}

			res, err := larkaction.RefreshAccessToken(token.RefreshToken, token.AppAccessToken)
			if err == nil && res.Code == 0 {

				token.AccessToken = res.AuthData.AccessToken
				token.RefreshToken = res.AuthData.RefreshToken
				token.RefreshTokenExpireTime = time.Now().Unix() + res.AuthData.RefreshExpiresIn
				token.AccessTokenExpireTime = time.Now().Unix() + res.AuthData.ExpiresIn

				tokentable.UpdateById(db, token.AppId, token)
				fmt.Println("update new access token from refresh token")
				break
			}

			fmt.Println("cannot get new access token from refresh token err: ", err, " msg: ", res.Msg)
			fmt.Println("try to get new refresh token")

			// check if app token is expired
			accessToken, err := larkaction.GetAppAccessToken(token.AppId, token.AppSecret)
			if err != nil || accessToken.Code != 0 {
				fmt.Println("get app access token failed", err, accessToken)
				continue
			}

			token.AppAccessToken = accessToken.AppAccessToken
			token.AppExpireTime = time.Now().Unix() + accessToken.Expire

			err = tokentable.UpdateById(db, token.AppId, token)
			if err != nil {
				fmt.Println("update app access token failed", err)
				continue
			}

			// check if refresh token is expired
			for {
				// renew token form database
				token, err := tokentable.GetById(db, botAppId)
				if err != nil {
					fmt.Println("get token err", err.Error())
					time.Sleep(10 * time.Second)
					continue
				}

				// get access token from refresh token, but first we need to check refresh token is alive
				if token.RefreshTokenExpireTime <= time.Now().Unix() {

					// get access token
					res, err := larkaction.GetAccessToken(token.AuthCode, token.AppAccessToken)
					if err == nil && res.Code == 0 {
						token.AccessToken = res.AuthData.AccessToken
						token.RefreshToken = res.AuthData.RefreshToken
						token.RefreshTokenExpireTime = time.Now().Unix() + res.AuthData.RefreshExpiresIn
						token.AccessTokenExpireTime = time.Now().Unix() + res.AuthData.ExpiresIn

						tokentable.UpdateById(db, token.AppId, token)
						fmt.Println("[bot] update new access token")
						break
					}

					fmt.Println("[bot] cannot get new access token, sleep 10s, err: ", err, " msg: ", res.Msg)
					// todo: notify
					time.Sleep(10 * time.Second)
				} else {
					break
				}

			}

		}
	}

	orders, err := ordercol.FindAllOrderHaveVnPostID(context.Background())
	if err != nil {
		return
	}

	fmt.Println(len(orders))
	var wg sync.WaitGroup

	for _, order := range orders {
		wg.Add(1)
		go func(item *ordercol.Order) {
			defer wg.Done()

			var journey *vnpost.Journey

			for {
				j, err := vnpost.GetJourney(item.VNPostID)
				if err == nil {
					journey = j
					break
				}

				time.Sleep(5 * time.Second)
			}

			if len(journey.OrderStatusHistoryDtoList) > len(item.Journey) {

				var newJourney []ordercol.OrderJourney

				//
				for _, journeyItem := range journey.OrderStatusHistoryDtoList {
					newJourney = append(newJourney, ordercol.OrderJourney{
						TraceDate:    journeyItem.TraceDate,
						StatusText:   journeyItem.StatusText,
						StatusDetail: journeyItem.StatusDetail,
						PostmanName:  journeyItem.PostmanName,
						PostmanTel:   journeyItem.PostmanTel,
						Lon:          journeyItem.Lon,
						Lat:          journeyItem.Lat,
					})
				}

				// update new journey
				item.Journey = newJourney
				_, err := ordercol.Update(context.Background(), item)
				if err == nil {
					fmt.Println("update new journey", item.OrderID)
				}

				fmt.Println(journey.OrderStatusHistoryDtoList[0])
				// check failed delivery
				if journey.OrderStatusHistoryDtoList[0].StatusText == "Giao hàng không thành công" {
					record := larkaction.AddBaseRecordResquest{
						AppId:   baseAppId,
						TableId: baseTableId,
						Fields: map[string]interface{}{
							"order_id":       item.OrderID,
							"delivery_code":  item.VNPostID,
							"order_status":   journey.OrderStatusHistoryDtoList[0].StatusText,
							"customer_name":  item.FirstName,
							"customer_phone": item.PhoneNumber,
							"postman":        journey.OrderStatusHistoryDtoList[0].PostmanName,
							"note":           journey.OrderStatusHistoryDtoList[0].StatusDetail,
							"update_time":    time.Now().Unix() * 1000,
						},
					}

					err = larkaction.AddBaseRecord(record, token.AccessToken)
					if err != nil {
						fmt.Println("add record error", err, item.OrderID)
						return
					}
				} else {
					page_token := ""
					is_update := false
					for {

						// get bitable data
						res, err := larkaction.GetBaseListRecord(larkaction.GetBaseListRecordResquest{
							AppId:     baseAppId,
							TableId:   baseTableId,
							ViewId:    baseViewId,
							PageToken: page_token,
						}, token.AccessToken)
						if err != nil {
							fmt.Println(err)
						}

						// for loop all records get from base
						for _, record := range res.Records.Items {
							fields := record.Fields.(map[string]interface{})

							if fmt.Sprintf("%v", fields["order_id"]) != item.OrderID {
								continue
							}

							// update record
							err = larkaction.UpdateBaseRecord(larkaction.UpdateBaseRecordResquest{
								AppId:    baseAppId,
								TableId:  baseTableId,
								RecordId: record.RecordId,
								Fields: map[string]interface{}{
									"order_status": journey.OrderStatusHistoryDtoList[0].StatusText,
									"note":         journey.OrderStatusHistoryDtoList[0].StatusDetail,
									"update_time":  time.Now().Unix() * 1000,
								},
							}, token.AccessToken)
							if err != nil {
								fmt.Println("update record error", err)
								continue
							}
							is_update = true

						}

						if res.Records.PageToken == "" || is_update {
							break
						} else {
							page_token = res.Records.PageToken
						}
					}
				}
			}
		}(order)

	}
	wg.Wait()
}
