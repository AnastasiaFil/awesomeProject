package domain

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
