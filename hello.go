package main

import (
	"fmt"
	"time"
)

func main(){
	fibo()
}

func fibo() {
	fmt.Println("Quantas casas de Fibonacci você quer ver?")

	var n0 int;

	fmt.Scanf("%d", &n0)

	start := time.Now()

	result := IFibonacci(n0)

	elapsed := time.Since(start)

	fmt.Printf("O valor de Fibonacci pra %d é: %d\n", n0, result)

	fmt.Printf("It took %v ms to calculate", elapsed.Milliseconds())
}

var cache[] uint64 = make([]uint64, 2000)

func fibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	if n >= len(cache) {
		fmt.Println("Cache overflow!")
		var newCache[] uint64 = make([]uint64, n + 1)
		cache = append(cache, newCache...)
	}

	if cache[n] != 0 {
		return cache[n]
	}

	cache[n] = fibonacci(n-1) + fibonacci(n-2)

	return cache[n]
}

func IFibonacci(n int) uint64 {
	var a uint64 = 0
	var b uint64 = 1
	var c uint64 = 0

	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}