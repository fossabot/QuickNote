package response

import (
	"fmt"
)

func New(message any, data ...any) *Response {
	switch len(data) {
	case 0:
		return &Response{
			Message: fmt.Sprint(message),
		}
	case 1:
		return &Response{
			Message: fmt.Sprint(message),
			Data:    data[0],
		}
	default:
		return &Response{
			Message: fmt.Sprint(message),
			Data:    data,
		}
	}
}
