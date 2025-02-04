package main

import "log"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := range 42 {
		res := fib(i)
		log.Printf("fib(%d) = %d\n", i, res)
	}
}
