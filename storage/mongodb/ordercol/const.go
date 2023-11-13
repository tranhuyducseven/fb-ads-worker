package ordercol

const (
	WeightUnit_KG string = "kg"
	WeightUnit_LB string = "lb"
	WeightUnit_OZ string = "oz"
	WeightUnit_G  string = "g"
)

const (
	Status_Draft      string = "draft"
	Status_Pending    string = "pending"
	Status_Paid       string = "paid"
	Status_Confirming string = "confirming"
	Status_Confirmed  string = "confirmed"

	Status_Packing string = "packing"
	Status_Packed  string = "packed"

	Status_Transported string = "transported"

	Status_Success string = "success"
	Status_Failed  string = "success"
	Status_Cancel  string = "cancel"
)

const (
	PaymentPaypal       int = 1
	PaymentTransferBank int = 2
)

func GetPaymentTitle(method int) string {
	if method == 1 {
		return "Paypal"
	} else if method == 2 {
		return "Bank Transfer"
	}

	return ""
}
