package client

type ErrorResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
