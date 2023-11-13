package ordercol

import (
	"time"

	"github.com/ponlv/go-kit/mongodb"
)

type OrderProduct struct {
	Id          string  `json:"id,omitempty" bson:"id,omitempty"`
	Quantity    int64   `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price       float64 `json:"price,omitempty" bson:"price,omitempty"`
	Discount    float64 `json:"discount,omitempty" bson:"discount,omitempty"`
	Img         string  `json:"img,omitempty" bson:"img,omitempty"`
	Title       string  `json:"title,omitempty" bson:"title,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`

	DiscountApplied []string `json:"discount_applied,omitempty" bson:"discount_applied,omitempty"`
}

type USPSTracking struct {
	TrackingNumber string `json:"tracking_number,omitempty" bson:"tracking_number,omitempty"`
}

type OrderJourney struct {
	TraceDate    string  `json:"traceDate"`
	StatusText   string  `json:"statusText"`
	StatusDetail string  `json:"statusDetail"`
	PostmanName  string  `json:"postmanName"`
	PostmanTel   string  `json:"postmanTel"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
}
type Order struct {
	mongodb.DefaultModel `json:",inline" bson:",inline,omitnested"`
	CreatedAt            time.Time `json:"created_at" bson:"created_at,omitempty" `
	UpdatedAt            time.Time `json:"updated_at" bson:"updated_at,omitempty" `

	TraceDate    string  `json:"traceDate"`
	StatusText   string  `json:"statusText"`
	StatusDetail string  `json:"statusDetail"`
	PostmanName  string  `json:"postmanName"`
	PostmanTel   string  `json:"postmanTel"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
	// user information
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`

	// international address
	Address     string `json:"address,omitempty" bson:"address,omitempty"`
	Apartment   string `json:"apartment,omitempty" bson:"apartment,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	Country     string `json:"country,omitempty" bson:"country,omitempty"`
	Zip         string `json:"zip,omitempty" bson:"zip,omitempty"`
	State       string `json:"state,omitempty" bson:"state,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Email       string `json:"email,omitempty" bson:"email,omitempty"`

	// domestically address
	CustomerId               int64  `json:"customer_id" bson:"customer_id"`
	ShipToCityLocationId     int64  `json:"ship_to_city_location_id" bson:"ship_to_city_location_id"`
	ShipToDistrictLocationId string `json:"ship_to_district_location_id" bson:"ship_to_district_location_id"`
	CustomerCityId           int64  `json:"customer_city_id" bson:"customer_city_id"`
	CustomerCity             string `json:"customer_city" bson:"customer_city"`
	CustomerDistrictId       int64  `json:"customer_district_id" bson:"customer_district_id"`
	CustomerDistrict         string `json:"customer_district" bson:"customer_district"`
	ShipToWardLocationId     int64  `json:"ship_to_ward_location_id" bson:"ship_to_ward_location_id"`
	CustomerWard             string `json:"customer_ward" bson:"customer_ward"`

	// order information
	OrderID    string         `json:"order_id,omitempty" bson:"order_id,omitempty"`
	Status     string         `json:"status,omitempty" bson:"status,omitempty"`
	Products   []OrderProduct `json:"products,omitempty" bson:"products,omitempty"`
	StatusName string         `json:"status_name" bson:"status_name"`

	// Payment
	PaymentMethodId int       `json:"payment_method,omitempty" bson:"payment_method,omitempty"`
	PaymentTitle    string    `json:"payment_title,omitempty" bson:"payment_title,omitempty"`
	Total           float64   `json:"total" bson:"total,omitempty"`
	Discount        float64   `json:"discount" bson:"discount"`
	ActualPayment   float64   `json:"actual_payment" bson:"actual_payment"`
	MoneyDeposit    float64   `json:"money_deposit" bson:"money_deposit"`
	MoneyTransfer   float64   `json:"money_transfer" bson:"money_transfer"`
	DepositAccount  string    `json:"deposit_account" bson:"deposit_account"`
	TransferAccount string    `json:"transfer_account" bson:"transfer_account"`
	UsedPoints      int64     `json:"used_points" bson:"used_points"`
	MoneyUsedPoints int64     `json:"money_used_points" bson:"money_used_points"`
	PaymentDate     time.Time `json:"payment_date" bson:"payment_date,omitempty"`

	// shipping
	ShippingFee        float64        `json:"shipping_fee,omitempty" bson:"shipping_fee"`
	ShippingId         string         `json:"shipping_id,omitempty" bson:"shipping_id,omitempty"`
	DeliveryDate       string         `json:"delivery_date" bson:"delivery_date"`
	CarrierId          int64          `json:"carrier_id" bson:"carrier_id"`
	CarrierName        string         `json:"carrier_name" bson:"carrier_name"`
	ServiceId          string         `json:"service_id" bson:"service_id"`
	ServiceName        string         `json:"service_name" bson:"service_name"`
	CarrierCode        string         `json:"carrier_code" bson:"carrier_code"`
	CodFee             float64        `json:"cod_fee" bson:"cod_fee"`
	CustomerShipFee    float64        `json:"customer_ship_fee" bson:"customer_ship_fee"`
	ReturnFee          float64        `json:"return_fee" bson:"return_fee"`
	OverWeightShipFee  float64        `json:"over_weight_ship_fee" bson:"over_weight_ship_fee"`
	DeclaredFee        float64        `json:"declared_fee" bson:"declared_fee"`
	Description        string         `json:"description" bson:"description"`
	PrivateDescription string         `json:"private_description" bson:"private_description"`
	SendCarrierDate    string         `json:"send_carrier_date" bson:"send_carrier_date"`
	Journey            []OrderJourney `json:"journey" bson:"journey"`

	// shipping - usps
	USPSTracking         USPSTracking `json:"usps_tracking,omitempty" bson:"usps_tracking,omitempty"`
	UpdateUSPSTrackingAt time.Time    `json:"update_usps_tracking_at,omitempty" bson:"update_usps_tracking_at,omitempty"`

	// shipping - vnpost
	VNPostID string `json:"vn_post_id" bson:"vn_post_id"`

	// note
	UserNote    string `json:"user_note" bson:"user_note,omitempty"`
	PrivateNote string `json:"private_note" bson:"private_note,omitempty"`

	// other channel info
	PrivateId              string `json:"privateId"`
	ShopOrderId            string `json:"shopOrderId"`
	Channel                int64  `json:"channel"`
	SaleChannel            int64  `json:"saleChannel"`
	MerchantTrackingNumber string `json:"merchantTrackingNumber"`

	// sale ID
	SaleId            int64  `json:"sale_id" bson:"sale_id"`
	SaleName          string `json:"sale_name" bson:"sale_name"`
	CreatedById       int64  `json:"created_by_id" bson:"created_by_id"`
	CreatedByUserName string `json:"created_by_user_name" bson:"created_by_user_name"`
	CreatedByName     string `json:"created_by_name" bson:"created_by_name"`

	// depot
	DepotId    int64  `json:"depot_id" bson:"depot_id"`
	HandoverId int64  `json:"handover_id" bson:"handover_id"`
	DepotName  string `json:"depot_name" bson:"depot_name"`

	TrafficSourceId   int64  `json:"traffic_source_id" bson:"traffic_source_id"`
	TrafficSourceName string `json:"traffic_source_name" bson:"traffic_source_name"`
}

// Pagination
type Pagination struct {
	Total  int64    `json:"total,omitempty" bson:"total,omitempty"`
	Orders []*Order `json:"orders,omitempty" bson:"orders,omitempty"`
}

func (Order) CollectionName() string {
	return "order"
}
