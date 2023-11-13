package customerproductordertable

type CustomerProductOrder struct {
	OrderId                int64
	CustomerId             int64
	ProductId              string
	CustomerProductOrderId string
	ProductCode            string
	MoneyDiscount          float64
	CustomerName           string
	CreatedDateTime        string
	DeliveryDate           string
	StatusName             string
	CalcTotalMoney         float64
	CusType                string
	CusLevel               string
	TotalMoney             string
	StartedDepotId         string
	StartedDate            string
	TotalBills             string
	LastBoughtDate         string
	Exp                    string
	ProductName            string
	Quantity               string
	Category               string
	Price                  int64

	RankingSuccessOrder       int64
	RankingProductOrder       int64
	RankingAllOrder           int64
	DaysSinceLastSuccessOrder int64
	DaysSinceLastOrder        int64
}

type Ranking struct {
	CustomerId                string
	OrderId                   string
	StatusName                string
	ProductId                 string
	ProductCode               string
	RankingAllOrder           int64
	DaysSinceLastSuccessOrder int64
	DaysSinceLastOrder        int64
	RankingSuccessOrder       int64
	DaysUseProduct            int64
	RankingProductOrder       int64
}
