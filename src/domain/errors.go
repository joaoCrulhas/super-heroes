package domain

import "errors"

func EmptyString() error {
	return errors.New("empty string is not allowed")
}

func InvalidCharacters() error {
	return errors.New("invalid characters are not allowed")
}
