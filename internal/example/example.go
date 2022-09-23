package example

import "context"

type Person interface {
	ExecutePerson()
	GetPerson() (*Person, error)
	PostPerson(person *Person, ctx context.Context) error
	UpdatePerson(person *Person, ctx context.Context) error
}

type person struct {
	name string
}

func ExecutePerson() {

}

func GetPerson() (*Person, error) {
	return nil, nil
}

func PostPerson(person *Person, ctx context.Context) error {
	return nil
}

func UpdatePerson(person *Person, ctx context.Context) error {
	return nil
}
