package main

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var primes []int

	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for _, num := range nums {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}

	println("Prime Numbers:")
	for _, prime := range primes {
		println(prime)
	}
}
