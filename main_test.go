package main

import "testing"

func benchmarkDemo(f func(), b *testing.B) {
	for n := 0; n < b.N; n++ {
		demo(f)
	}
}

func BenchmarkDemo1(b *testing.B) { benchmarkDemo(add, b) }
func BenchmarkDemo2(b *testing.B) { benchmarkDemo(mutexAdd, b) }
func BenchmarkDemo3(b *testing.B) { benchmarkDemo(atomicAdd, b) }
