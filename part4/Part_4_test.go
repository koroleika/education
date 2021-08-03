package part4

import "testing"

func Benchmark_Part4_factRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factRec(10)
	}
}
func Benchmark_Part4_factCycleWhile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factCycleWhile(10)
	}
}
func Benchmark_Part4_factCycleFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factCycleFor(10)
	}
}
func Benchmark_Part4_factTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factTree(10)
	}
}
func Benchmark_Part4_factDegree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factDegree(10)
	}
}
