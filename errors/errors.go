package errors

type ErrorType int

const (
    UNKNOWNERROR ErrorType = iota
    CONTENTNOTFOUND
    CONNECTIONERROR
)

