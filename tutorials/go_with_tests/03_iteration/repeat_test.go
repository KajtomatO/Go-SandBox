package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectReturn := func(t testing.TB, repeated, expected string) {
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Sanity test", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		assertCorrectReturn(t, repeated, expected)
	})

	t.Run("Zero rep", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectReturn(t, repeated, expected)
	})

	t.Run("Other char", func(t *testing.T) {
		repeated := Repeat("z", 6)
		expected := "zzzzzz"

		assertCorrectReturn(t, repeated, expected)
	})

	t.Run("Repeat word", func(t *testing.T) {
		repeated := Repeat("zonk", 3)
		expected := "zonkzonkzonk"

		assertCorrectReturn(t, repeated, expected)
	})

	t.Run("No string", func(t *testing.T) {
		repeated := Repeat("", 631)
		expected := ""

		assertCorrectReturn(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("samurai", 125)
	}
}

func ExampleRepeat() {
	s := Repeat("foBar", 3)
	fmt.Println(s)
	// Output: foBarfoBarfoBar
}
