package double

import "testing"

type SpySearcher struct {
	phone string
	searchWasCalled bool
}

func(ss SpySearcher) Search(people []*Person, firstname, lastname string) *Person {
	ss.searchWasCalled = true
	return &Person{
		Firstname: firstname,
		Lastname: lastname,
		Phone: ss.phone,
	}
}

func TestFindCallsSearchAndReurnsPerson(t *testing.T) {
	fakephone := "0123456789"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakephone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if phone != fakephone {
		t.Errorf("Want '%s', got '%s'", fakephone, phone)
	}
}