package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should multi parameters", func(t *testing.T) {
		got := sum([]int{}...)

		want := 3

		if got != want {
			t.Errorf("sum(1, 2) = %d; want 3", got)
		}
	})

	t.Run("including sign integer", func(t *testing.T) {
		want := 7
		xs := []int{2, 3, 3, -1}
		got := sum(xs...)
		if got != want {
			t.Error("Expected", 8, "Got, got")
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
