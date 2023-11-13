package response

type ErrorCode string

const (
	ErrInvalidParam ErrorCode = "INVALID_PARAM"
	ErrUserNotExist ErrorCode = "USER_NOT_EXIST"
)

var errorTextMap = map[ErrorCode]string{
	ErrUserNotExist: "User not exits",
	ErrInvalidParam: "Invalid Param",
}

func GetCode(e ErrorCode) string {
	if data, ok := errorTextMap[e]; ok {
		return data
	}
	return ""
}
