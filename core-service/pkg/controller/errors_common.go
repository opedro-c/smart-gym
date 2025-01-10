package controller

var (
	EmptyBodyError   = NewHTTPError(400, "empty body")
	InvalidBodyError = NewHTTPError(400, "invalid body")
)
