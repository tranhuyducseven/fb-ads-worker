package producttable

import (
	"data-pipeline/storage/postgreschema/depotproducttable"
)

type Product struct {
	Id                 string `json:"idNhanh" pg:",unique"`
	ParentId           string `json:"parentId"`
	BrandId            string `json:"brandId"`
	BrandName          string `json:"brandName"`
	TypeId             string `json:"typeId"`
	TypeName           string `json:"typeName"`
	AvgCost            string `json:"avgCost"`
	MerchantCategoryId string `json:"merchantCategoryId"`
	MerchantProductId  string `json:"merchantProductId"`
	CategoryId         string `json:"categoryId"`
	InternalCategoryId string `json:"internalCategoryId"`
	Code               string `json:"code"`
	Barcode            string `json:"barcode"`
	Name               string `json:"name"`
	OtherName          string `json:"otherName"`
	ImportPrice        string `json:"importPrice"`
	OldPrice           string `json:"oldPrice"`
	Price              string `json:"price"`
	WholesalePrice     string `json:"wholesalePrice"`
	Image              string `json:"image"`
	Status             string `json:"status"`
	ShowHot            int64  `json:"showHot"`
	ShowNew            int64  `json:"showNew"`
	ShowHome           int64  `json:"showHome"`
	Order              string `json:"order"`
	PreviewLink        string `json:"previewLink"`
	ShippingWeight     string `json:"shippingWeight"`
	Width              string `json:"width"`
	Length             string `json:"length"`
	Height             string `json:"height"`
	Vat                string `json:"vat"`
	CreatedDateTime    string `json:"createdDateTime"`
	UpdatedDateTime    string `json:"updatedDateTime"`
	WarrantyAddress    string `json:"warrantyAddress"`
	WarrantyPhone      string `json:"warrantyPhone"`
	CountryName        string `json:"countryName"`
	Unit               string `json:"unit"`

	Remain          int64 `json:"remain"`
	Shipping        int64 `json:"shipping"`
	Damaged         int64 `json:"damaged"`
	Holding         int64 `json:"holding"`
	Warranty        int64 `json:"warranty"`
	WarrantyHolding int64 `json:"warrantyHolding"`
	Available       int64 `json:"available"`

	Depots []depotproducttable.DepotProduct `pg:"rel:has-many"`
}
