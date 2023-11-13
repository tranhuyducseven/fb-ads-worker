package productexptable

type ProductExp struct {
	ProductName string `json:"productName"`
	ProductCode string `json:"productCode"`
	Package     string `json:"package"`
	Exp         string `json:"exp"`
	Category    string `json:"category"`
	Dosage      string `json:"dosage"`
}
