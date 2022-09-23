package repository

import (
	"context"
	"mocker/example/person"
)

type Repository interface {
	ExecutePerson()
	FindPerson() (*person.Person, error)
	PostPerson(person person.Person, ctx context.Context) error
	UpdatePerson(person person.Person, ctx context.Context) error
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (a *repository) ExecutePerson() {

}

func (a *repository) FindPerson() (*person.Person, error) {
	return nil, nil
}

func (a *repository) PostPerson(person person.Person, ctx context.Context) error {
	return nil
}

func (a *repository) UpdatePerson(person person.Person, ctx context.Context) error {
	return nil
}
