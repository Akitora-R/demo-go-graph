package main

import "testing"

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gen(getTxt(), getImg()[:1])
	}
}
