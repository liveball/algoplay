package main

import "testing"

func Benchmark_climb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climb(10)
	}
}

//goos: darwin
//goarch: amd64
//pkg: algoplay/core/dynamic_programming/climb_stairs
//Benchmark_climb-12    	10000000	       159 ns/op
//PASS

func Benchmark_climbMemo(b *testing.B) {
	tmap := make(map[int]int)
	for i := 0; i < b.N; i++ {
		climbMemo(10, tmap)
	}
}

//goos: darwin
//goarch: amd64
//pkg: algoplay/core/dynamic_programming/climb_stairs
//Benchmark_climbMemo-12    	200000000	         7.18 ns/op
//PASS

func Benchmark_climbByDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbByDynamic(10)
	}
}

//goos: darwin
//goarch: amd64
//pkg: algoplay/core/dynamic_programming/climb_stairs
//Benchmark_climbByDynamic-12    	300000000	         5.87 ns/op
//PASS

func Benchmark_climbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(10)
	}
}

//goos: darwin
//goarch: amd64
//pkg: algoplay/core/dynamic_programming/climb_stairs
//Benchmark_climbStairs-4   	30000000	        46.2 ns/op
//PASS

