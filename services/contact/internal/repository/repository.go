package repository

import "ass_3/services/contact/internal/domain"

type Repository interface {
	CreateContact(c domain.Contact) (domain.Contact, error)
	ReadContact(c domain.Contact) (domain.Contact, error)
	UpdateContact(c domain.Contact) (domain.Contact, error)
	DeleteContact(c domain.Contact) (domain.Contact, error)
	CreateGroup(g domain.Group) (domain.Group, error)
	ReadGroup(g domain.Group) (domain.Group, error)
	AddContactToGroup(g domain.Group, c domain.Contact) (domain.Group, error)
}

type RepositoryStruct struct {
}

func NewRepository() Repository {
	return &RepositoryStruct{}
}

func (r *RepositoryStruct) CreateContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (r *RepositoryStruct) ReadContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (r *RepositoryStruct) UpdateContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (r *RepositoryStruct) DeleteContact(c domain.Contact) (domain.Contact, error) {
	return c, nil
}
func (r *RepositoryStruct) CreateGroup(g domain.Group) (domain.Group, error) {
	return g, nil
}
func (r *RepositoryStruct) ReadGroup(g domain.Group) (domain.Group, error) {
	return g, nil
}
func (r *RepositoryStruct) AddContactToGroup(g domain.Group, c domain.Contact) (domain.Group, error) {
	return g, nil
}
