package response

// Response is the response that represents an error.
type Response struct {
	Success      bool        `json:"success"`
	Code         ErrorCode   `json:"code,omitempty"`
	DebugMessage string      `json:"status,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

type Paging struct {
	Current int  `json:"current"`
	HasNext bool `json:"hasNext"`
	ItemNum int  `json:"itemNum"`
}

type WithPaging struct {
	Success bool        `json:"success"`
	Paging  Paging      `json:"paging"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Error is required by the error imanga.
func (e Response) Error() string {
	return e.Message
}

func ErrorResponse(code ErrorCode) *Response {
	return &Response{
		Success: false,
		Code:    code,
		Message: GetCode(code),
	}
}

// SuccessResponse Response Success success and data
func SuccessResponse(data interface{}) Response {

	return Response{
		Success: true,
		Data:    data,
	}
}

// SuccessResponseWithPaging response data and paging info
func SuccessResponseWithPaging(data interface{}, page Paging) WithPaging {

	return WithPaging{
		Success: true,
		Paging:  page,
		Data:    data,
	}
}
