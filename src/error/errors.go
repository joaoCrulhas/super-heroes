package custom_errors

import "errors"

func EmptyString() error {
	return errors.New("empty string is not allowed")
}
