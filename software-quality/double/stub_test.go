package double

import "testing"

type StubSearcher struct {
	phone string
}

func(ss StubSearcher) Search(people []*Person, firstname, lastname string) *Person {
	return &Person{
		Firstname: firstname,
		Lastname: lastname,
		Phone: ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakephone := "0123456789"
	phonebook := &Phonebook{}

	ss := StubSearcher{
		phone: fakephone,
	}
	phone, _ := phonebook.Find(ss, "Jane", "Doe")

	if phone != fakephone {
		t.Errorf("Want '%s', got '%s'", fakephone, phone)
	}
}