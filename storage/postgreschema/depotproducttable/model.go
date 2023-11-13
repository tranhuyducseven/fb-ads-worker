package depotproducttable

type DepotProduct struct {
	DepotId         string `json:"depot_id"`
	ProductId       string `json:"product_id"`
	Remain          int64  `json:"remain"`
	Shipping        int64  `json:"shipping"`
	Damaged         int64  `json:"damaged"`
	Holding         int64  `json:"holding"`
	Warranty        int64  `json:"warranty"`
	WarrantyHolding int64  `json:"warrantyHolding"`
	Available       int64  `json:"available"`
}
