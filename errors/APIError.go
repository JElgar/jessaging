package errors

type APIError struct {
    Code    int           `json:"code"`
    Message string        `json:"message"`
    Type    ErrorType     `json:"type"`
}

func NewAPIError() APIError {
    return APIError {500, "Uh Oh! There has been an unknown error, we will get this sorted as soon as possible!", UNKNOWNERROR}
}
