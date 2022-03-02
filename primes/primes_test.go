package primes

import (
	"testing"
)

func TestIsItAPrime(t *testing.T) {
	simpleTestTable := map[int]bool{
		0: false,
		1: false,
		2: true,
		3: true,
		4: false,
		5: true,
		6: false,
		7: true,
	}

	mediumTestTable := map[int]bool{
		7817: true,
		7818: false,
		7823: true,
		7829: true,
		7841: true,
		7843: false,
		7853: true,
		7867: true,
		7873: true,
		7877: true,
		7879: true,
		7883: true,
		7885: false,
		7901: true,
		7907: true,
		7919: true,
	}

	hardTestTabe := map[int]bool{
		2147483647: true,
		2147483648: false,
		2147483649: false,
	}

	t.Run("simple test", func(t *testing.T) {
		for prime, want := range simpleTestTable {
			got := isItAPrime(prime)

			if got != want {
				t.Errorf("for a number %d got %v want %v", prime, got, want)
			}
		}
	})

	t.Run("medium test", func(t *testing.T) {
		for prime, want := range mediumTestTable {
			got := isItAPrime(prime)

			if got != want {
				t.Errorf("for a number %d got %v want %v", prime, got, want)
			}
		}
	})

	t.Run("hard test", func(t *testing.T) {
		for prime, want := range hardTestTabe {
			got := isItAPrime(prime)

			if got != want {
				t.Errorf("for a number %d got %v want %v", prime, got, want)
			}
		}
	})
}
