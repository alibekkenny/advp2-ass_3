package domain

import (
	"errors"
	"regexp"
)

type Contact struct {
	Id          int
	fullName    string
	PhoneNumber string
}

func NewContact(id int, fullName string, phoneNumber string) (*Contact, error) {
	correctPhoneNum := regexp.MustCompile(`^[\d]+$`).MatchString(phoneNumber)
	correctFullName := regexp.MustCompile(`^[A-Za-z]+\s[A-Za-z]+\s[A-Za-z]+$`).MatchString(fullName)

	if !correctPhoneNum || phoneNumber == "" ||
		!correctFullName || fullName == "" ||
		id == 0 {
		return nil, errors.New("not correct format of input data")
	}
	return &Contact{Id: id, fullName: fullName, PhoneNumber: phoneNumber}, nil
}

func (c *Contact) FullName() string {
	return c.fullName
}
