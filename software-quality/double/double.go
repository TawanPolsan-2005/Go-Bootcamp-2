package double

import "errors"

var (
	ErrMissingArgs = errors.New("FirstName and Lastname are mendatory arguments")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface {
	Search(people []*Person, firstname string, lastname string) *Person
}

type Person struct {
	Firstname string
	Lastname string
	Phone string
}

type Phonebook struct {
	People []*Person
}

func(p *Phonebook) Find(query Queryer, firstname, lastname string) (string, error) {
	if firstname == "" || lastname == "" {
		return "", ErrMissingArgs
	}

	person := query.Search(p.People, firstname, lastname)
	//person := &Person{}
	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}