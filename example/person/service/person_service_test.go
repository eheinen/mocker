package service

import (
	"context"
	"mocker/example/person"
	"mocker/example/person/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

const personName = "Eduardo Heinen"

var newPerson = person.Person{Name: personName}

func Test_FindPerson_Success(t *testing.T) {
	mockPersonRepository := new(mocks.Repository)
	mocks.MockFindPerson(&mockPersonRepository.Mock, newPerson, nil)

	actualPerson, err := New(mockPersonRepository).GetPerson()

	assert.Nil(t, err)
	assert.Equal(t, "Eheinen", &actualPerson.Name)

	mocks.AssertFindPerson(t, &mockPersonRepository.Mock, 1)
}

func Test_CreatePerson_Success(t *testing.T) {
	mockPersonRepository := new(mocks.Repository)
	mocks.MockPostPerson(&mockPersonRepository.Mock, nil)

	ctx := context.WithValue(context.Background(), "Key", "Value")
	err := New(mockPersonRepository).CreatePerson(newPerson, ctx)

	assert.Nil(t, err)

	mocks.AssertPostPerson(t, &mockPersonRepository.Mock, 1, newPerson, ctx)
}
