package helper

type Response struct {
	Meta meta `json:"meta"`
	Data any  `json:"data"`
}

type ResponseFail struct {
	Data interface{} `json:"data"`
}

type dataFail struct {
	Status   string `json:"status"`
	DataFail any    `json:"data_fail"`
	Message  string `json:"message"`
}

type meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ApiResponseSuccess(code int, message string, status string, data any) *Response {
	meta := meta{
		Code:    code,
		Message: message,
		Status:  status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return &response
}

func ApiResponseFail(message string, status string, data interface{}) *ResponseFail {
	response := ResponseFail{
		Data: dataFail{
			Status:   status,
			DataFail: data,
			Message:  message,
		},
	}

	return &response
}
