package domain

import (
	"errors"
	"unicode/utf8"
)

type Group struct {
	Id   int
	Name string
}

func NewGroup(id int, name string) (*Group, error) {
	nameLength := utf8.RuneCountInString(name)
	if nameLength > 250 || name == "" ||
		id == 0 {
		return nil, errors.New("name shouldn't be empty or greater than 250 symbols")
	}
	return &Group{Id: id, Name: name}, nil
}
