package main

import (
	"testing"
)

func BenchmarkFuncJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcjson(TestData)
	}
}

func BenchmarkFuncgoJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcgojson(TestData)
	}
}

func BenchmarkFuncJsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcjsoniter(TestData)
	}
}
