package orderproducttable

type OrderProduct struct {
	OrderId        int64         `json:"orderId"`
	ProductId      string        `json:"productId"`
	OrderProductId string        `json:"orderProductId" pg:",unique"`
	ProductName    string        `json:"productName"`
	ProductCode    string        `json:"productCode"`
	ProductBarcode string        `json:"productBarcode"`
	ProductImage   string        `json:"productImage"`
	Price          string        `json:"price"`
	Quantity       string        `json:"quantity"`
	Weight         string        `json:"weight"`
	Discount       string        `json:"discount"`
	Description    string        `json:"description"`
	Imei           string        `json:"imei"`
	GiftProducts   []interface{} `json:"giftProducts"`
}
