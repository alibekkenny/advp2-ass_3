package useCase

import (
	"ass_3/services/contact/internal/domain"
)

type UseCase interface {
	CreateContact(c domain.Contact) (domain.Contact, error)
	ReadContact(c domain.Contact) (domain.Contact, error)
	UpdateContact(c domain.Contact) (domain.Contact, error)
	DeleteContact(c domain.Contact) (domain.Contact, error)
	CreateGroup(g domain.Group) (domain.Group, error)
	ReadGroup(g domain.Group) (domain.Group, error)
	AddContactToGroup(g domain.Group, c domain.Contact) (domain.Group, error)
}

type UseCaseStruct struct {
}

func NewUseCase() UseCase {
	return &UseCaseStruct{}
}

func (u *UseCaseStruct) CreateContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (u *UseCaseStruct) ReadContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (u *UseCaseStruct) UpdateContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (u *UseCaseStruct) DeleteContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (u *UseCaseStruct) CreateGroup(g domain.Group) (domain.Group, error) {
	return g, nil
}
func (u *UseCaseStruct) ReadGroup(g domain.Group) (domain.Group, error) {
	return g, nil
}
func (u *UseCaseStruct) AddContactToGroup(g domain.Group, c domain.Contact) (domain.Group, error) {
	return g, nil
}
