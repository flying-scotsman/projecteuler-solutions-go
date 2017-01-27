package main

import "testing"

func BenchmarkHello(b *testing.B) {
	// Run palindromeSummer 10^8 times
	for i := 0; i < b.N; i++ {
		palindromeSummer(100000000)
	}
}
