package ordertable

import (
	"data-pipeline/storage/postgreschema/orderproducttable"
)

type Order struct {
	Id                       int64                            `json:"id"`
	PrivateId                string                           `json:"privateId"`
	ShopOrderId              string                           `json:"shopOrderId"`
	Channel                  int64                            `json:"channel"`
	SaleChannel              int64                            `json:"saleChannel"`
	MerchantTrackingNumber   string                           `json:"merchantTrackingNumber"`
	DepotId                  int64                            `json:"depotId"`
	HandoverId               int64                            `json:"handoverId"`
	DepotName                string                           `json:"depotName"`
	Type                     string                           `json:"type"`
	TypeId                   int64                            `json:"typeId"`
	MoneyDiscount            float64                          `json:"moneyDiscount"`
	MoneyDeposit             float64                          `json:"moneyDeposit"`
	MoneyTransfer            float64                          `json:"moneyTransfer"`
	DepositAccount           string                           `json:"depositAccount"`
	TransferAccount          string                           `json:"transferAccount"`
	UsedPoints               int64                            `json:"usedPoints"`
	MoneyUsedPoints          int64                            `json:"moneyUsedPoints"`
	CarrierId                int64                            `json:"carrierId"`
	CarrierName              string                           `json:"carrierName"`
	ServiceId                string                           `json:"serviceId"`
	ServiceName              string                           `json:"serviceName"`
	CarrierCode              string                           `json:"carrierCode"`
	ShipFee                  float64                          `json:"shipFee"`
	CodFee                   float64                          `json:"codFee"`
	CustomerShipFee          float64                          `json:"customerShipFee"`
	ReturnFee                float64                          `json:"returnFee"`
	OverWeightShipFee        float64                          `json:"overWeightShipFee"`
	DeclaredFee              float64                          `json:"declaredFee"`
	Description              string                           `json:"description"`
	PrivateDescription       string                           `json:"privateDescription"`
	CustomerId               int64                            `json:"customerId"`
	CustomerName             string                           `json:"customerName"`
	CustomerEmail            string                           `json:"customerEmail"`
	CustomerMobile           string                           `json:"customerMobile"`
	CustomerAddress          string                           `json:"customerAddress"`
	ShipToCityLocationId     int64                            `json:"shipToCityLocationId"`
	ShipToDistrictLocationId string                           `json:"shipToDistrictLocationId"`
	CustomerCityId           int64                            `json:"customerCityId"`
	CustomerCity             string                           `json:"customerCity"`
	CustomerDistrictId       int64                            `json:"customerDistrictId"`
	CustomerDistrict         string                           `json:"customerDistrict"`
	ShipToWardLocationId     int64                            `json:"shipToWardLocationId"`
	CustomerWard             string                           `json:"customerWard"`
	CreatedById              int64                            `json:"createdById"`
	CreatedByUserName        string                           `json:"createdByUserName"`
	CreatedByName            string                           `json:"createdByName"`
	SaleId                   int64                            `json:"saleId"`
	SaleName                 string                           `json:"saleName"`
	CreatedDateTime          string                           `json:"createdDateTime"`
	DeliveryDate             string                           `json:"deliveryDate"`
	SendCarrierDate          string                           `json:"sendCarrierDate"`
	StatusName               string                           `json:"statusName"`
	StatusCode               string                           `json:"statusCode"`
	CalcTotalMoney           float64                          `json:"calcTotalMoney"`
	TrafficSourceId          int64                            `json:"trafficSourceId"`
	TrafficSourceName        string                           `json:"trafficSourceName"`
	AffiliateCode            string                           `json:"affiliateCode"`
	AffiliateBonusCash       float64                          `json:"affiliateBonusCash"`
	AffiliateBonusPercent    float64                          `json:"affiliateBonusPercent"`
	Products                 []orderproducttable.OrderProduct `json:"products"`
}
