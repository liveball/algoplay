package main

import "testing"

func Benchmark_fbRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fbRecursion(10)
	}

	//b.Log(fbRecursion(10))
}

func Benchmark_fbMemo(b *testing.B) {
	tmap := make(map[int]int)
	for i := 0; i < b.N; i++ {
		fbMemo(10, tmap)
	}

	//b.Log(fbMemo(10, tmap))
}

func Benchmark_fbDp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fbDp(10)
	}

	//b.Log(fbDp(10))
}

func Benchmark_fbDp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fbDp2(10)
	}

	//b.Log(fbDp2(10))
}
