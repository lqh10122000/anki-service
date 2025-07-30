package model

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrInvalidID        = errors.New("invalid customer ID")
	ErrInvalidPassword  = errors.New("invalid password ID")
)
