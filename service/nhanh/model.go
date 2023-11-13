package nhanh

type Product struct {
	IdNhanh            string    `json:"idNhanh"`
	ParentId           string    `json:"parentId"`
	BrandId            string    `json:"brandId"`
	BrandName          string    `json:"brandName"`
	TypeId             string    `json:"typeId"`
	TypeName           string    `json:"typeName"`
	AvgCost            string    `json:"avgCost"`
	MerchantCategoryId string    `json:"merchantCategoryId"`
	MerchantProductId  string    `json:"merchantProductId"`
	CategoryId         string    `json:"categoryId"`
	InternalCategoryId string    `json:"internalCategoryId"`
	Code               string    `json:"code"`
	Barcode            string    `json:"barcode"`
	Name               string    `json:"name"`
	OtherName          string    `json:"otherName"`
	ImportPrice        string    `json:"importPrice"`
	OldPrice           string    `json:"oldPrice"`
	Price              string    `json:"price"`
	WholesalePrice     string    `json:"wholesalePrice"`
	Image              string    `json:"image"`
	Status             string    `json:"status"`
	ShowHot            int64     `json:"showHot"`
	ShowNew            int64     `json:"showNew"`
	ShowHome           int64     `json:"showHome"`
	Order              string    `json:"order"`
	PreviewLink        string    `json:"previewLink"`
	ShippingWeight     string    `json:"shippingWeight"`
	Width              string    `json:"width"`
	Length             string    `json:"length"`
	Height             string    `json:"height"`
	Vat                string    `json:"vat"`
	CreatedDateTime    string    `json:"createdDateTime"`
	Inventory          Inventory `json:"inventory"`
	WarrantyAddress    string    `json:"warrantyAddress"`
	WarrantyPhone      string    `json:"warrantyPhone"`
	Warranty           string    `json:"warranty"`
	CountryName        string    `json:"countryName"`
	Unit               string    `json:"unit"`
}

type Inventory struct {
	Remain          int64       `json:"remain"`
	Shipping        int64       `json:"shipping"`
	Damaged         int64       `json:"damaged"`
	Holding         int64       `json:"holding"`
	Warranty        int64       `json:"warranty"`
	WarrantyHolding int64       `json:"warrantyHolding"`
	Available       int64       `json:"available"`
	Depots          interface{} `json:"depots"`
}

type Order struct {
	Id                       int64          `json:"id"`
	PrivateId                string         `json:"privateId"`
	ShopOrderId              string         `json:"shopOrderId"`
	Channel                  int64          `json:"channel"`
	SaleChannel              int64          `json:"saleChannel"`
	MerchantTrackingNumber   string         `json:"merchantTrackingNumber"`
	DepotId                  int64          `json:"depotId"`
	HandoverId               int64          `json:"handoverId"`
	DepotName                string         `json:"depotName"`
	Type                     string         `json:"type"`
	TypeId                   int64          `json:"typeId"`
	MoneyDiscount            float64        `json:"moneyDiscount"`
	MoneyDeposit             float64        `json:"moneyDeposit"`
	MoneyTransfer            float64        `json:"moneyTransfer"`
	DepositAccount           string         `json:"depositAccount"`
	TransferAccount          string         `json:"transferAccount"`
	UsedPoints               int64          `json:"usedPoints"`
	MoneyUsedPoints          int64          `json:"moneyUsedPoints"`
	CarrierId                int64          `json:"carrierId"`
	CarrierName              string         `json:"carrierName"`
	ServiceId                interface{}    `json:"serviceId"`
	ServiceName              string         `json:"serviceName"`
	CarrierCode              string         `json:"carrierCode"`
	ShipFee                  float64        `json:"shipFee"`
	CodFee                   float64        `json:"codFee"`
	CustomerShipFee          float64        `json:"customerShipFee"`
	ReturnFee                float64        `json:"returnFee"`
	OverWeightShipFee        float64        `json:"overWeightShipFee"`
	DeclaredFee              float64        `json:"declaredFee"`
	Description              string         `json:"description"`
	PrivateDescription       string         `json:"privateDescription"`
	CustomerId               int64          `json:"customerId"`
	CustomerName             string         `json:"customerName"`
	CustomerEmail            string         `json:"customerEmail"`
	CustomerMobile           string         `json:"customerMobile"`
	CustomerAddress          string         `json:"customerAddress"`
	ShipToCityLocationId     int64          `json:"shipToCityLocationId"`
	ShipToDistrictLocationId interface{}    `json:"shipToDistrictLocationId"`
	CustomerCityId           int64          `json:"customerCityId"`
	CustomerCity             string         `json:"customerCity"`
	CustomerDistrictId       int64          `json:"customerDistrictId"`
	CustomerDistrict         string         `json:"customerDistrict"`
	ShipToWardLocationId     int64          `json:"shipToWardLocationId"`
	CustomerWard             string         `json:"customerWard"`
	CreatedById              int64          `json:"createdById"`
	CreatedByUserName        string         `json:"createdByUserName"`
	CreatedByName            string         `json:"createdByName"`
	SaleId                   int64          `json:"saleId"`
	SaleName                 string         `json:"saleName"`
	CreatedDateTime          string         `json:"createdDateTime"`
	DeliveryDate             string         `json:"deliveryDate"`
	SendCarrierDate          string         `json:"sendCarrierDate"`
	StatusName               string         `json:"statusName"`
	StatusCode               string         `json:"statusCode"`
	CalcTotalMoney           float64        `json:"calcTotalMoney"`
	TrafficSourceId          int64          `json:"trafficSourceId"`
	TrafficSourceName        string         `json:"trafficSourceName"`
	AffiliateCode            string         `json:"affiliateCode"`
	AffiliateBonusCash       float64        `json:"affiliateBonusCash"`
	AffiliateBonusPercent    float64        `json:"affiliateBonusPercent"`
	Products                 []OrderProduct `json:"products"`
}

