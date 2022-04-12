package main

import (
	"testing"
)

func BenchmarkFuncJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcjson(TestData2)
	}
}

func BenchmarkFuncgoJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcgojson(TestData2)
	}
}

func BenchmarkFuncJsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Funcjsoniter(TestData2)
	}
}

func BenchmarkFuncjsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncjsonMarshal(comData)
	}
}

func BenchmarkFuncGojsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncGojsonMarshal(comData)
	}
}

func BenchmarkFuncjsoniterMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncjsoniterMarshal(comData)
	}
}
