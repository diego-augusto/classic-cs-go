package fib

func fib(p int) int {
	if p < 2 {
		return p
	}
	last, next := 0, 1
	for i := 1; i < p; i++ {
		last, next = next, last+next
	}
	return next
}
