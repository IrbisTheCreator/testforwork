package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Info struct {
	Acces   string
	Refresh string
	Hash    string
}

type User struct {
	user_id int
	email   string
	Hash    string
}
