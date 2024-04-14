package testify_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Firstname string
	Lastname  string
	Phone     string
}

func TestSomething(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		t.Run("equal", func(t *testing.T) {
			want := 555
			got := 555

			assert.Equal(t, want, got, "they should equal")
		})
		
		pp := &Person{Firstname: "Tawan"}

		if assert.NotNil(t, pp) {
			assert.Equal(t, "Tawan", pp.Firstname)
		}
	})
}