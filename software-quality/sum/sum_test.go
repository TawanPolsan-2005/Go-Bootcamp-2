package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should return 3 when 1 and 2", func(t *testing.T) {
		//Arrange
		want := 3

		//Act
		got := sum(1, 2)

		//Assert
		if got != want {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("Should return 1 when 1 and 0", func(t *testing.T) {
		//Arrange
		want := 1

		//Act
		got := sum(1, 0)

		//Assert
		if got != want {
			t.Errorf("sum(1, 0) = %d; want 1", got)
		}
	})

	t.Run("Should return -2 when -1 and -1", func(t *testing.T) {
		//Arrange
		want := -2

		//Act
		got := sum(-1, -1)

		//Assert
		if got != want {
			t.Errorf("sum(-1, -1) = %d; want -2", got)
		}
	})
}
