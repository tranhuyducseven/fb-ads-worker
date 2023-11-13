package errors

type CONST_ERR struct {
	varName string
	value   string
}

var _errors = []CONST_ERR{
	// vaidate
	{
		varName: "USER_ADDRESS_NOT_EXIST",
		value:   "User address not exist",
	},
}

func GetError(err string) (string, string) {
	var varName = "INTERNAL_ERROR"
	var value = err
	for _, item := range _errors {
		if item.varName == err {
			varName = item.varName
			value = item.value
		}
	}
	return "_" + varName + "_", value
}
