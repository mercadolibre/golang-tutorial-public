package main

import (
	"testing"
)

// Benchmarks

func BenchmarkAdd(bt *testing.B) {
	for i := 0; i < bt.N; i++ {
		add(a, b)
	}
}

func BenchmarkAddSettingNLimit(bt *testing.B) {
	bt.N = 1000

	for i := 0; i < bt.N; i++ {
		add(a, b)
	}
}

func BenchmarkAddParallel(bt *testing.B) {
	bt.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				add(a, b)
			}
		})
}
