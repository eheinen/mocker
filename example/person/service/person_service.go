package service

import (
	"context"
	"mocker/example/person"
	"mocker/example/person/repository"
)

type Service interface {
	GetPerson() (*person.Person, error)
	CreatePerson(person person.Person, ctx context.Context) error
}

type service struct {
	repository repository.Repository
}

func New(repository repository.Repository) Service {
	return &service{repository}
}

func (a *service) GetPerson() (*person.Person, error) {
	return a.repository.FindPerson()
}

func (a *service) CreatePerson(person person.Person, ctx context.Context) error {
	return a.repository.PostPerson(person, ctx)
}
