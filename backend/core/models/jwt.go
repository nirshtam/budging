package models

type TokenProvider interface {
	GetToken() (string, error)
}
