package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func New(success bool, message string, data ...any) *Response {
	switch len(data) {
	case 0:
		return &Response{
			Success: success,
			Message: message,
		}
	case 1:
		return &Response{
			Success: success,
			Message: message,
			Data:    data[0],
		}
	default:
		return &Response{
			Success: success,
			Message: message,
			Data:    data,
		}
	}
}
