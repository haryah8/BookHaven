package models

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var ErrorInvalidRequest = ErrorResponse{
	Status:  "Error",
	Message: "Invalid Request",
}

var ErrorInternalServer = ErrorResponse{
	Status:  "Error",
	Message: "Internal Server Error",
}

var ErrorSuccess = ErrorResponse{
	Status:  "Success",
	Message: "Success",
}

func (er *ErrorResponse) SetStatus(m string) {
	er.Status = m
}
func (er *ErrorResponse) SetMessage(m string) {
	er.Message = m
}