type OrderProduct struct {
	ProductId      string        `json:"productId"`
	ProductName    string        `json:"productName"`
	ProductCode    string        `json:"productCode"`
	ProductBarcode string        `json:"productBarcode"`
	ProductImage   string        `json:"productImage"`
	Price          interface{}   `json:"price"`
	Quantity       string        `json:"quantity"`
	Weight         string        `json:"weight"`
	Discount       interface{}   `json:"discount"`
	Description    string        `json:"description"`
	Imei           string        `json:"imei"`
	GiftProducts   []interface{} `json:"giftProducts"`
}

type OrderGiftProduct struct {
}
type Depots struct {
	Remain          int64 `json:"remain"`
	Shipping        int64 `json:"shipping"`
	Damaged         int64 `json:"damaged"`
	Holding         int64 `json:"holding"`
	Warranty        int64 `json:"warranty"`
	WarrantyHolding int64 `json:"warrantyHolding"`
	Available       int64 `json:"available"`
}

type DepotInfo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	CityId       int64  `json:"cityId"`
	CityName     string `json:"cityName"`
	DistrictId   int64  `json:"districtId"`
	DistrictName string `json:"districtName"`
	WardId       int64  `json:"wardId"`
	WardName     string `json:"wardName"`
	Address      string `json:"address"`
}

type Customer struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	Mobile             string `json:"mobile"`
	Gender             string `json:"gender"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	Birthday           string `json:"birthday"`
	Code               string `json:"code"`
	Level              string `json:"level"`
	Group              string `json:"group"`
	LevelId            string `json:"levelId"`
	GroupId            string `json:"groupId"`
	CityLocationId     string `json:"cityLocationId"`
	DistrictLocationId string `json:"districtLocationId"`
	WardLocationId     string `json:"wardLocationId"`
	TotalMoney         interface{} `json:"totalMoney"`
	StartedDate        string `json:"startedDate"`
	StartedDepotId     string `json:"startedDepotId"`
	Points             interface{} `json:"points"`
	TotalBills         interface{} `json:"totalBills"`
	LastBoughtDate     string `json:"lastBoughtDate"`
	TaxCode            string `json:"taxCode"`
	BusinessName       string `json:"businessName"`
	BusinessAddress    string `json:"businessAddress"`
	Description        string `json:"description"`
}

type Staff struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	RoleName string `json:"roleName"`
}

// ================================ Get category ===================
type GetCategoryResponse struct {
	Code int64                     `json:"code"`
	Data []GetCategoryResponseData `json:"data"`
}

type GetCategoryResponseData struct {
	Id            string `json:"id"`
	ParentId      string `json:"parentId"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Order         string `json:"order"`
	ShowHome      string `json:"showHome"`
	ShowHomeOrder string `json:"showHomeOrder"`
	PrivateId     string `json:"privateId"`
	Status        string `json:"status"`
	Image         string `json:"image"`
}

// ================================ Search product ===================
type SearchProductResponse struct {
	Code int64                     `json:"code"`
	Data SearchProductResponseData `json:"data"`
}

type SearchProductResponseData struct {
	CurrentPage int64              `json:"currentPage"`
	TotalPages  int                `json:"totalPages"`
	Products    map[string]Product `json:"products"`
}

// ================================ Search order ===================
type SearchOrderResponse struct {
	Code int64                   `json:"code"`
	Data SearchOrderResponseData `json:"data"`
}

type SearchOrderResponseData struct {
	Page         int64            `json:"page"`
	TotalPages   int              `json:"totalPages"`
	TotalRecords int64            `json:"totalRecords"`
	Orders       map[string]Order `json:"orders"`
}

// ========== Customer =================
type SearchCustomerResponse struct {
	Code int64                      `json:"code"`
	Data SearchCustomerResponseData `json:"data"`
}

type SearchCustomerResponseData struct {
	Page       int64               `json:"page"`
	TotalPages int                 `json:"totalPages"`
	Customers  map[string]Customer `json:"customers"`
}

// ========== Customer =================
type GetDepotResponse struct {
	Code   int64                `json:"code"`
	Depots map[string]DepotInfo `json:"data"`
}

// ========== Staff =================
type SearchStaffResponse struct {
	Code int64                   `json:"code"`
	Data SearchStaffResponseData `json:"data"`
}

type SearchStaffResponseData struct {
	Page       int64            `json:"page"`
	TotalPages int              `json:"totalPages"`
	Staffs     map[string]Staff `json:"users"`
}

// ========== OrderDetail =================
type GetOrderDetailResponse struct {
	Code int64                      `json:"code"`
	Data GetOrderDetailResponseData `json:"data"`
}

type GetOrderDetailResponseData struct {
	Order map[string]Order `json:"orders"`
}

// ========== OrderDetail =================
type GetCustomerDetailResponse struct {
	Code int64                         `json:"code"`
	Data GetCustomerDetailResponseData `json:"data"`
}

type GetCustomerDetailResponseData struct {
	Customer map[string]Customer `json:"customers"`
}
