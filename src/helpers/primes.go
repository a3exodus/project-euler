package helpers

/*
input: n int
output: the sieve of Eratosthenes of length n, with sieve[p] iff p prime
*/
func SieveOfEratosthenes(n int) []bool {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	return primes
}
