package api

import "errors"

var ErrBadRequest = errors.New("Bad Request")
var ErrForbidden = errors.New("Forbidden")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrTooManyRequests = errors.New("Too Many Requests")
var ErrServerError = errors.New("Server Error")
