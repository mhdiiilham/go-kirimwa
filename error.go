package gokirimwa

type ErrorResponse struct {
	Message string `json:"message"`
}

func (er ErrorResponse) Error() string {
	return er.Message
}

func IsKirimWAError(err error) bool {
	if _, ok := err.(ErrorResponse); !ok {
		return false
	}

	return true
}
