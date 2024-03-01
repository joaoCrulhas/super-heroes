package domain

import "errors"

func EmptyString() error {
	return errors.New("empty string is not allowed")
}

func BadRequest(msg string) error {
	return errors.New(msg)
}
