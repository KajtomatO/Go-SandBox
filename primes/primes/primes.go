package primes

type listOfNumbers []int
type listOfBools []bool

func IsItAPrime(in int) bool {

	if in <= 1 {
		return false
	}

	if in == 2 {
		return true
	}

	if in%2 == 0 {
		return false
	}

	for i := (in - 1); i > 1; i-- {
		if in%i == 0 {
			return false
		}
	}

	return true
}
