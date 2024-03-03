package domain

// Here is a interface to be implemented by the authentication service
type Authentication[T any, R any] interface {
	Auth(input T) (R, error)
}
