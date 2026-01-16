package domain

import "errors"

var (
	ErrProductNotFound = errors.Join(errors.New("product not found"))
	ErrInvalidPrice    = errors.Join(errors.New("price must be greater than zero"))
	ErrInvalidQuantity = errors.Join(errors.New("quantity cannot be negative"))
)
