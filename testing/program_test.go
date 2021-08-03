package example

import "testing"

func Benchmark_Example_function(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yourFunction(10)
	}
}
